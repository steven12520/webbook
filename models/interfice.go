package models

import "fmt"

type Interfice struct {
	Id         int
	Iname      string
	Txt        string
	Ranges 		int
	Createtime string
	Orders int
}


func (br Interfice) GetInterficeList() []Interfice {

	sql := "SELECT `id`,`iname`,`txt`,`ranges`,`createtime` FROM interfice order by orders,id"

	rows, e := Db.Query(sql)
	if e != nil {
		fmt.Println("GetInterficeList error", e.Error())
	}
	list := make([]Interfice, 0)
	for rows.Next() {
		var br Interfice

		rows.Scan(&br.Id, &br.Iname, &br.Txt,&br.Ranges, &br.Createtime)
		list = append(list, br)
	}
	return list
}

func (br Interfice) Add() bool {

	sql := "INSERT INTO interfice(Iname,Txt,Ranges)VALUES(?,?,?)"

	_, e := Db.Exec(sql, br.Iname, br.Txt,br.Ranges)
	if e != nil {
		fmt.Println(" interfice add error", e.Error())
		return false
	} else {
		return true
	}
}

func (br Interfice) Update() bool {
	sql := "UPDATE interfice SET Iname=?,Txt=?,Ranges=? WHERE Id=?"

	_, e := Db.Exec(sql, br.Iname, br.Txt,br.Ranges, br.Id)
	if e != nil {
		fmt.Println(" interfice Update error", e.Error())
		return false
	} else {
		return true
	}
}

func (b *Interfice) Delete() bool {

	sql := "delete from interfice WHERE id=?"
	_, e := Db.Exec(sql, b.Id)

	if e == nil {
		return true
	} else {
		fmt.Println("删除失败", e.Error())
		return false
	}

}
func (br Interfice) GetInterficeById() (inmo Interfice, er error) {
	sql := "SELECT * FROM interfice WHERE id=?"
	list := make([]Interfice, 0)
	er = Db.Select(&list, sql, br.Id)
	if er != nil {
		fmt.Println("GetBorrowBookCount error ", er.Error())
	}
	if len(list) > 0 {
		inmo = list[0]
	} else {
		inmo.Id = 0
	}
	return
}