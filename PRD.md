# go-app-shell — Product Requirements Document

**Version:** 0.1

## Problem

Substantial Go web applications repeatedly need the same presentation foundation: a clear application frame, persistent navigation on large screens, usable mobile portrait behaviour, list/detail workspaces, accessible forms and status presentation, and an optional path to installation as a Progressive Web App.

Without a small reusable shell, each application either rebuilds these mechanics independently or imports a large frontend framework and its toolchain. Both choices create avoidable duplication, inconsistency and complexity.

The missing capability is presentation infrastructure, not application behaviour.

## Product objective

Provide a small reusable Go component that gives server-rendered business applications a polished, responsive and accessible application shell from phone portrait to large desktop, while remaining useful as an ordinary web application and as an optional Progressive Web App.

> **The shell owns presentation. The application owns meaning and behaviour.**

## Product principle

Use Go's strengths directly rather than recreating a JavaScript frontend framework in Go.

The component is a library called by the consuming application. It does not introduce a browser-side Go runtime, virtual DOM, hook system, hydration layer, client router or framework lifecycle simply to render an application shell.

## Users

Go applications that need a professional browser interface without adopting a JavaScript application framework, WebAssembly UI runtime, frontend build system or domain-specific UI platform.

The component is intended to be reusable across unrelated business applications such as operational tools, CRM, ERP and other substantial Go systems without knowing anything about those domains.

## In scope

- a public importable Go package;
- framework-neutral rendering through standard Go interfaces rather than framework-specific adapters;
- instance-based renderer/configuration with no required global engine, service container or mutable plugin registry;
- server-rendered HTML using Go's standard web/template stack;
- deterministic static HTML output suitable for direct tests and snapshots;
- explicit presentation slots for application header, navigation, page context, actions and workspace content;
- one responsive application shell from 320 CSS pixels upward;
- continuous reflow based on available space rather than separate mobile and desktop applications;
- compact phone portrait presentation;
- tablet and resized-window presentation;
- persistent desktop navigation with a large flexible workspace;
- wide-screen presentation without forcing readable content to stretch unnecessarily;
- single-pane and list/detail workspace layouts;
- a generic top application header;
- grouped primary navigation with current-item presentation;
- page title, supporting text and caller-supplied action areas;
- generic neutral/info/success/warning/danger presentation tones chosen explicitly by the consuming application;
- basic form, list and table presentation primitives;
- local overflow containment for wide tabular content rather than whole-page horizontal breakage;
- a deliberately small set of CSS custom properties for spacing, surfaces, typography, borders, focus and presentation tones;
- caller styling overrides without editing the component source;
- semantic HTML landmarks and keyboard operation;
- visible keyboard focus;
- practical 44 CSS pixel minimum sizing for normal interactive controls supplied by the shell;
- status and warning presentation that never relies on colour alone;
- core navigation, links, forms and page rendering usable with JavaScript disabled;
- first useful render delivered as HTML without waiting for a WebAssembly or client framework bootstrap;
- no external font, CDN or runtime-network dependency for the shell itself;
- embedded static assets so a consuming Go binary can remain self-contained;
- optional PWA manifest generation from caller-supplied application metadata;
- optional service-worker registration as progressive enhancement;
- optional offline fallback and precaching of an explicit same-origin static-asset allowlist;
- ordinary web operation when PWA support is disabled or unavailable;
- Go 1.27 baseline;
- standard-library-only runtime dependencies;
- no Node, npm, WebAssembly UI runtime, code generation or frontend build step required to consume the component.

## Out of scope

- business logic of any kind;
- application or domain vocabulary;
- CRM, ERP, incident, customer, invoice, risk, approval or workflow semantics;
- deciding what a status means or which presentation tone it should use;
- authentication or authorisation decisions;
- role, permission or entitlement models;
- application routing decisions;
- framework-specific router/context adapters;
- business validation;
- database access or persistence;
- application state management;
- background jobs, queues or workflow engines;
- a global application engine, dependency/service container or plugin runtime;
- theme plugins with version coupling to the shell runtime;
- browser-side Go application runtimes;
- WebAssembly rendering or hydration as a shell requirement;
- virtual DOM/reconciliation engines;
- React-style hooks or browser lifecycle abstractions;
- client-side routing/state systems;
- component code generation or custom component source formats;
- client-side business behaviour;
- a JavaScript SPA router or application runtime;
- React, Vue, Svelte, Angular or another frontend framework;
- Bootstrap, Tailwind or another external CSS framework;
- HTMX as a required dependency;
- a full widget, charting, data-grid or design-system library;
- application-specific icons, logos or branding assets;
- generic offline storage of authenticated application pages;
- generic offline business-data caching;
- offline mutation queues or write synchronisation;
- conflict resolution;
- deciding which business data is safe to retain on a device;
- guaranteeing PWA installation on browsers or platforms that do not support installation;
- sanitising arbitrary application HTML supplied outside Go's normal escaping boundary.

