package models

import "fmt"

type BookInfo struct {
	Id               int     `db:"b_id"`
	BookName         string  `db:"b_bookName"`
	PublishedTime    string  `db:"b_publishedTime"`
	Author           string  `db:"b_author"`
	ClassificationId int     `db:"b_classificationId"`
	Money            float64 `db:"b_money"`
	CreateTime       string  `db:"b_createTime"`
}

func (b BookInfo) GetBookList(page, pagesize int) []BookInfo {

	list := make([]BookInfo, 0)

	sql := "SELECT * FROM bookinfo "
	sql += fmt.Sprintf(" limit %d,%d", (page-1)*pagesize, pagesize)

	rows, e := Db.Query(sql)

	for rows.Next() {
		var book BookInfo
		rows.Scan(&book.Id, &book.BookName, &book.PublishedTime, &book.Author, &book.ClassificationId, &book.Money, &book.CreateTime)
		list = append(list, book)
	}

	if e != nil {
		fmt.Println("GetBookList error!")
	}
	return list
}

func (b BookInfo) GetBookCount() (count int64) {
	sql := "select count(*)as 'count' from bookinfo"
	rows, e := Db.Query(sql)
	for rows.Next() {
		rows.Scan(&count)
	}
	if e != nil {
		fmt.Sprintln("GetBookCount error")
	}
	return count
}

func (b BookInfo) Save() bool {

	sql := "INSERT INTO bookinfo(b_bookName,b_publishedTime,b_author,b_classificationId,b_money,b_createTime)VALUES(?,?,?,?,?,?)"
	_, e := Db.Exec(sql, b.BookName, b.PublishedTime, b.Author, b.ClassificationId, b.Money, b.CreateTime)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}
}

func (b BookInfo) Update() bool {

	sql := "UPDATE bookinfo SET b_bookName=?,b_publishedTime=?,b_author=?,b_classificationId=?,b_money=?,b_createTime=? WHERE b_id=?"
	_, e := Db.Exec(sql, b.BookName, b.PublishedTime, b.Author, b.ClassificationId, b.Money, b.CreateTime, b.Id)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}
}

func (b *BookInfo) GetbookById() (book BookInfo, er error) {

	sql := "SELECT * FROM  bookinfo WHERE b_id=?"
	list := make([]BookInfo, 0)
	er = Db.Select(&list, sql, b.Id)
	if er != nil {
		fmt.Println("GetbookById error", er)
	}
	if len(list) > 0 {
		book = list[0]
	}
	return book, er
}

func (b *BookInfo) GetBookByNameNoId() (book BookInfo, er error) {
	sql := "SELECT * FROM  bookinfo WHERE b_bookName=? and b_id!=?"
	list := make([]BookInfo, 0)
	er = Db.Select(&list, sql, b.BookName, b.Id)
	if er != nil {
		fmt.Println("GetbookById error", er)
	}
	if len(list) > 0 {
		book = list[0]
	}
	return book, er
}

func (b *BookInfo) Delete() bool {

	sql := "delete from bookinfo WHERE b_id=?"
	_, e := Db.Exec(sql, b.Id)

	if e == nil {
		return true
	} else {
		fmt.Println("保存失败", e.Error())
		return false
	}

}
