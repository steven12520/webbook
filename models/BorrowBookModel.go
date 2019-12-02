package models

import "fmt"

type BorrowBook struct {
	Id         int    `db:"b_id"`
	BookId     int    `db:"b_bookId"`
	UserId     int    `db:"b_userId"`
	CreateTime string `db:"b_createTime"`
	ToAlsoTime string `db:"b_toAlsoTime"`
	AlsoTime   string `db:"b_alsoTime"`
	Type       int    `db:"b_type"`
	BookName   string `db:"b_bookName"`
	Author     string `db:"b_author"`
	Loginname  string `db:"b_loginname"`
}

func (br BorrowBook) GetBorrowBookList(page, pagesize int) []BorrowBook {

	sql := "SELECT *FROM borrowbook LIMIT ?,?"

	rows, e := Db.Query(sql, (page-1)*pagesize, pagesize)
	if e != nil {
		fmt.Println("GetBorrowBookList error", e.Error())
	}
	list := make([]BorrowBook, 0)
	for rows.Next() {
		var br BorrowBook
		//b_bookId,b_userId,b_createTime,b_toAlsoTime,b_alsoTime,b_type,b_bookName,b_author,b_loginname
		rows.Scan(&br.Id, &br.BookId, &br.UserId, &br.CreateTime, &br.ToAlsoTime, &br.AlsoTime, &br.Type, &br.BookName, &br.Author, &br.Loginname)
		list = append(list, br)
	}
	return list
}

func (br BorrowBook) GetBorrowBookCount() (count int64) {
	sql := "SELECT COUNT(*)as count FROM borrowbook"
	rows, e := Db.Query(sql)
	if e != nil {
		fmt.Println("GetBorrowBookCount error ", e.Error())
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return
}
func (br BorrowBook) Add() bool {

	sql := "INSERT INTO borrowbook(b_bookId,b_userId,b_createTime,b_toAlsoTime,b_alsoTime,b_type,b_bookName,b_author,b_loginname)VALUES(?,?,?,?,?,?,?,?,?)"

	_, e := Db.Exec(sql, br.BookId, br.UserId, br.CreateTime, br.ToAlsoTime, br.AlsoTime, br.Type, br.BookName, br.Author, br.Loginname)
	if e != nil {
		fmt.Println(" BorrowBook add error", e.Error())
		return false
	} else {
		return true
	}
}

func (br BorrowBook) Update() bool {
	sql := "UPDATE borrowbook SET b_bookId=?,b_userId=?,b_createTime=?,b_toAlsoTime=?,b_alsoTime=?,b_type=?,b_bookName=?,b_author=?,b_loginname=? WHERE b_id=?"

	_, e := Db.Exec(sql, br.BookId, br.UserId, br.CreateTime, br.ToAlsoTime, br.AlsoTime, br.Type, br.BookName, br.Author, br.Loginname, br.Id)
	if e != nil {
		fmt.Println(" BorrowBook Update error", e.Error())
		return false
	} else {
		return true
	}
}
func (br BorrowBook) GetBorrowBookById() (book BorrowBook, er error) {
	sql := "SELECT * FROM borrowbook WHERE b_id=?"
	list := make([]BorrowBook, 0)
	er = Db.Select(&list, sql, br.Id)
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
func (b *BorrowBook) Delete() bool {

	sql := "delete from borrowbook WHERE b_id=?"
	_, e := Db.Exec(sql, b.Id)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}

}
