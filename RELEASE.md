# Release checklist

## Contract
- [ ] The shell contains no business logic or domain semantics.
- [ ] Rendering remains framework-neutral and instance-based.
- [ ] Core operation requires zero JavaScript and no browser-side Go/WASM runtime.
- [ ] No data/auth/routing/workflow responsibility has leaked into the package.
- [ ] Forms/lists/tables remain presentation conventions, not data-aware Go builders.
- [ ] Form-error, badge, key-value and responsive visibility helpers remain presentation-only.
- [ ] Dark mode remains CSS-token based with no theme runtime.
- [ ] PWA persistence remains explicit static same-origin URLs + offline fallback only.
- [ ] Default shell-owned caches upgrade deterministically and cleanup is limited to older shell-owned caches.
- [ ] Explicit caller cache names do not enable shell prefix cleanup.
- [ ] `VERSION`, demo version and release tag agree.

## Go support
- [ ] GitHub CI passes on the documented supported baseline, Go 1.27+.
- [ ] `go.mod` does not artificially require 1.27-specific features when a lower language/API floor is sufficient.
- [ ] Earlier compatible Go releases are described as best-effort only, not supported/tested.

## Engineering gate
- [ ] `gofmt` clean.
- [ ] `go vet ./...` passes.
- [ ] `go test -race -count=1 ./...` passes.
- [ ] native and Linux ARM64 demo builds pass.
- [ ] standard-library-only module dependency contract passes.
- [ ] first-party CSS + optional JS remains <= 32 KiB uncompressed.
- [ ] `./scripts/acceptance.sh` passes.
- [ ] form-error HTTP acceptance returns shell-rendered 422 output.
- [ ] render benchmark and exact asset bytes are recorded.

## Human browser gate
- [ ] Chromium and Firefox reviewed; WebKit/Safari where practical.
- [ ] 320 / 390 / 768 / 1024 / 1366 / 1920 / 2560 CSS-pixel widths reviewed.
- [ ] continuous resize and phone portrait reviewed.
- [ ] compact header does not show unusable truncation.
- [ ] list/detail compact-only/wide behaviour reviewed.
- [ ] keyboard-only and 200% zoom reviewed.
- [ ] light mode reviewed.
- [ ] OS/browser dark mode reviewed for contrast, hierarchy and control legibility.
- [ ] JavaScript-disabled core flow works.
- [ ] CSS-disabled page remains semantically usable.
- [ ] print output is sane and prints light.
- [ ] PWA standalone/offline-static fallback checked where supported.
- [ ] PWA upgrade tested by changing shell CSS/assets and confirming an installed client receives the new cache.

## Evidence
- [ ] PRD acceptance criteria map to executable/human evidence.
- [ ] generated EDD pins the implementation snapshot.
- [ ] generated comparison contains no unexplained mismatch.
