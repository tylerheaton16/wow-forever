// Waylaid Crates launcher: fetches live TSM prices for WoW Forever
// (US / Normal / Horde), serves the planner page on localhost, and
// opens it in the default browser. Double-click and go.
package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

//go:embed index.html
var page []byte

const csvURL = "https://public-data.tradeskillmaster.com/forever/us/realm/normal-horde/items.csv"

// buildPricesJS converts TSM's CSV into the window.PRICES blob the page expects.
// All items are included; the page only reads the ~126 it knows about.
func buildPricesJS() (string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(csvURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("TSM returned HTTP %d", resp.StatusCode)
	}
	r := csv.NewReader(resp.Body)
	r.FieldsPerRecord = -1
	header, err := r.Read()
	if err != nil || len(header) < 7 {
		return "", fmt.Errorf("unexpected CSV header: %v", err)
	}
	var b strings.Builder
	fmt.Fprintf(&b, "window.PRICES_META = { fetched: %q, market: \"US / Normal ruleset / Horde\" };\n", time.Now().UTC().Format(time.RFC3339))
	b.WriteString("window.PRICES = {\n")
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) < 7 {
			continue
		}
		name := strings.ReplaceAll(rec[1], `"`, `\"`)
		fmt.Fprintf(&b, "  %q: {name:\"%s\", mv:%s, min:%s, updated:%q},\n",
			rec[0], name, rec[2], rec[3], rec[6])
	}
	b.WriteString("};\n")
	return b.String(), nil
}

func openBrowser(url string) {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		exec.Command("xdg-open", url).Start()
	}
}

func main() {
	pricesJS, err := buildPricesJS()
	if err != nil {
		// Still serve the page; it shows its own "prices.js missing" state.
		pricesJS = fmt.Sprintf("console.error(%q);\n", "price fetch failed: "+err.Error())
	}

	// The embedded page's refresh hint mentions the shell script; not relevant here.
	served := strings.Replace(string(page),
		"Refresh with\n  <code>./refresh-prices.sh</code>.",
		"Prices are fetched live each time you launch this app.", 1)

	var lastHit atomic.Int64
	lastHit.Store(time.Now().Unix())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		lastHit.Store(time.Now().Unix())
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, served)
	})
	mux.HandleFunc("/prices.js", func(w http.ResponseWriter, req *http.Request) {
		lastHit.Store(time.Now().Unix())
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		io.WriteString(w, pricesJS)
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("could not open a local port:", err)
		return
	}
	url := fmt.Sprintf("http://%s/", ln.Addr().String())
	fmt.Println("Waylaid Crates planner:", url)
	openBrowser(url)

	// Exit once the page has been idle for a while; it's fully
	// client-side after load, so the server isn't needed anymore.
	go func() {
		for range time.Tick(time.Minute) {
			if time.Now().Unix()-lastHit.Load() > 30*60 {
				ln.Close()
				return
			}
		}
	}()
	http.Serve(ln, mux)
}
