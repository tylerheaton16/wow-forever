# Waylaid Crates — Merchant's Favor Planner

A single-page tool for WoW Forever (US · Normal ruleset · Horde) that shows every
Waylaid Crate on the auction house, the items that can fill it, the cheapest fill,
gold-per-Merchant's-Favor rankings, and a reputation planner for Notarized Shipping
Container turn-ins (Durotar Supply and Logistics).

## Use it

- **Hosted:** enable GitHub Pages (Settings → Pages → deploy from `main`, root) and open
  `https://<you>.github.io/<repo>/`. The included GitHub Action refreshes prices every
  6 hours automatically — no cloning needed.
- **Local:** `./refresh-prices.sh && open index.html`
- **Double-click app:** grab `waylaid-crates.exe` (Windows) or `waylaid-crates-mac`
  (macOS, Apple Silicon) from `dist/` or the releases page. It fetches live prices on
  launch and opens the page in your browser — no setup. (macOS will require
  right-click → Open the first time, since the binary is unsigned.)

## Files

- `index.html` — the page (all logic client-side; reads `prices.js`)
- `prices.js` — generated price snapshot (committed so the page works out of the box)
- `refresh-prices.sh` — regenerates `prices.js` from TSM's public CSV
- `app/` — Go source for the double-click launcher; build with
  `go build` (use `GOOS=windows GOARCH=amd64` for the .exe)
- `.github/workflows/refresh.yml` — scheduled price refresh for the hosted page

## Data & attribution

Prices come from [TradeSkillMaster public pricing data](https://tradeskillmaster.com/public-data)
(no key required; attribution requested, bulk republication prohibited). Crate fill
requirements and Merchant's Favor values are from community guides; Expert and Artisan
Favor payouts are community-reported and not yet confirmed on live servers.

Code is MIT licensed. The price data itself remains subject to TSM's terms.
