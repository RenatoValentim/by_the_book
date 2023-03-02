package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type authHanler struct {
	next http.Handler
}

func (h *authHanler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.next.ServeHTTP(w, r)
}

// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	rawSegs := strings.Split(r.URL.Path, "/")
	segs := []string{}
	for _, s := range rawSegs {
		if s != "" {
			segs = append(segs, s)
		}
	}

	if len(segs) != 3 {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	action := segs[1]
	provider := segs[2]

	supportedProviders := strings.Contains(provider, "facebook") || strings.Contains(provider, "github") || strings.Contains(provider, "google")

	if !supportedProviders {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHanler{next: handler}
}
