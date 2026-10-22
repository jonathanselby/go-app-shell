# go-app-shell — Design / Implementation Comparison

**Implementation evidence:** `dcd82bd1d343f8785956f592bf1c189218bd2dde`  
**Generated:** 2026-10-22

| Design area | Intended contract | Recovered implementation | Assessment |
|---|---|---|---|
| one-job boundary | presentation shell only | one public shell package; neutral demo outside reusable API | MATCH |
| business boundary | no business/domain logic | no DB/auth/workflow/domain subsystem | MATCH |
| framework breadth | standard Go boundary, no adapters | `Render(io.Writer, Page)`; demo uses plain `net/http` | MATCH |
| client runtime | zero required JS; no WASM/VDOM/hooks | server HTML + CSS; optional PWA registration only | MATCH |
| responsive model | one structure, available-width adaptation | compact/medium/desktop CSS from same template | MATCH |
| list/detail | one pane compact, both wide | explicit compact pane; both panes at 1024px+ | MATCH |
| presentation vocabulary | small and generic | forms/lists/tables/notices/badges/kv/visibility helpers only | MATCH |
| validation boundary | presentation only | application decides validation; shell styles errors | MATCH |
| error focus | application-controlled, zero-JS possible | demo uses native `tabindex=-1` + `autofocus`; shell owns no focus state | MATCH |
| selector boundary | do not restyle unrelated controls | invalid styling scoped beneath `.gas-field` | MATCH |
| dark mode | first-class, no runtime | CSS token override via `prefers-color-scheme` | MATCH |
| PWA persistence | static allowlist/fallback only | same-origin validation; GET-only exact precache + navigation fallback | MATCH |
| shell PWA upgrade | shell changes invalidate shell cache | version + embedded asset hash + cache inputs | MATCH |
| app PWA upgrade | stable app URLs can be invalidated explicitly | optional `CacheVersion` participates in cache identity | MATCH |
| caller cache isolation | never delete app-owned caches | explicit `CacheName` disables shell cleanup | MATCH |
| app asset ownership | shell must not become build system | no app-resource fetching/content hashing | MATCH |
| offline business data | none | no IndexedDB/sync/mutation queue/replay | MATCH |
| trusted dynamic HTML | explicit caller trust | demo dynamic fragments use `html/template` | MATCH |
| dependencies | stdlib only | no required modules; embedded assets | MATCH |
| Go support | support 1.27+ without artificial floor | CI Go 1.27.x; module language/API floor 1.22 | MATCH |
| asset budget | <=32 KiB CSS + optional JS | 11,607 bytes | MATCH |
| ARM64 | compile gate | Linux ARM64 build passes | MATCH |
| real browser acceptance | responsive/light-dark/accessibility/PWA matrix | documented; human acceptance pending | PENDING HUMAN ACCEPTANCE |

## Result

**MATCH for automated/code-verifiable architecture.** The second red-team round found one remaining PWA application-asset invalidation gap and two small presentation-boundary defects. They were fixed with one additive generic field (`CacheVersion`), scoped CSS and demo markup/accessibility changes. No new runtime, framework, state system, business logic or build-system ownership was introduced.
