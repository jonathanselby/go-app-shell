# go-app-shell — Code-Derived Engineering Evidence

**Recovered implementation:** `137b3d6b73be56e99a8819498094998d04f4ae5b`  
**Generated:** 2026-10-22

## Recovered runtime shape

```text
caller-owned route/auth/data/business behaviour
                  |
                  v
          shell.Page values
                  |
                  v
       Renderer.Render(io.Writer)
                  |
                  v
           html/template
                  |
       +----------+----------+
       |                     |
 semantic HTML          embedded assets
                       shell.css + pwa.js
       |                     |
       +----------+----------+
                  |
                browser
```

The reusable implementation is one public `shell` package plus a neutral `net/http` demonstration application. The shell owns no router, database, authentication service, application state store, workflow engine, plugin registry or client application runtime.

## Public package

The public package exposes explicit presentation data and an instance renderer:

- `Config`
- `Renderer`
- `AppHeader`
- `NavGroup` / `NavItem`
- `Page`
- `Workspace`
- `PWAHead`
- `PWAConfig` / `PWAIcon`
- `New`
- `Render`
- `AssetsFS`
- `Manifest`
- `ServiceWorker`

`Render` targets `io.Writer`, so framework compatibility is obtained without framework-specific adapters.

## State and rendering

Renderer configuration is instance state. Page/application state is caller input. There is no package-global mutable application state.

The renderer validates the workspace contract, executes the embedded `html/template` into a temporary buffer, then writes the completed document to the caller. Normal shell text is escaped; application-owned content slots are explicitly `template.HTML`.

Equivalent inputs have a deterministic-render regression test.

## Responsive presentation

One semantic document contains header, navigation and main workspace. CSS owns presentation changes:

```text
compact/mobile   < 600px
medium           600–1023px
desktop/wide     >= 1024px
```

Compact navigation uses a native checkbox/label control, so the core navigation requires no JavaScript. List/detail displays one caller-selected pane on compact widths and both panes at desktop width. Wide tables overflow inside their own wrapper rather than breaking the whole page.

## Presentation vocabulary

The stylesheet supplies generic classes for layout, notices, forms, lists, local table overflow, buttons/actions and empty states. There is no Go-side data/form/table builder and no business-record abstraction.

The default theme is controlled through namespaced `--gas-*` CSS custom properties and system fonts.

## PWA boundary

PWA behaviour is optional.

The manifest helper serialises caller-owned install metadata. The service-worker generator validates rooted same-origin inputs, precaches only the explicit static allowlist plus optional offline fallback, ignores non-GET requests and cross-origin requests, and uses the fallback only for failed navigation.

There is no mutation queue, IndexedDB business store, background-sync workflow or generic application-response cache.

## Degradation

The recovered implementation supports this deliberate degradation path:

```text
PWA available
 -> ordinary web
 -> JavaScript absent: server HTML + CSS
 -> CSS absent: semantic browser-default HTML
```

No degraded layer synthesises application/business state.

## Dependencies and budgets

- language baseline: Go 1.27;
- runtime Go dependencies: standard library only;
- required JavaScript: 0 bytes;
- WebAssembly: none;
- first-party `shell.css` + optional `pwa.js`: 9,532 bytes uncompressed at this snapshot;
- asset ceiling: 32 KiB.

The CI acceptance also records a non-gating render benchmark so allocation/time regression stays visible without turning machine-specific nanoseconds into an architecture contract.

## Executable evidence

Package tests cover:

- semantic landmarks and escaping;
- deterministic rendering;
- independent renderer instances;
- compact list/detail contract;
- zero-JavaScript core markup;
- optional PWA head integration;
- embedded asset/size contract;
- absence of selected business/client-runtime vocabulary in core Go source;
- manifest generation;
- refusal of external PWA persistence targets;
- GET-only/same-origin service-worker boundary;
- deterministic service-worker output.

`scripts/acceptance.sh` adds format, vet, race tests, dependency boundary, exact version, live neutral HTTP acceptance and benchmark evidence. GitHub CI additionally builds Linux ARM64.

## Human evidence still required before stable V0.1

Automated evidence does not prove browser aesthetics or every responsive/accessibility interaction. Stable V0.1 therefore still requires the human browser matrix in `RELEASE.md`: Chromium/Firefox and WebKit/Safari where practical, phone-to-wide widths, continuous resize, keyboard-only, 200% zoom, JavaScript disabled, CSS disabled, print and PWA/offline behaviour.
