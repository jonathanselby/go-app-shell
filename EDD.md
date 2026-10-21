# go-app-shell — Engineering Design Document

**Version:** 0.1

## 1. Purpose

Provide a small reusable Go presentation shell for substantial browser applications, with responsive behaviour from phone portrait to large desktop, accessible server-rendered navigation and workspace patterns, and optional Progressive Web App plumbing.

> **The shell owns presentation. The application owns meaning and behaviour.**

## 2. Design constraints

1. The component contains no business logic or domain semantics.
2. The consuming application owns routing decisions, permissions, validation, persistence, state and business behaviour.
3. The reusable package is publicly importable by other Go modules and therefore must not live solely under Go's `internal` boundary.
4. Integration is through standard Go values/interfaces; there are no framework-specific router/context adapters in V0.1.
5. Renderer/configuration state is instance-owned; there is no required global engine, mutable plugin registry or service/dependency container.
6. Core application navigation and form submission must remain useful without JavaScript.
7. V0.1 requires no Node, npm, frontend framework, code generation or frontend build step.
8. The runtime dependency set is the Go standard library only.
9. One server-rendered structure must adapt across compact and wide layouts; there is no separate mobile application or mobile URL tree.
10. Responsive behaviour is based on available layout width rather than device detection.
11. The shell may style application content but must not infer what application content means.
12. Optional PWA support may cache only an explicit static allowlist and offline fallback; it must not invent generic offline business-data behaviour.
13. The shell must work as an ordinary web application when PWA capability is disabled or unsupported.
14. The public API stays deliberately small. V0.1 is an application frame, not an admin framework or general widget library.

## 3. Architecture

```text
                  consuming Go application
                           |
                           | caller-owned meaning,
                           | routes, auth, data, actions
                           v
                 +---------------------+
                 |    go-app-shell     |
                 |                     |
                 | Renderer instance   |
                 | Page model          |
                 | Nav model           |
                 | Workspace model     |
                 | Theme tokens        |
                 | optional PWA config |
                 +----------+----------+
                            |
              +-------------+-------------+
              |                           |
              v                           v
      server-rendered HTML          embedded assets
      html/template                 shell.css
                                    optional pwa.js
                                    optional sw.js
                            |
                            v
                         browser
                            |
             +--------------+--------------+
             |                             |
         compact                       wide
       one useful pane        nav + flexible workspace
```

There is no application database, queue, state store, client-side router, component runtime, framework adapter layer, plugin registry, service container or JavaScript application state inside the shell.

## 4. Repository and package layout

The intended implementation layout is:

```text
go-app-shell/
├── .github/workflows/
│   ├── test.yml
│   └── release-tag.yml
├── CHANGELOG.md
├── EDD.md
├── LICENSE
├── PRD.md
├── README.md
├── RELEASE.md
├── VERSION
├── architecture.txt
├── go.mod
├── shell/
│   ├── shell.go
│   ├── render.go
│   ├── pwa.go
│   ├── shell_test.go
│   ├── pwa_test.go
│   ├── assets/
│   │   ├── shell.css
│   │   ├── pwa.js
│   │   └── service-worker.js
│   └── templates/
│       └── shell.html
├── cmd/
│   └── app-shell-demo/
│       └── main.go
├── examples/
│   └── neutral-demo/
└── generated/
    ├── EDD.md
    └── COMPARISON.md
```

The public package is `shell` so another repository can import it directly.

`cmd/app-shell-demo` is a neutral demonstration and acceptance surface. It is not part of the reusable API.

## 5. Public model

V0.1 should expose a small set of presentation-only types. Exact names may change during implementation if a simpler API emerges, but the semantic boundary is fixed.

Conceptually:

```go
type Renderer struct {
    // immutable compiled templates/assets plus instance configuration
}

type Page struct {
    App        AppHeader
    Navigation []NavItem
    Title      string
    Summary    string
    Actions    template.HTML
    Workspace  template.HTML
    Theme      Theme
    PWA        *PWAConfig
}

type NavItem struct {
    Label   string
    URL     string
    Current bool
    Group   string
}
```

Renderer construction/configuration returns an instance. Two differently configured renderer instances must be able to coexist in one process without shared mutable presentation state.

The component may also provide small generic presentation metadata such as tone names or layout mode.

It must not expose domain concepts such as customer, invoice, incident, approval, severity, owner, workflow state or entitlement.

The consuming application remains responsible for producing any trusted pre-rendered fragments it deliberately passes to the shell. Normal text fields remain escaped by `html/template`.

## 6. Rendering and integration model

The primary rendering path uses `html/template` from the standard library.

The shell renders:

