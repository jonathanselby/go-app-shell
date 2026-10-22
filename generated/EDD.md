# go-app-shell — Code-Derived Engineering Design

**Implementation evidence:** `762bc87d4643fc33e16a5c82f199f59583595f71`  
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
- closed compact drawer uses `visibility:hidden`/`pointer-events:none`, so off-screen navigation is not keyboard-reachable;
- open compact drawer locks background page scrolling where `:has()` is supported;
- tablet 768–1023px: persistent 11.5rem navigation rail and two-pane list/detail;
- desktop 1024px+: full 15rem navigation rail and wider list/detail split;
- detail pane is an inline-size container; key-value layout stacks when the pane itself is narrow rather than relying only on viewport width;
- same semantic page/navigation structure at every width;
- no device sniffing or resize controller.

## Recovered mobile/native-fit model

- `viewport-fit=cover` is paired with `safe-area-inset-*` variables;
- header, main content, compact navigation and skip link respect device safe areas;
- `100dvh` is preferred with `100vh` fallback for mobile browser chrome changes;
- compact controls with explicit pressed states suppress WebKit tap highlight so two feedback systems do not compete;
- background scroll lock, safe areas and dynamic viewport behaviour are CSS-only progressive enhancement.

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

Static-only same-origin precache/offline fallback remains intact. Default cache identity uses shell version/assets/cache inputs plus optional application `CacheVersion`. The demo uses an application-owned build version rather than `shell.Version` as the teaching pattern.

## Automated evidence

At implementation commit `762bc87d4643fc33e16a5c82f199f59583595f71`:

- GitHub Go 1.27 acceptance: PASS;
- race tests: PASS;
- Linux ARM64 build: PASS;
- standard-library-only dependency boundary: PASS;
- asset budget: PASS (< 32 KiB CSS + optional JS);
- tests assert the closed-drawer visibility contract, pane container-query contract, safe-area variables, dynamic viewport units and bounded tap-highlight suppression.

## Human acceptance still required

Human visual acceptance remains open for phone drawer focus order, background scroll lock, tablet narrow-pane behaviour, real-device safe areas, mobile browser chrome changes, desktop polish, primary-action hierarchy, badges/notices, light/dark browser chrome, transitions/reduced motion, keyboard/zoom, no-JS/no-CSS, print and installed-PWA behaviour.
