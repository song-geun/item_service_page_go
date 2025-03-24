package repository5

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var RepositoryInstance Repository

func repository_open() bool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	db, err := sql.Open("mysql", "song@/song")
	RepositoryInstance.DB = db
	RepositoryInstance.DB.Begin()
	if err != nil {

		return true
	}
	return false
}

func repository_close() bool {
	RepositoryInstance.DB.Close()
	return true
}

func repository_begin() bool {
	RepositoryInstance.DB.Begin()
	return true
}

func procedure_query(str string) *sql.Rows {

	str += ";"
	rows, err := RepositoryInstance.DB.Query(str)
	if err != nil {
		return nil
	}
	return rows
}

/*
	func procedure_exec(str string) sql.Result {
		RepositoryInstance.DB.Begin()
		str += ";"
		var result sql.Result
		result, err := RepositoryInstance.DB.ExecContext(context.Background(), str, sql.Named("str", sql.Out{Dest: &str}))
		if err != nil {
			return nil
		}
		RepositoryInstance.DB.Close()
		return result
	}
*/
func Repository1() {
	RepositoryInstance.Open = repository_open
	RepositoryInstance.Close = repository_close
	RepositoryInstance.Procedure_query = procedure_query
	//RepositoryInstance.Procedure_Exec = procedure_exec
	RepositoryInstance.Begin = repository_begin
	RepositoryInstance.Open()
	RepositoryInstance.Begin()
}
