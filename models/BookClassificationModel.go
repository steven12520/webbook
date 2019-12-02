package models

import "fmt"

type BookClassification struct {
	Id         int    `db:"b_id"`
	Name       string `db:"b_name"`
	CreateTime string `db:"b_createTime"`
	Type       int    `db:"b_type"`
}

func (b BookClassification) GetBookClassificationList(page, pagesize int) []BookClassification {

	sql := "SELECT *FROM bookclassification LIMIT ?,?"

	rows, e := Db.Query(sql, (page-1)*pagesize, pagesize)
	if e != nil {
		fmt.Println("GetBorrowBookList error", e.Error())
	}
	list := make([]BookClassification, 0)
	for rows.Next() {
		var br BookClassification
		rows.Scan(&br.Id, &br.Name, &br.CreateTime, &br.Type)
		list = append(list, br)
	}
	return list
}
func (b BookClassification) GetBookClassificationCount() (count int64) {
	sql := "SELECT COUNT(*)as count FROM bookclassification"
	rows, e := Db.Query(sql)
	if e != nil {
		fmt.Println("GetBorrowBookCount error ", e.Error())
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return
}

func (b BookClassification) GetBookClassificationById() (book BookClassification, er error) {
	sql := "SELECT * FROM bookclassification WHERE b_id=?"
	list := make([]BookClassification, 0)
	er = Db.Select(&list, sql, b.Id)
	if er != nil {
		fmt.Println("GetBorrowBookCount error ", er.Error())
	}
	if len(list) > 0 {
		book = list[0]
	} else {
		book.Id = 0
	}
	return
}

func (b BookClassification) Save() bool {

	sql := "INSERT INTO bookclassification(b_name,b_createTime,b_type)VALUES(?,?,?)"
	_, e := Db.Exec(sql, b.Name, b.CreateTime, b.Type)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}
}

func (b BookClassification) Update() bool {

	sql := "UPDATE bookclassification SET b_name=?,b_createTime=?,b_type=? WHERE b_id=?"
	_, e := Db.Exec(sql, b.Name, b.CreateTime, b.Type, b.Id)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}
}

func (br BookClassification) Delete() bool {
	sql := "DELETE FROM bookclassification WHERE b_id=?"
	_, e := Db.Exec(sql, br.Id)

	if e != nil {
		fmt.Println(" BorrowBook Delete error", e.Error())
		return false
	} else {
		return true
	}
}
