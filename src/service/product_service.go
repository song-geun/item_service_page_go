package service

import (
	"database/sql"
	"item_service_page/src/repository5"
	"strconv"
)

func Usp_productmanage(prstype string, p_id int64, p_name string, p_value int64, p_quatity int64) *sql.Rows {
	var s = ", "
	var querystr string = "call USP_ProductManage(" + "'" + prstype + "'" + s + strconv.FormatInt(p_id, 10) + s + "'" + p_name + "'" + s + strconv.FormatInt(p_value, 10) + s + strconv.FormatInt(p_quatity, 10) + ")"
	querystr += ";"
	var result = repository5.RepositoryInstance.Procedure_query(querystr)
	return result
}

func Product() {
	repository5.RepositoryInstance.USP_ProductManage = Usp_productmanage
}
