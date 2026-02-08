package router

import (
	"net/http"

	"go-server/models"

	"github.com/jaypipes/ghw"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	rInfo := models.RequesterInfo{}
	rInfo.ParseRequest(r)

	w.Write([]byte("<html>Hello</html>"))
}

func CPUInfo(w http.ResponseWriter, r *http.Request) {
	cpu, err := ghw.CPU()
	if err != nil {
		http.Error(w, "could not get cpu info", http.StatusInternalServerError)
		return
	}
	cpuJSON := cpu.JSONString(true)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(cpuJSON))
}
