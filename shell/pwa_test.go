package shell

import (
	"encoding/json"
	"strings"
	"testing"
)

func validPWA() PWAConfig {
	cfg := PWAConfig{}
	cfg.Name = "Example App"
	cfg.ShortName = "Example"
	cfg.StartURL = "/"
	cfg.Scope = "/"
	cfg.Display = "standalone"
	cfg.ThemeColor = "#ffffff"
	cfg.BackgroundColor = "#ffffff"
	cfg.Icons = []PWAIcon{{Src: "/icon-192.png", Sizes: "192x192", Type: "image/png"}}
	cfg.Precache = []string{"/_app-shell/shell.css", "/_app-shell/pwa.js"}
	cfg.OfflineURL = "/offline"
	cfg.CacheVersion = "build-1"
	return cfg
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
	for _, want := range []string{
		`event.request.method !== "GET"`,
		`requestURL.origin !== self.location.origin`,
		`PRECACHE.includes(requestKey)`,
		`event.request.mode === "navigate"`,
		`const CACHE_PREFIX = "go-app-shell-"`,
		`key.startsWith(CACHE_PREFIX) && key !== CACHE_NAME`,
	} {
		if !strings.Contains(s, want) {
			t.Fatalf("service worker boundary missing %q", want)
		}
	}
	for _, forbidden := range []string{"indexedDB", "backgroundSync", `method === "POST"`, "queue.push", "go-app-shell-v1"} {
		if strings.Contains(s, forbidden) {
			t.Fatalf("service worker contains forbidden generic offline behaviour/stale cache marker %q", forbidden)
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

func TestDefaultCacheNameChangesWithCacheInputs(t *testing.T) {
	a, err := defaultCacheName([]string{"/a"}, "build-1")
	if err != nil {
		t.Fatal(err)
	}
	b, err := defaultCacheName([]string{"/b"}, "build-1")
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatal("default cache name did not change with cache inputs")
	}
	if !strings.HasPrefix(a, defaultCachePrefix) || len(a) != len(defaultCachePrefix)+12 {
		t.Fatalf("unexpected cache name %q", a)
	}
}

func TestDefaultCacheNameChangesWithApplicationVersion(t *testing.T) {
	a, err := defaultCacheName([]string{"/icon.svg", "/offline"}, "build-1")
	if err != nil {
		t.Fatal(err)
	}
	b, err := defaultCacheName([]string{"/icon.svg", "/offline"}, "build-2")
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatal("default cache name did not change with application cache version")
	}
}

func TestCustomCacheNameDoesNotDeleteUnrelatedCaches(t *testing.T) {
	cfg := validPWA()
	cfg.CacheName = "consumer-owned-v7"
	b, err := ServiceWorker(cfg)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, `const CACHE_NAME = "consumer-owned-v7"`) {
		t.Fatal("custom cache name missing")
	}
	if !strings.Contains(s, `const CACHE_PREFIX = ""`) {
		t.Fatal("custom cache should not enable shell-owned cache cleanup")
	}
}
