package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jonathanselby/go-app-shell/shell"
)

// demoBuildVersion represents the consuming application's build version.
// Real applications can inject a Git SHA or release identifier with -ldflags.
var demoBuildVersion = "dev"

type app struct {
	renderer *shell.Renderer
	pwa      bool
}

type record struct {
	ID       string
	Label    string
	Selected bool
}

type formView struct {
	Label   string
	Invalid bool
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	version := flag.Bool("version", false, "print version")
	noPWA := flag.Bool("no-pwa", false, "disable PWA demonstration")
	flag.Parse()
	if *version {
		fmt.Printf("app-shell-demo %s\n", shell.Version)
		return
	}

	r, err := shell.New(shell.Config{AssetBase: "/_app-shell"})
	if err != nil {
		log.Fatal(err)
	}
	a := &app{renderer: r, pwa: !*noPWA}

	assets, err := shell.AssetsFS()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /_app-shell/", http.StripPrefix("/_app-shell/", http.FileServer(http.FS(assets))))
	mux.HandleFunc("GET /", a.overview)
	mux.HandleFunc("GET /records", a.records)
	mux.HandleFunc("GET /records/{id}", a.recordDetail)
	mux.HandleFunc("GET /form", a.form)
	mux.HandleFunc("POST /form", a.formPost)
	mux.HandleFunc("GET /offline", offline)
	mux.HandleFunc("GET /icon.svg", icon)
	mux.HandleFunc("GET /manifest.webmanifest", manifest)
	mux.HandleFunc("GET /service-worker.js", serviceWorker)

	log.Printf("go-app-shell demo %s at http://%s", shell.Version, *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

func (a *app) basePage(title, summary, current string) shell.Page {
	page := shell.Page{
		DocumentTitle: title + " — App Shell Demo",
		App:           shell.AppHeader{Name: "App Shell Demo", Context: "Neutral reference", Status: "Ready", UserLabel: "Example user"},
		Navigation: []shell.NavGroup{
			{Label: "Workspace", Items: []shell.NavItem{
				{Label: "Overview", URL: "/", Current: current == "overview"},
				{Label: "Records", URL: "/records", Current: current == "records"},
				{Label: "Form", URL: "/form", Current: current == "form"},
			}},
		},
		Title: title, Summary: summary,
	}
	if a.pwa {
		page.PWA = &shell.PWAHead{ManifestURL: "/manifest.webmanifest", RegistrationURL: "/_app-shell/pwa.js", ServiceWorkerURL: "/service-worker.js"}
	}
	return page
}

func (a *app) overview(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	page := a.basePage("Overview", "A deliberately neutral application showing the reusable shell rather than a business domain.", "overview")
	page.Actions = template.HTML(`<a class="gas-button gas-button--primary" href="/form">Open form</a>`)
	page.Workspace = shell.Workspace{Mode: shell.WorkspaceSingle, Content: template.HTML(`
<div class="gas-stack">
  <div class="gas-notice gas-notice--info">This page is server-rendered and remains useful without JavaScript.</div>
  <section class="gas-panel">
    <h2>Quiet workspace</h2>
    <p>The shell provides layout, navigation and presentation conventions. The consuming application supplies meaning and behaviour.</p>
  </section>
  <div class="gas-table-wrap">
    <table><thead><tr><th>Item</th><th>State</th><th>Updated</th></tr></thead>
    <tbody><tr><td>Alpha</td><td><span class="gas-badge gas-badge--success">Ready</span></td><td>Today</td></tr><tr><td>Bravo</td><td><span class="gas-badge gas-badge--warning">Review</span></td><td>Yesterday</td></tr></tbody></table>
  </div>
</div>`)}
	a.render(w, page)
}

func (a *app) records(w http.ResponseWriter, _ *http.Request) {
	page := a.basePage("Records", "Select an item. Tablet and desktop layouts show list and detail together; compact screens keep one useful pane.", "records")
	page.Workspace = shell.Workspace{Mode: shell.WorkspaceListDetail, CompactPane: shell.CompactList, List: recordList(""), Detail: template.HTML(`<div class="gas-empty">Select a record to view its detail.</div>`)}
	a.render(w, page)
}

func (a *app) recordDetail(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id != "alpha" && id != "bravo" && id != "charlie" {
		http.NotFound(w, r)
		return
	}
	label := strings.ToUpper(id[:1]) + id[1:]
	page := a.basePage("Records", "A generic list/detail workspace with URL-owned selection state.", "records")
	page.Workspace = shell.Workspace{
		Mode:        shell.WorkspaceListDetail,
		CompactPane: shell.CompactDetail,
		List:        recordList(id),
		Detail: mustFragment(`
<div class="gas-stack">
  <a class="gas-compact-only" href="/records">← Back to records</a>
  <section class="gas-panel gas-stack">
    <h2>{{.}}</h2>
    <dl class="gas-kv">
      <dt>Reference</dt><dd>{{.}}</dd>
      <dt>State</dt><dd><span class="gas-badge gas-badge--success">Available</span></dd>
    </dl>
    <div class="gas-notice gas-notice--success">The selected record is available.</div>
  </section>
</div>`, label),
	}
	a.render(w, page)
}

func recordList(selected string) template.HTML {
	records := make([]record, 0, 3)
	for _, id := range []string{"alpha", "bravo", "charlie"} {
		records = append(records, record{ID: id, Label: strings.ToUpper(id[:1]) + id[1:], Selected: id == selected})
	}
	return mustFragment(`<ul class="gas-list">{{range .}}<li><a href="/records/{{.ID}}"{{if .Selected}} aria-current="true"{{end}}>{{.Label}}</a></li>{{end}}</ul>`, records)
}

func (a *app) form(w http.ResponseWriter, _ *http.Request) {
	page := a.basePage("Form", "Ordinary semantic HTML styled by the shell; validation and business meaning remain upstream.", "form")
	page.Workspace = shell.Workspace{Mode: shell.WorkspaceSingle, Content: formContent(formView{})}
	a.render(w, page)
}

func (a *app) formPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	label := r.Form.Get("label")
	if strings.TrimSpace(label) == "" {
		page := a.basePage("Form", "Correct the highlighted field and submit again.", "form")
		page.DocumentTitle = "Error: " + page.DocumentTitle
		page.Workspace = shell.Workspace{Mode: shell.WorkspaceSingle, Content: formContent(formView{Label: label, Invalid: true})}
		a.renderStatus(w, page, http.StatusUnprocessableEntity)
		return
	}
	page := a.basePage("Form", "The demo accepted the ordinary server-side form submission.", "form")
	page.Workspace = shell.Workspace{Mode: shell.WorkspaceSingle, Content: template.HTML(`<div class="gas-stack"><div class="gas-notice gas-notice--success">Submission received by the reference application.</div><p><a href="/form">Return to form</a></p></div>`)}
	a.render(w, page)
}

func formContent(v formView) template.HTML {
	return mustFragment(`
<div class="gas-stack">
  {{if .Invalid}}<div class="gas-error-summary" role="alert" tabindex="-1" autofocus><strong>Check the form.</strong><a href="#label">Label is required.</a></div>{{end}}
  <form class="gas-form" method="post" action="/form">
    <div class="gas-field">
      <label for="label">Label</label>
      <input id="label" name="label" value="{{.Label}}" required{{if .Invalid}} aria-invalid="true" aria-describedby="label-error"{{end}}>
      {{if .Invalid}}<div id="label-error" class="gas-field-error">Enter a label.</div>{{else}}<div class="gas-help">Neutral demonstration text.</div>{{end}}
    </div>
    <div class="gas-field"><label for="kind">Kind</label><select id="kind" name="kind"><option>One</option><option>Two</option><option>Three</option></select></div>
    <div class="gas-field"><label for="notes">Notes</label><textarea id="notes" name="notes"></textarea></div>
    <div class="gas-cluster"><button class="gas-button--primary" type="submit">Submit</button><a class="gas-button" href="/">Cancel</a></div>
  </form>
</div>`, v)
}

func mustFragment(source string, data any) template.HTML {
	t := template.Must(template.New("fragment").Parse(source))
	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		panic(err)
	}
	return template.HTML(b.String())
}

