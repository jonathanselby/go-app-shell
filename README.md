# go-app-shell

[![test](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml/badge.svg)](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml)

**0.1.0-dev — release candidate. A small application shell for serious Go web applications.**

`go-app-shell` provides the presentation frame substantial Go applications repeatedly need: responsive navigation, flexible workspace layouts, accessible defaults, light/dark presentation and optional conservative PWA plumbing.

> **The shell owns presentation. The application owns meaning and behaviour.**

It was created after the Garage repeatedly encountered the same gap while building larger Go applications. It is deliberately generic and intended for anyone who finds the same primitive useful.

## Why this exists

The alternatives tend to grow in one of two directions:

- admin frameworks absorb databases, auth, routing, plugins and business/data concepts into the presentation layer;
- frontend frameworks replace JavaScript with another browser application runtime: WASM, state systems, virtual DOMs, hooks, hydration and routing.

This component does neither.

```text
application -> Page -> shell.Render -> semantic HTML + small CSS -> browser
```

## Properties

- public importable Go package;
- standard-library runtime only;
- framework-neutral `io.Writer` rendering;
- deterministic server-rendered HTML;
- 320px phone portrait through large desktop;
- persistent desktop navigation and zero-JS compact navigation;
- single and list/detail workspaces;
- semantic forms/lists/tables plus small generic error, badge and key-value presentation conventions;
- accessible landmarks, skip link, visible focus and practical touch targets;
- automatic dark mode through CSS tokens and `prefers-color-scheme`;
- zero JavaScript required for core use;
- no Node/npm/frontend build pipeline;
- no WebAssembly UI runtime;
- optional manifest/service worker with explicit static-only cache boundary;
- deterministic shell-owned PWA cache upgrades without deleting caller-owned caches;
- graceful degradation to useful semantic HTML if JS or CSS disappears.

## Go support

**Supported and tested:** Go 1.27 and later.

The implementation deliberately avoids Go 1.27-specific language/runtime features and the module declares the lowest language/API level currently required by the source. Earlier compatible Go 1.x releases may therefore work, but they are outside the supported test matrix and are best-effort only; adopters using them accept that compatibility risk.

## Install

```bash
go get github.com/jonathanselby/go-app-shell@latest
```

## Minimal use

```go
r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
if err != nil {
    return err
}

page := shell.Page{
    App: shell.AppHeader{Name: "Example"},
    Navigation: []shell.NavGroup{
        {Items: []shell.NavItem{{Label: "Overview", URL: "/", Current: true}}},
    },
    Title: "Overview",
    Workspace: shell.Workspace{
        Mode: shell.WorkspaceSingle,
        Content: trustedApplicationHTML,
    },
}

return r.Render(w, page)
```

Normal shell text is escaped by `html/template`. Application content slots use `template.HTML` deliberately: the consuming application owns that trust decision. The neutral demo renders dynamic fragment content through `html/template`, not string concatenation.

## Assets and theming

Assets are embedded and exposed as `fs.FS`. The application may mount them through any web stack; no framework adapter is required.

The default theme uses namespaced `--gas-*` CSS variables. Light values are the default; dark values follow the browser/OS `prefers-color-scheme: dark` setting with no JavaScript theme runtime. Consumers can override tokens in a stylesheet loaded after `shell.css`.

## Layout

```text
compact        320-599px
medium         600-1023px
desktop/wide   1024px+
```

List/detail selection remains application/URL state. The shell changes presentation only: one pane on compact screens, both panes on wide screens.

## Presentation classes

V0.1 intentionally keeps the vocabulary small:

- `gas-panel`, `gas-stack`, `gas-cluster`;
- `gas-notice` plus generic tones;
- `gas-form`, `gas-field`, `gas-help`, `gas-field-error`, `gas-error-summary`;
- `gas-list`, `gas-table-wrap`, `gas-kv`;
- `gas-badge` plus generic tones;
- `gas-compact-only`, `gas-wide-only`;
- `gas-button`, `gas-empty`, `gas-muted`.

These style ordinary semantic HTML. There is no data-aware Go table/form builder and no business validation logic.

For server-side validation, the consuming application decides when a submission failed and may render the error summary with `tabindex="-1" autofocus` so keyboard focus moves to the summary without JavaScript. The demo shows this pattern; the shell does not focus anything automatically.

## PWA boundary

PWA support is optional. The service worker may cache only explicitly configured same-origin static URLs plus an optional offline fallback. It does **not** cache arbitrary authenticated pages/API data and does not queue or replay mutations.

When `CacheName` is not supplied, the generated cache name is derived from the shell version, embedded shell asset bytes, configured static cache inputs and optional `CacheVersion`.

Use `CacheVersion` when the consuming application precaches assets whose **content can change while their URL stays the same** — for example an application logo, offline page or application CSS. A normal build/release identifier is sufficient:

```go
cfg := shell.PWAConfig{
    Name:         "Example",
    Precache:     []string{"/app.css", "/logo.svg"},
    OfflineURL:   "/offline",
    CacheVersion: buildVersion,
}
```

Changing `CacheVersion` generates a new shell-owned cache identity; activation removes only older `go-app-shell-` caches. Fingerprinted application URLs are equally valid and need no separate version field. If the application supplies its own `CacheName`, the shell respects it and performs no prefix cleanup.

The shell never fetches or hashes application-owned asset contents itself.

## Degraded operation

```text
PWA unavailable  -> ordinary web app
JavaScript absent -> core navigation/forms still work
CSS absent        -> semantic browser-default HTML remains usable
```

## Demo

```bash
go run ./cmd/app-shell-demo
```

Open `http://127.0.0.1:8080`.

Disable the optional PWA layer with:

```bash
go run ./cmd/app-shell-demo -no-pwa
```

## Engineering gates

```bash
./scripts/acceptance.sh
go test -race -count=1 ./...
go test -run '^$' -bench BenchmarkRender -benchmem ./shell
```

GitHub CI runs the supported Go 1.27 baseline and also proves Linux ARM64 build, exact version, standard-library-only module dependencies, PWA boundaries and first-party asset size.

Human release acceptance covers Chromium/Firefox and WebKit/Safari where practical across phone-to-wide widths, continuous resize, keyboard-only operation, 200% zoom, light/dark, no-JS, no-CSS, print, PWA/offline and PWA-upgrade behaviour.

## What this does not do

No business logic, domain semantics, auth/permissions, routing ownership, database access, workflow, client state framework, plugin engine, framework adapters, virtual DOM, hooks, hydration, WASM UI runtime, shell-owned partial rendering, generic offline business data or mutation replay.

## Licence

MIT. See `LICENSE`.
