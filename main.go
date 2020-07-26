package main

import (
	"github.com/common-phrases/model"
	"github.com/common-phrases/webservice"
	"github.com/go-martini/martini"
)

func main() {
	martiniClassic := martini.Classic()
	newText := &model.TextData{}
	webservice.RegisterWebService(newText, martiniClassic)
	martiniClassic.Run()
}
