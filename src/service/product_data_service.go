package service

import (
	_ "container/list"
	"database/sql"
	"item_service_page/src/repository5"
	"strconv"
)

func Usp_productdatamanage(prstype string, Pprod_data_id int64, p_id int64, p_name string, p_value int64, p_quatity int64, p_DATE string) *sql.Rows {
	const s = ", "
	var querystr string = "call USP_ProductDataManage(" + prstype + s + strconv.FormatInt(Pprod_data_id, 10) + s + strconv.FormatInt(p_id, 10) + s + p_name + s + strconv.FormatInt(p_value, 10) + s + strconv.FormatInt(p_quatity, 10) + s + p_DATE + ")"
	querystr += ";"
	var result *sql.Rows = repository5.RepositoryInstance.Procedure_query(querystr)
	return result
}

func Usp_productdatamanage2(prstype string, Pprod_data_id int64, p_id int64, p_name string, p_value int64, p_quatity int64, p_DATE1 string, p_DATE2 string) *sql.Rows {
	var s = ", "
	var querystr string = "call USP_ProductDataManage(" + prstype + s + strconv.FormatInt(Pprod_data_id, 10) + s + strconv.FormatInt(p_id, 10) + s + p_name + s + strconv.FormatInt(p_value, 10) + s + strconv.FormatInt(p_quatity, 10) + s + p_DATE1 + s + p_DATE2 + ")"
	querystr += ";"
	var result *sql.Rows = repository5.RepositoryInstance.Procedure_query(querystr)
	return result
}

func Productdata() {
	repository5.RepositoryInstance.USP_ProductDataManage = Usp_productdatamanage
	repository5.RepositoryInstance.USP_ProductDataManage2 = Usp_productdatamanage2
}
