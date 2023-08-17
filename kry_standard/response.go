package kry_standard

type ShopQueryStoreDetailResp struct {
	AddressProvince string `json:"addressProvince"` // 省
	Latitude        string `json:"latitude"`        // 纬度
	OrgMode         string `json:"orgMode"`         // 机构类型：单店SINGLE,连锁CHAIN
	ViceMeals       string `json:"viceMeals"`       // 附加业态
	OrgType         string `json:"orgType"`         // 机构类别
	AddressDetail   string `json:"addressDetail"`   // 地址详情
	ShopId          int64  `json:"shopId"`          // 客如云门店id
	CityName        string `json:"cityName"`        // 市
	AreaName        string `json:"areaName"`        // 区
	Contact         string `json:"contact"`         // 联系人
	Longitude       string `json:"longitude"`       // 经度
	AddressCity     string `json:"addressCity"`     // 城市
	MainMeal        string `json:"mainMeal"`        // 主营业态
	ContactMobile   string `json:"contactMobile"`   // 联系电话
	DetailPictures  string `json:"detailPictures"`  // 门店图片地址
	Name            string `json:"name"`            // 门店名称
	ServiceMobile   string `json:"serviceMobile"`   // 服务电话
	ProvinceName    string `json:"provinceName"`    // 省
	AddressArea     string `json:"addressArea"`     // 地区
	Status          string `json:"status"`          // 门店状态：正常NORMAL,禁用DISABLE
	ServerCode      int    `json:"serverCode"`      // 服务编码
	MessageId       string `json:"messageId"`       // 服务消息ID
}

type ShopQueryTableListResp struct {
	PageNum   int `json:"pageNum"`   // 当前分页页数
	PageSize  int `json:"pageSize"`  // 每页条目数
	TotalNum  int `json:"totalNum"`  // 总条目数
	TotalPage int `json:"totalPage"` // 总页数
	TableList []struct {
		TableId          string `json:"tableId"`          // 桌台ID
		TableName        string `json:"tableName"`        // 桌台名称
		AreaId           string `json:"areaId"`           // 桌台所属区域ID
		AreaName         string `json:"areaName"`         // 桌台所属区域名称
		TablePersonCount int    `json:"tablePersonCount"` // 桌台用餐人数
	} `json:"tableList"` // 桌台列表
	MessageId  string `json:"messageId"`  // 服务消息ID
	ServerCode int    `json:"serverCode"` // 服务编码
}

type ShopQueryServiceFeeListResp struct {
	MessageId       string `json:"messageId"`
	ExtraChargeList []struct {
		ExtraChargeId          string `json:"extraChargeId"`
		Name                   string `json:"name"`
		ExtraChargeType        string `json:"extraChargeType"`
		CalcWay                string `json:"calcWay"`
		CalcAmount             int    `json:"calcAmount"`
		AutoAddOrderFlag       string `json:"autoAddOrderFlag"`
		DiscountAfterCalcFlag  string `json:"discountAfterCalcFlag"`
		AllowDiscountShareFlag string `json:"allowDiscountShareFlag"`
		AllowDiscountFlag      string `json:"allowDiscountFlag"`
		EnabledFlag            string `json:"enabledFlag"`
	} `json:"extraChargeList"`
	ServerCode int `json:"serverCode"`
}