func (a *app) render(w http.ResponseWriter, page shell.Page) {
	a.renderStatus(w, page, http.StatusOK)
}

func (a *app) renderStatus(w http.ResponseWriter, page shell.Page, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	if err := a.renderer.Render(w, page); err != nil {
		if status == http.StatusOK {
			http.Error(w, "render failed", http.StatusInternalServerError)
		} else {
			log.Printf("render failed after status %d: %v", status, err)
		}
	}
}

func pwaConfig() shell.PWAConfig {
	return shell.PWAConfig{
		Name:            "App Shell Demo",
		ShortName:       "Shell Demo",
		StartURL:        "/",
		Scope:           "/",
		Display:         "standalone",
		ThemeColor:      "#ffffff",
		BackgroundColor: "#ffffff",
		Icons:           []shell.PWAIcon{{Src: "/icon.svg", Sizes: "any", Type: "image/svg+xml"}},
		Precache:        []string{"/_app-shell/shell.css", "/_app-shell/pwa.js", "/icon.svg"},
		OfflineURL:      "/offline",
		CacheVersion:    demoBuildVersion,
	}
}

func manifest(w http.ResponseWriter, _ *http.Request) {
	b, err := shell.Manifest(pwaConfig())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/manifest+json")
	_, _ = w.Write(b)
}

func serviceWorker(w http.ResponseWriter, _ *http.Request) {
	b, err := shell.ServiceWorker(pwaConfig())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Service-Worker-Allowed", "/")
	_, _ = w.Write(b)
}

func offline(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Offline</title><main><h1>Offline</h1><p>The application is unavailable. No business data has been invented or replayed.</p></main></html>`))
}

func icon(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	_, _ = w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><rect width="512" height="512" rx="96" fill="#2457a7"/><path d="M128 144h256v56H128zm0 96h160v56H128zm0 96h224v56H128z" fill="white"/></svg>`))
}
