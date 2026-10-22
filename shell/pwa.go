package shell

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type PWAIcon struct {
	Src     string `json:"src"`
	Sizes   string `json:"sizes,omitempty"`
	Type    string `json:"type,omitempty"`
	Purpose string `json:"purpose,omitempty"`
}

type PWAConfig struct {
	Name            string
	ShortName       string
	StartURL        string
	Scope           string
	Display         string
	ThemeColor      string
	BackgroundColor string
	Icons           []PWAIcon
	Precache        []string
	OfflineURL      string
	CacheName       string
}

type manifest struct {
	Name            string    `json:"name"`
	ShortName       string    `json:"short_name,omitempty"`
	StartURL        string    `json:"start_url"`
	Scope           string    `json:"scope,omitempty"`
	Display         string    `json:"display,omitempty"`
	ThemeColor      string    `json:"theme_color,omitempty"`
	BackgroundColor string    `json:"background_color,omitempty"`
	Icons           []PWAIcon `json:"icons,omitempty"`
}

func Manifest(cfg PWAConfig) ([]byte, error) {
	cfg, err := validatePWA(cfg)
	if err != nil {
		return nil, err
	}
	m := manifest{
		Name:            cfg.Name,
		ShortName:       cfg.ShortName,
		StartURL:        cfg.StartURL,
		Scope:           cfg.Scope,
		Display:         cfg.Display,
		ThemeColor:      cfg.ThemeColor,
		BackgroundColor: cfg.BackgroundColor,
		Icons:           cfg.Icons,
	}
	return json.MarshalIndent(m, "", "  ")
}

func ServiceWorker(cfg PWAConfig) ([]byte, error) {
	cfg, err := validatePWA(cfg)
	if err != nil {
		return nil, err
	}
	urls := append([]string(nil), cfg.Precache...)
	if cfg.OfflineURL != "" {
		urls = append(urls, cfg.OfflineURL)
	}
	urls = uniqueSorted(urls)
	precache, _ := json.Marshal(urls)
	offline, _ := json.Marshal(cfg.OfflineURL)
	cacheName, _ := json.Marshal(cfg.CacheName)

	source := fmt.Sprintf(`"use strict";
const CACHE_NAME = %s;
const PRECACHE = %s;
const OFFLINE_URL = %s;
self.addEventListener("install", event => {
  event.waitUntil(caches.open(CACHE_NAME).then(cache => cache.addAll(PRECACHE)));
  self.skipWaiting();
});
self.addEventListener("activate", event => {
  event.waitUntil(self.clients.claim());
});
self.addEventListener("fetch", event => {
  if (event.request.method !== "GET") return;
  const requestURL = new URL(event.request.url);
  if (requestURL.origin !== self.location.origin) return;
  const requestKey = requestURL.pathname + requestURL.search;
  if (PRECACHE.includes(requestKey)) {
    event.respondWith(caches.match(event.request).then(hit => hit || fetch(event.request)));
    return;
  }
  if (event.request.mode === "navigate" && OFFLINE_URL) {
    event.respondWith(fetch(event.request).catch(() => caches.match(OFFLINE_URL)));
  }
});
`, cacheName, precache, offline)
	return []byte(source), nil
}

func validatePWA(cfg PWAConfig) (PWAConfig, error) {
	if strings.TrimSpace(cfg.Name) == "" {
		return cfg, errors.New("PWA name is required")
	}
	if cfg.StartURL == "" {
		cfg.StartURL = "/"
	}
	if cfg.Scope == "" {
		cfg.Scope = "/"
	}
	if cfg.Display == "" {
		cfg.Display = "standalone"
	}
	if cfg.CacheName == "" {
		cfg.CacheName = "go-app-shell-v1"
	}
	if !localPWAURL(cfg.StartURL) || !localPWAURL(cfg.Scope) {
		return cfg, errors.New("PWA start URL and scope must be same-origin rooted URLs")
	}
	allowedDisplay := map[string]bool{"browser": true, "minimal-ui": true, "standalone": true, "fullscreen": true}
	if !allowedDisplay[cfg.Display] {
		return cfg, errors.New("unsupported PWA display mode")
	}
	for _, icon := range cfg.Icons {
		if !localPWAURL(icon.Src) {
			return cfg, errors.New("PWA icon source must be a same-origin rooted URL")
		}
	}
	for _, item := range cfg.Precache {
		if !localPWAURL(item) {
			return cfg, errors.New("PWA precache entries must be same-origin rooted URLs")
		}
	}
	if cfg.OfflineURL != "" && !localPWAURL(cfg.OfflineURL) {
		return cfg, errors.New("PWA offline URL must be a same-origin rooted URL")
	}
	return cfg, nil
}

func localPWAURL(s string) bool {
	if !localRootedPath(s) {
		return false
	}
	u, err := url.Parse(s)
	if err != nil || u.IsAbs() || u.Host != "" {
		return false
	}
	return true
}

func uniqueSorted(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	for _, s := range in {
		seen[s] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for s := range seen {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
