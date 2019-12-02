package models

import (
	"fmt"
)

//``````````````````
type UserInfo struct {
	B_id         int    `db:"b_id"`
	B_loginname  string `db:"b_loginname"`
	B_pwd        string `db:"b_pwd"`
	B_role       int    `db:"b_role"`
	B_type       int    `db:"b_type"`
	B_telephone  string `db:"b_telephone"`
	B_address    string `db:"b_address"`
	B_idCar      string `db:"b_idCar"`
	B_createTime string `db:"b_createTime"`
}

//登录 根据登录名获取用户信息
func AdminGetByName(loginName string) (u UserInfo, err error) {

	userinfo := make([]UserInfo, 0)

	err = Db.Select(&userinfo, "SELECT * FROM `userinfo` WHERE b_loginname=?", loginName)
	if err != nil {
		fmt.Println("exec QueryUserInfoById, ", err)
		return u, err
	}
	if len(userinfo) > 0 {
		return userinfo[0], err
	} else {
		return u, err
	}
}
func UserInfoGetByNameNoId(loginName string, id int) (u UserInfo, err error) {

	userinfo := make([]UserInfo, 0)

	err = Db.Select(&userinfo, "SELECT * FROM `userinfo` WHERE b_loginname=? and b_id !=?", loginName, id)
	if err != nil {
		fmt.Println("exec QueryUserInfoById, ", err)
		return u, err
	}
	if len(userinfo) > 0 {
		return userinfo[0], err
	} else {
		return u, err
	}
}

//根据id获取用户信息
func AdminGetById(id int) (u UserInfo, err error) {

	userinfo := make([]UserInfo, 0)

	err = Db.Select(&userinfo, "SELECT * FROM `userinfo` WHERE b_id=?", id)
	if err != nil {
		fmt.Println("exec QueryUserInfoById, ", err.Error())
	}
	return userinfo[0], err
}

//获取用户列表
func UserInfoList_Get(page int, pageSize int) []UserInfo {

	list := make([]UserInfo, 0)
	sql := "select * from userinfo "
	sql += fmt.Sprintf("LIMIT %d,%d ", ((page - 1) * pageSize), pageSize)
	rows, err := Db.Query(sql)
	for rows.Next() {
		var u UserInfo
		rows.Scan(&u.B_id, &u.B_loginname, &u.B_pwd, &u.B_role, &u.B_type, &u.B_telephone, &u.B_address, &u.B_idCar, &u.B_createTime)
		list = append(list, u)
	}

	if err != nil {
		fmt.Println("exec QuerybookList, ", err)
	}
	return list

}

//获取列表分页
func UserInfoCount_Get() int64 {
	var count int64
	sql := "select count(*) from userinfo "
	rows, err := Db.Query(sql)
	for rows.Next() {
		rows.Scan(&count)
	}
	if err != nil {
		fmt.Println("exec QuerybookList, ", err)
	}
	return count
}

//新建用户信息
func UserInfo_Save(user UserInfo) bool {

	sql := "INSERT INTO userinfo(b_loginname,b_pwd,b_role,b_type,b_telephone,b_address,b_idCar,b_createTime)VALUES(?,?,?,?,?,?,?,?)"
	_, e := Db.Exec(sql, user.B_loginname, user.B_pwd, user.B_role, user.B_type, user.B_telephone, user.B_address, user.B_idCar, user.B_createTime)
	if e == nil {
		return true
	} else {
		return false
	}
}

//更新用户信息
func UserInfo_Edite(user UserInfo) bool {
	sql := "UPDATE userinfo SET b_loginname=?,b_pwd=?,b_role=?,b_type=?,b_telephone=?,b_address=?,b_idCar=?,b_createTime=? WHERE b_id=?"
	_, e := Db.Exec(sql, user.B_loginname, user.B_pwd, user.B_role, user.B_type, user.B_telephone, user.B_address, user.B_idCar, user.B_createTime, user.B_id)
	if e == nil {
		return true
	} else {
		return false
	}
}

//删除用户信息
func UserInfo_Del(id int) bool {

	sql := "DELETE FROM userinfo WHERE b_id=?"
	_, e := Db.Exec(sql, id)
	if e == nil {
		return true
	} else {
		return false
	}
}
