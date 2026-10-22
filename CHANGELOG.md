# Changelog

## 0.1.0-dev — release candidate

- Added deterministic shell-owned PWA cache identities derived from version, embedded assets and configured static cache inputs.
- Added activation cleanup limited to older `go-app-shell-` caches; explicit caller cache names remain isolated.
- Added first-class dark mode through existing CSS tokens and `prefers-color-scheme`, with print forced back to light.
- Added presentation-only form error states, error summaries, generic badges/status chips, key-value `<dl>` styling and compact/wide visibility helpers.
- Improved compact header behaviour and removed redundant desktop back-navigation in the demo.
- Reworked demo dynamic fragments through `html/template` rather than string concatenation.
- Kept official support at Go 1.27+ while lowering the module language/API floor to Go 1.22 so earlier compatible releases are not rejected unnecessarily; earlier releases remain best-effort and unsupported.
- Expanded acceptance to cover form-error rendering and PWA cache-lifecycle boundaries.

## 0.1.0-dev — initial release candidate

- Added framework-neutral server-rendered application shell for Go.
- Added responsive compact, medium and desktop/wide layouts from one semantic structure.
- Added zero-JavaScript navigation, single/list-detail workspaces and accessible presentation defaults.
- Added semantic form/list/table presentation conventions and CSS design tokens.
- Added optional manifest/service-worker helpers with static-only same-origin offline boundary.
- Added neutral reference application, ARM64/CI acceptance and explicit asset/render budgets.
