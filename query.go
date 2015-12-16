package dbtest

import (
	"fmt"
	"errors"
	"database/sql"
)

func OutputQueryResult(rows *sql.Rows) (err error) {
	
	if rows == nil {
		return errors.New("rows is nil")
	}
	
	//get rows columns
	columns, err := rows.Columns()
	if err != nil {
		return errors.New("get rows columns error")
	}
	
	//make a slice for values
	values := make ([]sql.RawBytes, len(columns))
	scanArgs := make ([]interface{}, len(values))
	
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	for _, col := range columns {
		fmt.Printf("%s\t", col)
	}
	fmt.Println("")
	
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return errors.New("scan error")
		}
		
		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"				
			} else {
				value = string(col)
			}
			fmt.Printf("%s\t", value)
		}
		fmt.Println("")
	}
	
	if err = rows.Err(); err != nil {
		return err
	}
	
	return nil
}

func SkipQueryResult(rows *sql.Rows) (err error) {
	return nil
}

type CallBack func(rows *sql.Rows) (err error)

func SimpleQuery(sql string, testRetFunc CallBack) (err error) {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	
	defer rows.Close()
	return testRetFunc(rows)
}

func PreparedQuery(sql string, testRetFunc CallBack) (err error) {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	
	defer rows.Close()
	
	return testRetFunc(rows)
}

func TxSimpleQuery(sql string, testRetFunc CallBack) (err error) {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	
	defer db.Close()
	
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	
	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	
	defer rows.Close()
	
	err = testRetFunc(rows)
	if err != nil {
		return err
	}

	return tx.Commit()
}