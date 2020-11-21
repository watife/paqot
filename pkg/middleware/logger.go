package middleware

import (
	"log"
	"net/http"
	"os"
)

func LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime).Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
