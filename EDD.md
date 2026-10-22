# go-app-shell — Engineering Design Document

**Version:** 0.1

## 1. Purpose

Provide a small reusable Go presentation shell for substantial browser applications, responsive from phone portrait to large desktop, accessible by default, server-rendered, dark-mode capable and optionally PWA-capable.

> **The shell owns presentation. The application owns meaning and behaviour.**

## 2. Hard boundaries

- no business/domain logic;
- no DB/auth/workflow/routing ownership;
- public importable package;
- standard-library runtime only;
- zero required JavaScript;
- no WASM/VDOM/reconciler/hooks/hydration/client router;
- no framework adapters or global engine;
- no shell-owned speculative prefetch/navigation policy;
- no application-priority navigation model;
- no generic offline business data.

## 3. Go support

Go 1.27+ is supported/tested. `go.mod` declares the lowest language/API level currently required by source so earlier compatible releases are not rejected unnecessarily; they remain best-effort only.

## 4. Execution model

```text
request -> consuming application -> presentation values
        -> shell.Render(io.Writer, Page)
        -> html/template -> semantic HTML + embedded CSS
        -> optional manifest/service worker -> browser
```

## 5. Public API

`Renderer.Render(io.Writer, Page)` remains framework-neutral. Presentation data consists only of generic application header, navigation, page, action and workspace values.

`Config` includes `ThemeColorLight` and `ThemeColorDark` so browser/PWA chrome can track light/dark presentation without hard-coding application branding.

## 6. Trusted HTML boundary

Normal shell text is escaped by `html/template`. Application-owned fragment/action slots remain explicit `template.HTML` trust boundaries. The demo renders dynamic fragments through `html/template` rather than string concatenation.

## 7. Responsive layout

```text
320-767px   compact
768-1023px  tablet
1024px+     desktop/wide
```

### Compact

- sticky top header;
- zero-JS checkbox-controlled navigation;
- navigation is a fixed overlay drawer with scrim, not document-flow content;
- one useful list/detail pane at a time;
- full-width form/action treatment where useful.

### Tablet

- persistent narrow navigation rail (`--gas-nav-width-tablet`);
- list/detail shown together;
- narrower list column and smaller inter-pane gap than desktop;
- compact-only controls hidden.

### Desktop/wide

- full navigation rail (`--gas-nav-width`);
- larger list/detail split and spacing;
- readable content remains bounded by application markup rather than artificial shell cards.

CSS owns all adaptation; no device detection or resize controller exists.

## 8. Visual hierarchy

V0.1 provides generic secondary and primary action treatments. `gas-button--primary` supplies hierarchy only; the consuming application decides which action is primary.

Badges/notices use light tonal backgrounds plus text labels so they read as deliberate interface elements in both light and dark themes without making colour authoritative.

Panels normalise first/last-child vertical margins so headings and trailing content do not create accidental asymmetric whitespace.

Short CSS transitions and `:active` feedback apply to buttons/navigation/list links. Native browser/touch feedback is not disabled.

## 9. Navigation polish

Phone navigation uses the existing zero-JS toggle plus a fixed drawer and click/tap scrim. The shell deliberately does not implement bottom navigation or icon requirements because those require application-level prioritisation.

At 768px the drawer becomes a persistent rail using the same navigation data and markup.

## 10. Document transitions

The shell enables CSS cross-document view transitions through `@view-transition { navigation: auto; }`. This is progressive enhancement: browsers without support use normal MPA navigation.

Reduced-motion CSS suppresses meaningful transition duration. No JavaScript navigation runtime is introduced.

The shell does **not** emit speculation rules/prefetch instructions; only the application can know whether an authenticated destination is safe/useful to prefetch.

## 11. Light/dark and browser chrome

Namespaced `--gas-*` tokens define light presentation; `prefers-color-scheme: dark` overrides the same tokens. No JavaScript theme runtime exists.

The document head emits light/dark `theme-color` metadata from renderer configuration and adds `viewport-fit=cover`. Print restores light presentation.

## 12. Presentation conventions

V0.1 remains deliberately small: panels/stacks/clusters, empty states, notices/badges, form help/error/error-summary, ordinary lists/tables, key-value `<dl>`, compact/wide helpers and primary action styling.

There are no data-aware builders, timelines, pagination engines, filter frameworks, modals, icon registries or component runtimes.

## 13. Accessibility

The shell owns landmarks, skip link, logical source order, visible focus, `aria-current`, 44px practical targets, non-colour-only status presentation, scoped invalid-control styling, reduced-motion behaviour, responsive zoom and semantic no-CSS degradation.

Application code owns validation decisions, error-summary focus decisions, fragment content, image alternatives and custom widgets.

## 14. PWA boundary

PWA support remains optional and static-only. Default cache identity incorporates shell version, embedded shell assets, normalised cache URLs and optional application `CacheVersion`. Explicit caller cache names disable shell prefix cleanup.

The demo now uses a distinct application build variable (`demoBuildVersion`) for `CacheVersion`; it does not reuse `shell.Version` as application versioning guidance.

## 15. Resource budget and evidence

```text
CSS + optional PWA JS ceiling  <= 32 KiB uncompressed
required JS                    0 bytes
third-party runtime deps       0
WASM                           0
```

At implementation evidence `883028ebfa4a69139ad1261ea664699df4d01477`:

- Go 1.27 CI acceptance passed;
- race tests passed;
- Linux ARM64 build passed;
- CSS + optional PWA JS: 14,503 bytes;
- render benchmark: ~27.2 µs/op, 11,957 B/op, 148 allocs/op on the GitHub AMD runner.

Benchmark values are regression evidence, not brittle gates.

## 16. Human acceptance

Human review remains required for Chrome/Firefox and WebKit/Safari where practical at 320/390/768/820/1024/1366/1920/2560 widths, continuous resize, keyboard, 200% zoom, light/dark, phone drawer, tablet rail/two-pane layout, transitions, no-JS, no-CSS, print and installed-PWA behaviour.

## 17. Known limitations

No general component/design-system runtime; no bottom-tab navigation; no shell prefetch policy; no app icons; no charts/data-grid/modal/timeline/pagination/filter framework; no generic offline business data. Visual acceptance remains partly human by design.
