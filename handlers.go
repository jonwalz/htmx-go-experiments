package main

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"

	"github.com/user/project/templates"
	"github.com/user/project/templates/pages"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	scriptTags := pages.ScriptTags()

	// Define template body content.
	bodyContent := pages.BodyContent(
		"Welcome to example!",                // define h1 text
		"You're here because it worked out.", // define p text
	)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Welcome to example!", // define title text
		metaTags,              // define meta tags
		bodyContent,           // define body content
		scriptTags,
	)

	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, indexTemplate); err != nil {
		// If not, return HTTP 400 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}


// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, return HTTP 400 error.
		slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Write HTML content.
	w.Write([]byte("You clicked the button!"))

	// Send htmx response.
	htmx.NewResponse().Write(w)

	// Send log message.
	slog.Info("request API", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

func clickedHandler(w http.ResponseWriter, r *http.Request) {
	// Write HTML content
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<p>clicked</p>"))

	// Send log message
	slog.Info("clicked handler", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}