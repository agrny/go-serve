package router

import (
	"net/http"

	"go-server/models"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	rInfo := models.RequesterInfo{}
	rInfo.ParseRequest(r)

	w.Write([]byte("<html>Hello</html>"))
}

