# WoW Forever Tools

Small web tools for WoW Forever (US · Normal ruleset · Horde), hosted with GitHub Pages
and auto-updated with live auction house prices every 6 hours.

## Live pages

| Tool | Link | What it does |
|---|---|---|
| **Waylaid Crates — Merchant's Favor Planner** | **<https://tylerheaton16.github.io/wow-forever/waylaid-crates/>** | Every Waylaid Crate on the AH with its price, the items that can fill it, the cheapest fill highlighted, gold-per-Favor rankings, and a reputation planner for Durotar Supply and Logistics (Honored / Revered / Exalted costs via Notarized Shipping Containers). |

## How it stays fresh

A scheduled GitHub Action ([`refresh.yml`](.github/workflows/refresh.yml)) pulls
[TradeSkillMaster public pricing data](https://tradeskillmaster.com/public-data) every
6 hours and commits the new snapshot; GitHub Pages redeploys automatically. Each page
shows its "snapshot fetched" time in the header.

See [`waylaid-crates/README.md`](waylaid-crates/README.md) for local usage, the
double-click desktop launchers, and data attribution details.
