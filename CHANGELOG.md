# Changelog

## 0.1.0-dev — release candidate

### Visual hardening
- Added tablet app layout from 768px with persistent narrow navigation rail and two-pane list/detail.
- Changed compact navigation from document-flow expansion to a zero-JS overlay drawer with scrim.
- Added generic primary-action styling and demo action hierarchy.
- Added tinted notices/badges and balanced panel first/last-child spacing.
- Added short hover/pressed feedback while retaining native browser/touch feedback.
- Added progressive CSS cross-document View Transitions with reduced-motion handling.
- Added caller-configurable light/dark browser `theme-color` metadata and `viewport-fit=cover`.
- Corrected the demo to use its own application build version for `PWAConfig.CacheVersion`.
- Deliberately rejected shell-owned speculative prefetching, bottom-tab navigation and icon-priority APIs.

### Earlier hardening
- Added deterministic PWA cache identities and optional application `CacheVersion`.
- Added dark mode through CSS tokens.
- Added form-error, error-summary, generic badge, key-value and responsive visibility conventions.
- Scoped invalid-control styling beneath `.gas-field`.
- Reworked demo dynamic fragments through `html/template`.
- Kept official Go support at 1.27+ while retaining the lower source/API floor for best-effort compatibility.

### Initial release candidate
- Added framework-neutral server-rendered application shell for Go.
- Added responsive navigation, single/list-detail workspaces and accessible defaults.
- Added optional static-only PWA plumbing, ARM64/CI acceptance and explicit resource budgets.
