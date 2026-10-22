# go-app-shell

[![test](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml/badge.svg)](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml)

**A small application shell for serious Go web applications.**

`go-app-shell` provides the presentation frame that substantial Go applications repeatedly need: responsive navigation, a flexible workspace, list/detail presentation, accessible defaults and optional conservative PWA plumbing.

> **The shell owns presentation. The application owns meaning and behaviour.**

It was created after the Garage repeatedly encountered the same gap while building larger Go applications. It is deliberately generic and intended for anyone who finds the same Go presentation primitive useful.

## Why this exists

Admin frameworks tend to absorb databases, auth, routing, plugins and data concepts into presentation. Frontend frameworks often replace JavaScript with another browser runtime: WASM, state systems, virtual DOMs, hooks, hydration and routing.

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
- ordinary semantic form/list/table conventions rather than data-aware builders;
- accessible landmarks, skip link, visible focus and practical touch targets;
- documented CSS custom properties for theming;
- zero JavaScript required for core use;
- no Node/npm/frontend build pipeline;
- no WebAssembly UI runtime;
- optional manifest/service worker with explicit static-only cache boundary;
- graceful degradation to useful semantic HTML if JS or CSS disappears.

## Install

Requires Go 1.27.

```bash
go get github.com/jonathanselby/go-app-shell@latest
```

## Minimal use

```go
r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
if err != nil { return err }

page := shell.Page{
    App: shell.AppHeader{Name: "Example"},
    Navigation: []shell.NavGroup{{Items: []shell.NavItem{{Label: "Overview", URL: "/", Current: true}}}},
    Title: "Overview",
    Workspace: shell.Workspace{Mode: shell.WorkspaceSingle, Content: trustedApplicationHTML},
}

return r.Render(w, page)
```

Normal text values are escaped by `html/template`. Application content slots use `template.HTML` deliberately: the consuming application owns that trust decision.

## Assets

```go
assets, _ := shell.AssetsFS()
mux.Handle("GET /_app-shell/", http.StripPrefix("/_app-shell/", http.FileServer(http.FS(assets))))
```

The application may mount them through any web stack. No framework adapter is required.

## Layout

```text
compact        320-599px
medium         600-1023px
desktop/wide   1024px+
```

List/detail selection remains application/URL state. The shell only changes presentation: one pane on compact screens, both panes on wide screens.

## Presentation classes

V0.1 intentionally keeps the vocabulary small: `gas-panel`, `gas-stack`, `gas-cluster`, generic notices, `gas-form`, `gas-field`, `gas-list`, `gas-table-wrap`, `gas-button`, `gas-empty` and `gas-muted`.

These style ordinary semantic HTML. There is no data-aware Go table/form builder.

## Theming

Override documented `--gas-*` CSS custom properties in a stylesheet loaded after `shell.css`. No theme plugin/runtime is required.

## PWA boundary

PWA support is optional. The service worker may cache only explicitly configured same-origin static URLs plus an optional offline fallback. It does **not** cache arbitrary authenticated pages/API data and does not queue or replay mutations.

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

## Engineering gates

```bash
./scripts/acceptance.sh
go test -race -count=1 ./...
go test -run '^$' -bench BenchmarkRender -benchmem ./shell
```

CI also proves native and Linux ARM64 builds, exact version, standard-library-only module dependencies, PWA boundaries and first-party asset size.

Human release acceptance covers Chromium/Firefox and WebKit/Safari where practical across phone-to-wide widths, continuous resize, keyboard-only operation, 200% zoom, no-JS, no-CSS, print and PWA/offline behaviour.

## What this does not do

No business logic, domain semantics, auth/permissions, routing ownership, database access, workflow, client state framework, plugin engine, framework adapters, virtual DOM, hooks, hydration, WASM UI runtime, generic offline business data or mutation replay.

## Engineering artefacts

- `PRD.md` — product boundary and acceptance contract;
- `EDD.md` — engineering design and degradation boundaries;
- `architecture.txt` — one-screen architecture view;
- `shell/` — reusable package;
- `cmd/app-shell-demo/` — neutral reference application;
- `scripts/acceptance.sh` — executable acceptance;
- `generated/EDD.md` — code-derived implementation evidence;
- `generated/COMPARISON.md` — intended design versus implementation;
- `RELEASE.md`, `CHANGELOG.md`, `VERSION` — release discipline.

## Licence

MIT. See `LICENSE`.
