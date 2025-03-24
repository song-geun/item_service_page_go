package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"item_service_page/src/dto"
	"net/http"
)

func Controller(w http.ResponseWriter, req *http.Request) {
	var products = dto.FindAllD()
	data, _ := json.Marshal(products)
	buff := bytes.NewBuffer(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buff)
}
