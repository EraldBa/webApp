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

func NewRendererConf(a *config.AppConfig) {
	app = a
}

func AddDefaultData(r *http.Request, tmplData *models.TemplateData) {
	tmplData.CSRFToken = nosurf.Token(r)
	tmplData.Success = app.Session.PopString(r.Context(), "success")
	tmplData.Error = app.Session.PopString(r.Context(), "error")
	tmplData.Flash = app.Session.PopString(r.Context(), "flash")
	tmplData.IsAuthenticated = app.Session.Exists(r.Context(), "user_id")
}

func Template(w http.ResponseWriter, r *http.Request, tmplName string, tmplData *models.TemplateData) {
	var tmplCache map[string]*template.Template
	var err error

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
	tmpl, ok := tmplCache[tmplName]
	if !ok {
		app.ErrorLog.Fatal("Templates not accessible")
	}

	buf := new(bytes.Buffer)

	AddDefaultData(r, tmplData)
	// rendering template
	if err = tmpl.Execute(buf, tmplData); err != nil {
		app.ErrorLog.Println("Problem with rendering template:", err)
	}

	if _, err = buf.WriteTo(w); err != nil {
		app.ErrorLog.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := make(map[string]*template.Template)
	// Get all names of page templates from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		// Get only base name of template (/templates/home.page.gohtml -> home.page.gohtml)
		tmplName := filepath.Base(page)
		tmpl, err := template.New(tmplName).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}
		// Parsing all layout files to tmpl
		tmpl, err = tmpl.ParseGlob("./templates/*.layout.gohtml")
		if err != nil {
			return tmplCache, err
		}

		tmplCache[tmplName] = tmpl
	}
	return tmplCache, nil
}
