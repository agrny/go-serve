package main

import (
	"fmt"
	"net/http"
	"strings"

	"go-server/router"
)

func main() {
	port := ":2000"

	fileServer := http.FileServer(http.Dir("./file"))

	HTTPRouter := router.NewHTTPServer("0.0.0.0")
	HTTPRouter.Handle("/file/", http.StripPrefix("/file", fileServer))
	HTTPRouter.HandleFunc("/hwinfo/cpu", router.CPUInfo)
	HTTPRouter.HandleFunc("/hwinfo/gpu", router.GPUInfo)
	HTTPRouter.HandleFunc("/", router.RootHandler)

	fmt.Printf("Listening on port %s....\n", strings.Split(port, ":")[1])
	err := http.ListenAndServe(port, HTTPRouter)
	if err != nil {
		println(err.Error())
		return
	}
}
