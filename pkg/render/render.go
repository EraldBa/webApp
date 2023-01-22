package render

import (
	"bytes"
	"github.com/EraldBa/webApp/pkg/config"
	"github.com/EraldBa/webApp/pkg/models"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(r *http.Request, td *models.TemplateData) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	td.Success = app.Session.PopString(r.Context(), "success")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}

	return td
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var err error
	var tmplCache map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tmplCache = app.TemplateCache
	} else {
		tmplCache, err = CreateTemplateCache()
		if err != nil {
			app.ErrorLog.Fatal(err)
		}
	}

	// get requested template from cache
	t, ok := tmplCache[tmpl]
	if !ok {
		app.ErrorLog.Fatal("Templates not accessible")
	}
	// for extra security
	buf := new(bytes.Buffer)

	td = AddDefaultData(r, td)
	// rendering template
	_ = t.Execute(buf, td)

	if _, err := buf.WriteTo(w); err != nil {
		app.ErrorLog.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}
	// Get all names of page templates from ./templates
	pages, err := filepath.Glob("./templates/*page.gohtml")
	if err != nil {
		return tc, err
	}
	// Get only base name of template (/templates/home.page.gohtml -> home.page.gohtml)
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}
		// Get all names of layout templates from ./templates
		matches, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return tc, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return tc, err
			}
		}
		tc[name] = ts
	}
	return tc, nil
}
