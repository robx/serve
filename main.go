package main

import (
	"flag"
	"time"
	"log"
	"net/http"
	"os"

	"github.com/dchest/uniuri"
	"github.com/julienschmidt/httprouter"
)

var (
	baseUrl = flag.String("base", "http://localhost:7777/", "base URL")
	listen = flag.String("listen", ":7777", "http listen")
	expire = flag.Duration("expire", 6*time.Hour, "time to keep serving")
	static = flag.String("static", ".", "directory to serve")
)

func main() {
	flag.Parse()

	s := uniuri.New()
	r := httprouter.New()
	r.ServeFiles("/" + s + "/*filepath", http.Dir(*static))
	log.Printf("serving files for %s at %s%s", *expire, *baseUrl, s)
	go func() {
		time.Sleep(*expire)
		os.Exit(0)
	}()
	log.Fatal(http.ListenAndServe(*listen, r))
}
