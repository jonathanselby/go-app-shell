# go-app-shell — Code-Derived Engineering Design

**Implementation evidence:** `dcd82bd1d343f8785956f592bf1c189218bd2dde`  
**Generated:** 2026-10-22

## Recovered implementation

The implementation is a public Go package named `shell` plus a neutral demonstration application. The reusable package owns presentation only and contains no database, authentication, authorisation, workflow, routing framework or business-domain subsystem.

### Core execution path

```text
caller application
  -> shell.Page / presentation values
  -> Renderer.Render(io.Writer, Page)
  -> html/template
  -> semantic HTML
  -> embedded shell.css
  -> optional manifest/service worker
  -> browser
```

There is no client application runtime, virtual DOM, hooks, hydration, WASM renderer or shell-owned router.

## Recovered layout and presentation

- compact: 320–599 CSS px;
- medium: 600–1023 CSS px;
- desktop/wide: 1024 CSS px and above;
- same semantic structure at all sizes;
- compact list/detail shows one useful pane; desktop exposes both;
- compact navigation remains zero-JavaScript;
- light/dark presentation uses CSS tokens and `prefers-color-scheme` only;
- generic presentation includes notices, badges, form error/error summary, key-value `<dl>` and compact/wide visibility helpers;
- invalid-control styling is scoped to `.gas-field`;
- no Go-side business/data form or table builders exist.

## Recovered validation boundary

The application decides whether validation failed. The shell supplies presentation hooks only. The neutral demo returns a shell-rendered 422 page and demonstrates native zero-JS error-summary focus with `tabindex="-1" autofocus`; the shell itself performs no focus or validation behaviour.

## Recovered PWA boundary and upgrade model

The generated service worker:

- accepts only same-origin rooted precache/offline inputs;
- intercepts GET only;
- precaches only configured static URLs and optional offline fallback;
- does not use IndexedDB, queues, background sync or mutation replay;
- derives the automatic cache identity from `shell.Version`, embedded shell asset bytes, normalised cache inputs and optional `CacheVersion`;
- uses `CacheVersion` only as caller-supplied invalidation input for app-owned resources with stable URLs;
- deletes only older caches with the shell-owned `go-app-shell-` prefix;
- disables prefix cleanup when a caller provides an explicit `CacheName`.

The shell never fetches or hashes application-owned resource contents.

## Recovered support model

`go.mod` declares `go 1.22`, matching the lowest language/API level used by the current source/demo. GitHub CI runs Go 1.27.x, which is the documented supported/tested baseline. Earlier compatible releases remain best-effort and unsupported.

## Automated evidence

At implementation commit `dcd82bd1d343f8785956f592bf1c189218bd2dde` on GitHub Go 1.27.3:

- format, vet and race tests passed;
- standard-library-only dependency boundary passed;
- neutral HTTP/form acceptance passed;
- Linux ARM64 build passed;
- shell CSS + optional PWA JS measured **11,607 bytes** uncompressed against the 32 KiB ceiling;
- `BenchmarkRender` measured approximately **24,592 ns/op, 10,580 B/op and 138 allocs/op** on the GitHub AMD EPYC runner;
- PWA tests prove same-origin/static-only boundaries, deterministic cache identities, shell-owned cleanup, app `CacheVersion` invalidation and caller-cache isolation.

## Human acceptance still required

Human browser/visual acceptance remains open for responsive layout, light/dark visual quality, keyboard/zoom, CSS-off/no-JS, print, Android presentation and installed PWA behaviour.
