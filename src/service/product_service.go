package service

import (
	"database/sql"
	"item_service_page/src/repository5"
	"strconv"
)

func Usp_productmanage(prstype string, p_id int, p_name string, p_value int, p_quatity int) *sql.Rows {
	var s = ", "
	var querystr string = "call USP_ProductManage(" + "'" + prstype + "'" + s + strconv.Itoa(p_id) + s + "'" + p_name + "'" + s + strconv.Itoa(p_value) + s + strconv.Itoa(p_quatity) + ")"
	querystr += ";"
	var result = repository5.RepositoryInstance.Procedure_query(querystr)
	return result
}

func Product() {
	repository5.RepositoryInstance.USP_ProductManage = Usp_productmanage
}