- document structure and metadata;
- skip-to-content link;
- top application header;
- primary navigation;
- page heading/context;
- caller-provided action area;
- workspace container;
- optional PWA metadata/registration hook.

The shell does not register application routes and does not know a framework request context.

The primary API should render to an `io.Writer` (or equivalent standard-library boundary) from explicit page data. This keeps the shell usable from `net/http`, Gin, Echo, Chi, Fiber, Beego or another stack without carrying adapters or their dependencies.

Optional convenience helpers may return `http.Handler` only for shell-owned static/PWA assets because `net/http` is standard library. They must not become an application router.

No client-side templating is required.

## 7. Why there is no adapter layer

Framework-specific adapters appear to increase compatibility but actually couple the component to framework APIs, versions and transitive dependencies.

`go-app-shell` instead targets the smaller common denominator:

```text
application/framework handler
          |
          | construct presentation Page
          v
shell.Render(io.Writer, Page)
```

A consuming application's existing router remains untouched.

This provides broader compatibility with less code than maintaining adapters for each framework.

## 8. Why there is no plugin/runtime engine

V0.1 has no global `Engine`, plugin registry, dependency/service container or runtime extension mechanism.

Extensibility comes from ordinary Go composition:

- the application decides what page fragments to render;
- the application owns its handlers/services/data;
- shell configuration is explicit;
- theming uses CSS custom-property overrides;
- optional richer client behaviour can be added by the consuming application independently.

This prevents presentation infrastructure becoming a second application framework.

## 9. Responsive layout

The layout is one semantic document structure with CSS adaptation.

Target behaviour:

```text
320-599px   compact
600-1023px  medium
1024px+     desktop/wide
```

These are layout ranges, not device identities.

### Compact

- top application header remains visible;
- primary navigation is not permanently width-consuming;
- one useful workspace pane is shown at a time;
- list/detail mode becomes sequential list-or-detail presentation;
- forms use the available width;
- tables or intrinsically wide content scroll locally inside their content region if necessary;
- the shell itself must not create whole-page horizontal scrolling.

### Medium

- navigation may use compact/overlay presentation;
- workspace remains primary;
- layouts can use additional horizontal space where it materially improves usability.

### Desktop and wide

- primary navigation may remain persistent;
- main workspace receives remaining width;
- list/detail can show two panes simultaneously;
- readable prose/form regions may retain sensible maximum widths while data-heavy workspaces may use more available space.

CSS Grid/Flexbox provide macro layout. CSS custom properties provide tokens. Container queries may be used where a reusable region should react to its own allocated space rather than the whole viewport.

There is no JavaScript resize controller.

## 10. Navigation model

Navigation is presentation metadata supplied by the application.

The shell may render:

- grouped primary navigation;
- current-item state;
- compact navigation trigger where required by layout;
- accessible labels/expanded state where JavaScript enhancement is present.

The shell does not decide which items a user may see. The application passes only the navigation entries appropriate for that request.

With JavaScript unavailable, all supplied navigation remains reachable through ordinary links. Compact presentation must degrade to an accessible static navigation mechanism rather than making navigation dependent on script.

## 11. Workspace modes

V0.1 needs only two workspace concepts:

### Single

One primary content region. Suitable for forms, reports, dashboards, settings and ordinary pages.

### List/detail

A collection/list region plus a selected detail region.

At sufficient width both may be visible. At compact width the application determines whether the current route represents the list or the detail; the shell presents the current pane full width.

The shell does not maintain selected-record state. Selection remains URL/request/application state.

This preserves normal browser behaviour, deep links, back/forward navigation and no-JavaScript operation.

## 12. Presentation vocabulary

V0.1 deliberately avoids becoming a full design system.

The reusable presentation vocabulary is limited to:

- shell/header/navigation/workspace;
- page heading/context/actions;
- generic notice/status tones;
- standard form spacing and controls;
- simple list/table treatment;
- selected/current states;
- empty-state presentation;
- local overflow containment;
- focus/keyboard presentation.

The component may define generic tone names:

```text
neutral
info
success
warning
danger
```

These are presentation choices only. The shell never maps business values to them.

## 13. Design tokens and theming

A deliberately small CSS custom-property set should cover the shell:

```text
--gas-header-height
--gas-nav-width
--gas-space-1 .. --gas-space-5
--gas-radius
--gas-border
--gas-text
--gas-text-muted
--gas-surface
--gas-surface-raised
--gas-accent
--gas-focus
--gas-success
--gas-warning
--gas-danger
```

Names may be shortened if the implementation becomes clearer, but tokens must be namespaced enough not to collide casually with consuming applications.

The default theme uses system fonts only.

