# go-app-shell — Product Requirements Document

**Version:** 0.1

## Problem

Substantial Go web applications repeatedly need the same presentation foundation: a clear application frame, responsive navigation, credible phone and tablet behaviour, a large desktop workspace, accessible forms and data presentation, light/dark presentation, and an optional path to installation as a Progressive Web App.

Rebuilding that foundation for every application wastes effort. Importing a large frontend or admin framework solves much more than the problem and introduces avoidable runtime, tooling and architectural coupling.

The missing capability is presentation infrastructure, not application behaviour.

## Product objective

Provide a small reusable Go component that gives server-rendered applications a polished, responsive and accessible application shell from 320 CSS pixels through large desktop displays, while remaining useful as an ordinary web application and as an optional Progressive Web App.

> **The shell owns presentation. The application owns meaning and behaviour.**

## Product principles

1. Use Go and the browser directly rather than recreating a JavaScript frontend framework in Go.
2. Be broadly reusable by doing less: no domain, data, auth, workflow or routing ownership.
3. Keep the execution path learnable: application data -> page model -> renderer -> semantic HTML/CSS.
4. Make budgets explicit: standard-library runtime, zero required JavaScript and a small asset ceiling.
5. Degrade downward into simpler web behaviour rather than fail when optional capability disappears.
6. Treat visual credibility as part of the engineering contract, not post-build decoration.
7. Add presentation primitives only when they solve repeated application needs without acquiring business meaning.

## Users

Go applications that need a professional browser interface without adopting a JavaScript/WASM application runtime, frontend build system or domain-specific admin platform.

## Go support and compatibility

Go **1.27 and later** is the supported and tested baseline for V0.1.

The module deliberately avoids Go 1.27-specific language/runtime features and declares the lowest language/API level currently required by the implementation. Earlier compatible Go 1.x releases may work, but are outside the supported test matrix and are used on a best-effort basis at the adopter's own risk.

## In scope

- public importable Go package;
- framework-neutral rendering through standard Go interfaces;
- independent renderer instances with no mutable global engine;
- deterministic server-rendered HTML using `html/template`;
- explicit application-header, navigation, page-context, action and workspace slots;
- single-pane and list/detail workspace modes;
- one responsive structure from 320 CSS pixels upward;
- compact phone layout with zero-JS overlay navigation drawer and scrim;
- tablet app layout from 768 CSS pixels with persistent narrow navigation rail;
- tablet list/detail layout showing both panes where space supports it;
- wider desktop layout with full navigation rail and larger list/detail split;
- clear primary-action styling without adding workflow semantics;
- short hover/pressed presentation feedback while retaining native browser/touch behaviour;
- cross-document view transitions as progressive CSS enhancement where supported;
- semantic forms, lists and tables styled through CSS conventions rather than Go-side business builders;
- generic field-error, error-summary, badge/status, key-value detail and compact/wide visibility presentation conventions;
- tinted notices/badges and predictable panel content rhythm in light and dark modes;
- local overflow containment for intrinsically wide content;
- small namespaced CSS design-token set and caller theme overrides;
- automatic dark mode from browser/OS preference with no JavaScript theme runtime;
- browser chrome theme colours for light and dark presentation;
- semantic landmarks, skip link, visible focus and keyboard operation;
- practical 44 CSS pixel minimum sizing for normal shell controls;
- embedded local assets and system fonts by default;
- optional PWA manifest/service worker with conservative static-only cache boundary;
- deterministic default PWA cache identity and optional caller `CacheVersion`;
- ordinary web operation when PWA support is disabled or unavailable;
- meaningful semantic HTML if CSS is unavailable;
- standard-library-only runtime dependencies;
- no Node, npm, code generation or frontend build step required to consume V0.1.

## Out of scope

