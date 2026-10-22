package shell

import (
	"encoding/json"
	"strings"
	"testing"
)

func validPWA() PWAConfig {
	return PWAConfig{
		Name: "Example App", ShortName: "Example", StartURL: "/", Scope: "/", Display: "standalone",
		ThemeColor: "#ffffff", BackgroundColor: "#ffffff",
		Icons:    []PWAIcon{{Src: "/icon-192.png", Sizes: "192x192", Type: "image/png"}},
		Precache: []string{"/_app-shell/shell.css", "/_app-shell/pwa.js"}, OfflineURL: "/offline",
	}
}

func TestManifest(t *testing.T) {
	b, err := Manifest(validPWA())
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got["name"] != "Example App" || got["start_url"] != "/" || got["display"] != "standalone" {
		t.Fatalf("unexpected manifest: %v", got)
	}
}

func TestPWARefusesExternalPersistenceTargets(t *testing.T) {
	cases := []PWAConfig{}
	base := validPWA()
	a := base
	a.Precache = []string{"https://example.com/app.js"}
	cases = append(cases, a)
	b := base
	b.Precache = []string{"//example.com/app.js"}
	cases = append(cases, b)
	c := base
	c.OfflineURL = "https://example.com/offline"
	cases = append(cases, c)
	d := base
	d.Icons = []PWAIcon{{Src: "https://example.com/icon.png"}}
	cases = append(cases, d)
	for i, cfg := range cases {
		if _, err := Manifest(cfg); err == nil && i == 3 {
			t.Fatalf("case %d accepted external target", i)
		}
		if _, err := ServiceWorker(cfg); err == nil {
			t.Fatalf("case %d accepted external target", i)
		}
	}
}

func TestServiceWorkerBoundary(t *testing.T) {
	b, err := ServiceWorker(validPWA())
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	for _, want := range []string{`event.request.method !== "GET"`, `requestURL.origin !== self.location.origin`, `PRECACHE.includes(requestKey)`, `event.request.mode === "navigate"`} {
		if !strings.Contains(s, want) {
			t.Fatalf("service worker boundary missing %q", want)
		}
	}
	for _, forbidden := range []string{"indexedDB", "backgroundSync", `method === "POST"`, "queue.push"} {
		if strings.Contains(s, forbidden) {
			t.Fatalf("service worker contains forbidden generic offline behaviour %q", forbidden)
		}
	}
}

func TestServiceWorkerDeterministic(t *testing.T) {
	cfg := validPWA()
	cfg.Precache = []string{"/b", "/a", "/b"}
	a, err := ServiceWorker(cfg)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ServiceWorker(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != string(b) {
		t.Fatal("service worker output is not deterministic")
	}
	if strings.Count(string(a), `"/b"`) != 1 {
		t.Fatal("precache entries were not deduplicated")
	}
}