Consumers override documented tokens after loading the embedded/default stylesheet. They should not need to fork the shell CSS for ordinary branding.

There is no theme plugin interface and no shell/theme version negotiation. Theme customisation is data/CSS, not a runtime subsystem.

## 14. Accessibility boundary

The shell targets WCAG 2.2 AA for the markup and defaults it owns.

The shell owns:

- semantic `header`, `nav`, `main` and heading structure;
- a skip-to-main-content link;
- current navigation indication;
- visible `:focus-visible` treatment;
- logical document/focus order;
- non-colour status cues;
- practical 44px sizing for normal shell-supplied controls;
- responsive reflow that remains usable at 200% browser zoom;
- reduced-motion respect if motion is introduced.

The consuming application owns:

- accessible names/labels for its controls;
- application content hierarchy inside supplied fragments;
- image alternative text;
- application-specific keyboard interactions;
- contrast of custom theme overrides;
- accessibility of third-party/application widgets embedded inside the workspace.

## 15. JavaScript boundary

Core shell operation requires zero JavaScript.

Optional JavaScript is restricted to genuine browser enhancement, primarily:

- PWA service-worker registration;
- optional enhancement of compact navigation if the no-script fallback remains fully usable.

There is no shell client-side state model, SPA router, validation engine, data-fetch layer or component framework.

The V0.1 first-party uncompressed CSS plus optional JavaScript budget is at most 32 KiB, excluding icons and application content.

## 16. PWA model

PWA support is optional and configured by the consuming application.

A `PWAConfig` supplies presentation/install metadata such as:

- application name/short name;
- start URL;
- scope;
- theme/background colours;
- display preference;
- icon metadata;
- explicit static precache URLs;
- optional offline fallback URL.

The shell may expose standard handlers/helpers for:

```text
manifest.webmanifest
service-worker.js
pwa-registration.js
```

The consuming application decides where these are mounted.

PWA support must not change the application's routing, authentication or data model.

## 17. Service-worker safety boundary

The service worker is intentionally conservative.

It may:

- precache explicitly configured same-origin static URLs;
- serve those exact cached static resources while offline;
- serve a configured offline fallback for navigation failure where appropriate.

It must not generically:

- cache arbitrary authenticated HTML pages;
- cache API responses;
- intercept or queue POST/PUT/PATCH/DELETE;
- replay mutations;
- store business records;
- implement conflict resolution;
- infer which application responses are safe to persist.

Any richer offline capability belongs to the consuming application's threat model and business design.

## 18. Security and trust boundary

The shell is not a security or authorisation component.

Trusted inputs include:

- consuming application configuration;
- caller-selected navigation items;
- caller-selected presentation tone;
- explicitly configured PWA static URLs;
- application-owned trusted HTML fragments deliberately passed across the rendering boundary.

Untrusted user text should be passed as ordinary strings so `html/template` escapes it.

The component must not expose helpers that encourage arbitrary untrusted strings to be cast to `template.HTML`.

Security headers, CSRF, authentication, sessions, authorisation and content-security policy remain application/server responsibilities because deployment needs differ.

## 19. Failure behaviour

- invalid shell configuration: renderer construction/configuration returns an explicit error;
- invalid or unsafe PWA precache URL: refuse configuration rather than broaden cache scope;
- template execution failure: return an error to the caller; never write a fabricated success page;
- unsupported PWA browser: ordinary web application remains functional;
- JavaScript disabled: core server-rendered UI remains usable;
- service worker unavailable/fails registration: ordinary web application remains usable;
- offline and uncached business route requested: offline fallback or browser failure is shown; shell must not synthesize stale business data.

## 20. Dependencies and runtime

V0.1 baseline:

- Go 1.27;
- standard library only;
- embedded templates/CSS/optional JS via `embed`;
- no runtime CDN/network dependency;
- no Node/npm/toolchain dependency;
- no frontend code-generation dependency;
- no router/framework adapters;
- no database drivers;
- no auth/RBAC package;
- no theme package dependency.

The component is linked into the consuming application binary. It is not a service and has no daemon/runtime of its own.

## 21. Neutral reference application

`cmd/app-shell-demo` exists to prove the shell rather than to become a product.

It should contain neutral fictional data such as generic records/items and demonstrate:

- standard `net/http` integration;
- top header;
- grouped navigation;
- single workspace;
- list/detail workspace;
- forms;
- notices/status tones;
- table/list behaviour;
- compact navigation;
- theme override;
- optional PWA mode;
- no-JavaScript operation.

It must not embed CRM, ERP, cyber, risk, finance, approval or workflow semantics.

The demo is proof of the smallest integration surface, not a required hosting framework.

## 22. Test strategy

### Go package tests

