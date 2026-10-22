# go-app-shell — Engineering Design Document

**Version:** 0.1

## 1. Purpose

Provide a small reusable Go presentation shell for substantial browser applications, responsive from phone portrait to large desktop, accessible by default, server-rendered, dark-mode capable, and optionally PWA-capable.

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
13. Dark mode is a CSS-token presentation state, not an application runtime.

## 3. Go support and compatibility

Go 1.27+ is the supported/tested baseline. CI remains authoritative on Go 1.27.

`go.mod` declares the lowest language/API level currently required by the source (`go 1.22`) so earlier compatible toolchains are not rejected for policy reasons alone. Earlier Go 1.x releases are outside the support matrix and are best-effort only.

## 4. Execution model

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

## 5. Public rendering API

The core remains an explicit renderer:

```go
r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
err = r.Render(w, page)
```

`Render` accepts `io.Writer`, keeping it independent of `net/http` and third-party frameworks.

Workspace modes remain `single` and `list-detail`; selected-record state remains URL/request/application state.

## 6. Trusted HTML boundary

Normal shell text is escaped by `html/template`. Application-owned content/action slots remain `template.HTML`, making the trust decision explicit at the caller boundary. The reference demo renders dynamic fragments through `html/template` rather than string concatenation, demonstrating the safe consumption pattern without adding a component framework.

## 7. Asset boundary and budgets

CSS and optional registration JavaScript are embedded and exposed as `fs.FS`.

```text
shell.css + optional pwa.js <= 32 KiB uncompressed
required JavaScript            0 bytes
third-party runtime deps       0
WASM                           0
```

System fonts remain the default.

## 8. Responsive layout

```text
320-599px   compact
600-1023px  medium
1024px+     desktop/wide
```

Compact uses zero-JS navigation, one useful list/detail pane, full-width forms and local table overflow. Desktop exposes persistent left navigation, a flexible workspace and simultaneous list/detail panes. `gas-compact-only` / `gas-wide-only` are presentation-only visibility helpers.

## 9. Presentation conventions

The shell styles ordinary semantic HTML through a deliberately small class vocabulary.

V0.1 includes panels/stacks/clusters, empty states, generic notices/badges, form help/error/error-summary presentation, ordinary lists/tables, key-value `<dl>` presentation and compact/wide visibility helpers.

There are intentionally no business/data builders such as `NewTable`, record binding, workflow state or validation logic.

Invalid-control styling is scoped beneath `.gas-field` so the shell does not restyle unrelated controls. The consuming application decides when validation fails. It may focus an application-rendered error summary using native HTML (`tabindex="-1" autofocus`); the shell never performs focus changes automatically.

## 10. Light/dark theming

A small namespaced `--gas-*` token set covers dimensions, spacing, borders, text, surfaces, accent/focus and generic notice tones.

Default light values live at `:root`. Dark values are supplied under `@media (prefers-color-scheme: dark)` using the same tokens. No JavaScript state, theme registry or plugin system exists. Consumers may override tokens after `shell.css` if they require branded or explicit theme behaviour.

Print explicitly returns to a light palette.

## 11. Accessibility

The shell owns semantic landmarks, skip link, logical order, visible `:focus-visible`, current-navigation semantics, 44px practical target sizing, non-colour-only notices/badges, scoped `aria-invalid` styling hooks, 200% zoom reflow, reduced-motion respect and useful semantic markup when CSS is missing.

The application owns validation decisions, error-summary focus decisions, field messages, content hierarchy inside fragments, image alternatives, custom widgets and contrast of its overrides.

## 12. Progressive Web App boundary

PWA support remains optional. `PWAConfig` supplies manifest metadata and an explicit static precache list.

The generated service worker may precache explicit same-origin static URLs and an offline fallback. It never caches arbitrary authenticated HTML/API responses, queues/replays mutations, stores business records, synchronises writes or resolves conflicts.

### Default cache lifecycle

If the caller does not provide `CacheName`, the shell derives a deterministic cache name from:

