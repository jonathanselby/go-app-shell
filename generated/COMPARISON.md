# go-app-shell — Design / Implementation Comparison

**Implementation evidence:** `883028ebfa4a69139ad1261ea664699df4d01477`  
**Generated:** 2026-10-22

| Design area | Intended contract | Recovered implementation | Assessment |
|---|---|---|---|
| one-job boundary | presentation shell only | one public shell package + neutral demo | MATCH |
| business boundary | no business/domain logic | no DB/auth/workflow/domain subsystem | MATCH |
| framework breadth | standard Go boundary, no adapters | `Render(io.Writer, Page)` | MATCH |
| client runtime | zero required JS; no WASM/VDOM/hooks | server HTML + CSS; optional PWA registration only | MATCH |
| compact navigation | app-like without JS | fixed drawer + scrim using native checkbox/label | MATCH |
| tablet layout | persistent rail + two panes from 768px | 11.5rem rail + tablet list/detail CSS | MATCH |
| desktop layout | larger rail/workspace split | 15rem rail + wider list/detail | MATCH |
| action hierarchy | generic primary presentation only | `gas-button--primary`; app chooses meaning | MATCH |
| status presentation | visually legible, non-authoritative | tinted notices/badges with text labels | MATCH |
| panel rhythm | predictable content spacing | first/last-child margin normalisation | MATCH |
| interaction feedback | short CSS feedback, zero runtime | hover/active transitions; native touch feedback retained | MATCH |
| document transitions | progressive enhancement only | CSS `@view-transition`; normal fallback | MATCH |
| prefetch policy | application-owned | no shell speculation rules/prefetch | MATCH |
| mobile nav semantics | no app-priority assumptions | overlay drawer; no bottom-tab/icon API | MATCH |
| dark/browser chrome | CSS dark mode + matching chrome | `prefers-color-scheme` + configurable theme-color meta | MATCH |
| PWA cache version | app build version, not shell version | demo `demoBuildVersion` -> `CacheVersion` | MATCH |
| asset budget | <=32 KiB | 14,503 bytes | MATCH |
| Go/ARM64 gates | supported CI + cross-build | Go 1.27 and Linux ARM64 pass | MATCH |
| real browser visual acceptance | human multi-device review | documented, not automated | PENDING HUMAN ACCEPTANCE |

## Result

**MATCH for automated/code-verifiable architecture.** Visual hardening improved phone, tablet and desktop presentation without adding business logic, a frontend runtime, speculative network behaviour or application navigation semantics. Human visual acceptance is the remaining release-candidate gate.
