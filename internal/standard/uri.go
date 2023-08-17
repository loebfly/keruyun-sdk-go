package standard

const (
	/*
		门店相关
	*/

	// UriShopQueryStoreDetail 根据门店id获取门店详情 POST
	UriShopQueryStoreDetail = "/open/kry_standard/shop/MerchantOrgReadService/queryById"
	// UriShopQueryTableList 获取门店桌台列表 POST
	UriShopQueryTableList = "/open/kry_standard/shop/ShopTableClient/pageQuery"
	// UriShopQueryServiceFeeList 查询门店服务费列表 POST
	UriShopQueryServiceFeeList = "/open/kry_standard/shop/ShopExtraChargeQueryClient/listQuery"
	// UriShopQueryTableInfo 根据桌台ID获取桌台信息 POST
	UriShopQueryTableInfo = "/open/kry_standard/shop/ShopTableClient/getShopTable"
	// UriShopQueryServiceFeeInfo 根据服务费ID获取服务费详情 POST
	UriShopQueryServiceFeeInfo = "/open/kry_standard/shop/ShopExtraChargeQueryClient/getExtraChargeById"

	/*
		菜品相关
	*/
	// UriDishQueryPage 门店分页查询菜品基本信息 POST
	UriDishQueryPage = "/open/kry_standard/dish/shop/pageQueryBaseDish"
	// UriDishQueryDetail 门店查询菜品详情 POST
	UriDishQueryDetail = "/open/kry_standard/dish/shop/listQueryDetailDish"
	// UriDishQueryCategory 门店查询分类信息 POST
	UriDishQueryCategory = "/open/kry_standard/dish/shop/listQueryCategory"

	/*
		会员相关
	*/
	// UriCrmCreate 会员创建 POST
	UriCrmCreate = "/open/kry_standard/crm/CustomerOpenFacade/createCustomer"
	// UriCrmQueryByMobile 根据手机号查询会员 POST
	UriCrmQueryByMobile = "/open/kry_standard/crm/CustomerOpenFacade/queryByMobile"
	// UriCrmQueryProperty 会员资产查询 POST
	UriCrmQueryProperty = "/open/kry_standard/crm/CustomerOpenFacade/queryCustomerProperty"
	// UriCrmDirectCharge 储值充值接口 POST
	UriCrmDirectCharge = "/open/kry_standard/crm/RechargeOpenFacade/directCharge"

	/*
		订单相关
	*/
	// UriOrderQueryDetail 订单详情查询接口 POST
	UriOrderQueryDetail = "/open/kry_standard/order/queryDetail"
	// UriOrderQueryList 订单列表查询接口 POST
	UriOrderQueryList = "/open/kry_standard/order/queryList"

	/*
		供应链2.0
	*/
	// UriStockLossOutQueryList 报损出库单列表查询 POST
	UriStockLossOutQueryList = "/open/kry_standard/stock/lossOut/findByCondition"
	// UriStockLossOutQueryDetail 报损出库单详情 POST
	UriStockLossOutQueryDetail = "/open/kry_standard/stock/lossOut/findDetailByBillId"
	// UriStockOtherInQueryDetail 其他入库单详情 POST
	UriStockOtherInQueryDetail = "/open/kry_standard/stock/otherIn/findDetailByBillId"
	// UriStockOtherInQueryList 其他入库单列表查询 POST
	UriStockOtherInQueryList = "/open/kry_standard/stock/otherIn/findByCondition"
	// UriStockOtherOutQueryList 其他出库单列表查询 POST
	UriStockOtherOutQueryList = "/open/kry_standard/stock/otherOut/findByCondition"
	// UriStockOtherOutQueryDetail 其他出库单详情 POST
	UriStockOtherOutQueryDetail = "/open/kry_standard/stock/otherOut/findDetailByBillId"
	// UriStockTransferInternalQueryList 组织内调拨订单列表查询 POST
	UriStockTransferInternalQueryList = "/open/kry_standard/stock/transfer/internal/findByCondition"
	// UriStockTransferInternalQueryDetail 组织内调拨订单详情 POST
	UriStockTransferInternalQueryDetail = "/open/kry_standard/stock/transfer/internal/findDetailByBillId"
	// UriStockCheckQueryList 盘点单列表查询 POST
	UriStockCheckQueryList = "/open/kry_standard/stock/check/findByCondition"
	// UriStockCheckQueryDetail 盘点单详情 POST
	UriStockCheckQueryDetail = "/open/kry_standard/stock/check/findDetailByBillId"
	// UriStockTransferOutQueryDetail 组织间调拨出库单详情 POST
	UriStockTransferOutQueryDetail = "/open/kry_standard/stock/transfer/out/findDetailByBillId"
	// UriStockTransferOutStoreToStoreQueryList 组织间调拨出库单列表查询 POST
	UriStockTransferOutStoreToStoreQueryList = "/open/kry_standard/stock/transfer/out/store-to-store/findByCondition"
	// UriStockTransferInStoreToStoreQueryList 组织间调拨入库单列表 POST
	UriStockTransferInStoreToStoreQueryList = "/open/kry_standard/stock/transfer/in/store-to-store/findByCondition"
	// UriStockTransferInQueryDetail 组织间调拨入库单详情 POST
	UriStockTransferInQueryDetail = "/open/kry_standard/stock/transfer/in/findDetailByBillId"
	// UriStockSaleOutQueryDetail 消耗出库单详情 POST
	UriStockSaleOutQueryDetail = "/open/kry_standard/stock/sale-out/findDetailByBillId"
	// UriStockSaleOutQueryList 消耗出库单列表查询 POST
	UriStockSaleOutQueryList = "/open/kry_standard/stock/stock/sale-out/list"
	// UriStockOtherInReceive 接收其他入库单 POST
	UriStockOtherInReceive = "/open/kry_standard/stock/otherIn/audit/add"
	// UriStockOtherOutReceive 接收其他出库单 POST
	UriStockOtherOutReceive = "/open/kry_standard/stock/otherOut/audit/add"
	// UriStockPurchaseInReceive 接收采购入库(门店) POST
	UriStockPurchaseInReceive = "/open/kry_standard/purchase/in/store/audit"
	// UriStockPurchaseBackReceive 接收采购退货出库（门店）POST
	UriStockPurchaseBackReceive = "/open/kry_standard/purchase/back/store/auditAPI"
	// UriStockQuery 库存信息查询 POST
	UriStockQuery = "/open/kry_standard/stock/goods/findByCondition"
	// UriOrgQuery 查询组织机构信息 POST
	UriOrgQuery = "/open/kry_standard/basic/supply/prodbase/orgInfoCommon/simple-info"
)
