# Release checklist

## Contract
- [ ] No business/domain/data/auth/workflow logic in the shell.
- [ ] Framework-neutral, instance-based rendering remains intact.
- [ ] Core operation requires zero JavaScript and no browser-side Go/WASM runtime.
- [ ] Forms/lists/tables/status/action treatments remain presentation-only.
- [ ] Dark mode remains CSS-token based.
- [ ] PWA persistence remains static-only and caller cache isolation is preserved.
- [ ] `VERSION`, demo version and release tag agree.

## Engineering gate
- [ ] Go 1.27 CI passes format, vet and race tests.
- [ ] Native and Linux ARM64 demo builds pass.
- [ ] Standard-library-only dependency contract passes.
- [ ] CSS + optional JS remains <= 32 KiB uncompressed.
- [ ] `./scripts/acceptance.sh` passes.
- [ ] closed compact navigation is hidden from keyboard/accessibility reach until opened.
- [ ] detail-pane presentation can respond to narrow container width without viewport assumptions.
- [ ] safe-area inset and dynamic-viewport CSS contracts remain present.
- [ ] render benchmark and exact asset bytes are recorded.

## Human browser gate
- [ ] Chromium and Firefox reviewed; WebKit/Safari where practical.
- [ ] 320 / 390 / 768 / 820 / 1024 / 1366 / 1920 / 2560 widths reviewed.
- [ ] compact phone navigation opens as overlay drawer and scrim closes it.
- [ ] with compact drawer closed, Tab moves from the menu control directly into visible page content rather than hidden navigation links.
- [ ] with compact drawer open, background page scrolling is locked where `:has()` is supported.
- [ ] tablet 768/820 layout shows persistent narrow rail and useful two-pane list/detail.
- [ ] narrow tablet detail panes stack key-value content without crossing panel borders.
- [ ] desktop proportions remain calm and spacious.
- [ ] primary versus secondary action hierarchy is immediately clear.
- [ ] badges/notices have sufficient visual weight in light and dark modes.
- [ ] panel vertical rhythm is balanced.
- [ ] taps/clicks provide visible pressed feedback without feeling over-animated.
- [ ] native safe areas are respected on notched/home-indicator devices, especially landscape and installed PWA mode.
- [ ] mobile browser chrome expansion/collapse does not strand navigation below the usable viewport.
- [ ] cross-document transitions feel subtle where supported and ordinary navigation remains good where unsupported.
- [ ] reduced-motion preference removes meaningful animation.
- [ ] browser/status-bar chrome matches light/dark theme where supported.
- [ ] keyboard-only and 200% zoom reviewed.
- [ ] JavaScript-disabled core flow works.
- [ ] CSS-disabled page remains semantically usable.
- [ ] print output is sane and light.
- [ ] installed PWA/offline/update behaviour checked where supported.

## Evidence
- [ ] PRD acceptance criteria map to executable/human evidence.
- [ ] generated EDD pins the implementation snapshot.
- [ ] generated comparison contains no unexplained mismatch.
