# go-app-shell — Code-Derived Engineering Design

**Implementation evidence:** `ba797b0a384168750f92044ec793ff268a3783a8`  
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

## Recovered layout model

- compact: 320–599 CSS px;
- medium: 600–1023 CSS px;
- desktop/wide: 1024 CSS px and above;
- same semantic HTML structure at all sizes;
- compact list/detail shows one useful pane;
- desktop list/detail exposes both panes;
- compact/wide visibility helpers are CSS only;
- compact navigation remains zero-JavaScript.

## Recovered presentation vocabulary

The CSS remains deliberately small and generic:

- panels/stacks/clusters/empty states;
- generic notices and badges;
- semantic form layout, help, field-error and error-summary presentation;
- list and table wrappers;
- key-value `<dl>` styling;
- compact/wide visibility helpers.

No Go-side data-aware form/table builder was introduced.

## Recovered light/dark model

Light tokens are defined at `:root`. Dark tokens override the same variables under `@media (prefers-color-scheme: dark)`. No JavaScript or theme registry is present. Print restores a light palette.

## Recovered PWA boundary

The generated service worker:

- accepts only same-origin rooted precache/offline inputs;
- intercepts GET only;
- precaches only configured static URLs and optional offline fallback;
- does not use IndexedDB, queues, background sync or mutation replay;
- derives the default cache identity from `shell.Version`, embedded `shell.css`/`pwa.js` bytes and normalised cache inputs;
- deletes only older caches with the shell-owned `go-app-shell-` prefix;
- disables prefix cleanup when a caller provides an explicit `CacheName`.

This corrects the initial release-candidate stale-static-cache risk without broadening offline responsibility.

## Recovered compatibility/support model

`go.mod` declares `go 1.22`, matching the lowest language/API level used by the current source/demo. GitHub CI remains configured for Go 1.27.x, which is the documented supported/tested baseline. Earlier compatible releases are not part of the support matrix.

## Recovered demo safety pattern

Dynamic record/form fragments are rendered with `html/template` before being passed into the explicitly trusted `template.HTML` shell slots. The shell trust boundary is unchanged; the demo no longer teaches dynamic string concatenation.

## Automated evidence

At implementation commit `ba797b0a384168750f92044ec793ff268a3783a8`:

- local acceptance passed on an available Go 1.23 runner;
- GitHub Go 1.27 acceptance passed;
- Linux ARM64 build passed;
- race tests passed;
- standard-library-only dependency boundary passed;
- shell CSS + optional PWA JS measured 11,607 bytes uncompressed, below the 32 KiB ceiling;
- HTTP acceptance proved list/detail, manifest, GET-only service-worker boundary and shell-rendered form validation error state;
- cache tests prove deterministic default identity, stale shell-cache cleanup boundary and caller-cache isolation.

## Human acceptance still required

Human browser/visual acceptance remains open for responsive layout, light/dark visual quality, keyboard/zoom, CSS-off/no-JS, print and installed PWA upgrade behaviour.
