# go-app-shell — Design / Implementation Comparison

**Implementation evidence:** `762bc87d4643fc33e16a5c82f199f59583595f71`  
**Generated:** 2026-10-22

| Design area | Intended contract | Recovered implementation | Assessment |
|---|---|---|---|
| one-job boundary | presentation shell only | one public shell package + neutral demo | MATCH |
| business boundary | no business/domain logic | no DB/auth/workflow/domain subsystem | MATCH |
| framework breadth | standard Go boundary, no adapters | `Render(io.Writer, Page)` | MATCH |
| client runtime | zero required JS; no WASM/VDOM/hooks | server HTML + CSS; optional PWA registration only | MATCH |
| compact navigation | app-like without JS | fixed drawer + scrim using native checkbox/label | MATCH |
| closed drawer accessibility | hidden nav must not receive focus | `visibility:hidden` + `pointer-events:none`, restored when opened | MATCH |
| compact background scroll | drawer should behave modally | `body:has(.gas-nav-toggle:checked)` scroll lock where supported | MATCH |
| tablet layout | persistent rail + two panes from 768px | 11.5rem rail + tablet list/detail CSS | MATCH |
| narrow pane layout | local pane width, not viewport, controls detail layout | detail pane container + key-value container query | MATCH |
| desktop layout | larger rail/workspace split | 15rem rail + wider list/detail | MATCH |
| safe-area handling | viewport-fit must not hide content under device chrome | header/main/nav/skip-link use safe-area insets | MATCH |
| mobile viewport height | collapsing browser chrome must not strand nav | `100dvh` with `100vh` fallback | MATCH |
| tap feedback | explicit pressed feedback without double flash | bounded WebKit tap-highlight suppression on shell controls | MATCH |
| action hierarchy | generic primary presentation only | `gas-button--primary`; app chooses meaning | MATCH |
| status presentation | visually legible, non-authoritative | tinted notices/badges with text labels | MATCH |
| panel rhythm | predictable content spacing | first/last-child margin normalisation | MATCH |
| document transitions | progressive enhancement only | CSS `@view-transition`; normal fallback | MATCH |
| prefetch policy | application-owned | no shell speculation rules/prefetch | MATCH |
| mobile nav semantics | no app-priority assumptions | overlay drawer; no bottom-tab/icon API | MATCH |
| dark/browser chrome | CSS dark mode + matching chrome | `prefers-color-scheme` + configurable theme-color meta | MATCH |
| PWA cache version | app build version, not shell version | application build version -> `CacheVersion` | MATCH |
| asset budget | <=32 KiB | PASS | MATCH |
| Go/ARM64 gates | supported CI + cross-build | Go 1.27 and Linux ARM64 pass | MATCH |
| real browser visual acceptance | human multi-device review | documented, not automated | PENDING HUMAN ACCEPTANCE |

## Result

**MATCH for automated/code-verifiable architecture.** The accessibility/layout hardening fixes concrete mobile and tablet defects using browser-native CSS only: visibility, container queries, safe-area environment variables and dynamic viewport units. No business logic, JavaScript runtime, framework adapter, app-state system or navigation semantics were added. Human visual acceptance remains the release-candidate gate.
