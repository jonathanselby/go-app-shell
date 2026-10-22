# go-app-shell

[![test](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml/badge.svg)](https://github.com/jonathanselby/go-app-shell/actions/workflows/test.yml)

**0.1.0-dev — release candidate. A small application shell for serious Go web applications.**

`go-app-shell` provides the presentation frame substantial Go applications repeatedly need: responsive navigation, polished phone/tablet/desktop layouts, accessible defaults, light/dark presentation and optional conservative PWA plumbing.

> **The shell owns presentation. The application owns meaning and behaviour.**

## Why this exists

Admin frameworks often absorb data/auth/routing/business concepts. Frontend frameworks often replace JavaScript complexity with another browser runtime. `go-app-shell` does neither.

```text
application -> Page -> shell.Render -> semantic HTML + small CSS -> browser
```

## Properties

- public importable Go package;
- standard-library runtime only;
- framework-neutral `io.Writer` rendering;
- deterministic server-rendered HTML;
- 320px phone portrait through large desktop;
- zero-JS overlay navigation drawer on compact screens;
- persistent narrow tablet rail from 768px;
- tablet and desktop list/detail layouts;
- generic primary-action hierarchy;
- tinted notices/badges and balanced panel rhythm;
- accessible landmarks, focus and practical touch targets;
- automatic dark mode plus light/dark browser chrome;
- CSS cross-document view transitions where supported;
- zero JavaScript required for core use;
- no Node/npm/frontend build pipeline or WASM UI runtime;
- optional static-only PWA support with deterministic cache upgrades;
- graceful degradation to useful semantic HTML if JS or CSS disappears.

## Go support

**Supported and tested:** Go 1.27 and later. Earlier compatible Go 1.x releases may work, but are outside the supported matrix and are best-effort only.

## Minimal use

```go
r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
if err != nil { return err }

page := shell.Page{
    App: shell.AppHeader{Name: "Example"},
    Navigation: []shell.NavGroup{{Items: []shell.NavItem{{Label: "Overview", URL: "/", Current: true}}}},
    Title: "Overview",
    Actions: template.HTML(`<a class="gas-button gas-button--primary" href="/new">New</a>`),
    Workspace: shell.Workspace{Mode: shell.WorkspaceSingle, Content: trustedApplicationHTML},
}
return r.Render(w, page)
```

Application content slots use `template.HTML` deliberately: the consuming application owns that trust decision. Dynamic demo fragments use `html/template`, not string concatenation.

## Layout

```text
compact        320-767px   overlay navigation + one useful pane
tablet         768-1023px  persistent narrow rail + two-pane list/detail
desktop/wide   1024px+     full rail + wider workspace split
```

The shell changes presentation only. Selection, permissions and workflow remain application state.

## Presentation classes

V0.1 intentionally keeps the vocabulary small: panels/stacks/clusters, notices, forms/errors, lists/tables, key-value details, badges, compact/wide helpers, buttons and `gas-button--primary`.

No data-aware Go table/form builder exists.

## Light/dark and browser chrome

Light/dark presentation is CSS-token based and follows `prefers-color-scheme`. Renderer config also exposes `ThemeColorLight` / `ThemeColorDark` for matching browser/PWA chrome. There is no JavaScript theme engine.

## Navigation feel

Compact navigation is a zero-JS overlay drawer with scrim. Tablet/desktop use the same navigation data as a persistent rail. The shell intentionally does not provide bottom tabs or icon priority because those are application decisions.

Buttons and navigation have short CSS hover/pressed feedback. Native browser/touch feedback remains enabled.

## Page transitions

The stylesheet opts into same-origin cross-document View Transitions where the browser supports them. Unsupported browsers simply perform ordinary navigation. Reduced-motion preferences are respected.

The shell deliberately does **not** prefetch application pages. Authenticated/staleness/network policy belongs to the application.

## PWA boundary

PWA support is optional and static-only. `CacheVersion` lets a consuming app invalidate stable-URL precached assets with its own build/release identifier; fingerprinted URLs are equally valid. The shell never owns application data, mutations or build outputs.

## Demo

```bash
go run ./cmd/app-shell-demo
```

Open `http://127.0.0.1:8080`.

The demo uses its own `demoBuildVersion` for application cache invalidation rather than teaching callers to reuse `shell.Version`.

## Engineering gates

```bash
./scripts/acceptance.sh
go test -race -count=1 ./...
go test -run '^$' -bench BenchmarkRender -benchmem ./shell
```

At the current visual-hardening evidence snapshot, CSS + optional PWA JS is **14,503 bytes**, below the 32 KiB ceiling.

Human release acceptance covers Chromium/Firefox and WebKit/Safari where practical across phone, tablet and wide layouts, keyboard, 200% zoom, light/dark, no-JS, no-CSS, transitions, print and installed-PWA behaviour.

## What this does not do

No business logic, domain semantics, auth/permissions, routing ownership, DB access, workflow, client state framework, plugin engine, framework adapters, virtual DOM, hooks, hydration, WASM UI runtime, shell-owned partial rendering, speculative prefetch policy, bottom-tab application navigation, generic offline business data or mutation replay.

## Licence

MIT. See `LICENSE`.
