package text

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
)

func (t *TextData) GetPath() string {
	return "/text"
}

func (t *TextData) WebPost(params martini.Params, w http.ResponseWriter, req *http.Request) (int, string) {
	defer req.Body.Close()

	requestBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return 500, "internal error"
	}

	if len(params) != 0 {
		return 405, "method not allowed"
	}

	var text TextData
	err = json.Unmarshal(requestBody, &text)

	if err != nil {
		return 400, "invalid JSON data"
	}

	text.GetCommonPhrases()
	fmt.Println("webservice")

	w.Header().Set("content-type", "application/JSON")
	json.NewEncoder(w).Encode(text)

	return 200, "success"
}
