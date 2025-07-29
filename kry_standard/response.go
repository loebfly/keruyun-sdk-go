package kry_standard

type ShopStoreDetailResp struct {
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

type ShopTableListResp struct {
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

type ShopServiceFeeListResp struct {
	MessageId       string `json:"messageId"` // 服务消息ID
	ExtraChargeList []struct {
		ExtraChargeId          string `json:"extraChargeId"`          // 服务费ID
		Name                   string `json:"name"`                   // 服务费名称
		ExtraChargeType        string `json:"extraChargeType"`        // 服务费类型
		CalcWay                string `json:"calcWay"`                // 计算方式,PERCENT按消费比例收取,PERSON按用餐人数收取,FIXED按固定金额收取
		CalcAmount             int    `json:"calcAmount"`             // 计算数额:消费比例/固定金额/每人金额
		AutoAddOrderFlag       string `json:"autoAddOrderFlag"`       // 是否自动加入订单，ON是，OFF否
		DiscountAfterCalcFlag  string `json:"discountAfterCalcFlag"`  // 是否优惠后计算，ON是，OFF否
		AllowDiscountShareFlag string `json:"allowDiscountShareFlag"` // 是否参与优惠分摊，ON是，OFF否
		AllowDiscountFlag      string `json:"allowDiscountFlag"`      // 是否参与折扣，ON是，OFF否
		EnabledFlag            string `json:"enabledFlag"`            // 启用状态，ON启用，OFF禁用
	} `json:"extraChargeList"` // 服务费列表
	ServerCode int `json:"serverCode"` // 服务编码
}

type ShopTableInfoResp struct {
	MessageId        string `json:"messageId"`        // 服务消息ID
	TableId          string `json:"tableId"`          // 桌台ID
	TableName        string `json:"tableName"`        // 桌台名称
	AreaId           string `json:"areaId"`           // 桌台所属区域ID
	AreaName         string `json:"areaName"`         // 桌台所属区域名称
	TablePersonCount int    `json:"tablePersonCount"` // 桌台用餐人数
	ServerCode       int    `json:"serverCode"`       // 服务编码
}

type DishPageResp struct {
	Value struct {
		Total    int `json:"total"` // 数据总数
		DataList []struct {
			DishId           string   `json:"dishId"`           // 菜品ID
			DishName         string   `json:"dishName"`         // 菜品名称
			DishDesc         string   `json:"dishDesc"`         // 菜品描述
			CategoryId       string   `json:"categoryId"`       // 分类ID
			Sort             int      `json:"sort"`             // 排序值
			HelpCode         string   `json:"helpCode"`         // 助记码
			DishType         string   `json:"dishType"`         // 菜品类型， SINGLE：单菜 ，COMBO：套餐， SIDE：配料
			State            string   `json:"state"`            // 菜品状态
			WeighFlag        string   `json:"weighFlag"`        // 称重菜标识, Y:是，N:否
			DishImageUrlList []string `json:"dishImageUrlList"` //菜品图片列表
		} `json:"dataList"` // 具体数据
	} `json:"value"` // 业务结果
	Success    string `json:"success"`    // 是否成功, true:成功，false:失败
	ErrorDesc  string `json:"errorDesc"`  // 错误信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type DishDetailResp struct {
	Value []struct {
		DishId       string `json:"dishId"`       //菜品ID
		DishName     string `json:"dishName"`     //菜品名称
		DishCode     string `json:"dishCode"`     //菜品编码
		DishType     string `json:"dishType"`     //菜品类型。SINGLE：单菜 ，COMBO：套餐， SIDE：配料
		CategoryId   string `json:"categoryId"`   //菜品分类ID
		CategoryName string `json:"categoryName"` //菜品分类名称
		DishSkuList  []struct {
			SpecName       string `json:"specName"`       //规格名称
			DefaultSkuFlag string `json:"defaultSkuFlag"` //是否为默认规格
			SellPrice      int    `json:"sellPrice"`      //售卖价，单位：分
			BarCode        string `json:"barCode"`        //条形码
			SkuId          string `json:"skuId"`          //菜品SKU ID
			DishSkuCode    string `json:"dishSkuCode"`    //菜品SKU Code
		} `json:"dishSkuList"` //菜品SKU列表
		WeighFlag      string `json:"weighFlag"` //是否为称重菜 Y:是；N:否
		Weight         int    `json:"weight"`    //重量，单位：毫克
		UnitName       string `json:"unitName"`  //单位名称
		HelpCode       string `json:"helpCode"`  //助记码
		ComboGroupList []struct {
			GroupName            string `json:"groupName"`  //分组名称
			MaxChoose            int    `json:"maxChoose"`  //套餐分组子菜的最大选择数
			MinChoose            int    `json:"minChoose"`  //套餐分组子菜的最小选择数
			Sort                 int    `json:"sort"`       //排序值
			Repeatable           string `json:"repeatable"` //是否可重复选 Y:是；N:否
			ComboGroupDetailList []struct {
				SingleDishId  string `json:"singleDishId"`  //子菜菜品ID
				MaxChoose     int    `json:"maxChoose"`     //最大可选数量
				MinChoose     int    `json:"minChoose"`     //最小可选数量
				FixChoose     int    `json:"fixChoose"`     //固定数量
				DishSkuId     string `json:"dishSkuId"`     //子菜SKU ID
				DishSkuPrice  int    `json:"dishSkuPrice"`  //套餐分组为可选分组时的子菜加价金额,单位：分
				OptType       string `json:"optType"`       //可选类型(OPTIONAL-可选/REQUIRED-必选)
				DefaultFlag   string `json:"defaultFlag"`   //是否为默认子菜
				DishName      string `json:"dishName"`      //子菜名称
				SpecName      string `json:"specName"`      //子菜规格名
				SellPrice     int    `json:"sellPrice"`     //子菜售卖价，单位：分
				MultiSpecFlag string `json:"multiSpecFlag"` //子菜是否为多规格 Y:是，N:否
				Sort          int    `json:"sort"`          //排序号
			} `json:"comboGroupDetailList"` //套餐分组子菜列表
			GroupType string `json:"groupType"` //套餐组类型（FIXED:固定，OPTIONAL:可选）
		} `json:"comboGroupList"` //菜品为套餐时的套餐分组信息
		ComboPriceIncludeChildDishSideDishPrice   string   `json:"comboPriceIncludeChildDishSideDishPrice"`   //套餐价格是否包含子菜加料价格， Y:是；N:否
		ComboPriceIncludeChildDishCookingWayPrice string   `json:"comboPriceIncludeChildDishCookingWayPrice"` //套餐价格是否包含子菜做法价格, Y:是；N:否
		StatsLabelNameList                        []string `json:"statsLabelNameList"`                        //统计标签名称列表
		SaleLabelNameList                         []string `json:"saleLabelNameList"`                         //售卖标签名称列表
		RemarkNameList                            []string `json:"remarkNameList"`                            //备注名称列表
		CookingWayGroupList                       []struct {
			CookingWayGroupName string `json:"cookingWayGroupName"` //做法组名
			OptType             string `json:"optType"`             //可选类型(OPTIONAL-可选/REQUIRED-必选)
			CookingWayList      []struct {
				CookingWayName string `json:"cookingWayName"` //做法名称
				AddPrice       int    `json:"addPrice"`       //加价，单位：分
				DefaultFlag    string `json:"defaultFlag"`    //是否为默认
				CookingWayId   string `json:"cookingWayId"`   //做法ID
			} `json:"cookingWayList"` //做法列表
			CookingWayGroupId string `json:"cookingWayGroupId"` //做法组ID
		} `json:"cookingWayGroupList"` //做法组信息
		SideDishGroupList []struct {
			GroupName               string `json:"groupName"` //分组名称
			SideDishGroupDetailList []struct {
				SideDishName string `json:"sideDishName"` //加料名称
				SideDishId   string `json:"sideDishId"`   //加料ID
				AddPrice     int    `json:"addPrice"`     //配料价格 单位： 分
				Sort         int    `json:"sort"`         //排序值
			} `json:"sideDishGroupDetailList"` //加料分组明细列表
		} `json:"sideDishGroupList"` //加料组信息
		State             string   `json:"state"`            //售卖状态(ONLINE-售卖, PAUSE-停售)
		StartNumber       int      `json:"startNumber"`      //起售数量
		StartInterval     int      `json:"startInterval"`    //增售数量
		ModifyPriceFlag   string   `json:"modifyPriceFlag"`  //是否可以手动改价
		DiscountFlag      string   `json:"discountFlag"`     //是否可以手动打折
		OrderSingleFlag   string   `json:"orderSingleFlag"`  //是否可以单点
		DishImageUrlList  []string `json:"dishImageUrlList"` //菜品图片列表
		SpicyLevel        int      `json:"spicyLevel"`       //辣度等级
		Sort              int      `json:"sort"`             //排序值
		DishStockInfoList []struct {
			DishId          string `json:"dishId"`          //菜品SPU ID
			SkuId           string `json:"skuId"`           //菜品SKU ID
			ResidualDecimal int64  `json:"residualDecimal"` //剩余售卖量
			SoldOutFlag     string `json:"soldOutFlag"`     //是否售罄（Y：售罄，N：未售罄）
		} `json:"dishStockInfoList"` //菜品库存信息列表
		UnitId string `json:"unitId"` //菜品单位ID
	} `json:"value"` //业务数据
	Success    bool   `json:"success"`    //是否成功,true:成功,false:失败
	ErrorDesc  string `json:"errorDesc"`  //错误描述
	ServerCode int    `json:"serverCode"` //服务编码
	MessageId  string `json:"messageId"`  //服务消息ID
}

type DishCategoryResp struct {
	Value []struct {
		CategoryId   string `json:"categoryId"`   // 分类ID
		ParentId     string `json:"parentId"`     // 父级分类ID
		CategoryName string `json:"categoryName"` // 分类名称
		Sort         int    `json:"sort"`         // 排序值
		CategoryType string `json:"categoryType"` // 分类类型，DISH：菜品分类 SIDE_DISH_GROUP:加料分组
	} `json:"value"` // 业务结果
	Success    bool   `json:"success"`    // 是否成功, true:成功，false:失败
	ErrorDesc  string `json:"errorDesc"`  // 错误信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmCreateResp struct {
	Data struct {
		CustomerId string `json:"customerId"` // 用户id
	} `json:"data"`
	Success    bool   `json:"success"`    // 是否成功, true成功, false失败
	MsgCode    string `json:"msgCode"`    // success为false时的异常码
	MsgInfo    string `json:"msgInfo"`    // success为false时的异常信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmUpdateResp struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	MessageUuid string `json:"messageUuid"`
	Result      struct {
		CustomerId string `json:"customerId"`
		Success    string `json:"success"`
		BizMessage string `json:"bizMessage"`
		ServerCode int    `json:"serverCode"`
		MessageId  string `json:"messageId"`
	} `json:"result"`
}

type CrmCustomerInfoResp struct {
	Data struct {
		Birthday   string `json:"birthday"`   // 生日日期,yyyy-MM-dd格式
		CustomerId string `json:"customerId"` // 用户id
		Gender     int    `json:"gender"`     // 用户性别,0代表女,1代表男,2代表其他
		LevelDTO   struct {
			LevelName string `json:"levelName"` // 等级名称
			LevelNo   int    `json:"levelNo"`   // 等级级别
		} `json:"levelDTO"` // 等级对象
		Mobile string `json:"mobile"` // 手机号
		Name   string `json:"name"`   // 用户名称
		State  int    `json:"state"`  // 用户状态,1代表启用,0代表停用
	} `json:"data"` // 数据体对象
	Success    bool   `json:"success"`    // 请求请求结果, true代表成功, false代表失败
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmCustomerIdsResp struct {
	List       []Customer `json:"list"`
	Success    bool       `json:"success"`
	BizMessage string     `json:"bizMessage"`
	ServerCode string     `json:"serverCode"`
	MessageId  string     `json:"messageId"`
}

type Customer struct {
	Birthday   int64 `json:"birthday"`
	Gender     int   `json:"gender"`
	MemberTime int64 `json:"memberTime"`
	LevelDTO   struct {
		LevelId   string `json:"levelId"`
		LevelNo   int    `json:"levelNo"`
		LevelName string `json:"levelName"`
	} `json:"levelDTO"`
	Remark       string `json:"remark"`
	CustomerType int    `json:"customerType"`
	CustomerId   string `json:"customerId"`
	State        int    `json:"state"`
	Email        string `json:"email"`
	CustomerTime int64  `json:"customerTime"`
	Address      string `json:"address"`
	Mobile       string `json:"mobile"`
	Name         string `json:"name"`
	Growth       string `json:"growth"`
	Invoice      string `json:"invoice"`
}

type CrmCustomerPropertyResp struct {
	Data struct {
		CustomerDTO struct {
			CustomerId string `json:"customerId"` // 用户id
			State      int    `json:"state"`      // 用户状态
			Mobile     string `json:"mobile"`     // 手机号
		} `json:"customerDTO"` // 会员基础信息
		PosCardDTOList []struct {
			PosRechargeAccountList []struct {
				RemainAvailableValue struct {
					TotalValue int `json:"totalValue"` // 当前剩余可用储值总金额 = 当前剩余可用实储总金额 + 当前剩余可用赠储总金额，单位：分
					RealValue  int `json:"realValue"`  // 当前剩余可用实储总金额，单位：分
					GiftValue  int `json:"giftValue"`  // 当前剩余可用赠储总金额，单位：分
				} `json:"remainAvailableValue"` // 当前剩余可用储值
			} `json:"posRechargeAccountList"` // 储值账户
			Status   string `json:"status"`   // 卡状态，SOLD：已出售；ACTIVED：已激活；STOP：已停用；INVALID：已作废；EXPIRED：已过期；REFUND：已退卡
			CardId   string `json:"cardId"`   // 卡号id
			CardType string `json:"cardType"` // 卡类型，MEMBER：会员卡；GIFT：礼品卡；PAY_MEMBER：付费会员卡
		} `json:"posCardDTOList"` // 卡列表
		NormalVoucherInstanceCount int `json:"normalVoucherInstanceCount"` // 可使用有效券张数
		PointAccountDTO            struct {
			RemainAvailableValue int `json:"remainAvailableValue"`
		} `json:"pointAccountDTO"` // 积分账户
	} `json:"data"` // 数据体
	Success    bool   `json:"success"`    // 是否成功,true:成功,false:失败
	ServerCode int    `json:"serverCode"` // 服务端返回码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmDirectChargeResp struct {
	Success bool   `json:"success"` // 是否成功,true:成功，false:失败
	MsgCode string `json:"msgCode"` // success为false时的异常码
	MsgInfo string `json:"msgInfo"` // success为false时的异常信息
	Data    struct {
		AccountId        string `json:"accountId"` // 储值账户id
		RemainTotalValue struct {
			RealValue  int64 `json:"realValue"`  // 实储余额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 赠储余额，单位：分
			PreValue   int64 `json:"preValue"`   // 预储余额，单位：分
			TotalValue int64 `json:"totalValue"` // 总余额=实储余额 + 赠储余额 + 预储余额，单位：分
		} `json:"remainTotalValue"` // 储值账户余额 = 储值账户可用余额 + 储值账户预扣金额
		RemainAvailableValue struct {
			RealValue  int64 `json:"realValue"`  // 可用实储余额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 可用赠储余额，单位：分
			PreValue   int64 `json:"preValue"`   // 可用预储余额，单位：分
			TotalValue int64 `json:"totalValue"` // 可用总余额=可用实储余额 + 可用赠储余额 +可用预储余额，单位：分
		} `json:"remainAvailableValue"` // 储值账户可用余额
		PreDeductValue struct {
			RealValue  int64 `json:"realValue"`  // 预扣实储金额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 预扣赠储金额，单位：分
			PreValue   int64 `json:"preValue"`   // 预扣预储金额，单位：分
			TotalValue int64 `json:"totalValue"` // 预扣总金额=预扣实储金额 + 预扣赠储金额 + 预扣预储金额，单位：分
		} `json:"preDeductValue"` // 储值账户预扣金额
		TotalValue struct {
			RealValue  int64 `json:"realValue"`  // 累计实储金额，单位：分
			GiftValue  int64 `json:"giftValue"`  // 累计赠储金额，单位：分
			PreValue   int64 `json:"preValue"`   // 累计预储金额，单位：分
			TotalValue int64 `json:"totalValue"` // 累计储值总金额=累计实储金额 +累计赠储金额 + 累计预储金额，单位：分
		} `json:"totalValue"` // 储值累计总额
	} `json:"data"` // 业务数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type CrmTemplateListResp struct {
	Result     []CouponTemplate `json:"result"`
	TotalSize  int              `json:"totalSize"`
	Success    bool             `json:"success"`
	MsgInfo    string           `json:"msgInfo"`
	MsgCode    string           `json:"msgCode"`
	ServerCode int              `json:"serverCode"`
	MessageId  string           `json:"messageId"`
}

// CouponTemplate 是 result 数组中的单个元素
type CouponTemplate struct {
	MinChargeType      string   `json:"minChargeType"`
	UnavailableTime    string   `json:"unavailableTime"` // 这里是嵌套的JSON字符串
	Inventory          int      `json:"inventory"`
	Type               string   `json:"type"`
	MinCharge          int      `json:"minCharge"`
	AvailableTime      string   `json:"availableTime"` // 这里是嵌套的JSON字符串
	EffectiveAfterDays string   `json:"effectiveAfterDays"`
	ValidDayCount      int      `json:"validDayCount"`
	StartTime          int64    `json:"startTime"`
	Id                 string   `json:"id"`
	UserLimit          int      `json:"userLimit"`
	ShopCoverage       string   `json:"shopCoverage"`
	ValidDateType      string   `json:"validDateType"`
	ItemCoverage       string   `json:"itemCoverage"`
	Instruction        string   `json:"instruction"`
	UseChannels        []string `json:"useChannels"`
	Name               string   `json:"name"`
	EndTime            int64    `json:"endTime"`
	Status             string   `json:"status"`
}

type CrmTemplateInfoResp struct {
	Data       InfoCouponTemplate `json:"data"`
	Success    bool               `json:"success"`
	MsgInfo    string             `json:"msgInfo"`
	MsgCode    string             `json:"msgCode"`
	ServerCode int                `json:"serverCode"`
	MessageId  string             `json:"messageId"`
}

type InfoCouponTemplate struct {
	CouponTemplate
	DiscountSetting struct {
		Discount string `json:"discount"`
	} `json:"discountSetting"`
	GiftSetting struct {
		ExchangeType string `json:"exchangeType"`
		Type         int    `json:"type"`
		Name         string `json:"name"`
		Cost         int64  `json:"cost"` // 使用指针类型，因为JSON中可能为null
	} `json:"giftSetting"`
	BoughtGiftSetting struct {
		ExchangeRuleType string `json:"exchangeRuleType"`
		ItemBoughtNum    int    `json:"itemBoughtNum"`
		ItemExchangeNum  int    `json:"itemExchangeNum"`
		LimitDiscountNum int    `json:"limitDiscountNum"`
	} `json:"boughtGiftSetting"`
	ConditionDiscountSetting struct {
		ItemBoughtNum    int    `json:"itemBoughtNum"`
		Discount         int    `json:"discount"`
		ExchangeItemType string `json:"exchangeItemType"`
	} `json:"conditionDiscountSetting"`
	CashSetting struct {
		Denomination int `json:"denomination"`
	} `json:"cashSetting"`
}

type CrmTemplateShopResp struct {
	Result     []ShopInfo `json:"result"`
	TotalSize  int        `json:"totalSize"`
	MsgInfo    string     `json:"msgInfo"`
	MsgCode    string     `json:"msgCode"`
	ServerCode int        `json:"serverCode"`
	MessageId  string     `json:"messageId"`
	Success    bool       `json:"success"`
}
type ShopInfo struct {
	ShopId string `json:"shopId"`
}

type CrmTemplateDishResp struct {
	Result     []DishInfo `json:"result"`
	TotalSize  int        `json:"totalSize"`
	Success    bool       `json:"success"`
	MsgInfo    string     `json:"msgInfo"`
	MsgCode    string     `json:"msgCode"`
	ServerCode int        `json:"serverCode"`
	MessageId  string     `json:"messageId"`
}

type DishInfo struct {
	ItemId        string `json:"itemId"`
	ItemApplyType string `json:"itemApplyType"`
	ItemType      string `json:"itemType"`
	DishId        string `json:"dishId"`
}

type CrmCouponSendResp struct {
	Data struct {
		ResourceCode string `json:"resourceCode"`
	} `json:"data"`
	Success    bool   `json:"success"`
	MsgInfo    string `json:"msgInfo"`
	MsgCode    string `json:"msgCode"`
	ServerCode int    `json:"serverCode"`
	MessageId  string `json:"messageId"`
}

type CrmCouponInvalidResp struct {
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
	Success    bool   `json:"success"`
	MsgInfo    string `json:"msgInfo"`
	MsgCode    string `json:"msgCode"`
	ServerCode int    `json:"serverCode"`
	MessageId  string `json:"messageId"`
}

type CrmCouponQueryResp struct {
	Result     []TemplateInfo `json:"result"`
	TotalSize  int            `json:"totalSize"`
	Success    bool           `json:"success"`
	MsgInfo    string         `json:"msgInfo"`
	MsgCode    string         `json:"msgCode"`
	ServerCode int            `json:"serverCode"`
	MessageId  string         `json:"messageId"`
}

type TemplateInfo struct {
	TemplateType string `json:"templateType"`
	ResourceCode string `json:"resourceCode"`
	TemplateName string `json:"templateName"`
	GmtCreated   int64  `json:"gmtCreated"`
	StartTime    int64  `json:"startTime"`
	EndTime      int64  `json:"endTime"`
	TemplateId   string `json:"templateId"`
	Status       string `json:"status"`
}

type CrmCouponInfoResp struct {
	Data       Data   `json:"data"`
	Success    bool   `json:"success"`
	MsgInfo    string `json:"msgInfo"`
	MsgCode    string `json:"msgCode"`
	ServerCode int    `json:"serverCode"`
	MessageId  string `json:"messageId"`
}

type Data struct {
	ResourceCode    string          `json:"resourceCode"`
	GmtCreated      int64           `json:"gmtCreated"`
	StartTime       int64           `json:"startTime"`
	VoucherTemplate VoucherTemplate `json:"voucherTemplate"`
	EndTime         int64           `json:"endTime"`
	Status          string          `json:"status"`
}
type VoucherTemplate struct {
	MinChargeType   string   `json:"minChargeType"`
	ItemCoverage    string   `json:"itemCoverage"`
	Type            string   `json:"type"`
	MinCharge       int      `json:"minCharge"`
	AvailableTime   string   `json:"availableTime"` // 嵌套JSON字符串
	Instruction     string   `json:"instruction"`
	UseChannels     []string `json:"useChannels"`
	Name            string   `json:"name"`
	ID              string   `json:"id"`
	Status          string   `json:"status"`
	ShopCoverage    string   `json:"shopCoverage"`
	DiscountSetting struct {
		Discount string `json:"discount"`
	} `json:"discountSetting"`
	GiftSetting struct {
		ExchangeType string  `json:"exchangeType"`
		Type         int     `json:"type"`
		Name         string  `json:"name"`
		Cost         *string `json:"cost,omitempty"`
	} `json:"giftSetting"`
	BoughtGiftSetting struct {
		ExchangeRuleType string `json:"exchangeRuleType"`
		ItemBoughtNum    int    `json:"itemBoughtNum"`
		ItemExchangeNum  int    `json:"itemExchangeNum"`
		LimitDiscountNum int    `json:"limitDiscountNum"`
	} `json:"boughtGiftSetting"`
	ConditionDiscountSetting struct {
		ItemBoughtNum    int    `json:"itemBoughtNum"`
		Discount         int    `json:"discount"`
		ExchangeItemType string `json:"exchangeItemType"`
	} `json:"conditionDiscountSetting"`
	CashSetting struct {
		Denomination int `json:"denomination"`
	} `json:"cashSetting"`
	UnavailableTime string `json:"unavailableTime"` // 嵌套JSON字符串
}

type OrderDetailResp struct {
	Success  bool   `json:"success"`  // 是否成功,true:成功，false:失败
	MsgCode  string `json:"msgCode"`  // success为false时的异常码
	MsgInfo  string `json:"msgInfo"`  // success为false时的异常信息
	CanRetry bool   `json:"canRetry"` // 可以重试
	Data     struct {
		OrderBaseVO struct {
			ShopId             string      `json:"shopId"`             // 门店ID
			ShopName           string      `json:"shopName"`           // 门店名称
			BrandId            string      `json:"brandId"`            // 品牌ID
			BrandName          string      `json:"brandName"`          // 品牌名称
			OrderId            string      `json:"orderId"`            // 订单ID
			BusiOrderNo        string      `json:"busiOrderNo"`        // 业务订单号
			ThirdOrderNo       string      `json:"thirdOrderNo"`       // 三方订单号
			OpenTime           string      `json:"openTime"`           // 下单时间
			SettleTime         string      `json:"settleTime"`         // 结账时间
			FinishBusiDate     string      `json:"finishBusiDate"`     // 营业日
			OrderSource        string      `json:"orderSource"`        // 订单来源
			OrderType          string      `json:"orderType"`          // 订单类型
			OrderStatus        string      `json:"orderStatus"`        // 订单状态
			OrderAmt           interface{} `json:"orderAmt"`           // 订单金额
			PromoAmt           string      `json:"promoAmt"`           // 优惠金额
			OrderReceivedAmt   string      `json:"orderReceivedAmt"`   // 订单收入
			OpenOperatorName   string      `json:"openOperatorName"`   // 开单人
			SettleOperatorName string      `json:"settleOperatorName"` // 结账人
			MemberId           string      `json:"memberId"`           // 会员ID
			MemberPhone        string      `json:"memberPhone"`        // 会员手机号
			MemberName         string      `json:"memberName"`         // 会员姓名
			OrderPeopleCnt     string      `json:"orderPeopleCnt"`     // 就餐人数
			SerialNo           string      `json:"serialNo"`           // 流水号
			Subject            string      `json:"subject"`            // 整单备注
			RelatedOrderId     string      `json:"relatedOrderId"`     // 关联订单ID，如整单退或反结时的原单订单ID
			ThirdSerialNo      string      `json:"thirdSerialNo"`      // 第三方订单流水号
		} `json:"orderBaseVO"` // 主单信息
		OrderTableVoList []struct {
			TableName string `json:"tableName"` // 桌台名称
			TableId   string `json:"tableId"`   // 桌台ID
		} `json:"orderTableVoList"` // 桌台信息
		OrderItemVoList []struct {
			ItemType           string `json:"itemType"`           // 商品类型
			GiftFlag           bool   `json:"giftFlag"`           // 是否赠送
			WeighFlag          bool   `json:"weighFlag"`          // 是否是称重商品
			TempFlag           bool   `json:"tempFlag"`           // 是否是临时菜
			PromoFlag          bool   `json:"promoFlag"`          // 是否是优惠菜
			BigTypeName        string `json:"bigTypeName"`        // 商品大类名称
			MidTypeName        string `json:"midTypeName"`        // 商品中类名称
			ItemCode           string `json:"itemCode"`           // 商品编码
			ItemName           string `json:"itemName"`           // 商品名称
			SaleStatusType     string `json:"saleStatusType"`     // 商品售卖状态类型
			SaleStatusTypeCode string `json:"saleStatusTypeCode"` // 商品售卖状态类型编码
			Id                 string `json:"id"`                 // 商品id
			ParentId           string `json:"parentId"`           // 父商品id
			ProductionDeptId   string `json:"productionDeptId"`   // 出品部门ID
			UnitName           string `json:"unitName"`           // 单位名称
			SpecName           string `json:"specName"`           // 规格
			SpecNameConcat     string `json:"specNameConcat"`     // 规格名称全称
			ItemPrice          string `json:"itemPrice"`          // 商品原始单价
			SalePrice          string `json:"salePrice"`          // 商品售价
			PracticeVoList     []struct {
				PracticeName string `json:"practiceName"` // 做法名称
				PracticeAmt  int    `json:"practiceAmt"`  // 做法金额
			} `json:"practiceVoList"` // 做法
			ItemSubject           string `json:"itemSubject"`           // 商品备注
			Quantity              string `json:"quantity"`              // 销售数量、退菜数量或赠菜数量 1. 是否是赠菜通过giftFlag判断 2. 是否是退菜通过saleStatusTypeCode判断，DISCARD表示退菜，NORMAL表示销售
			ItemSaleAmt           string `json:"itemSaleAmt"`           // 商品销售金额
			ExtraFeeApportionAmt  string `json:"extraFeeApportionAmt"`  // 服务费分摊金额
			ItemPromoApportionAmt string `json:"itemPromoApportionAmt"` // 商品优惠分摊
			ItemReceivedAmt       string `json:"itemReceivedAmt"`       // 商品收入
			Children              []struct {
				ItemType              string        `json:"itemType"`
				GiftFlag              bool          `json:"giftFlag"`
				WeighFlag             bool          `json:"weighFlag"`
				TempFlag              bool          `json:"tempFlag"`
				PromoFlag             bool          `json:"promoFlag"`
				BigTypeName           string        `json:"bigTypeName"`
				ItemCode              string        `json:"itemCode"`
				ItemName              string        `json:"itemName"`
				SaleStatusType        string        `json:"saleStatusType"`
				SaleStatusTypeCode    string        `json:"saleStatusTypeCode"`
				Id                    string        `json:"id"`
				ParentId              string        `json:"parentId"`
				ItemId                string        `json:"itemId"`
				ItemSkuId             string        `json:"itemSkuId"`
				ProductionDeptId      string        `json:"productionDeptId"`
				UnitName              string        `json:"unitName"`
				SpecNameConcat        string        `json:"specNameConcat"`
				ItemPrice             string        `json:"itemPrice"`
				SalePrice             string        `json:"salePrice"`
				PracticeVoList        []interface{} `json:"practiceVoList"`
				Quantity              string        `json:"quantity"`
				ItemSaleAmt           string        `json:"itemSaleAmt"`
				ExtraFeeApportionAmt  string        `json:"extraFeeApportionAmt"`
				ItemPromoApportionAmt string        `json:"itemPromoApportionAmt"`
				ItemReceivedAmt       string        `json:"itemReceivedAmt"`
			} `json:"children"` // 子节点
		} `json:"orderItemVoList"` // 菜品明细
		OpenOrderPromoVoList []struct {
			PromoName    string `json:"promoName"`    // 优惠名称
			PromoType    string `json:"promoType"`    // 优惠类型
			PromoTime    string `json:"promoTime"`    // 优惠时间
			OperatorName string `json:"operatorName"` // 操作人姓名
			PromoAmt     int    `json:"promoAmt"`     // 优惠金额
		} `json:"openOrderPromoVoList"` // 订单优惠信息
		OpenExtraFeeVoList []struct {
			ExtraFeeName              string `json:"extraFeeName"`              // 服务费名称
			ExtraFeeType              string `json:"extraFeeType"`              // 服务费类型
			ExtraFeeAmt               int    `json:"extraFeeAmt"`               // 服务费金额
			ExtraFeePromoApportionAmt int    `json:"extraFeePromoApportionAmt"` // 服务费优惠分摊金额
			ExtraFeeReceivedAmt       int    `json:"extraFeeReceivedAmt"`       // 服务费收入
			ApportionedAmt            int    `json:"apportionedAmt"`            // 被分摊金额
			PromoTotalAmt             int    `json:"promoTotalAmt"`             // 服务费优惠总金额
		} `json:"openExtraFeeVoList"` // 服务费信息
		OpenPaymentDetailVoList []struct {
			PayDetailNo         string `json:"payDetailNo"`         // 支付单号
			FaceAmt             string `json:"faceAmt"`             // 面额
			PayAmt              string `json:"payAmt"`              // 支付金额
			ShopPromoAmt        string `json:"shopPromoAmt"`        // 商户优惠金额
			PlatformServiceAmt  string `json:"platformServiceAmt"`  // 平台抽佣/服务费
			ActualReceiveAmt    string `json:"actualReceiveAmt"`    // 商户实收
			PayMethodName       string `json:"payMethodName"`       // 支付方式名称
			PayDetailStatus     string `json:"payDetailStatus"`     // 支付状态
			PayDetailStartTime  string `json:"payDetailStartTime"`  // 创建时间/支出开始时间
			PayDetailEndTime    string `json:"payDetailEndTime"`    // 支付/退款完成时间
			PayPromoAmt         string `json:"payPromoAmt"`         // 支付优惠
			CouponCnt           string `json:"couponCnt"`           // 券数量
			CouponName          string `json:"couponName"`          // 券名称
			OperatorName        string `json:"operatorName"`        // 操作人姓名/收银员
			ActualPayAmt        string `json:"actualPayAmt"`        // 实际支付金额
			PlatformPromoAmt    string `json:"platformPromoAmt"`    // 平台优惠金额
			PayMethodId         int    `json:"payMethodId"`         // 支付方式ID，-3:现金，-4:银行卡，-129:扫码支付，-130:收银码，-5:微信，-6:支付宝，-37:云闪付，-1:会员卡，-15:实体卡，-20:匿名卡，-127:储值补录，-128:挂账，-24:美团团购券，-36:口碑团购券，0:抵用券
			CouponReconcileFlag bool   `json:"couponReconcileFlag"` // 团购券是否已对账 true:已对账 false:未对账
		} `json:"openPaymentDetailVoList"` // 支付信息
	} `json:"data"` // 数据
	ExtInfo struct {
		TraceId                string `json:"traceId"`                // 跟踪标识
		ServerCurrentTimestamp string `json:"serverCurrentTimestamp"` // 服务器当前时间戳
		ServerCurrentTime      string `json:"serverCurrentTime"`      // 服务器当前时间
	} `json:"extInfo"` // 附加信息
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type OrderListResp struct {
	Success    bool `json:"success"`    // 是否成功,true:成功,false:失败
	ServerCode int  `json:"serverCode"` // 服务编码
	Data       struct {
		TotalCount string `json:"totalCount"` // 总数
		PageNo     string `json:"pageNo"`     // 页号
		PageSize   string `json:"pageSize"`   // 页大小
		TotalPage  string `json:"totalPage"`  // 总页数
		List       []struct {
			BrandId          string `json:"brandId"`          // 品牌ID
			ShopId           string `json:"shopId"`           // 门店ID
			OrderId          string `json:"orderId"`          // 订单ID
			BusiOrderNo      string `json:"busiOrderNo"`      // 业务订单号
			OrderStatus      string `json:"orderStatus"`      // 订单状态{1:待处理(WAIT_PROCESSED) 2:已接单(ORDER_RECEIVED) 3:已完成(SUCCESS) 4:待结账(WAIT_SETTLED) 5:已结账(SETTLED) 6:已退单(REFUND) 7:已关闭(CLOSED) 8:已作废(INVALID) 9:已取消(CANCELLED) 10:已拒绝(REJECTED) 11:已反结账(ANTI_SETTLED)}
			OrderType        string `json:"orderType"`        // 订单类型{1:堂食(FOR_HERE) 2:平台外卖(PLATFORM_TAKE_OUT) 3:自营外卖(SELF_TAKE_OUT) 4:自提(SELF_TAKE) 5:无单收银(NO_ORDER_CASHIER) 6:会员充值(MEMBER_STORE) 7:会员补录(MEMBER_MANUAL_STORE) 8:销账订单(REPAYMENT_ORDER)}
			OrderAmt         string `json:"orderAmt"`         // 订单金额(单位/分)
			PromoAmt         string `json:"promoAmt"`         // 优惠金额(单位/分)
			OrderReceivedAmt string `json:"orderReceivedAmt"` // 订单收入/订单实收(单位/分)
			OpenTime         string `json:"openTime"`         // 下单时间
			FinishTime       string `json:"finishTime"`       // 完结时间
			ThirdOrderNo     string `json:"thirdOrderNo"`     // 第三方订单号
			SerialNo         string `json:"serialNo"`         // 订单流水号
			ThirdSerialNo    string `json:"thirdSerialNo"`    // 第三方流水号
		} `json:"list"` // 列表
		PrevPage     string `json:"prevPage"`     // 上一页
		NextPage     string `json:"nextPage"`     // 下一页
		EmptyForList bool   `json:"emptyForList"` // 当前list是否为null
	} `json:"data"` // 	数据
	MessageId string `json:"messageId"` // 服务消息ID
}

type StockInOutListResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		List []struct {
			Id             string `json:"id"`             // 单据id
			OrgId          string `json:"orgId"`          // 机构id
			OrgName        string `json:"orgName"`        // 机构名字
			DepotId        string `json:"depotId"`        // 仓库id
			DepotName      string `json:"depotName"`      // 仓库名字
			BussDate       string `json:"bussDate"`       // 业务日期
			BillNo         string `json:"billNo"`         // 单号
			Status         int    `json:"status"`         // 状态 961：暂存，963：提交，962：已审核，964：已驳回，965：已反审，966：已作废
			TotalAmt       int    `json:"totalAmt"`       // 成本总金额
			TotalNum       int    `json:"totalNum"`       // 总数量
			Remarks        string `json:"remarks"`        // 备注
			CreateUserName string `json:"createUserName"` // 创建人
			UpdateUserName string `json:"updateUserName"` // 修改人
			CreateTime     string `json:"createTime"`     // 创建时间
			UpdateTime     string `json:"updateTime"`     // 修改时间
		} `json:"list"` // 单据列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockInOutDetailResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    []struct {
		Id         string `json:"id"` // 单据id
		DetailList []struct {
			GoodsId       string `json:"goodsId"`       // 物品id
			GoodsName     string `json:"goodsName"`     // 物品名字
			GoodsCode     string `json:"goodsCode"`     // 物品编码
			GoodsSpec     string `json:"goodsSpec"`     // 物品规格
			UnitId        string `json:"unitId"`        // 标准单位id
			UnitName      string `json:"unitName"`      // 标准单位名字
			DualUnitName  string `json:"dualUnitName"`  // 辅助单位名字
			DualUnitId    string `json:"dualUnitId"`    // 辅助单位id
			GoodsQty      int    `json:"goodsQty"`      // 出库数量
			DualGoodsQty  int    `json:"dualGoodsQty"`  // 辅助单位出库数量
			UnitPrice     int    `json:"unitPrice"`     // 成本单价
			GoodsAmt      int    `json:"goodsAmt"`      // 成本金额
			ProDate       string `json:"proDate"`       // 生产日期
			ExpDate       string `json:"expDate"`       // 有效期至
			QualityPeriod int    `json:"qualityPeriod"` // 保质期天数
			GoodsBatchNo  string `json:"goodsBatchNo"`  // 批号
			Remarks       string `json:"remarks"`       // 备注
		} `json:"detailList"` // 单据明细列表
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockTransferInternalListResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		List []struct {
			Id             string `json:"id"`             // 单据id
			StoreId        string `json:"storeId"`        // 机构id
			StoreName      string `json:"storeName"`      // 机构名字
			BillNo         string `json:"billNo"`         // 单号
			BussDate       string `json:"bussDate"`       // 业务日期
			OutDepotId     string `json:"outDepotId"`     // 出库仓库id
			OutDepotName   string `json:"outDepotName"`   // 出库仓库名字
			InDepotId      string `json:"inDepotId"`      // 入库仓库id
			InDepotName    string `json:"inDepotName"`    // 入库仓库名字
			TotalAmt       int    `json:"totalAmt"`       // 成本总金额
			Status         int    `json:"status"`         // 状态：961：暂存，963：提交，962：已审核，964：已驳回，965：已反审
			CreateUserName string `json:"createUserName"` // 创建人
			UpdateUserName string `json:"updateUserName"` // 修改人
			CreateTime     string `json:"createTime"`     // 创建时间
			UpdateTime     string `json:"updateTime"`     // 修改时间
			Remarks        string `json:"remarks"`        // 备注
		} `json:"list"` // 单据列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockTransferInOutDetailResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    []struct {
		Id         string `json:"id"` // 单据id
		DetailList []struct {
			GoodsId       string `json:"goodsId"`       // 物品id
			GoodsCode     string `json:"goodsCode"`     // 物品编码
			GoodsName     string `json:"goodsName"`     // 物品名字
			GoodsSpec     string `json:"goodsSpec"`     // 物品规格
			UnitId        string `json:"unitId"`        // 标准单位id
			UnitName      string `json:"unitName"`      // 标准单位名字
			UnitPrice     int    `json:"unitPrice"`     // 成本单价
			GoodsQty      int    `json:"goodsQty"`      // 数量
			GoodsAmt      int    `json:"goodsAmt"`      // 成本金额
			DualGoodsQty  int    `json:"dualGoodsQty"`  // 辅助单位数量
			DualUnitName  string `json:"dualUnitName"`  // 辅助单位名字
			DualUnitId    string `json:"dualUnitId"`    // 辅助单位id
			Remarks       string `json:"remarks"`       // 备注
			ProDate       string `json:"proDate"`       // 生产日期
			ExpDate       string `json:"expDate"`       // 有效期至
			QualityPeriod int    `json:"qualityPeriod"` // 保质期天数
			GoodsBatchNo  string `json:"goodsBatchNo"`  // 批号
		} `json:"detailList"` // 单据明细
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockCheckListResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		List []struct {
			Id             string `json:"id"`             // 盘点单id
			StoreId        string `json:"storeId"`        // 机构id
			StoreName      string `json:"storeName"`      // 机构名称
			DepotId        string `json:"depotId"`        // 仓库id
			DepotName      string `json:"depotName"`      // 仓库名称
			BussDate       string `json:"bussDate"`       // 盘点业务日期
			BillNo         string `json:"billNo"`         // 盘点单号
			CheckMode      int    `json:"checkMode"`      // 盘点类型（941日盘，942周盘，943月盘，944其他）
			IsProfLoss     int    `json:"isProfLoss"`     // 盘盈盘亏（1盘盈，0盘亏）
			Status         int    `json:"status"`         // 状态：961：暂存，963：提交，962：已审核，964：已驳回，965：已反审
			TotalCheckAmt  int    `json:"totalCheckAmt"`  // 实盘总金额
			TotalProfitAmt int    `json:"totalProfitAmt"` // 总盈亏金额
			TotalAccAmt    int    `json:"totalAccAmt"`    // 总账面金额
			CreateUserName string `json:"createUserName"` // 创建人名字
			CreateTime     string `json:"createTime"`     // 创建时间
			UpdateUserName string `json:"updateUserName"` // 更新人名字
			UpdateTime     string `json:"updateTime"`     // 更新时间
		} `json:"list"` // 盘点单列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockCheckDetailResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		Id         string `json:"id"` // 单据id
		DetailList []struct {
			GoodsId              string `json:"goodsId"`              // 物品id
			GoodsName            string `json:"goodsName"`            // 物品名字
			GoodsCode            string `json:"goodsCode"`            // 物品编码
			GoodsSpec            string `json:"goodsSpec"`            // 物品规格
			UnitId               string `json:"unitId"`               // 标准单位id
			UnitName             string `json:"unitName"`             // 标准单位名字
			CheckQty             int    `json:"checkQty"`             // 实盘数量
			AccAmt               int    `json:"accAmt"`               // 账面金额
			AccQty               int    `json:"accQty"`               // 账面数量
			ProfLossQty          int    `json:"profLossQty"`          // 盈亏数量
			ProfLossAmt          int    `json:"profLossAmt"`          // 盈亏金额
			InventoryProfitQty   int    `json:"inventoryProfitQty"`   // 盘盈数量
			InventoryProfitPrice int    `json:"inventoryProfitPrice"` // 盘盈单价
			InventoryProfitAmt   int    `json:"inventoryProfitAmt"`   // 盘盈金额
			InventoryLossQty     int    `json:"inventoryLossQty"`     // 盘亏数量
			InventoryLossPrice   int    `json:"inventoryLossPrice"`   // 盘亏单价
			InventoryLossAmt     int    `json:"inventoryLossAmt"`     // 盘亏金额
			Remarks              string `json:"remarks"`              // 备注
			DualGoodsQty         int    `json:"dualGoodsQty"`         // 辅助数量
			DualUnitName         string `json:"dualUnitName"`         // 辅助单位名字
			DualUnitId           string `json:"dualUnitId"`           // 辅助单位id
			DualProfLossQty      int    `json:"dualProfLossQty"`      // 辅助单位盈亏数
			QualityPeriodFlag    int    `json:"qualityPeriodFlag"`    // 保质期开关 0：未开启 1：已开启
			GoodsBatchFlag       int    `json:"goodsBatchFlag"`       // 批次开关 0：未开启 1:已开启
			ProDate              string `json:"proDate"`              // 生产日期
			ExpDate              string `json:"expDate"`              // 有效期至
			QualityPeriod        int    `json:"qualityPeriod"`        // 保质期天数
			GoodsBatchNo         string `json:"goodsBatchNo"`         // 批号
			CheckAmt             int    `json:"checkAmt"`             // 实盘金额
			DualAccQty           int    `json:"dualAccQty"`           // 辅助单位账面数
		} `json:"detailList"` // 单据明细
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockSaleInOutDetailResp struct {
	Code string `json:"code"` // 业务响应码 10000：正常
	Msg  string `json:"msg"`  // 响应信息
	Data []struct {
		Id         string `json:"id"` // 单据id
		DetailList []struct {
			BillId       string `json:"billId"`       // 单据id
			GoodsId      string `json:"goodsId"`      // 物品id
			GoodsName    string `json:"goodsName"`    // 物品名称
			GoodsCode    string `json:"goodsCode"`    // 物品编码
			GoodsSpec    string `json:"goodsSpec"`    // 物品规格
			GoodsQty     int    `json:"goodsQty"`     // 出库数量
			GoodsAmt     int    `json:"goodsAmt"`     // 出库金额
			UnitId       string `json:"unitId"`       // 单位id
			UnitName     string `json:"unitName"`     // 单位名称
			UnitPrice    int    `json:"unitPrice"`    // 单价
			GoodsBatchNo string `json:"goodsBatchNo"` // 批号
			Remarks      string `json:"remarks"`      // 备注
		} `json:"detailList"` // 单据物品明细
	} `json:"data"`
	Success    bool   `json:"success"`    // 是否成功, true:成功，false:失败
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockSaleOutInOutListResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：正常
	Msg     string `json:"msg"`     // 响应信息
	Success bool   `json:"success"` // 是否成功, true:成功，false:失败
	Data    struct {
		List []struct {
			Id         string `json:"id"`         // 单据id
			TenantId   string `json:"tenantId"`   // 商户id
			StoreId    string `json:"storeId"`    // 门店id
			StoreName  string `json:"storeName"`  // 门店名称
			DepotId    string `json:"depotId"`    // 仓库id
			DepotName  string `json:"depotName"`  // 仓库名称
			BussDate   string `json:"bussDate"`   // 业务日期
			TotalAmt   int    `json:"totalAmt"`   // 总金额
			Remarks    string `json:"remarks"`    // 备注
			CreateTime string `json:"createTime"` // 创建时间
			UpdateTime string `json:"updateTime"` // 修改时间
		} `json:"list"` // 单据列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 返回数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type StockOtherInOutReceiveResp struct {
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
	Success    bool   `json:"success"`    // 是否成功 true：处理成功 false：处理失败
	Code       string `json:"code"`       // 响应码
	Msg        string `json:"msg"`        // 错误信息
}

type StockPurchaseInOutReceiveResp struct {
	ServerCode int    `json:"serverCode"` // 服务编码
	Msg        string `json:"msg"`        // 结果说明
	Data       struct {
		BillNo string `json:"billNo"` // 单据编号
		Id     string `json:"id"`     // 单据id
	} `json:"data"` // 执行结果
	Success   bool   `json:"success"`   // 执行结果
	MessageId string `json:"messageId"` // 服务消息ID
}

type StockQueryResp struct {
	Code    string `json:"code"`    // 业务响应码 10000：成功
	Msg     string `json:"msg"`     // 响应消息
	Success bool   `json:"success"` // 业务执行状态 true：成功 false：失败
	Data    struct {
		List []struct {
			OrgId        string `json:"orgId"`        // 组织ID
			OrgName      string `json:"orgName"`      // 组织名称
			DepotId      string `json:"depotId"`      // 仓库ID
			DepotName    string `json:"depotName"`    // 仓库名称
			CateId       string `json:"cateId"`       // 物品类别ID
			CateName     string `json:"cateName"`     // 物品类别名称
			GoodsId      string `json:"goodsId"`      // 物品ID
			GoodsCode    string `json:"goodsCode"`    // 物品编码
			GoodsName    string `json:"goodsName"`    // 物品名称
			GoodsSpec    string `json:"goodsSpec"`    // 物品规格
			UnitId       string `json:"unitId"`       // 标准单位ID
			UnitName     string `json:"unitName"`     // 标准单位名称
			CurrQty      int    `json:"currQty"`      // 标准数量
			DualUnitId   string `json:"dualUnitId"`   // 辅助单位ID
			DualUnitName string `json:"dualUnitName"` // 辅助单位名称
			DualCurrQty  int    `json:"dualCurrQty"`  // 辅助单位数量
			AverPrice    int    `json:"averPrice"`    // 成本单价
			CurrAmt      int    `json:"currAmt"`      // 金额
		} `json:"list"` // 单据明细列表
		Total int `json:"total"` // 总条数
	} `json:"data"` // 数据
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type OrgQueryResp struct {
	Code    string `json:"code"`    // 业务响应码 10000代表成功，其他都是失败。
	Msg     string `json:"msg"`     // 接口返回说明，成功时返回SUCCESS，失败时返回具体错误信息
	Success bool   `json:"success"` // 接口是否成功，对应code的10000
	Data    []struct {
		OrgId     string `json:"orgId"`     // 机构ID，始终有值
		OrgName   string `json:"orgName"`   // 机构名称
		OrgType   int    `json:"orgType"`   // 机构类型，同入参的类型枚举
		DepotId   string `json:"depotId"`   // 仓库ID，只有查询仓库信息时才有值
		DepotCode string `json:"depotCode"` // 仓库编码
		DepotName string `json:"depotName"` // 仓库名称
		TenantId  string `json:"tenantId"`  // 商户ID
	} `json:"data"` // 接口返回数据体
	ServerCode int    `json:"serverCode"` // 服务编码
	MessageId  string `json:"messageId"`  // 服务消息ID
}

type BookQueryResponse struct {
	Data []struct {
		PeriodId   string `json:"periodId"`
		PeriodName string `json:"periodName"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	} `json:"data"`
	MessageId  string `json:"messageId"`
	ServerCode int    `json:"serverCode"`
}

type BookSaveResp struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	MessageUuid string `json:"messageUuid"`
	Result      struct {
		BookOrderNo string `json:"bookOrderNo"`
		ReqId       string `json:"reqId"`
		Code        string `json:"code"`
		Msg         string `json:"msg"`
		Success     string `json:"success"`
		MsgCode     string `json:"msgCode"`
		MsgInfo     string `json:"msgInfo"`
		MessageId   string `json:"messageId"`
		ServerCode  int    `json:"serverCode"`
	} `json:"result"`
}

type BookQueryPeriodTimeResp struct {
	Data []struct {
		PeriodId   string `json:"periodId"`
		PeriodName string `json:"periodName"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	} `json:"data"`
	MessageId  string `json:"messageId"`
	ServerCode int    `json:"serverCode"`
}

type BookQueryOrderResp struct {
	MessageId      string `json:"messageId"`
	CreateId       string `json:"createId"`
	CreatorName    string `json:"creatorName"`
	UpdateId       string `json:"updateId"`
	UpdateName     string `json:"updateName"`
	OrderNo        string `json:"orderNo"`
	DinnerTime     string `json:"dinnerTime"`
	TableNum       int    `json:"tableNum"`
	DinerNum       int    `json:"dinerNum"`
	CustomerName   string `json:"customerName"`
	CustomerSex    string `json:"customerSex"`
	CustomerPhone  string `json:"customerPhone"`
	OutBizNo       string `json:"outBizNo"`
	BusinessType   string `json:"businessType"`
	Note           string `json:"note"`
	Status         string `json:"status"`
	BookingCnt     int    `json:"bookingCnt"`
	CancelCnt      int    `json:"cancelCnt"`
	OverdueCnt     int    `json:"overdueCnt"`
	GmtCreate      string `json:"gmtCreate"`
	Source         string `json:"source"`
	SourceName     string `json:"sourceName"`
	PeriodTimeList []struct {
		PeriodId   string `json:"periodId"`
		PeriodName string `json:"periodName"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	} `json:"periodTimeList"`
	CancelReasonDto struct {
		Id            string `json:"id"`
		ReasonContent string `json:"reasonContent"`
	} `json:"cancelReasonDto"`
	OpenInfoDto struct {
		OpenSourceId   string `json:"openSourceId"`
		OpenSourceName string `json:"openSourceName"`
	} `json:"openInfoDto"`
	TableRecordList []struct {
		AreaId          string `json:"areaId"`
		AreaName        string `json:"areaName"`
		TableId         string `json:"tableId"`
		TableName       string `json:"tableName"`
		DinnerTime      string `json:"dinnerTime"`
		Status          string `json:"status"`
		BookingTimeType string `json:"bookingTimeType"`
		PeriodTimeList  []any  `json:"periodTimeList"`
	} `json:"tableRecordList"`
	CancelTime      string `json:"cancelTime"`
	BookingTimeType string `json:"bookingTimeType"`
	ServerCode      int    `json:"serverCode"`
}

type BookTbaleInfoResp struct {
	PageNum   int `json:"pageNum"`
	PageSize  int `json:"pageSize"`
	TotalNum  int `json:"totalNum"`
	TotalPage int `json:"totalPage"`
	PageList  []struct {
		TableId        string `json:"tableId"`
		TableName      string `json:"tableName"`
		AreaId         string `json:"areaId"`
		AreaName       string `json:"areaName"`
		TableStatus    string `json:"tableStatus"`
		DinersNum      int    `json:"dinersNum"`
		TableNum       int    `json:"tableNum"`
		OpenTime       string `json:"openTime"`
		Sort           int    `json:"sort"`
		TableNameIndex string `json:"tableNameIndex"`
		DiningFlag     bool   `json:"diningFlag"`
		TableTypeCode  string `json:"tableTypeCode"`
	} `json:"pageList"`
	MessageId  string `json:"messageId"`
	ServerCode int    `json:"serverCode"`
}

type BookConfirmResp struct {
	ReqId      string `json:"reqId"`
	Data       bool   `json:"data"`
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	Success    string `json:"success"`
	MsgCode    string `json:"msgCode"`
	MsgInfo    string `json:"msgInfo"`
	MessageId  string `json:"messageId"`
	ServerCode int    `json:"serverCode"`
}

type BookCancelResp struct {
	ReqId      string `json:"reqId"`
	Data       bool   `json:"data"`
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	Success    string `json:"success"`
	MsgCode    string `json:"msgCode"`
	MsgInfo    string `json:"msgInfo"`
	MessageId  string `json:"messageId"`
	ServerCode int    `json:"serverCode"`
}

type BusinessIncomePromoResp struct {
	Data struct {
		TotalSize int `json:"totalSize"`
		List      []struct {
			BrandId             int    `json:"brandId"`
			ShopName            string `json:"shopName"`
			ShopId              int    `json:"shopId"`
			AddressProvince     string `json:"addressProvince"`
			AddressProvinceName string `json:"addressProvinceName"`
			AddressCity         string `json:"addressCity"`
			AddressCityName     string `json:"addressCityName"`
			AddressArea         string `json:"addressArea"`
			AddressAreaName     string `json:"addressAreaName"`
			Date                string `json:"date"`
			OrderPromoItems     struct {
				ItemList []PromoItem `json:"itemList"`
				SubTotal string      `json:"subTotal"`
			} `json:"orderPromoItems"`
			PaymentPromoItems struct {
				ItemList []PromoItem `json:"itemList"`
				SubTotal string      `json:"subTotal"`
			} `json:"paymentPromoItems"`
			OrderExpenseItems struct {
				ItemList []PromoItem `json:"itemList"`
				SubTotal string      `json:"subTotal"`
			} `json:"orderExpenseItems"`
		} `json:"list"`
	} `json:"data"`
	Success bool   `json:"success"`
	MsgCode string `json:"msgCode"`
	MsgInfo string `json:"msgInfo"`
}

type PromoItem struct {
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	Amount     string      `json:"amount"`
	TextVal    string      `json:"textVal"`
	DefaultVal string      `json:"defaultVal"`
	ItemList   []PromoItem `json:"itemList"`
}

type BusinessIncomeResp struct {
	Data struct {
		TotalSize int `json:"totalSize"`
		List      []struct {
			BrandId                     int    `json:"brandId"`
			ShopName                    string `json:"shopName"`
			ShopId                      int    `json:"shopId"`
			AddressProvince             string `json:"addressProvince"`
			AddressProvinceName         string `json:"addressProvinceName"`
			AddressCity                 string `json:"addressCity"`
			AddressCityName             string `json:"addressCityName"`
			AddressArea                 string `json:"addressArea"`
			AddressAreaName             string `json:"addressAreaName"`
			Date                        string `json:"date"`
			SaleAmt                     string `json:"saleAmt"`
			TotalPromoAmt               string `json:"totalPromoAmt"`
			PromoAmtProportion          string `json:"promoAmtProportion"`
			BusinessIncomeAmt           string `json:"businessIncomeAmt"`
			ItemActualReceivedAmt       string `json:"itemActualReceivedAmt"`
			ExtraFeeActualAmt           string `json:"extraFeeActualAmt"`
			OrderCnt                    int    `json:"orderCnt"`
			AvgTradeAmtPreDiscount      string `json:"avgTradeAmtPreDiscount"`
			AvgTradeAmtAfterDiscount    string `json:"avgTradeAmtAfterDiscount"`
			OrderPeopleCnt              int    `json:"orderPeopleCnt"`
			AvgCustomerAmtPreDiscount   string `json:"avgCustomerAmtPreDiscount"`
			AvgCustomerAmtAfterDiscount string `json:"avgCustomerAmtAfterDiscount"`
			AvgDiningDuration           string `json:"avgDiningDuration"`
			OpenTableCnt                int    `json:"openTableCnt"`
			OpenTableRate               string `json:"openTableRate"`
			ReopenTableRate             string `json:"reopenTableRate"`
			OrderTypeItems              struct {
				ItemList []OrderTypeItem `json:"itemList"`
				SubTotal string          `json:"subTotal"`
			} `json:"orderTypeItems"`
		} `json:"list"`
	} `json:"data"`
}

// 定义 OrderTypeItem 结构体
type OrderTypeItem struct {
	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Amount     string          `json:"amount"`
	TextVal    string          `json:"textVal"`
	DefaultVal string          `json:"defaultVal"`
	ItemList   []OrderTypeItem `json:"itemList"`
}

type BusinessIncomePromoStatisticsResp struct {
	Data struct {
		Values []struct {
			FinishBusiDate     string `json:"finishBusiDate"`
			PromoName          string `json:"promoName"`
			PromoSecType       string `json:"promoSecType"`
			PromoSecTypeName   string `json:"promoSecTypeName"`
			PromoThirdTypeName string `json:"promoThirdTypeName"`
			PromoType          string `json:"promoType"`
			PromoTypeName      string `json:"promoTypeName"`
			PromoAmt           string `json:"promoAmt"`
		} `json:"values"`
		TotalSize int `json:"totalSize"`
	} `json:"data"`
	Success bool   `json:"success"`
	MsgCode string `json:"msgCode"`
	MsgInfo string `json:"msgInfo"`
}

type BusinessIncomeConstituteResp struct {
	Data struct {
		TotalSize int `json:"totalSize"`
		List      []struct {
			BrandId             int    `json:"brandId"`
			ShopName            string `json:"shopName"`
			ShopId              int    `json:"shopId"`
			AddressProvince     string `json:"addressProvince"`
			AddressProvinceName string `json:"addressProvinceName"`
			AddressCity         string `json:"addressCity"`
			AddressCityName     string `json:"addressCityName"`
			AddressArea         string `json:"addressArea"`
			AddressAreaName     string `json:"addressAreaName"`
			Date                string `json:"date"`
			BusinessIncomeItems struct {
				ItemList []BusinessIncomeItem `json:"itemList"`
				SubTotal string               `json:"subTotal"`
			} `json:"businessIncomeItems"`
		} `json:"list"`
	} `json:"data"`
	Success bool   `json:"success"`
	MsgCode string `json:"msgCode"`
	MsgInfo string `json:"msgInfo"`
}

// 定义 BusinessIncomeItem 结构体
type BusinessIncomeItem struct {
	Code       string               `json:"code"`
	Name       string               `json:"name"`
	Amount     string               `json:"amount"`
	TextVal    string               `json:"textVal"`
	DefaultVal string               `json:"defaultVal"`
	ItemList   []BusinessIncomeItem `json:"itemList"`
}

type KposLocalResp struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Data      struct {
		CartDTO struct {
			ActualFee                      int64             `json:"actualFee"`
			ActualReceive                  int64             `json:"actualReceive"`
			AfterDiscountFee               int64             `json:"afterDiscountFee"`
			AfterDiscountFeeStr            string            `json:"afterDiscountFeeStr"`
			AreaId                         string            `json:"areaId"`
			AreaName                       string            `json:"areaName"`
			AvailableVoucherNum            int               `json:"availableVoucherNum"`
			BusinessDate                   string            `json:"businessDate"`
			BusinessFormat                 string            `json:"businessFormat"`
			CartId                         string            `json:"cartId"`
			CustomerFlag                   string            `json:"customerFlag"`
			DepositFee                     int64             `json:"depositFee"`
			DeviceId                       string            `json:"deviceId"`
			DinersNum                      int               `json:"dinersNum"`
			ExtMap                         map[string]string `json:"extMap"`
			ExtraPromoFee                  int64             `json:"extraPromoFee"`
			ExtraTotalFee                  int64             `json:"extraTotalFee"`
			GmtCreate                      int64             `json:"gmtCreate"`
			GmtModified                    int64             `json:"gmtModified"`
			HangCartNum                    int               `json:"hangCartNum"`
			HasPassword                    bool              `json:"hasPassword"`
			HasPointFlag                   string            `json:"hasPointFlag"`
			HasVoucherFlag                 string            `json:"hasVoucherFlag"`
			JoinTableAttachFeeFlag         bool              `json:"joinTableAttachFeeFlag"`
			JoinTableMemberFlag            bool              `json:"joinTableMemberFlag"`
			JoinTablePrepaymentFlag        bool              `json:"joinTablePrepaymentFlag"`
			JoinTablePromoFlag             bool              `json:"joinTablePromoFlag"`
			JointTableFlag                 string            `json:"jointTableFlag"`
			JointTableInfos                []interface{}     `json:"jointTableInfos"`
			MantissaFee                    int64             `json:"mantissaFee"`
			MantissaFeeStr                 string            `json:"mantissaFeeStr"`
			ModifyCartKeyInfo              bool              `json:"modifyCartKeyInfo"`
			ModifyNumFlag                  bool              `json:"modifyNumFlag"`
			MultiTableFlag                 string            `json:"multiTableFlag"`
			OpenTime                       string            `json:"openTime"`
			OrderId                        string            `json:"orderId"`
			OrderPrintType                 []string          `json:"orderPrintType"`
			OrderStatus                    string            `json:"orderStatus"`
			OrderType                      string            `json:"orderType"`
			PayFlag                        bool              `json:"payFlag"`
			PointPromoFee                  int64             `json:"pointPromoFee"`
			PreMantissaFee                 int64             `json:"preMantissaFee"`
			PreMantissaFeeUseScope         string            `json:"preMantissaFeeUseScope"`
			PrepaymentFee                  int64             `json:"prepaymentFee"`
			PromoFee                       int64             `json:"promoFee"`
			PromoFeeStr                    string            `json:"promoFeeStr"`
			PromoSaveSource                string            `json:"promoSaveSource"`
			QueryMemberCardFlag            string            `json:"queryMemberCardFlag"`
			QueryVoucherFlag               string            `json:"queryVoucherFlag"`
			RefreshGapTime                 int               `json:"refreshGapTime"`
			ReverseFlag                    bool              `json:"reverseFlag"`
			ShopId                         string            `json:"shopId"`
			SupportMemDayPrice             bool              `json:"supportMemDayPrice"`
			SupportMemPrice                bool              `json:"supportMemPrice"`
			SupportSpecialItemWithRecharge bool              `json:"supportSpecialItemWithRecharge"`
			TableId                        string            `json:"tableId"`
			TableName                      string            `json:"tableName"`
			TogetherTableNum               int               `json:"togetherTableNum"`
			TotalFee                       int64             `json:"totalFee"`
			TotalFeeStr                    string            `json:"totalFeeStr"`
			TotalPromoFee                  int64             `json:"totalPromoFee"`
			UseMemberDayPrice              string            `json:"useMemberDayPrice"`
			UseMemberPrice                 string            `json:"useMemberPrice"`
			ValidateItemTotalFee           int64             `json:"validateItemTotalFee"`
			VerifyFlag                     bool              `json:"verifyFlag"`
			WholePackFlag                  string            `json:"wholePackFlag"`
		} `json:"cartDTO"`
		CartItemDTOList     []CartItemDTO     `json:"cartItemDTOList"`
		DcOrderBatchDTOList []DcOrderBatchDTO `json:"dcOrderBatchDTOList"`
		OrderDTO            OrderDTO          `json:"orderDTO"`
	} `json:"data"`
	RequestCloud   bool `json:"requestCloud"`
	RequestCloudRt int  `json:"requestCloudRt"`
}

type CartItemDTO struct {
	ActualFee               int64             `json:"actualFee"`
	AddSpiceTotalFee        int64             `json:"addSpiceTotalFee"`
	AdditionActualFee       int64             `json:"additionActualFee"`
	AliasName               string            `json:"aliasName"`
	BoolChangePrice         string            `json:"boolChangePrice"`
	CartItemId              string            `json:"cartItemId"`
	ChangePriceFlag         string            `json:"changePriceFlag"`
	CookAttachTotalFee      int64             `json:"cookAttachTotalFee"`
	CostPrice               int64             `json:"costPrice"`
	CustomExtMap            map[string]string `json:"customExtMap"`
	CustomerFlag            string            `json:"customerFlag"`
	DeviceId                string            `json:"deviceId"`
	DishDiscountFlag        string            `json:"dishDiscountFlag"`
	DishSkuWeight           int64             `json:"dishSkuWeight"`
	DoubleUnitWeighDishFlag string            `json:"doubleUnitWeighDishFlag"`
	GiveFlag                string            `json:"giveFlag"`
	ItemCategoryId          string            `json:"itemCategoryId"`
	ItemCategoryInfo        ItemCategoryInfo  `json:"itemCategoryInfo"`
	ItemCategoryName        string            `json:"itemCategoryName"`
	ItemGroupFee            int64             `json:"itemGroupFee"`
	ItemGroupFeeStr         string            `json:"itemGroupFeeStr"`
	ItemId                  string            `json:"itemId"`
	ItemName                string            `json:"itemName"`
	ItemNum                 int               `json:"itemNum"`
	ItemNumStr              string            `json:"itemNumStr"`
	ItemOriginGroupFee      int64             `json:"itemOriginGroupFee"`
	ItemOriginPrice         int64             `json:"itemOriginPrice"`
	ItemPrice               int64             `json:"itemPrice"`
	ItemPriceStr            string            `json:"itemPriceStr"`
	ItemType                string            `json:"itemType"`
	ModifyItemNum           int               `json:"modifyItemNum"`
	MultiSpecFlag           string            `json:"multiSpecFlag"`
	SaasItemId              string            `json:"saasItemId"`
	SaasSkuId               string            `json:"saasSkuId"`
	ServingNo               int               `json:"servingNo"`
	SkuCode                 string            `json:"skuCode"`
	SkuId                   string            `json:"skuId"`
	SkuName                 string            `json:"skuName"`
	SkuUuid                 string            `json:"skuUuid"`
	SupportDecimalFlag      string            `json:"supportDecimalFlag"`
	TotalFee                int64             `json:"totalFee"`
	TotalFeeStr             string            `json:"totalFeeStr"`
	UnitId                  string            `json:"unitId"`
	UnitName                string            `json:"unitName"`
	WeighDishFlag           string            `json:"weighDishFlag"`
}

type ItemCategoryInfo struct {
	ItemCategoryId     string `json:"itemCategoryId"`
	ItemCategoryName   string `json:"itemCategoryName"`
	ItemCategorySort   int    `json:"itemCategorySort"`
	ItemDishCategoryId string `json:"itemDishCategoryId"`
}

// ---------------- DcOrderBatchDTO ----------------
type DcOrderBatchDTO struct {
	HasFulfillBatch bool        `json:"hasFulfillBatch"`
	OrderBatch      OrderBatch  `json:"orderBatch"`
	OrderId         string      `json:"orderId"`
	OrderItemList   []OrderItem `json:"orderItemList"`
	OrderNo         string      `json:"orderNo"`
	ShopId          string      `json:"shopId"`
}

type OrderBatch struct {
	BatchCreatorName   string            `json:"batchCreatorName"`
	BatchId            string            `json:"batchId"`
	BatchOrderTime     string            `json:"batchOrderTime"`
	BatchOrderTimeDate int64             `json:"batchOrderTimeDate"`
	BatchStatus        string            `json:"batchStatus"`
	BatchType          string            `json:"batchType"`
	Deleted            int               `json:"deleted"`
	Env                string            `json:"env"`
	ExtMap             map[string]string `json:"extMap"`
	GmtCreate          int64             `json:"gmtCreate"`
	GmtModified        int64             `json:"gmtModified"`
	OrderId            string            `json:"orderId"`
	OrderNo            string            `json:"orderNo"`
	ShopId             string            `json:"shopId"`
}

type OrderItem struct {
	// 以下仅列出常用字段，可按需要继续补充
	AbolishRefundFlag        string            `json:"abolishRefundFlag"`
	ActualFee                int64             `json:"actualFee"`
	AddSpiceTotalActualFee   int64             `json:"addSpiceTotalActualFee"`
	AddSpiceTotalFee         int64             `json:"addSpiceTotalFee"`
	AdditionActualFee        int64             `json:"additionActualFee"`
	AliasName                string            `json:"aliasName"`
	BatchId                  string            `json:"batchId"`
	BoolChangePrice          string            `json:"boolChangePrice"`
	BuffetMealVoucher        bool              `json:"buffetMealVoucher"`
	BuyTime                  string            `json:"buyTime"`
	CartItemId               string            `json:"cartItemId"`
	ChangePriceFlag          string            `json:"changePriceFlag"`
	CookAttachTotalActualFee int64             `json:"cookAttachTotalActualFee"`
	CookAttachTotalFee       int64             `json:"cookAttachTotalFee"`
	CostPrice                int64             `json:"costPrice"`
	CustomExtMap             map[string]string `json:"customExtMap"`
	CustomerFlag             string            `json:"customerFlag"`
	Deleted                  int               `json:"deleted"`
	DishDiscountFlag         string            `json:"dishDiscountFlag"`
	DishSkuWeight            int64             `json:"dishSkuWeight"`
	DisplayNo                string            `json:"displayNo"`
	DoubleUnitWeighDish      bool              `json:"doubleUnitWeighDish"`
	DoubleUnitWeighDishFlag  string            `json:"doubleUnitWeighDishFlag"`
	Env                      string            `json:"env"`
	FinishedDishFlag         string            `json:"finishedDishFlag"`
	GiveFlag                 string            `json:"giveFlag"`
	GmtCreate                int64             `json:"gmtCreate"`
	GmtModified              int64             `json:"gmtModified"`
	GoodsOriginType          string            `json:"goodsOriginType"`
	InventoryNum             int               `json:"inventoryNum"`
	ItemCategoryId           string            `json:"itemCategoryId"`
	ItemCategoryInfo         ItemCategoryInfo  `json:"itemCategoryInfo"`
	ItemCategoryName         string            `json:"itemCategoryName"`
	ItemFulfillStatus        string            `json:"itemFulfillStatus"`
	ItemGroupFee             int64             `json:"itemGroupFee"`
	ItemGroupRefundFee       int64             `json:"itemGroupRefundFee"`
	ItemId                   string            `json:"itemId"`
	ItemName                 string            `json:"itemName"`
	ItemNum                  int               `json:"itemNum"`
	ItemNumStr               string            `json:"itemNumStr"`
	ItemOriginGroupFee       int64             `json:"itemOriginGroupFee"`
	ItemOriginPrice          int64             `json:"itemOriginPrice"`
	ItemPrice                int64             `json:"itemPrice"`
	ItemSortIndex            int               `json:"itemSortIndex"`
	ItemType                 string            `json:"itemType"`
	JoinTableUniKey          string            `json:"joinTableUniKey"`
	LastOperateUserId        string            `json:"lastOperateUserId"`
	ManualDiscountReason     string            `json:"manualDiscountReason"`
	MultiSpecFlag            string            `json:"multiSpecFlag"`
	MustOrderFlag            string            `json:"mustOrderFlag"`
	NetDishFlag              string            `json:"netDishFlag"`
	OpenTableDishFlag        string            `json:"openTableDishFlag"`
	OrderAddSpiceList        []interface{}     `json:"orderAddSpiceList"`
	OrderId                  string            `json:"orderId"`
	OrderItemId              string            `json:"orderItemId"`
	OrderItemOperateList     []interface{}     `json:"orderItemOperateList"`
	OrderItemPackageList     []interface{}     `json:"orderItemPackageList"`
	OrderNo                  string            `json:"orderNo"`
	OrderPackingBoxList      []interface{}     `json:"orderPackingBoxList"`
	PackageFee               int64             `json:"packageFee"`
	PackingBoxTotalFee       int64             `json:"packingBoxTotalFee"`
	PromoFee                 int64             `json:"promoFee"`
	RefundFlag               string            `json:"refundFlag"`
	RequestSource            string            `json:"requestSource"`
	RowId                    string            `json:"rowId"`
	SaasItemId               string            `json:"saasItemId"`
	SaasSkuId                string            `json:"saasSkuId"`
	ServingNo                int               `json:"servingNo"`
	ShopId                   string            `json:"shopId"`
	SideDishPricingMode      string            `json:"sideDishPricingMode"`
	SingleItemGroupFee       int64             `json:"singleItemGroupFee"`
	SingleItemNum            int               `json:"singleItemNum"`
	SinglePackFlag           string            `json:"singlePackFlag"`
	SkuCode                  string            `json:"skuCode"`
	SkuId                    string            `json:"skuId"`
	SkuName                  string            `json:"skuName"`
	SkuSaleDisplayFlag       string            `json:"skuSaleDisplayFlag"`
	SkuUuid                  string            `json:"skuUuid"`
	SpecList                 []Spec            `json:"specList"`
	SpuName                  string            `json:"spuName"`
	StepNum                  int               `json:"stepNum"`
	SupportDecimalFlag       string            `json:"supportDecimalFlag"`
	TempDishFlag             string            `json:"tempDishFlag"`
	TimePeriod               bool              `json:"timePeriod"`
	TimePeriodFlag           string            `json:"timePeriodFlag"`
	TotalFee                 int64             `json:"totalFee"`
	TotalFeeStr              string            `json:"totalFeeStr"`
	UnitId                   string            `json:"unitId"`
	UnitName                 string            `json:"unitName"`
	UrgeDishFlag             string            `json:"urgeDishFlag"`
	WaitCallState            string            `json:"waitCallState"`
	WaiterId                 string            `json:"waiterId"`
	WaiterName               string            `json:"waiterName"`
	WeighDishFlag            string            `json:"weighDishFlag"`
}

type Spec struct {
	SpecId   string `json:"specId"`
	SpecName string `json:"specName"`
}

// ---------------- OrderDTO ----------------
// 这里仅列出少量常用字段，可按需要继续补全
type OrderDTO struct {
	ActualFee        int64  `json:"actualFee"`
	AfterDiscountFee int64  `json:"afterDiscountFee"`
	AreaId           string `json:"areaId"`
	AreaName         string `json:"areaName"`
	BusinessDate     string `json:"businessDate"`
	DeviceId         string `json:"deviceId"`
	DinersNum        int    `json:"dinnersNum"`
	OrderId          string `json:"orderId"`
	OrderNo          string `json:"orderNo"`
	OrderStatus      string `json:"orderStatus"`
	PayStatus        string `json:"payStatus"`
	ShopId           string `json:"shopId"`
	TableId          string `json:"tableId"`
	TableName        string `json:"tableName"`
	TotalFee         int64  `json:"totalFee"`
	// ... 其余字段可按业务补全
}
