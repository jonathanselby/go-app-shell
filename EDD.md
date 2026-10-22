# go-app-shell — Engineering Design Document

**Version:** 0.1

## 1. Purpose

Provide a small reusable Go presentation shell for substantial browser applications, responsive from phone portrait to large desktop, accessible by default, server-rendered, and optionally PWA-capable.

> **The shell owns presentation. The application owns meaning and behaviour.**

## 2. Design constraints

1. No business logic or domain semantics.
2. Routing, auth, permissions, validation, persistence and application state remain upstream.
3. Public importable package; not hidden under `internal/`.
4. Standard-library runtime only.
5. Core UI works with zero JavaScript.
6. No browser-side Go/WASM runtime, virtual DOM, reconciler, hooks, hydration or client router.
7. No Node/npm, frontend framework, code generation or frontend build step.
8. One semantic server-rendered structure adapts by available width; no separate mobile application.
9. No framework-adapter layer. Broad compatibility comes from standard Go boundaries.
10. No global engine/service/plugin runtime. Renderer instances are explicit and independent.
11. V0.1 is an application frame, not a widget/data/admin framework.
12. Optional PWA persistence is static-shell only; business data remains application-owned.

## 3. Execution model

```text
request
  |
  v
consuming application
routes / auth / data / business behaviour
  |
  | constructs presentation values
  v
shell.Render(io.Writer, Page)
  |
  v
html/template
  |
  +--> semantic HTML
  +--> embedded shell.css
  +--> optional manifest / tiny PWA registration / service worker
  |
  v
browser
```

A developer should be able to explain this path without hidden lifecycle or client-runtime machinery.

## 4. Repository layout

```text
go-app-shell/
├── .github/workflows/{test.yml,release-tag.yml}
├── PRD.md / EDD.md / README.md / RELEASE.md / CHANGELOG.md / VERSION
├── architecture.txt / LICENSE / go.mod
├── scripts/acceptance.sh
├── shell/
│   ├── shell.go / pwa.go / tests
│   ├── assets/{shell.css,pwa.js}
│   └── templates/shell.html
├── cmd/app-shell-demo/main.go
├── examples/neutral-demo/README.md
└── generated/{EDD.md,COMPARISON.md}
```

The reusable API is public package `shell`. The demo is acceptance evidence, not a dependency.

## 5. Public rendering API

The core is an explicit renderer:

```go
r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
err = r.Render(w, page)
```

`Render` accepts `io.Writer`, keeping it independent of `net/http` and third-party frameworks.

The model contains generic presentation concepts only: `AppHeader`, `NavGroup`, `NavItem`, `Page`, `Workspace`, `PWAHead`.

Workspace modes are `single` and `list-detail`. Selected-record state remains URL/request/application state; the shell stores none.

## 6. Trusted HTML boundary

Normal shell text is escaped by `html/template`. Application-owned content/action slots are `template.HTML`, making the trust decision explicit at the caller boundary. The shell exposes no helper that converts arbitrary untrusted strings into trusted HTML.

## 7. Asset boundary and budgets

CSS and optional registration JavaScript are embedded and exposed as `fs.FS`; the consuming application chooses how to mount them.

```text
shell.css + optional pwa.js <= 32 KiB uncompressed
required JavaScript            0 bytes
third-party runtime deps       0
WASM                           0
```

System fonts are the default.

## 8. Responsive layout

```text
320-599px   compact
600-1023px  medium
1024px+     desktop/wide
```

These are available-space ranges, not device identities.

Compact uses zero-JS navigation, one useful list/detail pane, full-width forms and local table overflow. Desktop exposes persistent left navigation, a flexible workspace and simultaneous list/detail panes. CSS Grid/Flexbox/media queries own adaptation; there is no resize controller or device sniffing.

## 9. Navigation

Navigation is caller-supplied metadata. The application decides what a user may see. Compact navigation uses native HTML/CSS state and current navigation uses `aria-current="page"`.

## 10. Presentation conventions, not business components

The shell styles ordinary semantic application HTML through a small class vocabulary for panels, stacks, notices, fields, lists, table wrappers, actions and empty states.

