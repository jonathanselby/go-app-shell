# Release checklist

## Contract
- [x] No business/domain/data/auth/workflow logic in the shell.
- [x] Framework-neutral, instance-based rendering remains intact.
- [x] Core operation requires zero JavaScript and no browser-side Go/WASM runtime.
- [x] Forms/lists/tables/status/action treatments remain presentation-only.
- [x] Dark mode remains CSS-token based.
- [x] PWA persistence remains static-only and caller cache isolation is preserved.
- [x] `VERSION`, demo version and release tag target agree at `0.1.0` / `v0.1.0`.

## Engineering gate
- [x] Go 1.27 CI passes format, vet and race tests.
- [x] Native and Linux ARM64 demo builds pass.
- [x] Standard-library-only dependency contract passes.
- [x] CSS + optional JS remains <= 32 KiB uncompressed.
- [x] `./scripts/acceptance.sh` passes.
- [x] Closed compact navigation is hidden from keyboard/accessibility reach until opened.
- [x] Detail-pane presentation responds to narrow container width without viewport assumptions.
- [x] Safe-area inset and dynamic-viewport CSS contracts remain present.
- [x] Render benchmark and exact asset bytes are recorded in generated evidence.

## Human browser gate
- [x] Linux Chrome reviewed and passed.
- [x] Linux Firefox reviewed and passed.
- [x] Android Chrome reviewed and passed.
- [x] Compact phone navigation opens as overlay drawer and scrim closes it.
- [x] With compact drawer closed, keyboard focus does not enter hidden navigation links.
- [x] Tablet-width list/detail layout remains contained and readable.
- [x] Desktop proportions remain calm and spacious.
- [x] Primary versus secondary action hierarchy is clear.
- [x] Badges/notices have sufficient visual weight in light and dark modes.
- [x] Panel and stack vertical rhythm is consistent.
- [x] Taps/clicks provide visible pressed feedback without feeling over-animated.
- [x] Cross-document transitions remain progressive enhancement only.
- [x] Browser/status-bar chrome follows light/dark theme where supported.
- [x] Keyboard and responsive behaviour reviewed during UAT.
- [x] Installed/mobile behaviour reviewed on the available UAT devices.

WebKit/Safari and iOS-specific notch/home-indicator behaviour were not available in this UAT environment. The shell retains standards-based safe-area handling for those platforms, but V0.1 does not claim device-specific Safari certification.

## Evidence
- [x] PRD acceptance criteria map to executable/human evidence.
- [x] Generated EDD pins the implementation snapshot.
- [x] Generated comparison contains no unexplained mismatch.
- [x] Human UAT outcome recorded for V0.1.
