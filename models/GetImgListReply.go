package models

type ResultPublicDate struct {
	Msg string	`json:"msg"`
	Status int`json:"status"`
}

//图片列表
type ResultDate struct {
	Data GetImgListReply
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
//拉取自己的订单
type PretrialPush struct {
	Data PretrialOrderInfoAll
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
type AddFJ struct {
	ReturnReason string
	TitleText  string
	SampleImg string
	FileName string
	PicID int
} 
//历史订单
type OrderHistoryDate struct {
	Data OrderHistory
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
//操作历史
type OperateLogDate struct {
	Data OperateLogModel
	Msg string	`json:"msg"`
	Status int`json:"status"`
}


//操作历史
type OperateLogModel struct {
	Oper       string
	Op         string
	OpSource   string
	Remark     string
	CreateTime string
}
//图片详情
type GetImgDetailReplyDate struct {
	Data GetImgDetailReply
	Msg string	`json:"msg"`
	Status int`json:"status"`
}

//图片详情
type GetImgDetailReply struct {
	/// <summary>
	/// 图片名称
	/// </summary>
	ImageName string
	/// <summary>
	/// 审核标准
	/// </summary>
	CheckPassDesc string

	/// <summary>
	/// 是否通过 1通过  0未通过  -1未审核过 2 重新上传图片
	/// </summary>
	IsPass int

	/// <summary>
	/// 车牌地区
	/// </summary>
	Carlicensetitle string

	/// <summary>
	/// 车牌号码
	/// </summary>
	Carlicensetxt string

	/// <summary>
	/// 登记地区
	/// </summary>
	RegisterArea string

	/// <summary>
	/// 是否显示车主地址
	/// </summary>
	IsShowLikeAddr bool

	/// <summary>
	/// 产品类型
	/// </summary>
	TaskTypeName string

	/// <summary>
	/// 文本退回原因列表
	/// </summary>
	TxtReturnList []TaskReturnLogVo

	/// <summary>
	/// 获取Redis中缓存的json数据并序列化为对象
	/// </summary>
	RedisPretrail RedisPretrailModelV2
}

//获取订单基本信息
type OrderInfoModelDate struct {
	Data GetImgDetailReply
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
//获取城市列表
type ProvincesAndCitysVoDate struct {
	Data ProvincesAndCitysVo
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
//根据车牌定位城市
type ProvinceCityModelDate struct {
	Data ProvinceCityModel
	Msg string	`json:"msg"`
	Status int`json:"status"`
}
type ProvinceCityModel struct {
	/// <summary>
	/// 省ID
	/// </summary>
	ProvinceID int
	/// <summary>
	/// 省份名称
	/// </summary>
	ProvinceName string
}

type UploadPic struct {
	Status          int
	Msg             string
	PicId           int
	PicName         string
	FastDFSBasePath string
	PicPath         string
}


//获取城市列表
type ProvincesAndCitysVo struct {
	ProvincesAndCitys []ProvincesAndCitysVoItem
}

type ProvincesAndCitysVoItem struct {
	  Value string
	  Label string
	  Children []ProvincesAndCitysVoItem
}

type OrderInfoModel struct {
	StatusName string
	Status     int
	Taskid     int
	/// <summary>
	/// 下单账号
	/// </summary>
	CreateUserId int
	/// <summary>
	/// 下单账号+名称
	/// </summary>
	CreateUser string
	/// <summary>
	/// 机构id
	/// </summary>
	SourceID int
	/// <summary>
	/// 机构名称
	/// </summary>
	SourceName string
	/// <summary>
	/// 下单地区
	/// </summary>
	CreateCity string
	/// <summary>
	///下单人机构端填写
	/// </summary>
	LikeMan string
	/// <summary>
	/// 下单张数和类型
	/// </summary>
	TaskTypeName string
	/// <summary>
	/// App下单备注
	/// </summary>
	AppRemark string
	/// <summary>
	/// 预审备注备注
	/// </summary>
	PretrialRemark string
	/// <summary>
	/// 评估师备注
	/// </summary>
	PGSRemark string
	/// <summary>
	/// 系统提示
	/// </summary>
	SysRemark string

	/// <summary>
	/// VIN疑似造假描述
	/// </summary>
	VinSuspectmsg string

	/// <summary>
	/// 订单编号
	/// </summary>
	OrderNo string

	/// <summary>
	/// 是否挂起,true:挂起
	/// </summary>
	IsForked bool

	/// <summary>
	/// 是否显示挂起按钮: 0: 挂起解挂都不显示；1：显示挂起 2：显示解卦
	/// </summary>
	ShowForkBtnFlag int
}

type TaskReturnLogVo struct {
	TaskReturnLogModel
	/// <summary>
	/// 0,不选中，1 选中
	/// </summary>

	IsChecked int

	/// <summary>
	/// 是否默认示例照片
	/// </summary>

	IsDefaultAttach bool

	/// <summary>
	/// 默认示例照片
	/// </summary>
	DefaultAttachUrl string
}

//历史订单
type OrderHistory struct {
	Id int
	/// <summary>
	/// 评估日期
	/// </summary>
	Datestr string
	/// <summary>
	/// 机构名称
	/// </summary>
	SourceName string
	/// <summary>
	/// 品牌车型
	/// </summary>
	CarFullName string
	/// <summary>
	/// 收车价
	/// </summary>
	AssessmentPrace string
	/// <summary>
	/// 售车价
	/// </summary>
	SalePrice string
	/// <summary>
	/// 状态
	/// </summary>
	Status string
	/// <summary>
	/// 报告地址
	/// </summary>
	Reporturl string
	/// <summary>
	/// 详情地址
	/// </summary>
	Detailurl string
}

type PretrialOrderInfoAll struct {
	UserId   int
	UserName string

	/// <summary>
	/// 接单状态,1:接单中（页面显示停止接单按钮）；0:停止接单（页面显示接单按钮）；-1:无权限
	/// </summary>
	UserReceiptOrderStatus int

	WebSocketOnline map[int]string

	/// <summary>
	/// 新分派的订单，待领单
	/// </summary>
	Neworder []PretrialOrderInfo

	/// <summary>
	/// 自己领取的订单，预审中和挂起
	/// </summary>
	Selforder []PretrialOrderInfo

	/// <summary>
	/// 评估师退回的订单
	/// </summary>
	Returnorder []PretrialOrderInfo

	/// <summary>
	/// 被退回数量
	/// </summary>
	ReturnOrder int
	/// <summary>
	/// 排队订单数量
	/// </summary>
	QueueOrder int
	/// <summary>
	/// 预审中
	/// </summary>
	PretrialOrder int
	/// <summary>
	/// 待修改 
	/// </summary>
	UpdateOrder int
	/// <summary>
	/// 待处理 
	/// </summary>
	NeedFixed int
	/// <summary>
	/// 新订单 
	/// </summary>
	NewOrder int

	/// <summary>
	/// 本月已签收
	/// </summary>
	MonthEvaluated int
	/// <summary>
	/// 本day已签收
	/// </summary>
	DayEvaluated int
}




type PretrialOrderInfo struct {
	TaskID        int
	Vin           string
	SourceName    string
	CreateTime    string
	CreateTimeStr string
	Status        int
	TbStatus      int
	StatusName    string
	StatusStr     string
}





//图片结合模型
type GetImgListReply struct {
	/// <summary>
	/// 当前要预审的图片在carPicList中的索引, [-1标识所有图片或者视频审核通过]
	/// </summary>
	CurrentPretrialPicIndex int `json:"currentPretrialPicIndex"`
	/// <summary>
	/// 是否显示退回按钮
	/// </summary>
	ShowPretrailSubmitBack bool `json:"showPretrailSubmitBack"`
	/// <summary>
	/// 是否显示退回到机构按钮
	/// </summary>
	ShowPretrailSubmitBackOrg bool `json:"showPretrailSubmitBackOrg"`
	/// <summary>
	/// 预审按钮显示逻辑 -2不再显示审核按钮  -1完成审核  0 审核图片  1 继续审核  
	/// </summary>
	PretrailBtnStatus int `json:"pretrailBtnStatus"`
	/// <summary>
	/// 图片列表(id,Path)
	/// </summary>
	CarPicList []CarPicModelSimpleVo `json:"carPicList"`
	/// <summary>
	/// 视频信息
	/// </summary>
	VedioInfo VedioInfo `json:"vedioInfo"`
	/// <summary>
	/// 获取Redis中缓存的json数据并序列化为对象
	/// </summary>
	RedisPretrail RedisPretrailModelV2 `json:"redisPretrail"`
	/// <summary>
	/// 服务端内部使用，【客户端不应该使用】
	/// </summary>
	Tc TaskCarBasicEPModel `json:"tc"`
	/// <summary>
	/// 行驶证照片
	/// </summary>
	VehicleLicensePic string `json:"vehicleLicensePic"`
	/// <summary>
	/// 审核通过标准列表
	/// </summary>
	ListCheckPassDesc []CheckPassDescGroup `json:"listCheckPassDesc"`
	/// <summary>
	/// 所有拒评原因
	/// </summary>
	RejectReasons []ConfigRejectReasonModel `json:"rejectReasons"`
	/// <summary>
	/// 审核补充项目Pop说明
	/// </summary>
	PromptTip []RejectPromptMessage `json:"PromptTip"`
}
type VedioInfo struct {
	/// <summary>
	/// ItemId field
	/// </summary>
	  ItemId  int

	/// <summary>
	/// ItemId field
	/// </summary>
	 ItemName string


	/// <summary>
	/// Path field
	/// </summary>
	 Path string

	// <summary>
	/// 是否通过 1通过  0未通过  -1未审核过 2 重新上传图片
	/// </summary>
	 Status int

	/// <summary>
	/// 退回原因
	/// </summary>
	PicReturnReason []TaskReturnLogModel
}
type RedisPretrailModelV2 struct {
	/// <summary>
	/// 订单图片基本信息；(服务端内部缓存图片基本信息，包含附加照片； 只用于索引查询)
	/// </summary>
	 CarPicList []CarPicModelSimple
	/// <summary>
	/// 预审图片键值对(key:预审图片ID,如267表示左前45，详情参见AssessmentItem表)
	/// </summary>
	 PicReturnReason map[string]PretrialPicV2
	/// <summary>
	/// 视频退回原因
	/// </summary>
	  VideoReturnReason PretrialPicV2
	/// <summary>
	/// 文本退回原因，【订单关闭(拒评)，退回机构】
	/// </summary>
    TxtReturnReason []ConfigRejectReasonModel
	/// <summary>
	/// 订单信息
	/// </summary>
	  TaskCarBasic TaskCarBasicEPModel `json:"taskCarBasic"`

	/// <summary>
	/// 新增加附加
	/// </summary>
	AddPic []string
	/// <summary>
	/// 新增加附加信息
	/// </summary>
    AddPicTemp map[string]CarPicModelSimple

	/// <summary>
	/// 工单备注信息
	/// </summary>
	  YsyOrderRemark string
	/// <summary>
	/// 预审上传的附件列表
	/// </summary>
	 YsyOrderAttachFiles []TaskReconsiderationModel
	/// <summary>
	/// 驳回信息-退回给用户的备注（app可见）
	/// </summary>
	  YsyReturnSummaryRemark string
}
type PretrialPicV2 struct {
	/// <summary>
	/// 是否通过  -1未审核过 0未通过 1通过   2重新上传图片(废弃!重新上传后改为未审核)
	/// </summary>
	 IsPass int
	/// <summary>
	/// 修改照片
	/// </summary>
	UpdatePic UpdatePicdetail
	/// <summary>
	/// 图片退回原因列表(图片名称,如左前45, 详情参见AssessmentItem表)
	/// 该数据仅用于在预审总览页以及预审详情页中进行鼠标弹层展示
	/// </summary>
    PicReturn []TaskReturnLogModel
}
type UpdatePicdetail struct {
	/// <summary>
	/// 照片名称
	/// </summary>
	  Picname string
	/// <summary>
	/// 0未修改 1 已修改
	/// </summary>
	 Picstatus int
}
type CarPicModelSimple struct {
	/// <summary>
	/// ID field
	/// </summary>
	Id int
	/// <summary>
	/// ItemId field
	/// </summary>
	ItemId int
	/// <summary>
	/// ItemId field
	/// </summary>
	ItemName   string
	ItemName18 string
	/// <summary>
	/// Path field
	/// </summary>
	Path string
	/// <summary>
	/// Path field
	/// </summary>
	PathBig string
	Index   int
	/// <summary>
	/// 是否是附加图片 false否 true是
	/// </summary>
	IsAnnex bool
}
type TaskReconsiderationModel struct {
	 Id int `json:"id"`
	 Opid int `json:"opid"`
	 Types int `json:"types"`
	 Url string `json:"url"`
	 Timelenth int `json:"timelenth"`
	 Taskid int `json:"taskid"`
	 ItemId int `json:"itemId"`
	 ItemName string `json:"itemName"`
	 OrderKey int `json:"orderKey"`
}
type TaskCarBasicEPModel struct {
	/// <summary>
	/// 车主名称
	/// </summary>
	 TaskOwnerName string
	 ShowArea string
	 ShowArea_v string

	/// <summary>
	/// 特殊照片套餐ID
	/// </summary>
	 ProgramId string
	/// <summary>
	/// ID field
	/// </summary>
	 Id int
	/// <summary>
	/// OrderNo field
	/// </summary>
	 OrderNo string
	/// <summary>
	/// SourceID field
	/// </summary>
	 SourceID int
	/// <summary>
	///下单城市
	/// </summary>
	 CityID int
	/// <summary>
	/// Des field
	/// </summary>
	 Des string
	/// <summary>
	/// LikeMan field
	/// </summary>
	 LikeMan string
	/// <summary>
	/// LikeTel field
	/// </summary>
	 LikeTel string
	/// <summary>
	/// LikeAddr field
	/// </summary>
	 LikeAddr string
	/// <summary>
	/// vin field
	/// </summary>
	 Vin string
	/// <summary>
	/// CarLicense field
	/// </summary>
	 CarLicense string
	/// <summary>
	/// RecordBrand field
	/// </summary>
	 RecordBrand string
	/// <summary>
	/// EngineNum field
	/// </summary>
	 EngineNum string
	/// <summary>
	/// RecordDate field
	/// </summary>
	 RecordDate string

	/// <summary>
	/// MakeID field
	/// </summary>
	 MakeID int
	/// <summary>
	/// ModelID field
	/// </summary>
	 ModelID int
	/// <summary>
	/// StyleID field
	/// </summary>
	 StyleID int
	/// <summary>
	/// Color field
	/// </summary>
	 Color int
	/// <summary>
	/// Mileage field
	/// </summary>
	 Mileage int
	/// <summary>
	/// Service field
	/// </summary>
	 Service int
	/// <summary>
	/// 收车价
	/// </summary>B
	  AssessmentPrace float32
	/// <summary>
	/// 售车价
	/// </summary>
	 SalePrice float32
	/// <summary>
	/// AssessmentDes field
	/// </summary>
	 AssessmentDes string
	/// <summary>
	/// UserID field
	/// </summary>
	 UserID int
	/// <summary>
	/// -1 已删除，0 待检测，1 已确认，2已检测，3超时，4被退回，5已审核，6已签收, 7预审认领, 8预审通过, 9拒评,10退回给预审
	/// </summary>
	 Status int
	 StatusName string
	/// <summary>
	/// 1 待特批2 特批中3 已特批4 待审查5 审查中6 已审无问题7 已审有问题8 待抽查9 抽查中10已查无问题11已查有问题 12待复审 13复审中
	/// </summary>
	 OrderStatus int
	/// <summary>
	/// CreateTime field
	/// </summary>
	 CreateTime string
	/// <summary>
	/// UpdateTime field
	/// </summary>
	 UpdateTime string
	/// <summary>
	/// StartTime field
	/// </summary>
	 StartTime string
	/// <summary>
	/// EndTime field
	/// </summary>
	 EndTime string
	/// <summary>
	/// Exhaust field
	/// </summary>
	 Exhaust string
	/// <summary>
	/// Seating field
	/// </summary>
	 Seating int
	 PerfSeatNum string
	/// <summary>
	/// CarType field
	/// </summary>
	 CarType string
	/// <summary>
	/// DrivingMode field
	/// </summary>
	 DrivingMode int
	/// <summary>
	/// Transmission field
	/// </summary>
	 Transmission int
	/// <summary>
	/// FuelType field
	/// </summary>
	 FuelType int
	/// <summary>
	/// ProductionTime field
	/// </summary>
	 ProductionTime string
	/// <summary>
	/// Certificates field
	/// </summary>
	 Certificates int
	/// <summary>
	/// ManufacturerPrice field
	/// </summary>
	 ManufacturerPrice int
	/// <summary>
	/// 交易价
	/// </summary>
	 BusinessPrice int
	/// <summary>
	/// SetGroupID field
	/// </summary>
	 SetGroupID int

	/// <summary>
	/// 1-18张 2-6张 3-线下 5-9张
	/// </summary>
	 TaskType int
	/// <summary>
	/// TaskBackNum field
	/// </summary>
	 TaskBackNum int
	/// <summary>
	/// TaskBackReason field
	/// </summary>
	 TaskBackReason string
	/// <summary>
	/// AppraiseBackNum field
	/// </summary>
	 AppraiseBackNum int
	/// <summary>
	/// AppraiseBackReason field
	/// </summary>
	 AppraiseBackReason string
	/// <summary>
	/// TransferCount field
	/// </summary>
	 TransferCount int

	/// <summary>
	/// Insurance field
	/// </summary>
	 Insurance string
	/// <summary>
	/// Inspection field
	/// </summary>
	 Inspection string
	/// <summary>
	/// CarDes field
	/// </summary>
	//   CarDes 
	 CreateUserId int
	/// <summary>
	/// 下单省份ID
	/// </summary>
	 ProvID int
	 ProName string
	 CityName string

	 YXOrderNo string


	 VideoPath string
	/// <summary>
	/// 上牌省份
	/// </summary>
	 RegisterProvID int
	 RegisterProvName string
	/// <summary>
	/// 上牌城市
	/// </summary>
	 RegisterCityID int
	 RegisterCityname string
	///// <summary>
	///// 复核收车价，C2B
	///// </summary>
	//32 CompositeProPrice 

	///// <summary>
	///// 复核售车价,B2C
	///// </summary>
	//32 SuggestSellPrice 
	/// <summary>
	/// 下单机构名称
	/// </summary>
	 SourceName string
	/// <summary>
	/// 驱动形式
	/// </summary>
	 Perf_DriveType string
	/// <summary>
	/// 变速器类型
	/// </summary>
	 TransmissionType string
	/// <summary>
	/// 排气量
	/// </summary>
	 Engine_Exhaust string
	/// <summary>
	/// 燃料名称
	/// </summary>
	 Fuel string
	///// <summary>
	///// 发单人
	///// </summary>
	// TaskOwnerName 
	 Tasktel string
	 ProductType int
	 AppraiseBackReasonNew string
	///// <summary>
	///// 数据模型车商收车价
	///// </summary>
	// JZGAssessmentPrice 
	///// <summary>
	///// 数据模型车商售车价
	///// </summary>
	// JZGSalePrice 
	/// <summary>
	/// 是否显示机构名称 0-否 1-是
	/// </summary>
	 ShowSourceName int
	 ProgrammeId int
	 IsComplete int
	/// <summary>
	/// 审车类型 1预警有维保；
	/// 2预警无维保；
	/// 3未预警有维保；
	/// 4未预警无维保；
	/// 5：峰值预警有维保；
	/// 6：峰值预警无维保；
	/// </summary>
	 ReViewType int
	 ReportPcLink string
	 ReportMLink string
	 ReportPrintLink string
	 AutoStar int
	/// <summary>
	/// 预计完成时间
	/// </summary>
	 EstimatedTime string
	/// <summary>
	/// 数据版本 1：1.0流程 2:2.0流程 3:3.0流程
	/// </summary>
	 TaskVersion int
	/// <summary>
	/// 下单手机号
	/// </summary>
	 OrderTelphone string
	/// <summary>
	/// 3.0流程: 预审员ID
	/// </summary>
	 PretrialUser int

	/// <summary>
	/// 车辆(车型)全称
	/// </summary>
	 CarFullName string
	 IsXing int
	 Channel int
	 CreateOrderName string
	 IsForTransfer int
	 FirstDate string
	 SecondDate string

	 AccidentBasis string
	/// <summary>
	/// 1照片 2维保 3其他 多个用,分割保存
	/// </summary>
	 AccidentBasisType string
	 Hdtjyy string
	 Type string

	/// <summary>
	/// 预审抽查 0 待抽查 1抽查中 2已查无问题 3已查有问题
	/// </summary>
	 RandomYS int
	/// <summary>
	/// 评估师抽查 0 待抽查 1抽查中 2已查无问题 3已查有问题
	/// </summary>
	 RandomPGS int
	/// <summary>
	/// 用户名
	/// </summary>
	 IdUserName string
	/// <summary>
	/// 身份证号
	/// </summary>
	  IdNumber string

	 IsMortgage int
	/// <summary>
	/// 是否为事故车 1，事故车-有残值，2，事故车-无残值，3，否，费事故车
	/// </summary>
	 ScrapValue int
	/// <summary>
	/// 保养状态
	/// </summary>
	 MaintainStatus int


	 AssessmentPraceScopeStart float32
	 AssessmentPraceScopeEnd float32
	 SalePriceScopeStart float32
	 SalePriceScopeEnd float32

	/// <summary>
	/// 评估师售车价
	/// </summary>
	 PgsSalePrice float32

	 RejectReasons []ConfigRejectReasonModel

	//2C报告链接
	 SignatureReport string
	 NewEdition int

	/// <summary>
	/// 复议状态: 1 待复议,2 复议中,3 复议通过 4复议驳回
	/// </summary>
	 Reconsideration int
}
type CheckPassDescGroup struct {
	Category string `json:"category"`
	Text     string `json:"text"`
}
type ConfigRejectReasonModel struct {
	SourceName   string
	SourceID     int
	ProductType  int
	ProductName  string
	ItemID       int
	ItemName     string
	ItemDesc     string
	ReasonID     int
	Reason       string
	CreateUserID int
	CreateTime   string
	Status       int
	RejectID     int
	UpdateTime   string
}
type RejectPromptMessage struct {
	ItemName string
	Message  string
}
type CarPicModelSimpleVo struct {
	Id     int
	ItemId int
	ItemName   string
	ItemName18 string
	Path string
	PathBig string
	Index   int
	PicName string `json:"picName"`
	/// <summary>
	/// 是否通过 1通过  0未通过  -1未审核过 2 重新上传图片
	/// </summary>
	Status int
	/// <summary>
	/// 是否是附加图片 false否 true是
	/// </summary>
	IsAnnex bool
	/// <summary>
	/// 退回原因-列表
	/// </summary>
	PicReturnReason []TaskReturnLogModel
}
type TaskReturnLogModel struct {
	/// <summary>
	/// 自增
	/// </summary>
	ID int
	/// <summary>
	/// 订单ID
	/// </summary>
	TaskID int
	/// <summary>
	/// 类型：1图片 2文本基本信息 3视频
	/// </summary>
	Type int
	/// <summary>
	/// 图片ID
	/// </summary>
	ItemID int
	/// <summary>
	/// 新增图片的ID （对应type=3）
	/// </summary>
	PicID int
	/// <summary>
	/// 退回原因ID(关联退回原因表id)
	/// </summary>
	ReturnID int
	/// <summary>
	/// Return(关联退回原因表Name)
	/// </summary>
	ReturnName string
	/// <summary>
	/// 退回原因
	/// </summary>
	ReturnReason string
	/// <summary>
	/// 创建时间
	/// </summary>
	CreateTime string
	/// <summary>
	/// 状态(1:有效，0:无效)
	/// </summary>
	Status int
	/// <summary>
	/// 预审员或评估师ID
	/// </summary>
	CreateUserID int
	/// <summary>
	/// 预审员或评估师名称
	/// </summary>
	CreateUserName string
	/// <summary>
	/// 标题
	/// </summary>
	TitleText string
	/// <summary>
	/// 样例图片
	/// </summary>
	SampleImg string
	/// <summary>
	/// 样例图片全地址
	/// </summary>
	FastDFSBasePath string
	/// <summary>
	/// 附件名称
	/// </summary>
	FileName string
}