Automated tests should cover at minimum:

- successful render with escaped caller text;
- semantic shell landmarks;
- skip-link presence;
- current navigation semantics;
- two independent renderer instances do not share mutable state;
- renderer writes to standard `io.Writer`/standard-library boundary;
- no framework adapter packages or framework dependencies enter the module;
- no global plugin/service registry is required;
- no accidental domain vocabulary in public API/neutral fixtures;
- theme-token output contract;
- embedded asset availability;
- manifest generation from caller metadata;
- PWA disabled path emits no manifest/service-worker integration;
- service-worker configuration rejects non-same-origin/external precache targets;
- service-worker source contains no non-GET mutation queue/replay behaviour;
- asset size budget;
- deterministic rendering where inputs are equal.

### CI

CI should prove:

```text
gofmt clean
go vet ./...
go test -race -count=1 ./...
native demo build
exact version contract
GOOS=linux GOARCH=arm64 build
rendering acceptance
framework-neutral integration acceptance
PWA boundary acceptance
no-business-logic boundary acceptance
```

A no-JavaScript acceptance can be exercised through direct HTTP requests/HTML contracts because all core routes are ordinary server-rendered links/forms.

### Human acceptance

Browser acceptance is required because responsive presentation cannot be proved completely through Go unit tests.

Review approximately:

```text
320
390
768
1024
1366
1920
2560 CSS px
```

Also check:

- continuous resize between those widths;
- phone portrait;
- desktop landscape;
- keyboard-only navigation;
- 200% zoom;
- browser print sanity where applicable;
- PWA standalone launch where supported;
- offline static shell/fallback where service workers are supported;
- JavaScript disabled for core web flow.

## 23. Known limitations

- V0.1 is not a complete visual component library;
- no generic data grid, charts, modal system or rich client widgets;
- no generic offline business-data mode;
- no framework-specific convenience adapter in V0.1;
- PWA install UX varies by browser/platform and cannot be normalised by the library;
- shell CSS cannot guarantee accessibility of arbitrary caller content or custom theme overrides;
- visual/browser acceptance remains partly human because CSS reflow and browser rendering are not fully represented by Go tests;
- 320 CSS pixels is the V0.1 lower acceptance width, not a promise for arbitrarily narrower displays;
- advanced interactions may be added by consuming applications independently, including optional HTMX, provided they preserve their own accessibility and security contracts.

## 24. Acceptance-to-evidence map

| Requirement | Planned evidence |
|---|---|
| public cross-repository reuse | external/import compilation test or demo module build |
| no frontend toolchain/runtime dependency | `go.mod`, CI build and repository inspection |
| framework-neutral integration | standard `net/http` demo + dependency/source contract test |
| no adapter proliferation | module/repository source contract |
| independent renderer instances | concurrency/state-isolation test |
| no plugin/service runtime | public API/source contract |
| 320px compact usability | human responsive acceptance + shell overflow contract |
| phone portrait one-pane behaviour | human acceptance + CSS contract |
| tablet/window reflow | human responsive acceptance |
| desktop persistent navigation | human acceptance + CSS contract |
| list/detail adaptive behaviour | neutral demo + human responsive acceptance |
| continuous resize/no device detection | CSS/source contract + human resize test |
| 200% zoom | human accessibility acceptance |
| semantic landmarks/current nav/skip link | render tests |
| 44px shell control sizing | CSS contract test + human touch review |
| visible focus/non-colour states | CSS/render contract + human keyboard review |
| core works without JS | neutral demo HTTP acceptance with script disabled/absent |
| no external runtime assets | asset/source test |
| documented theming | token contract test + neutral demo theme override |
| optional valid manifest | manifest tests |
| PWA disabled cleanly | render/handler tests |
| static allowlist only | PWA configuration/service-worker tests |
| no mutation queue/replay | service-worker source contract test |
| offline static fallback only | service-worker contract + human offline check |
| zero required JS | render/source contract |
| <=32 KiB first-party asset budget | CI size test |
| WCAG-targeted defaults | render/CSS tests + human acceptance |
| neutral/no-domain reference | source vocabulary boundary test + review |
| format/vet/race/native/version/ARM64 | `.github/workflows/test.yml` |
| code/design alignment | `generated/EDD.md` + `generated/COMPARISON.md` pinned to source |

## 25. Definition of done

V0.1 is complete when all PRD acceptance criteria have evidence, the public package remains a framework-neutral presentation component rather than an admin framework, two independent renderer instances can coexist without global state, the neutral demo proves compact and wide operation plus optional PWA mode, browser/accessibility acceptance passes, and code-derived comparison contains no unexplained mismatch or leaked business/framework responsibility.
