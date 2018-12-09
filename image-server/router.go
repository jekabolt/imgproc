package server

import (
	"net/http"
	"time"

	log "github.com/InVisionApp/go-logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kardianos/osext"
)

type ProcRouter struct {
	HTTPPort       string
	ImageDimension int
}

var folderPath, _ = osext.ExecutableFolder()
var logger = log.NewSimple().WithFields(log.Fields{"(" + folderPath: ")"})

func (pr ProcRouter) InitRouter() {
	r := chi.NewRouter()

	// Init middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		//TODO: server status
		w.Write([]byte("ok"))
	})

	//TODO: rate limit by ip
	r.Route("/dot", func(r chi.Router) {
		r.Post("/", pr.handleDotImage)
	})

	go http.ListenAndServe(pr.HTTPPort, r)

}