There are intentionally no Go APIs such as `NewTable`, `AddColumn`, `BindRecord` or database-aware field builders. The application writes ordinary `<form>`, `<table>`, `<ul>` and related HTML and retains all semantics.

## 11. Theming

A small namespaced `--gas-*` custom-property set covers dimensions, spacing, borders, text, surfaces, accent/focus and generic notice tones. Consumers override tokens in CSS loaded after `shell.css`. There is no theme plugin/runtime/version system.

## 12. Accessibility

The shell owns semantic landmarks, skip link, logical order, visible `:focus-visible`, current-navigation semantics, 44px practical target sizing, non-colour-only notices, 200% zoom reflow, reduced-motion respect and useful semantic markup when CSS is missing.

The application owns labels/content hierarchy inside its fragments, image alternatives, custom widgets and contrast of its overrides.

## 13. Progressive Web App boundary

PWA support is optional. `PWAConfig` supplies manifest metadata and an explicit static precache list. The generated service worker may precache explicit same-origin static URLs and an offline fallback. It does not cache arbitrary authenticated HTML/API responses, queue/replay mutations, store business records, synchronise writes or resolve conflicts.

Failed/unsupported service workers leave the ordinary web application unchanged.

## 14. Failure and degradation model

```text
full web + optional PWA
        -> ordinary web
        -> no JavaScript: server HTML + CSS
        -> no CSS: semantic browser-default HTML
```

The shell never invents stale business data or application state. Invalid renderer/PWA configuration returns errors rather than silently broadening scope.

## 15. Security boundary

The shell is not an authentication, authorisation, CSRF or security-header component. PWA precache/icon/offline targets must be rooted same-origin URLs. Security headers, sessions, auth and application content policy remain with the consuming server.

## 16. Neutral reference application

`app-shell-demo` uses generic records/forms only to prove header/navigation, single/list-detail workspaces, ordinary server POST, form/table/list presentation, generic notices, responsive layout and optional PWA plumbing. It contains no reusable business logic.

## 17. Test strategy

Automated evidence covers escaped text, semantic landmarks, deterministic rendering, independent renderer instances, no-JS navigation, list/detail compact mode, embedded assets/budget, absence of domain/client-runtime vocabulary in core source, manifest generation, same-origin PWA targets, GET-only service-worker behaviour and deterministic service-worker generation.

CI adds `gofmt`, `go vet`, race tests, native build, exact version, Linux ARM64 build, HTTP acceptance, standard-library-only module contract, render benchmark and exact asset bytes. Benchmark numbers are evidence for regression visibility, not brittle nanosecond gates.

Human release acceptance covers Chromium/Firefox and WebKit/Safari where practical at 320/390/768/1024/1366/1920/2560 widths, continuous resize, keyboard-only, 200% zoom, JavaScript disabled, CSS disabled, print sanity and PWA/offline behaviour where supported.

## 18. Known limitations

Not a general widget/design-system library; no charts/data-grid/modal framework; no generic offline business data; PWA install UX varies by platform; caller content/themes can still harm accessibility; visual acceptance remains partly human; 320 CSS px is the V0.1 lower acceptance width; richer interaction including optional HTMX remains application-owned.

## 19. Acceptance-to-evidence map

| Contract | Evidence |
|---|---|
| framework-neutral deterministic render | package API, `TestRenderDeterministic` |
| escaped text/semantic landmarks | `TestRenderSemanticShellAndEscapesText` |
| independent instances | `TestIndependentRenderers` |
| no-JS navigation | `TestCoreMarkupDoesNotDependOnJavaScript`, HTTP demo acceptance |
| list/detail adaptive contract | `TestListDetailCompactPaneContract`, browser matrix |
| asset/dependency budgets | `TestAssetsAndSizeBudget`, CI byte/module checks |
| no business/client runtime in core | `TestNoBusinessVocabularyOrClientRuntimeInCoreSource` |
| optional/static-only PWA | PWA tests and service-worker boundary tests |
| degraded CSS/no-JS operation | semantic markup tests + human acceptance |
| ARM64 | CI cross-build |
| render/resource visibility | `BenchmarkRender` + CI asset bytes |
| real-browser behaviour | release acceptance matrix |
| implementation matches design | `generated/EDD.md`, `generated/COMPARISON.md` |
