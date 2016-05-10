package phalgo

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

func OpenPprof() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()
}