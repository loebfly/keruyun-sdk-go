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
			DishId     string `json:"dishId"`     // 菜品ID
			DishName   string `json:"dishName"`   // 菜品名称
			DishDesc   string `json:"dishDesc"`   // 菜品描述
			CategoryId string `json:"categoryId"` // 分类ID
			Sort       int    `json:"sort"`       // 排序值
			HelpCode   string `json:"helpCode"`   // 助记码
			DishType   string `json:"dishType"`   // 菜品类型， SINGLE：单菜 ，COMBO：套餐， SIDE：配料
			State      string `json:"state"`      // 菜品状态
			WeighFlag  string `json:"weighFlag"`  // 称重菜标识, Y:是，N:否
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
