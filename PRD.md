# go-app-shell — Product Requirements Document

**Version:** 0.1

## Problem

Substantial Go web applications repeatedly need the same presentation foundation: a clear application frame, persistent navigation on large screens, usable mobile portrait behaviour, list/detail workspaces, accessible forms and status presentation, and an optional path to installation as a Progressive Web App.

Without a small reusable shell, each application either rebuilds these mechanics independently or imports a large frontend framework and its toolchain. Both choices create avoidable duplication, inconsistency and complexity.

The missing capability is presentation infrastructure, not application behaviour.

## Product objective

Provide a small reusable Go component that gives server-rendered business applications a polished, responsive and accessible application shell from phone portrait to large desktop, while remaining useful as an ordinary web application and as an optional Progressive Web App.

> **The shell owns presentation. The application owns meaning and behaviour.**

## Users

Go applications that need a professional browser interface without adopting a JavaScript application framework, frontend build system or domain-specific UI platform.

The component is intended to be reusable across unrelated business applications such as operational tools, CRM, ERP and other substantial Go systems without knowing anything about those domains.

## In scope

- a public importable Go package;
- server-rendered HTML using Go's standard web/template stack;
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
- no external font, CDN or runtime-network dependency for the shell itself;
- embedded static assets so a consuming Go binary can remain self-contained;
- optional PWA manifest generation from caller-supplied application metadata;
- optional service-worker registration as progressive enhancement;
- optional offline fallback and precaching of an explicit same-origin static-asset allowlist;
- ordinary web operation when PWA support is disabled or unavailable;
- Go 1.27 baseline;
- standard-library-only runtime dependencies;
- no Node, npm, code generation or frontend build step required to consume the component.

## Out of scope

- business logic of any kind;
- application or domain vocabulary;
- CRM, ERP, incident, customer, invoice, risk, approval or workflow semantics;
- deciding what a status means or which presentation tone it should use;
- authentication or authorisation decisions;
- role, permission or entitlement models;
- application routing decisions;
- business validation;
- database access or persistence;
- application state management;
- background jobs, queues or workflow engines;
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
3. the reference application remains usable at 320 CSS pixels wide without shell-generated whole-page horizontal scrolling;
4. at common phone portrait widths, primary navigation is compact/collapsible and the workspace presents one useful pane at a time rather than squeezing desktop columns together;
5. at tablet/small-window widths, navigation and workspace reflow without a separate mobile template or URL;
6. at desktop widths, primary navigation can remain visible beside a flexible main workspace;
7. when list/detail mode is requested and enough width is available, the list and selected detail can be shown together; at compact widths the same content becomes a sequential list-or-detail experience;
8. resizing continuously between compact and wide layouts does not require a page reload or device detection;
9. the reference application remains functional at 200% browser zoom;
10. shell-generated markup uses semantic header/navigation/main structure, exposes the current navigation item, provides a skip-to-content path and preserves logical keyboard focus order;
11. shell-generated normal interactive controls use a practical minimum target size of 44 by 44 CSS pixels unless the control is naturally inline text;
12. keyboard focus is plainly visible and presentation states are not conveyed by colour alone;
13. all substantive reference-application navigation and form submission still works with JavaScript disabled;
14. the default shell loads no external font, CDN stylesheet, CDN script or other presentation dependency;
15. a consuming application can override the documented design tokens without forking or editing shell source;
16. enabling PWA mode can produce a standards-shaped web-app manifest using caller-supplied name, start URL, scope, colours and icons;
17. PWA mode is optional: disabling it removes the manifest/service-worker integration without changing ordinary application behaviour;
18. the optional service worker caches only explicitly configured same-origin static URLs and an optional offline fallback; it does not automatically cache authenticated pages, arbitrary GET responses, API responses or mutations;
19. non-GET requests are never intercepted for offline queuing or replay by the shell;
20. if the network is unavailable, configured static shell assets and the configured offline fallback remain available where the browser supports service workers, while business-data behaviour remains the consuming application's responsibility;
21. the shell's required core JavaScript budget is zero; any shipped JavaScript exists only for optional browser/PWA enhancement;
22. the uncompressed first-party shell CSS plus optional PWA registration/service-worker JavaScript remains deliberately small, with a combined V0.1 budget of no more than 32 KiB excluding caller-supplied icons and application content;
23. the default theme and reference markup target WCAG 2.2 AA, with the component documenting that caller content and caller theme overrides remain the consuming application's responsibility;
24. the reference application demonstrates shell behaviour using neutral fictional content and contains no dependency on a specific business domain;
25. CI proves `gofmt`, `go vet`, race-enabled tests, native build, exact version, Linux ARM64 build, rendering contracts, PWA boundary contracts and the component's no-business-logic boundary;
26. release acceptance includes human browser review at approximately 320, 390, 768, 1024, 1366, 1920 and 2560 CSS-pixel widths, plus keyboard-only and 200% zoom review;
27. code-derived EDD evidence is pinned to the released implementation and the design comparison contains no unexplained mismatch.

## Definition of done

V0.1 is done when the shell can be imported by an unrelated Go application, the neutral reference application demonstrates compact and wide layouts from one server-rendered structure, core operation remains useful without JavaScript, optional PWA support stays inside the static-shell/offline-fallback boundary, all automated gates pass, the human responsive/accessibility matrix is accepted, and code-derived design evidence confirms that no business logic or application-specific semantics entered the component.