- business logic or domain vocabulary;
- authentication, authorisation, roles or permissions;
- application routing decisions or framework adapters;
- database access, persistence or business validation;
- application/client state management;
- workflow engines, queues or background jobs;
- a global service container or plugin runtime;
- browser-side Go/WASM UI runtime;
- virtual DOM, reconciler, hooks or hydration;
- client-side router or SPA runtime;
- shell-owned speculative navigation/prefetch policy;
- custom component source formats or code generation;
- React, Vue, Svelte, Angular, Bootstrap, Tailwind or another required UI framework;
- HTMX as a required dependency;
- partial-page rendering as a shell-owned runtime;
- generic widget, charting, timeline, pagination, filter-toolbar, modal or data-grid framework;
- application-priority navigation such as bottom tabs;
- application-specific icons, logos or branding;
- Go APIs that turn ordinary forms/tables into data/business abstractions;
- generic caching of authenticated pages or API responses;
- offline business-data stores, mutation queues, replay, synchronisation or conflict resolution;
- support guarantees for Go versions outside the documented support matrix.

## Acceptance criteria

1. an unrelated Go module can import the public package directly;
2. a consumer builds with `go build` and no frontend toolchain or third-party runtime dependency;
3. supported CI uses Go 1.27+ while the module does not artificially require 1.27-specific syntax/features;
4. integration requires no web-framework adapter;
5. independently configured renderer instances coexist without shared mutable presentation state;
6. equivalent inputs render deterministic static HTML;
7. the first useful page is server-rendered and requires no client runtime bootstrap;
8. the reference application is usable at 320, 390, 768, 820, 1024, 1366, 1920 and 2560 CSS-pixel widths;
9. below 768px navigation is an overlay drawer rather than pushing the page down;
10. at 768px+ a persistent navigation rail is available and list/detail can show both panes;
11. at 1024px+ navigation/list/detail expand to the wider desktop proportions;
12. continuous resizing requires no reload or device detection;
13. 200% browser zoom remains usable;
14. action hierarchy includes an explicit generic primary action style;
15. generic badges/notices have readable visual weight in light and dark modes without relying on colour alone;
16. panel first/last-child spacing is visually balanced;
17. taps/clicks receive visible pressed feedback without requiring JavaScript;
18. same-origin document navigation may use CSS cross-document view transitions where supported and degrades to ordinary navigation elsewhere;
19. the shell does not prefetch application pages on its own;
20. shell markup provides semantic header/nav/main structure, skip-to-content and current-navigation semantics;
21. keyboard focus is clearly visible and normal shell controls meet the practical 44px target;
22. core navigation and ordinary form submission work with JavaScript disabled;
23. if CSS is unavailable, semantic content/navigation/forms remain readable and operable;
24. default assets require no CDN, external font or runtime network dependency;
25. documented CSS tokens can be overridden without editing shell source;
26. dark mode follows `prefers-color-scheme: dark` and requires no JavaScript;
27. light/dark `theme-color` metadata is emitted and caller-configurable;
28. generic form-error styling supports `aria-invalid`, field-error text and an application-supplied error summary without performing validation itself;
29. PWA mode is opt-in; disabling/failing it does not alter ordinary web behaviour;
30. service-worker persistence remains limited to explicit same-origin static URLs plus optional offline fallback;
31. default cache identity changes when shell assets/version/cache inputs or caller `CacheVersion` change;
32. non-GET requests are never queued or replayed by the shell;
33. required core JavaScript is zero and V0.1 ships no WASM/VDOM/hooks/hydration/client router;
34. first-party uncompressed shell CSS plus optional PWA JavaScript remains <= 32 KiB;
35. default shell markup/light/dark themes target WCAG 2.2 AA within the shell boundary;
36. the neutral reference app demonstrates safe templating and uses an application-owned build version for `CacheVersion`;
37. CI proves format, vet, race tests, native build, exact version, Linux ARM64, rendering/PWA boundaries and architecture boundaries;
38. CI publishes a non-gating render benchmark and exact first-party asset bytes;
39. release acceptance includes Chromium/Firefox and WebKit/Safari where practical across responsive, keyboard, zoom, light/dark, no-JS, no-CSS, transitions, PWA/offline and PWA-upgrade checks;
40. code-derived EDD evidence is pinned to the released implementation with no unexplained mismatch.

## Definition of done

V0.1 is done when an unrelated Go application can reuse the shell without reworking its visual foundation; phone, tablet, desktop and installed-PWA presentation all feel deliberately designed rather than merely functional; the shell remains zero-business-logic and zero-required-JavaScript; automated and human acceptance pass; and code-derived evidence matches the intended design.
