package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jazzopaul/Hotel-Bookings-GoLang/pkg/config"
	"github.com/jazzopaul/Hotel-Bookings-GoLang/pkg/handlers"
	"github.com/jazzopaul/Hotel-Bookings-GoLang/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8081"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
