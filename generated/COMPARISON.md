# go-app-shell — Design / Implementation Comparison

**Implementation evidence:** `ba797b0a384168750f92044ec793ff268a3783a8`  
**Generated:** 2026-10-22

| Design area | Intended contract | Recovered implementation | Assessment |
|---|---|---|---|
| one-job boundary | presentation shell only | one public shell package; neutral demo outside reusable API | MATCH |
| business boundary | no business/domain logic | no DB/auth/workflow/domain subsystem; core vocabulary remains generic | MATCH |
| framework breadth | standard Go boundary, no adapters | `Render(io.Writer, Page)`; demo uses plain `net/http` | MATCH |
| state | caller/application owns business state | renderer config only; list/detail selection remains request/URL state | MATCH |
| client runtime | zero required JS; no WASM/VDOM/hooks | server HTML + CSS; optional tiny PWA registration only | MATCH |
| responsive model | one structure, available-width adaptation | CSS media layout; compact/medium/desktop from same template | MATCH |
| compact navigation | usable without JavaScript | native checkbox/label + CSS navigation state | MATCH |
| list/detail | one pane compact, both wide | explicit compact pane; desktop CSS exposes both panes | MATCH |
| forms/lists/tables | presentation conventions only | semantic HTML styled by small `gas-*` vocabulary | MATCH |
| validation presentation | show errors, do not validate | `aria-invalid`, field error and summary styling; demo validation stays in app | MATCH |
| badges/detail | generic presentation only | generic badges and `<dl>` key-value styling; no status semantics | MATCH |
| theming | CSS tokens, no theme runtime | namespaced `--gas-*` tokens | MATCH |
| dark mode | first-class, no JS runtime | `prefers-color-scheme: dark` token override; print light | MATCH |
| accessibility | semantic shell, focus, 44px controls | landmarks/skip link/aria-current/focus/2.75rem controls/error hooks | MATCH; human review pending |
| degraded operation | PWA -> web -> no JS -> no CSS | optional PWA; zero-JS nav/forms; semantic source remains | MATCH; human review pending |
| PWA persistence | static allowlist/fallback only | same-origin validation; GET-only exact precache + navigation fallback | MATCH |
| PWA upgrade | shell asset changes must not stay stale | cache identity hashes version/assets/cache inputs; old shell caches cleaned | MATCH |
| caller cache isolation | do not delete app-owned caches | explicit `CacheName` disables shell prefix cleanup | MATCH |
| offline business data | none | no IndexedDB/sync/mutation queue/replay | MATCH |
| trusted dynamic HTML | explicit caller trust, safe demo pattern | shell retains `template.HTML`; demo dynamic fragments use `html/template` | MATCH |
| dependencies | stdlib only, no frontend toolchain | `go.mod` has no required modules; embedded assets | MATCH |
| Go support | official support 1.27+; avoid artificial minimum | CI Go 1.27.x; module language/API floor 1.22 | MATCH |
| asset budget | <=32 KiB CSS + optional JS | 11,607 bytes at evidence snapshot | MATCH |
| ARM64 | compile gate | GitHub CI Linux ARM64 build | MATCH |
| real browser acceptance | multi-browser responsive/light-dark/accessibility/PWA-upgrade matrix | documented in `RELEASE.md`; not automated | PENDING HUMAN ACCEPTANCE |

## Red-team changes accepted

The hardening pass accepted only findings that improved correctness or broadly reusable presentation without acquiring application logic:

- fixed stale PWA static-cache lifecycle;
- separated documented Go support from best-effort compatibility;
- added CSS-token dark mode;
- added form-error, generic badge, key-value and responsive visibility conventions;
- improved compact header behaviour;
- changed demo dynamic fragments from string concatenation to `html/template`.

## Red-team changes deliberately rejected/deferred

No partial-render runtime, HTMX dependency, timeline framework, pagination/filter framework, data-aware form/table builders, theme engine, package rename or additional client state machinery was added.

## Result

**MATCH for automated/code-verifiable architecture.** The hardening pass corrected genuine defects/gaps without changing the core architecture or introducing business logic/framework bloat. Human browser/visual/PWA-upgrade acceptance remains the final release-candidate gate.
