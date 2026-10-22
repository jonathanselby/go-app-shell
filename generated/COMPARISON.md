# go-app-shell — Design / Implementation Comparison

**Implementation evidence:** `137b3d6b73be56e99a8819498094998d04f4ae5b`  
**Generated:** 2026-10-22

| Design area | Intended contract | Recovered implementation | Assessment |
|---|---|---|---|
| one-job boundary | presentation shell only | one public shell package; neutral demo outside reusable API | MATCH |
| business boundary | no business/domain logic | no DB/auth/workflow/domain subsystem; core vocabulary generic | MATCH |
| framework breadth | standard Go boundary, no adapters | `Render(io.Writer, Page)`; demo uses plain `net/http` | MATCH |
| state | caller/application owns business state | renderer instance config only; selection remains request/URL state | MATCH |
| client runtime | zero required JS; no WASM/VDOM/hooks | server HTML + CSS; optional tiny PWA registration only | MATCH |
| responsive model | one structure, available-width adaptation | CSS media layout; compact and desktop from same template | MATCH |
| compact navigation | usable without JavaScript | native checkbox/label + CSS navigation state | MATCH |
| list/detail | one pane compact, both wide | explicit `CompactPane`; desktop CSS exposes both panes | MATCH |
| forms/lists/tables | presentation conventions only | semantic HTML styled by small `gas-*` class vocabulary | MATCH |
| theming | small CSS token set | namespaced `--gas-*` custom properties; no theme runtime | MATCH |
| accessibility | semantic shell, focus, 44px practical controls | landmarks/skip link/aria-current/focus/2.75rem controls | MATCH; human review pending |
| degraded operation | PWA -> web -> no JS -> no CSS | optional PWA; zero-JS navigation/forms; semantic source remains | MATCH; CSS-off human review pending |
| PWA | manifest + static allowlist/fallback only | same-origin validation; GET-only exact precache + navigation fallback | MATCH |
| offline business data | none | no IndexedDB/sync/mutation queue/replay | MATCH |
| dependencies | stdlib only, no frontend toolchain | `go.mod` has no required modules; embedded assets | MATCH |
| asset budget | <=32 KiB CSS + optional JS | 9,532 bytes at evidence snapshot | MATCH |
| ARM64 | compile gate | GitHub CI Linux ARM64 build | MATCH |
| real browser acceptance | multi-browser responsive/accessibility matrix | documented in `RELEASE.md`; not automated | PENDING HUMAN ACCEPTANCE |

## Corrections made during build

The first Git publication tree accidentally reformatted four Go source files while transferring the already-tested local source into Git objects. The format gate rejected that tree immediately. Commit `05d1cf306986294fc75f30ee518dfb318148c25b` replaced those blobs with the actual `gofmt` output; the complete Go 1.27 acceptance and ARM64 gate then passed.

Before promotion, release state was also corrected from an implied stable `0.1.0` to `0.1.0-dev`. Stable tagging is manual and refuses prerelease versions, so automated code acceptance cannot accidentally bypass the required human browser review.

The renderer was hardened to execute templates into a buffer before writing to the caller, avoiding intentional partial-document output if template execution fails.

## Result

**MATCH for automated/code-verifiable architecture.** No unexplained architectural mismatch was found. The only open acceptance is the deliberately human browser/visual/accessibility matrix required before changing `VERSION` from `0.1.0-dev` to stable `0.1.0` and creating the release tag.