## Acceptance criteria

1. another Go module can import the reusable package directly; the public implementation is not hidden behind Go's `internal` import boundary;
2. a consuming application can build and run with `go build` alone, with no Node/npm/code-generation step and no runtime third-party dependency required by the shell;
3. the reference integration uses only standard Go rendering/HTTP interfaces and needs no Gin, Echo, Fiber, Chi, Beego or other framework adapter package;
4. two independently configured shell renderer instances can coexist in one process without shared mutable global presentation state;
5. application content is supplied through explicit presentation slots/values and does not require plugin registration or a service container;
6. shell output is deterministic for equivalent inputs and can be tested as static HTML without launching a browser runtime;
7. the first useful page is server-rendered HTML and does not depend on downloading or starting a WebAssembly/client rendering runtime;
8. the reference application remains usable at 320 CSS pixels wide without shell-generated whole-page horizontal scrolling;
9. at common phone portrait widths, primary navigation is compact/collapsible and the workspace presents one useful pane at a time rather than squeezing desktop columns together;
10. at tablet/small-window widths, navigation and workspace reflow without a separate mobile template or URL;
11. at desktop widths, primary navigation can remain visible beside a flexible main workspace;
12. when list/detail mode is requested and enough width is available, the list and selected detail can be shown together; at compact widths the same content becomes a sequential list-or-detail experience;
13. resizing continuously between compact and wide layouts does not require a page reload or device detection;
14. the reference application remains functional at 200% browser zoom;
15. shell-generated markup uses semantic header/navigation/main structure, exposes the current navigation item, provides a skip-to-content path and preserves logical keyboard focus order;
16. shell-generated normal interactive controls use a practical minimum target size of 44 by 44 CSS pixels unless the control is naturally inline text;
17. keyboard focus is plainly visible and presentation states are not conveyed by colour alone;
18. all substantive reference-application navigation and form submission still works with JavaScript disabled;
19. the default shell loads no external font, CDN stylesheet, CDN script or other presentation dependency;
20. a consuming application can override the documented design tokens without forking or editing shell source;
21. enabling PWA mode can produce a standards-shaped web-app manifest using caller-supplied name, start URL, scope, colours and icons;
22. PWA mode is optional: disabling it removes the manifest/service-worker integration without changing ordinary application behaviour;
23. the optional service worker caches only explicitly configured same-origin static URLs and an optional offline fallback; it does not automatically cache authenticated pages, arbitrary GET responses, API responses or mutations;
24. non-GET requests are never intercepted for offline queuing or replay by the shell;
25. if the network is unavailable, configured static shell assets and the configured offline fallback remain available where the browser supports service workers, while business-data behaviour remains the consuming application's responsibility;
26. the shell's required core JavaScript budget is zero; any shipped JavaScript exists only for optional browser/PWA enhancement;
27. no `.wasm`, virtual-DOM runtime, hook runtime, hydration runtime or client router is required or shipped by V0.1;
28. the uncompressed first-party shell CSS plus optional PWA registration/service-worker JavaScript remains deliberately small, with a combined V0.1 budget of no more than 32 KiB excluding caller-supplied icons and application content;
29. the default theme and reference markup target WCAG 2.2 AA, with the component documenting that caller content and caller theme overrides remain the consuming application's responsibility;
30. the reference application demonstrates shell behaviour using neutral fictional content and contains no dependency on a specific business domain;
31. CI proves `gofmt`, `go vet`, race-enabled tests, native build, exact version, Linux ARM64 build, rendering contracts, PWA boundary contracts and the component's no-business-logic/no-client-runtime boundary;
32. release acceptance includes human browser review at approximately 320, 390, 768, 1024, 1366, 1920 and 2560 CSS-pixel widths, plus keyboard-only and 200% zoom review;
33. code-derived EDD evidence is pinned to the released implementation and the design comparison contains no unexplained mismatch.

## Definition of done

V0.1 is done when the shell can be imported by an unrelated Go application, integrated without a framework adapter or browser-side UI runtime, the neutral reference application demonstrates compact and wide layouts from one deterministic server-rendered structure, core operation remains useful without JavaScript, optional PWA support stays inside the static-shell/offline-fallback boundary, all automated gates pass, the human responsive/accessibility matrix is accepted, and code-derived design evidence confirms that no business logic, application-specific semantics or recreated frontend framework entered the component.
