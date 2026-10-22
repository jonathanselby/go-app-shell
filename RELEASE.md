# Release checklist

## Contract
- [ ] The shell contains no business logic or domain semantics.
- [ ] Rendering remains framework-neutral and instance-based.
- [ ] Core operation requires zero JavaScript and no browser-side Go/WASM runtime.
- [ ] No data/auth/routing/workflow responsibility has leaked into the package.
- [ ] Forms/lists/tables remain presentation conventions, not data-aware Go builders.
- [ ] PWA persistence remains explicit static same-origin URLs + offline fallback only.
- [ ] `VERSION`, demo version and release tag agree.

## Engineering gate
- [ ] `gofmt` clean.
- [ ] `go vet ./...` passes.
- [ ] `go test -race -count=1 ./...` passes.
- [ ] native and Linux ARM64 demo builds pass.
- [ ] standard-library-only module dependency contract passes.
- [ ] first-party CSS + optional JS remains <= 32 KiB uncompressed.
- [ ] `./scripts/acceptance.sh` passes.
- [ ] render benchmark and exact asset bytes are recorded.

## Human browser gate
- [ ] Chromium and Firefox reviewed; WebKit/Safari where practical.
- [ ] 320 / 390 / 768 / 1024 / 1366 / 1920 / 2560 CSS-pixel widths reviewed.
- [ ] continuous resize and phone portrait reviewed.
- [ ] keyboard-only and 200% zoom reviewed.
- [ ] JavaScript-disabled core flow works.
- [ ] CSS-disabled page remains semantically usable.
- [ ] print output is sane.
- [ ] PWA standalone/offline-static fallback checked where supported.

## Evidence
- [ ] PRD acceptance criteria map to executable/human evidence.
- [ ] generated EDD pins the implementation snapshot.
- [ ] generated comparison contains no unexplained mismatch.
