package models

import (
	"fmt"
	"strconv"
)

type PublicUsersModel struct {
	ID                   int
	UserName             string
	Password             string
	Status               int
	CreateTime           string
	HeadPic              string
	Nickname             string
	GUID                 string
	UserID               string
	AuthorityID          int
	telephone            string
	ProvinceID           int
	CityID               int
	Address              string
	Type                 int
	TaskSourceId         int
	SetAppriasePrice     string
	IsWork               int
	UserType             int
	RegionId             int
	TaskToAllArea        int
	ShowPrice            int
	TaskToAffiliatedArea int
	CanEditPassword      int
	UserProductType      int
	CanQueryCarClaims    int
	AppraiseRange        int
	UserMD5              string
	FastDFS              int
	PicParameter         int
	IsOrderCheck         int
	IsPalyCheckOrder     int
	PricePrecision       int
	ShowExpression       int
	ShowProgrammeName    int
	ShowProgrammeTime    string
	AppLogin             int
	PlaceOrder           int
	ExtendedWarranty     int
	OrderList            int
	ExtendedWarrantyList int
	UserIdCS             int
	TokenId              string
	UserProperty         int
	UpdateTime           string
	PreUpdateTime        string
	PreOpUser            string
	ShowPGTime           int
}

func (p *PublicUsersModel)GetPublicUsers() []PublicUsersModel {

	list := make([]PublicUsersModel, 0)

	sql := "SELECT * from PublicUsers where id="+strconv.Itoa(p.ID)

	rows, e := Dbsql.Query(sql)

	for rows.Next() {
		var b PublicUsersModel
		rows.Scan(&b.ID,&b.UserName,&b.Password,&b.Status,&b.CreateTime,&b.HeadPic,&b.Nickname,&b.GUID,&b.UserID,&b.AuthorityID,&b.telephone,&b.ProvinceID,&b.CityID,&b.Address,&b.Type,&b.TaskSourceId,&b.SetAppriasePrice,&b.IsWork,&b.UserType,&b.RegionId,&b.TaskToAllArea,&b.ShowPrice,&b.TaskToAffiliatedArea,&b.CanEditPassword,&b.UserProductType,&b.CanQueryCarClaims,&b.AppraiseRange,&b.UserMD5,&b.FastDFS,&b.PicParameter,&b.IsOrderCheck,&b.IsPalyCheckOrder,&b.PricePrecision,&b.ShowExpression,&b.ShowProgrammeName,&b.ShowProgrammeTime,&b.AppLogin,&b.PlaceOrder,&b.ExtendedWarranty,&b.OrderList,&b.ExtendedWarrantyList,&b.UserIdCS,&b.TokenId,&b.UserProperty,&b.UpdateTime,&b.PreUpdateTime,&b.PreOpUser,&b.ShowPGTime)

		list = append(list,b)
	}
	if e != nil {
		fmt.Println("GetPublicUsers error!")
	}
	return list
}

func (p *PublicUsersModel) DeleteUser()bool  {

	sql := "delete from PublicUsers where id=?"
	_, e := Dbsql.Exec(sql,p.ID)

	if e == nil {
		return true
	} else {
		fmt.Println("删除失败", e.Error())
		return false
	}
}