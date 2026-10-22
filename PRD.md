# go-app-shell — Product Requirements Document

**Version:** 0.1

## Problem

Substantial Go web applications repeatedly need the same presentation foundation: a clear application frame, responsive navigation, usable phone portrait behaviour, a large desktop workspace, accessible forms and data presentation, and an optional path to installation as a Progressive Web App.

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

## Users

Go applications that need a professional browser interface without adopting a JavaScript/WASM application runtime, frontend build system or domain-specific admin platform.

The component is intended for unrelated corporate and operational applications and for anyone else who needs the same Go presentation primitive.

## In scope

- public importable Go package;
- framework-neutral rendering through standard Go interfaces;
- independent renderer instances with no mutable global application engine;
- deterministic server-rendered HTML using `html/template`;
- explicit application-header, navigation, page-context, action and workspace slots;
- single-pane and list/detail workspace modes;
- one responsive structure from 320 CSS pixels upward;
- phone portrait, tablet/resized-window and large desktop behaviour;
- persistent desktop navigation and compact navigation without required JavaScript;
- semantic forms, lists and tables styled through documented CSS conventions rather than Go-side business component builders;
- local overflow containment for intrinsically wide content;
- small namespaced CSS design-token set and caller theme overrides;
- semantic landmarks, skip link, visible focus and keyboard operation;
- practical 44 CSS pixel minimum sizing for normal shell controls;
- non-colour-only notice/status presentation;
- embedded local assets and system fonts only by default;
- optional PWA manifest generation;
- optional service-worker registration;
- optional explicit same-origin static precache plus offline fallback;
- ordinary web operation when PWA support is disabled or unavailable;
- meaningful semantic HTML if the stylesheet is unavailable;
- Go 1.27 baseline;
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
- custom component source formats or component code generation;
- React, Vue, Svelte, Angular, Bootstrap, Tailwind or another required UI framework;
- HTMX as a required dependency;
- generic widget, charting, modal or data-grid framework;
- Go APIs that turn ordinary forms/tables into data/business abstractions;
- application-specific icons, logos or branding;
- generic caching of authenticated pages or API responses;
- offline business-data stores, mutation queues, replay, synchronisation or conflict resolution;
- deciding which application data is safe to retain on a device;
- guaranteeing PWA installation where the browser/platform does not support it.

## Acceptance criteria

1. an unrelated Go module can import the public package directly;
2. a consumer builds with `go build` and no frontend toolchain or third-party runtime dependency;
3. integration requires no Gin/Echo/Fiber/Chi/Beego or other framework adapter;
4. independently configured renderer instances coexist without shared mutable presentation state;
5. equivalent inputs render deterministic static HTML;
6. the first useful page is server-rendered and requires no client runtime bootstrap;
7. the reference application is usable at 320, 390, 768, 1024, 1366, 1920 and 2560 CSS-pixel widths;
8. compact layouts show one useful list/detail pane rather than squeezed desktop columns;
9. desktop layouts can show persistent navigation and simultaneous list/detail panes;
10. continuous resizing requires no reload or device detection;
11. 200% browser zoom remains usable;
12. shell markup provides semantic header/nav/main structure, skip-to-content and current-navigation semantics;
13. keyboard focus is clearly visible and normal shell controls meet the practical 44px target;
14. presentation states are not conveyed by colour alone;
15. core navigation and ordinary form submission work with JavaScript disabled;
16. if shell CSS is unavailable, semantic content, navigation and forms remain readable/operable in browser defaults;
17. default shell assets require no CDN, external font or runtime network dependency;
18. documented CSS tokens can be overridden without editing shell source;
19. PWA mode is opt-in and can emit a manifest from caller metadata;
20. disabling/failing PWA support does not alter ordinary web behaviour;
21. service-worker persistence is limited to explicit same-origin static URLs and an optional offline fallback;
22. non-GET requests are never queued or replayed by the shell;
23. required core JavaScript is zero;
24. V0.1 ships no `.wasm`, virtual-DOM runtime, hook runtime, hydration runtime or client router;
25. first-party uncompressed shell CSS plus optional PWA JavaScript is <= 32 KiB, excluding caller content/icons;
26. default shell markup/theme targets WCAG 2.2 AA within the boundary it owns;
27. the neutral reference application contains no dependency on a specific business domain;
28. CI proves format, vet, race tests, native build, exact version, Linux ARM64 build, rendering/PWA boundaries and no-business-logic/no-client-runtime boundaries;
29. CI publishes a non-gating render benchmark (`ns/op`, `B/op`, `allocs/op`) and exact first-party asset bytes for regression visibility;
30. release acceptance includes Chromium and Firefox, plus WebKit/Safari where practical, across the responsive/keyboard/zoom matrix;
31. code-derived EDD evidence is pinned to the released implementation and the comparison contains no unexplained mismatch.

## Definition of done

V0.1 is done when an unrelated Go application can import the shell, render a useful responsive application from one deterministic server-side structure without a framework adapter or browser UI runtime, degrade cleanly through loss of PWA/JavaScript/CSS, pass automated and human acceptance, stay inside its resource budgets, and show through code-derived evidence that no business logic or recreated frontend framework entered the component.
