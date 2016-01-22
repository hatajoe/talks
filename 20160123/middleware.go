package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/zenazn/goji/web/mutil"
)

func main() {
	mux := web.New()
	goji.DefaultMux.Abandon(middleware.Logger)
	mux.Use(Logger)
}

func Logger(c *web.C, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := middleware.GetReqID(*c)
		t1 := time.Now()
		lw := mutil.WrapWriter(w)
		h.ServeHTTP(lw, r)
		t2 := time.Now()
		log.Printf("%s\t%f\t%d\t%s\t%s\t%s", id, t2.Sub(t1).Seconds(), lw.Status, r.Method, r.URL.RequestURI(), r.UserAgent())
	})
}
