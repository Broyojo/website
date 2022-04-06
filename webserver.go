package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/caddyserver/certmagic"
	"golang.org/x/net/publicsuffix"
)

func main() {
	whitelist := []string{"broyojo.com", "www.broyojo.com"}
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = "dha@xoba.com"
	log.Println("Starting Webserver...")
	log.Fatal(certmagic.HTTPS(whitelist, http.HandlerFunc(serve)))
}

func serve(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	setHeaders(w)
	log.Printf("%s %s %s %s", r.RemoteAddr, r.Host, r.RequestURI, r.Header.Get("user-agent"))
	tld, err := normalizeHost(r.Host)

	if err != nil {
		http.Error(w, "can't parse host", http.StatusBadRequest)
	}

	file := r.URL.Path[1:]

	if file == "" {
		file = "index.html"
	}

	http.ServeFile(w, r, filepath.Clean(tld+"/"+file))
}

func setHeaders(w http.ResponseWriter) {
	r := func(k, v string) {
		w.Header().Set(k, v)
	}
	r("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0, private")
	r("Expires", "0")
	r("Pragma", "no-cache")
	r("Strict-Transport-Security", "max-age=86400; includeSubDomains")
	r("Vary", "*")
	r("X-Content-Type-Options", "nosniff")
	r("X-Frame-Options", "DENY")
	r("X-Xss-Protection", "1; mode=block")
}

func normalizeHost(host string) (string, error) {
	tld, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return "", err
	}
	return strings.ToLower(tld), nil
}