- `shell.Version`;
- embedded `shell.css` and `pwa.js` bytes;
- the normalised precache/offline URL set;
- optional caller-supplied `CacheVersion`.

`CacheVersion` exists only for application-owned precached resources whose contents can change while their URLs remain stable. A build or release identifier is sufficient. Fingerprinted URLs remain an equally valid alternative.

The shell deliberately does **not** fetch/hash application-owned resource contents or own the application's build system.

During activation the worker deletes only older caches with the shell-owned `go-app-shell-` prefix. If the caller provides an explicit `CacheName`, the shell respects it and disables prefix cleanup.

## 13. Failure and degradation model

```text
full web + optional PWA
        -> ordinary web
        -> no JavaScript: server HTML + CSS
        -> no CSS: semantic browser-default HTML
```

The shell never invents stale business data or application state. Failed/unsupported service workers leave ordinary web behaviour unchanged.

## 14. Neutral reference application

The demo proves header/navigation, single/list-detail workspaces, ordinary server POST, generic form validation presentation, native error-summary focus after invalid POST, badges, key-value detail, responsive visibility, safe dynamic fragment templating, light/dark CSS and optional PWA plumbing. It contains no reusable business logic.

The demo sets `CacheVersion` to the application version, demonstrating the expected stable-URL application-asset upgrade pattern.

## 15. Test strategy

Automated evidence covers deterministic rendering, escaping, independent renderer instances, no-JS navigation, list/detail compact mode, dark-mode/presentation CSS contracts, asset budget, absence of domain/client-runtime vocabulary, manifest generation, same-origin PWA targets, GET-only service-worker behaviour, deterministic cache naming, cache-version invalidation, old shell-cache cleanup boundaries and explicit custom-cache isolation.

CI on Go 1.27 adds `gofmt`, `go vet`, race tests, native build, exact version, Linux ARM64 build, HTTP/form acceptance, standard-library-only module contract, render benchmark and exact asset bytes.

Human release acceptance covers Chromium/Firefox and WebKit/Safari where practical at 320/390/768/1024/1366/1920/2560 widths, continuous resize, keyboard-only, 200% zoom, light/dark, JavaScript disabled, CSS disabled, print sanity, PWA install/offline and PWA upgrade behaviour.

## 16. Known limitations

Not a general widget/design-system library; no charts/data-grid/modal/timeline/pagination/filter framework; no generic offline business data; no shell-owned partial-render runtime; PWA install UX varies by platform; caller content/themes can still harm accessibility; visual acceptance remains partly human; 320 CSS px is the V0.1 lower acceptance width.

## 17. Acceptance-to-evidence map

| Contract | Evidence |
|---|---|
| framework-neutral deterministic render | package API, `TestRenderDeterministic` |
| escaped text/semantic landmarks | `TestRenderSemanticShellAndEscapesText` |
| independent instances | `TestIndependentRenderers` |
| no-JS navigation | `TestCoreMarkupDoesNotDependOnJavaScript`, HTTP demo acceptance |
| list/detail adaptive contract | `TestListDetailCompactPaneContract`, browser matrix |
| light/dark + small presentation vocabulary | `TestAssetsAndSizeBudget`, browser matrix |
| form-error presentation | demo 422 HTTP acceptance + browser review |
| asset/dependency budgets | `TestAssetsAndSizeBudget`, CI byte/module checks |
| no business/client runtime in core | `TestNoBusinessVocabularyOrClientRuntimeInCoreSource` |
| shell-asset PWA upgrade | default cache hash + cleanup tests |
| app-asset PWA upgrade | `TestDefaultCacheNameChangesWithApplicationVersion` + manual installed upgrade |
| caller cache isolation | `TestCustomCacheNameDoesNotDeleteUnrelatedCaches` |
| degraded CSS/no-JS operation | semantic markup tests + human acceptance |
| ARM64 | CI cross-build |
| supported Go baseline | GitHub CI Go 1.27 |
| render/resource visibility | `BenchmarkRender` + CI asset bytes |
| real-browser behaviour | release acceptance matrix |
| implementation matches design | `generated/EDD.md`, `generated/COMPARISON.md` |
