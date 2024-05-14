package database

import (
	"database/sql"
)

func ExecuteQuery(DB *sql.DB,query string, params []interface{})(*sql.Rows, error){

	rows, err := DB.Query(query, params...)
    if err != nil {
		return nil, err
    }
	return rows, nil

}