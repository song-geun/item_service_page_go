package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"item_service_page/src/dto"
	"item_service_page/src/repository5"
	"net/http"
)

func FinDAll(w http.ResponseWriter, req *http.Request) {

	var products = dto.ProductDto(repository5.RepositoryInstance.USP_ProductManage("S1", 0, " ", 0, 0))

	data, _ := json.Marshal(products)
	buff := bytes.NewBuffer(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, buff)
}

func Insert(w http.ResponseWriter, req *http.Request) {

	/***********************************************************/
	scanner := bufio.NewScanner(req.Body)
	var str string = ""
	for i := 0; scanner.Scan() && i < 5; i++ {
		str += scanner.Text()
	}
	dto.ProductInsert(str, 1, len(str)-1)
	//repository5.RepositoryInstance.USP_ProductManage("I1",)
	//
	//
	/***********************************************************/
}
