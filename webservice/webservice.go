package webservice

import (
	"net/http"

	"github.com/go-martini/martini"
)

type WebService interface {
	GetPath() string
	WebPost(params martini.Params, w http.ResponseWriter, req *http.Request) (int string)
}

func RegisterWebService(webService WebService, classicMartini *martini.ClassicMartini) {
	path := webService.GetPath()
	classicMartini.Post(path, webService.WebPost)
}
