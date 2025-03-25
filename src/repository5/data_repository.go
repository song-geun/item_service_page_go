package repository5

import (
	"database/sql"
)

type Repository struct {
	DB    *sql.DB
	Open  func() bool
	Close func() bool
	Begin func() bool

	Procedure_query func(str string) *sql.Rows
	//Procedure_Exec  func(string) sql.Result

	USP_ProductDataManage  func(prstype string, Pprod_data_id int64, p_id int64, p_name string, p_value int64, p_quatity int64, p_DATE string) *sql.Rows
	USP_ProductDataManage2 func(prstype string, Pprod_data_id int64, p_id int64, p_name string, p_value int64, p_quatity int64, p_DATE1 string, p_DATE2 string) *sql.Rows
	USP_ProductManage      func(prstype string, p_id int, p_name string, p_value int, p_quatity int) *sql.Rows
}
