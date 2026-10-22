# go-app-shell — Code-Derived Engineering Design

**Implementation evidence:** `883028ebfa4a69139ad1261ea664699df4d01477`  
**Generated:** 2026-10-22

## Recovered implementation

The implementation remains one public `shell` package plus a neutral demo. It owns presentation only: no DB, auth, workflow, business-domain subsystem, web-framework adapter or browser application runtime.

## Recovered execution path

```text
application -> shell.Page -> Renderer.Render(io.Writer, Page)
            -> html/template -> semantic HTML + embedded CSS
            -> optional manifest/service worker -> browser
```

## Recovered responsive model

- compact below 768px: sticky header, zero-JS fixed navigation drawer + scrim, one useful list/detail pane;
- tablet 768–1023px: persistent 11.5rem navigation rail and two-pane list/detail;
- desktop 1024px+: full 15rem navigation rail and wider list/detail split;
- same semantic page/navigation structure at every width;
- no device sniffing or resize controller.

## Recovered visual system

- generic secondary and primary actions;
- short CSS hover/pressed feedback;
- tinted generic notices and badges;
- panel first/last-child margin normalisation;
- light/dark token themes;
- light/dark browser `theme-color` metadata;
- CSS cross-document view transitions as progressive enhancement;
- reduced-motion override;
- no speculative prefetching, bottom-tab navigation or icon-priority system.

## Recovered PWA boundary

Static-only same-origin precache/offline fallback remains intact. Default cache identity uses shell version/assets/cache inputs plus optional application `CacheVersion`. The demo now uses an application-owned `demoBuildVersion`, not `shell.Version`, as the teaching pattern.

## Automated evidence

At `883028ebfa4a69139ad1261ea664699df4d01477`:

- GitHub Go 1.27 acceptance: PASS;
- race tests: PASS;
- Linux ARM64 build: PASS;
- standard-library-only dependency boundary: PASS;
- CSS + optional PWA JS: **14,503 bytes** (budget <= 32 KiB);
- render benchmark: approximately **27.2 µs/op, 11,957 B/op, 148 allocs/op** on GitHub AMD EPYC runner.

## Human acceptance still required

Human visual acceptance remains open for phone drawer behaviour, tablet rail/two-pane layout, desktop polish, primary-action hierarchy, badges/notices, light/dark browser chrome, transitions/reduced motion, keyboard/zoom, no-JS/no-CSS, print and installed-PWA behaviour.
