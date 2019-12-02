package models

type Menu struct {
	Id         int    `db:"b_id"`
	Parent     int    `db:"b_parent"`
	Name       string `db:"b_name"`
	Url        string `db:"b_url"`
	Createtime string `db:"b_createtime"`
	Type       int    `db:"b_type"`
	Icon       string `db:"b_icon"`
}

func MenuList() (list []Menu, err error) {

	err = Db.Select(&list, "SELECT * FROM menu")

	return list, err

}
