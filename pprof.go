package phalgo

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

func OpenPprof(port string) {
	go func() {
		log.Println(http.ListenAndServe("localhost:" + port, nil))
	}()
}