package shell

import (
	"bytes"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func testPage() Page {
	return Page{
		DocumentTitle: "Neutral demo",
		App:           AppHeader{Name: "Example", Context: "Workspace", Status: "Ready", UserLabel: "Operator"},
		Navigation:    []NavGroup{{Label: "Main", Items: []NavItem{{Label: "Overview", URL: "/", Current: true}, {Label: "Records", URL: "/records"}}}},
		Title:         "Overview",
		Summary:       "A neutral demonstration.",
		Actions:       template.HTML(`<a class="gas-button" href="/form">Add item</a>`),
		Workspace:     Workspace{Mode: WorkspaceSingle, Content: template.HTML(`<p>Useful content.</p>`)},
	}
}

func TestRenderSemanticShellAndEscapesText(t *testing.T) {
	r, err := New(Config{})
	if err != nil {
		t.Fatal(err)
	}
	page := testPage()
	page.Title = `<script>alert(1)</script>`
	var b bytes.Buffer
	if err := r.Render(&b, page); err != nil {
		t.Fatal(err)
	}
	html := b.String()
	for _, want := range []string{`<header class="gas-header">`, `<nav class="gas-nav" aria-label="Primary">`, `<main id="main-content"`, `class="gas-skip-link"`, `aria-current="page"`} {
		if !strings.Contains(html, want) {
			t.Fatalf("missing %q", want)
		}
	}
	if strings.Contains(html, `<script>alert(1)</script>`) {
		t.Fatal("untrusted text was not escaped")
	}
	if !strings.Contains(html, `&lt;script&gt;alert(1)&lt;/script&gt;`) {
		t.Fatal("escaped title missing")
	}
}

func TestRenderDeterministic(t *testing.T) {
	r, _ := New(Config{})
	var a, b bytes.Buffer
	page := testPage()
	if err := r.Render(&a, page); err != nil {
		t.Fatal(err)
	}
	if err := r.Render(&b, page); err != nil {
		t.Fatal(err)
	}
	if a.String() != b.String() {
		t.Fatal("equivalent input produced different output")
	}
}

func TestIndependentRenderers(t *testing.T) {
	a, err := New(Config{AssetBase: "/a"})
	if err != nil {
		t.Fatal(err)
	}
	b, err := New(Config{AssetBase: "/b"})
	if err != nil {
		t.Fatal(err)
	}
	var ao, bo bytes.Buffer
	if err := a.Render(&ao, testPage()); err != nil {
		t.Fatal(err)
	}
	if err := b.Render(&bo, testPage()); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(ao.String(), `href="/a/shell.css"`) || strings.Contains(ao.String(), `href="/b/shell.css"`) {
		t.Fatal("renderer A leaked state")
	}
	if !strings.Contains(bo.String(), `href="/b/shell.css"`) {
		t.Fatal("renderer B configuration missing")
	}
}

func TestListDetailCompactPaneContract(t *testing.T) {
	r, _ := New(Config{})
	page := testPage()
	page.Workspace = Workspace{Mode: WorkspaceListDetail, List: template.HTML(`<p>list</p>`), Detail: template.HTML(`<p>detail</p>`), CompactPane: CompactDetail}
	var b bytes.Buffer
	if err := r.Render(&b, page); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(b.String(), `data-compact-pane="detail"`) {
		t.Fatal("compact pane missing")
	}
}

func TestCoreMarkupDoesNotDependOnJavaScript(t *testing.T) {
	r, _ := New(Config{})
	var b bytes.Buffer
	if err := r.Render(&b, testPage()); err != nil {
		t.Fatal(err)
	}
	html := b.String()
	if strings.Contains(html, `<script`) {
		t.Fatal("PWA-disabled page unexpectedly requires script")
	}
	for _, want := range []string{`type="checkbox"`, `for="gas-nav-toggle"`, `href="/records"`, `<main`} {
		if !strings.Contains(html, want) {
			t.Fatalf("no-script navigation contract missing %q", want)
		}
	}
}

func TestPWAHeadIsOptIn(t *testing.T) {
	r, _ := New(Config{})
	page := testPage()
	page.PWA = &PWAHead{ManifestURL: "/manifest.webmanifest", RegistrationURL: "/_app-shell/pwa.js", ServiceWorkerURL: "/service-worker.js"}
	var b bytes.Buffer
	if err := r.Render(&b, page); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`rel="manifest" href="/manifest.webmanifest"`, `data-worker-url="/service-worker.js"`} {
		if !strings.Contains(b.String(), want) {
			t.Fatalf("missing PWA opt-in %q", want)
		}
	}
}

func TestAssetsAndSizeBudget(t *testing.T) {
	assets, err := AssetsFS()
	if err != nil {
		t.Fatal(err)
	}
	var total int64
	for _, name := range []string{"shell.css", "pwa.js"} {
		info, err := fs.Stat(assets, name)
		if err != nil {
			t.Fatal(err)
		}
		total += info.Size()
	}
	if total > 32*1024 {
		t.Fatalf("first-party asset budget exceeded: %d bytes", total)
	}
	css, err := fs.ReadFile(assets, "shell.css")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"@media (min-width: 64rem)", "overflow-x: auto", "min-height: 2.75rem", "@media print"} {
		if !strings.Contains(string(css), want) {
			t.Fatalf("CSS contract missing %q", want)
		}
	}
}

func TestNoBusinessVocabularyOrClientRuntimeInCoreSource(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	banned := []string{"customer", "invoice", "incident", "severity", "entitlement", "virtualdom", "usestat", "hydrate("}
	entries, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range entries {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		lower := strings.ToLower(string(data))
		for _, term := range banned {
			if strings.Contains(lower, term) {
				t.Fatalf("core source %s contains prohibited vocabulary/runtime marker %q", filepath.Base(path), term)
			}
		}
	}
}

func BenchmarkRender(b *testing.B) {
	r, _ := New(Config{})
	page := testPage()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		if err := r.Render(&out, page); err != nil {
			b.Fatal(err)
		}
	}
}
