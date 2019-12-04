package models

import (
	"fmt"
	"strconv"
)

type ProgrammeConfigurationModel struct {
	Id            int
	ConfigureId   int
	ProgrammeId   int
	Usertype      int
	TaskType      int
	RealVehicle   int
	Formalities   int
	SpecialPic    int
	Video         int
	ProgrammeName string
	MainTaskType  int
}


type PicSpecialUserSetModel struct {
	ID int
	SourceID int
	UserID int
	CityID int
	PicSetID int
	IsRequired  int
	OpUserID int
	CreateTime string
}

type PicSpecialSetModel struct {
	ItemID int
	Name string
	ID int
	Path string
	PathAssist string
	ImageDirection int
	PictureUrl string
	Description string
	VideoUrl string
	PhotoTips string
}

type UserProductTypeModel struct {
	ID          int
	UserId      int
	ProductId   int
	ProductName string
	Type        int
	AlterUserId string

	//
}
//读取用户线上产品类型
func (u *UserProductTypeModel) GetUserProductType(userid int) []UserProductTypeModel {
	list := make([]UserProductTypeModel, 0)
	sql := "SELECT ID,UserId,ProductId,ProductName,Type,AlterUserId FROM UserProductType WHERE UserId="+strconv.Itoa(userid) +" and Type=0"

	rows, e := Dbsql.Query(sql)

	for rows.Next() {
		var b UserProductTypeModel
		rows.Scan(&b.ID,&b.UserId,&b.ProductId,&b.ProductName,&b.Type,&b.AlterUserId)
		list = append(list,b)
	}
	if e != nil {
		fmt.Println("GetPublicUsers error!")
	}
	return list

}


//获取配置的 方案
func (p *ProgrammeConfigurationModel) Get(userid string) []ProgrammeConfigurationModel {

	list := make([]ProgrammeConfigurationModel, 0)
	sql := "IF EXISTS(SELECT 1 FROM dbo.ProgrammeConfiguration config WITH(NOLOCK) WHERE config.ConfigureId = "+userid+" AND Usertype = 2) "
	sql+=" SELECT config.Id, config.ConfigureId, config.ProgrammeId, config.Usertype, config.TaskType,config.RealVehicle,config.Formalities,config.SpecialPic,config.Video, CASE config.TaskType WHEN 1 THEN oline.Name ELSE ofline.ProgrammeName END 'ProgrammeName',CASE config.TaskType WHEN 1 THEN oline.TaskType ELSE 3 END 'MainTaskType' FROM dbo.ProgrammeConfiguration config  WITH(NOLOCK) "
	sql+=" LEFT JOIN dbo.OnLineProgramme oline WITH(NOLOCK) ON  config.TaskType =1  AND oline.Id = config.ProgrammeId "
	sql+=" LEFT JOIN dbo.Programme ofline WITH(NOLOCK) ON config.TaskType = 2 AND ofline.Id = config.ProgrammeId "
	sql+=" WHERE (oline.Status = 1 OR oline.Status IS NULL) AND config.ConfigureId = "+userid+" AND config.Usertype = 2  and config.TaskType=1 "
	sql+=" ELSE "
	sql+=" SELECT config.Id, config.ConfigureId, config.ProgrammeId, config.Usertype, config.TaskType,config.RealVehicle,config.Formalities,config.SpecialPic,config.Video, CASE config.TaskType WHEN 1 THEN oline.Name ELSE ofline.ProgrammeName END 'ProgrammeName',CASE config.TaskType WHEN 1 THEN oline.TaskType ELSE 3 END 'MainTaskType' FROM dbo.ProgrammeConfiguration config  WITH(NOLOCK) "
	sql+=" LEFT JOIN dbo.OnLineProgramme oline WITH(NOLOCK) ON  config.TaskType =1  AND oline.Id = config.ProgrammeId "
	sql+=" LEFT JOIN dbo.Programme ofline WITH(NOLOCK) ON config.TaskType = 2 AND ofline.Id = config.ProgrammeId "
	sql+=" WHERE (oline.Status = 1 OR oline.Status IS NULL) AND config.ConfigureId = (SELECT TaskSourceId FROM dbo.PublicUsers WITH(NOLOCK) WHERE ID = "+userid+") AND config.Usertype =1  and config.TaskType=1 "

	rows, e := Dbsql.Query(sql)

	for rows.Next() {
		var b ProgrammeConfigurationModel
		rows.Scan(&b.Id,&b.ConfigureId,&b.ProgrammeId,&b.Usertype,&b.TaskType,&b.RealVehicle,&b.Formalities,&b.SpecialPic,&b.Video,&b.ProgrammeName,&b.MainTaskType)
		list = append(list,b)
	}
	if e != nil {
		fmt.Println("GetPublicUsers error!")
	}
	return list

}

func getPicSpecialUserSet(userid string)[] PicSpecialUserSetModel  {

	list := make([]PicSpecialUserSetModel, 0)
	sql := "SELECT  pssu.ID,pssu.SourceID,pssu.UserID,pssu.CityID,pssu.PicSetID,pssu.IsRequired,pssu.OpUserID,pssu.CreateTime "
	sql += " FROM PicSpecialUserSet as pssu with(nolock) inner join PicSet as ps with(nolock) on pssu.PicSetID = ps.ID "
	sql += " WHERE pssu.UserID="+userid+"  and pssu.Enabled=1 and ps.Enabled=1 and pssu.IsRequired=1 "

	rows, e := Dbsql.Query(sql)

	for rows.Next() {
		var b PicSpecialUserSetModel
		rows.Scan(&b.ID,&b.SourceID,&b.UserID,&b.CityID,&b.PicSetID,&b.IsRequired,&b.OpUserID,&b.CreateTime)
		list = append(list,b)
	}
	if e != nil {
		fmt.Println("(p *PicSpecialUserSetModel)get error!")
	}
	return list
}
//获取配置的特殊粘片
func (p *PicSpecialSetModel)GetPicSpecialSet(userid string)[]PicSpecialSetModel  {

	list := make([]PicSpecialSetModel, 0)

	plist:=getPicSpecialUserSet(userid)
	if len(plist)==0 {
		return list
	}
	id:=strconv.Itoa(plist[0].PicSetID)
	sql := "SELECT p.ID,r.ItemID,p.Name,p.Path,p.PathAssist,p.ImageDirection,p.PictureUrl,p.Description,p.VideoUrl,p.PhotoTips "
	sql +=" FROM PicSet ps with(nolock) "
	sql +=" inner join PicSpecialSet r  with(nolock) on ps.id = r.PicSetID "
	sql +=" inner join PicSpecial p  with(nolock) on r.ItemID=p.ItemID "
	sql +=" where ps.id="+id+" and ps.Enabled=1 and p.Enabled=1   ORDER BY ps.Id,r.OrderNum"

	rows, e := Dbsql.Query(sql)

	for rows.Next() {
		var b PicSpecialSetModel
		rows.Scan(&b.ID,&b.ItemID,&b.Name,&b.Path,&b.PathAssist,&b.ImageDirection,&b.PictureUrl,&b.Description,&b.VideoUrl,&b.PhotoTips)
		list = append(list,b)
	}
	if e != nil {
		fmt.Println("(p *PicSpecialUserSetModel)get error!")
	}
	return list
}

