package router

import (
	// "encoding/json"
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
	cpuInfo := cpu.JSONString(false)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(cpuInfo))
}
func GPUInfo(w http.ResponseWriter, r *http.Request) {
	gpu, err := ghw.GPU()
	if err != nil {
		http.Error(w, "could not get cpu info", http.StatusInternalServerError)
		return
	}
	gpuInfo := gpu.JSONString(false)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(gpuInfo))
}
