package main

import (
	"fmt"
	"net/http"
	"strings"

	"go-server/logger"
	"go-server/router"
)

func main() {
	port := ":2000"

	fileServer := http.FileServer(http.Dir("./file"))
	logger := logger.NewLogHandler()
	defer logger.Close()

	server := router.NewHTTPServer("0.0.0.0")
	server.Handle("/file/", http.StripPrefix("/file", fileServer))
	server.HandleFunc("/hwinfo/cpu", router.CPUInfo)
	server.HandleFunc("/hwinfo/gpu", router.GPUInfo)
	server.HandleFunc("/", router.RootHandler)

	fmt.Printf("Listening on port %s....\n", strings.Split(port, ":")[1])
	logger.Info("server starting", "port", port)

	err := http.ListenAndServe(port, server)
	if err != nil {
		println(err.Error())
		return
	}
}
