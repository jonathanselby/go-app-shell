package shell

import (
	"bytes"
	"embed"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

const Version = "0.1.0"

//go:embed templates/*.html assets/*
var embedded embed.FS

type Config struct {
	AssetBase             string
	Language              string
	AdditionalStylesheets []string
	ThemeColorLight       string
	ThemeColorDark        string
}

type Renderer struct {
	config Config
	tmpl   *template.Template
}

type AppHeader struct {
	Name      string
	Context   string
	Status    string
	UserLabel string
}

type NavGroup struct {
	Label string
	Items []NavItem
}

type NavItem struct {
	Label   string
	URL     string
	Current bool
}

type WorkspaceMode string

const (
	WorkspaceSingle     WorkspaceMode = "single"
	WorkspaceListDetail WorkspaceMode = "list-detail"
)

type CompactPane string

const (
	CompactList   CompactPane = "list"
	CompactDetail CompactPane = "detail"
)

type Workspace struct {
	Mode        WorkspaceMode
	Content     template.HTML
	List        template.HTML
	Detail      template.HTML
	CompactPane CompactPane
}

type PWAHead struct {
	ManifestURL      string
	RegistrationURL  string
	ServiceWorkerURL string
}

type Page struct {
	DocumentTitle string
	App           AppHeader
	Navigation    []NavGroup
	Title         string
	Summary       string
	Actions       template.HTML
	Workspace     Workspace
	PWA           *PWAHead
}

type renderData struct {
	Config Config
	Page   Page
}

func New(cfg Config) (*Renderer, error) {
	if cfg.AssetBase == "" {
		cfg.AssetBase = "/_app-shell"
	}
	if !localRootedPath(cfg.AssetBase) {
		return nil, errors.New("asset base must be a same-origin rooted path")
	}
	cfg.AssetBase = strings.TrimSuffix(cfg.AssetBase, "/")
	if cfg.Language == "" {
		cfg.Language = "en"
	}
	if cfg.ThemeColorLight == "" {
		cfg.ThemeColorLight = "#ffffff"
	}
	if cfg.ThemeColorDark == "" {
		cfg.ThemeColorDark = "#0d1117"
	}
	for _, u := range cfg.AdditionalStylesheets {
		if strings.TrimSpace(u) == "" {
			return nil, errors.New("additional stylesheet URL must not be empty")
		}
	}

	t, err := template.ParseFS(embedded, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &Renderer{config: cfg, tmpl: t}, nil
}

func (r *Renderer) Render(w io.Writer, page Page) error {
	if r == nil || r.tmpl == nil {
		return errors.New("nil shell renderer")
	}
	if strings.TrimSpace(page.Title) == "" {
		return errors.New("page title is required")
	}
	if page.DocumentTitle == "" {
		page.DocumentTitle = page.Title
	}
	if page.Workspace.Mode == "" {
		page.Workspace.Mode = WorkspaceSingle
	}
	switch page.Workspace.Mode {
	case WorkspaceSingle:
	case WorkspaceListDetail:
		if page.Workspace.CompactPane == "" {
			page.Workspace.CompactPane = CompactList
		}
		if page.Workspace.CompactPane != CompactList && page.Workspace.CompactPane != CompactDetail {
			return errors.New("list-detail compact pane must be list or detail")
		}
	default:
		return errors.New("unsupported workspace mode")
	}
	var buf bytes.Buffer
	if err := r.tmpl.ExecuteTemplate(&buf, "shell.html", renderData{Config: r.config, Page: page}); err != nil {
		return err
	}
	_, err := buf.WriteTo(w)
	return err
}

func AssetsFS() (fs.FS, error) {
	return fs.Sub(embedded, "assets")
}

func localRootedPath(s string) bool {
	return strings.HasPrefix(s, "/") && !strings.HasPrefix(s, "//") && !strings.Contains(s, "\\") && !strings.ContainsAny(s, "\r\n\x00")
}
