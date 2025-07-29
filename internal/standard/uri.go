package standard

const (
	/*
		门店相关
	*/

	// UriShopQueryStoreDetail 根据门店id获取门店详情 POST
	UriShopQueryStoreDetail = "/open/standard/shop/MerchantOrgReadService/queryById"
	// UriShopQueryTableList 获取门店桌台列表 POST
	UriShopQueryTableList = "/open/standard/shop/ShopTableClient/pageQuery"
	// UriShopQueryServiceFeeList 查询门店服务费列表 POST
	UriShopQueryServiceFeeList = "/open/standard/shop/ShopExtraChargeQueryClient/listQuery"
	// UriShopQueryTableInfo 根据桌台ID获取桌台信息 POST
	UriShopQueryTableInfo = "/open/standard/shop/ShopTableClient/getShopTable"
	// UriShopQueryServiceFeeInfo 根据服务费ID获取服务费详情 POST
	UriShopQueryServiceFeeInfo = "/open/standard/shop/ShopExtraChargeQueryClient/getExtraChargeById"

	/*
		菜品相关
	*/
	// UriDishQueryPage 门店分页查询菜品基本信息 POST
	UriDishQueryPage = "/open/standard/dish/shop/pageQueryBaseDish"
	// UriDishQueryDetail 门店查询菜品详情 POST
	UriDishQueryDetail = "/open/standard/dish/shop/listQueryDetailDish"
	// UriDishQueryCategory 门店查询分类信息 POST
	UriDishQueryCategory = "/open/standard/dish/shop/listQueryCategory"

	/*
		会员相关
	*/
	// UriCrmCreate 会员创建 POST
	UriCrmCreate = "/open/standard/crm/CustomerOpenFacade/createCustomer"
	// UriCrmCreate 会员创建 POST
	UriCrmUpdate = "/open/standard/crm/CustomerOpenFacade/updateCustomer"
	// UriCrmQueryByMobile 根据手机号查询会员 POST
	UriCrmQueryByMobile = "/open/standard/crm/CustomerOpenFacade/queryByMobile"
	// UriCrmQueryByMobile 根据手机号查询会员 POST
	UriCrmQueryByIds = "/open/standard/crm/CustomerQueryFacade/queryByCustomerIds"
	// UriCrmQueryProperty 会员资产查询 POST
	UriCrmQueryProperty = "/open/standard/crm/CustomerOpenFacade/queryCustomerProperty"
	// UriCrmDirectCharge 储值充值接口 POST
	UriCrmDirectCharge = "/open/standard/crm/RechargeOpenFacade/directCharge"
	// UriCrmTemplateList
	UriCrmTemplateList = "/open/standard/crm/OpenVoucherTemplateQueryFacade/page"
	// UriCrmTemplateInfo
	UriCrmTemplateInfo = "/open/standard/crm/OpenVoucherTemplateQueryFacade/queryOne"
	// UriCrmTemplateShop
	UriCrmTemplateShopList = "/open/standard/crm/OpenVoucherTemplateShopQueryFacade/page"
	// UriCrmTemplateDishList
	UriCrmTemplateDishList = "/open/standard/crm/OpenVoucherTemplateItemQueryFacade/page"
	// UriCrmCouponSend
	UriCrmCouponSend = "/open/standard/crm/OpenVoucherInstanceOperateFacade/send"
	//UriCrmCouponInvalid
	UriCrmCouponInvalid = "/open/standard/crm/OpenVoucherInstanceOperateFacade/invalid"
	// UriCrmCouponQuery
	UriCrmCouponQueryList = "/open/standard/crm/OpenVoucherInstanceQueryFacade/page"
	// UriCrmCouponQueryInfo
	UriCrmCouponQueryInfo = "/open/standard/crm/OpenVoucherInstanceQueryFacade/queryOne"
	/*
		订单相关
	*/
	// UriOrderQueryDetail 订单详情查询接口 POST
	UriOrderQueryDetail = "/open/standard/order/queryDetail"
	// UriOrderQueryList 订单列表查询接口 POST
	UriOrderQueryList = "/open/standard/order/queryList"

	UriKposLocalOrderDetail = "/open/standard/kposlocal/K003004"

	/*
		供应链2.0
	*/
	// UriStockLossOutQueryList 报损出库单列表查询 POST
	UriStockLossOutQueryList = "/open/standard/stock/lossOut/findByCondition"
	// UriStockLossOutQueryDetail 报损出库单详情 POST
	UriStockLossOutQueryDetail = "/open/standard/stock/lossOut/findDetailByBillId"
	// UriStockOtherInQueryDetail 其他入库单详情 POST
	UriStockOtherInQueryDetail = "/open/standard/stock/otherIn/findDetailByBillId"
	// UriStockOtherInQueryList 其他入库单列表查询 POST
	UriStockOtherInQueryList = "/open/standard/stock/otherIn/findByCondition"
	// UriStockOtherOutQueryList 其他出库单列表查询 POST
	UriStockOtherOutQueryList = "/open/standard/stock/otherOut/findByCondition"
	// UriStockOtherOutQueryDetail 其他出库单详情 POST
	UriStockOtherOutQueryDetail = "/open/standard/stock/otherOut/findDetailByBillId"
	// UriStockTransferInternalQueryList 组织内调拨订单列表查询 POST
	UriStockTransferInternalQueryList = "/open/standard/stock/transfer/internal/findByCondition"
	// UriStockTransferInternalQueryDetail 组织内调拨订单详情 POST
	UriStockTransferInternalQueryDetail = "/open/standard/stock/transfer/internal/findDetailByBillId"
	// UriStockCheckQueryList 盘点单列表查询 POST
	UriStockCheckQueryList = "/open/standard/stock/check/findByCondition"
	// UriStockCheckQueryDetail 盘点单详情 POST
	UriStockCheckQueryDetail = "/open/standard/stock/check/findDetailByBillId"
	// UriStockTransferOutQueryDetail 组织间调拨出库单详情 POST
	UriStockTransferOutQueryDetail = "/open/standard/stock/transfer/out/findDetailByBillId"
	// UriStockTransferOutStoreToStoreQueryList 组织间调拨出库单列表查询 POST
	UriStockTransferOutStoreToStoreQueryList = "/open/standard/stock/transfer/out/store-to-store/findByCondition"
	// UriStockTransferInStoreToStoreQueryList 组织间调拨入库单列表 POST
	UriStockTransferInStoreToStoreQueryList = "/open/standard/stock/transfer/in/store-to-store/findByCondition"
	// UriStockTransferInQueryDetail 组织间调拨入库单详情 POST
	UriStockTransferInQueryDetail = "/open/standard/stock/transfer/in/findDetailByBillId"
	// UriStockSaleOutQueryDetail 消耗出库单详情 POST
	UriStockSaleOutQueryDetail = "/open/standard/stock/sale-out/findDetailByBillId"
	// UriStockSaleOutQueryList 消耗出库单列表查询 POST
	UriStockSaleOutQueryList = "/open/standard/stock/stock/sale-out/list"
	// UriStockOtherInReceive 接收其他入库单 POST
	UriStockOtherInReceive = "/open/standard/stock/otherIn/audit/add"
	// UriStockOtherOutReceive 接收其他出库单 POST
	UriStockOtherOutReceive = "/open/standard/stock/otherOut/audit/add"
	// UriStockPurchaseInReceive 接收采购入库(门店) POST
	UriStockPurchaseInReceive = "/open/standard/purchase/in/store/audit"
	// UriStockPurchaseBackReceive 接收采购退货出库（门店）POST
	UriStockPurchaseBackReceive = "/open/standard/purchase/back/store/auditAPI"
	// UriStockQuery 库存信息查询 POST
	UriStockQuery = "/open/standard/stock/goods/findByCondition"
	// UriOrgQuery 查询组织机构信息 POST
	UriOrgQuery = "/open/standard/basic/supply/prodbase/orgInfoCommon/simple-info"

	/*
		预定的
	*/
	// UriBookSave 预定单保存 POST
	UriBookSave = "/open/standard/book/save"
	// UriBookQueryShopBookPeriodTime 获取预订营业日信息 POST
	UriBookQueryShopBookPeriodTime = "/open/standard/book/queryShopBookPeriodTime"
	// UriBookQuery 预订单详情查询 POST
	UriBookQuery = "/open/standard/book/queryBookOrderDetailByOrderNo"
	// UriBookQueryBookTableInfo 查询桌台信息与预订单聚合信息 POST
	UriBookQueryBookTableInfo = "/open/standard/book/queryBookingTableInfo"
	// UriBookConfirm 预订到店 POST
	UriBookConfirm = "/open/standard/book/confirmArrival"
	// UriBookCancel 取消预订 POST
	UriBookCancel = "/open/standard/book/cancel"

	// /*
	// 	数据报表
	// */
	// UriBusinessIncomePromoList 优惠构成统计 POST
	UriBusinessIncomePromoList = "/open/standard/report/business/income/promo/v3/list"
	// UriBusinessIncomeList 店内营收统计 POST
	UriBusinessIncomeList = "/open/standard/report/business/income/v3/list"
	// UriBusinessIncomePromoStatistics 收入优惠统计报表
	UriBusinessIncomePromoStatistics = "/open/standard/report/income/promo/statistics"
	// UriBusinessIncomeConstituteList 收入构成统计
	UriBusinessIncomeConstituteList = "/open/standard/report/business/income/constitute/v3/list"
	// UriPaidIncomeList 营业收款统计
	UriPaidIncomeList = "/open/standard/report/paid/income/v6/list"
	// UriPaymentReconciliationList 支付结算统计
	UriPaymentReconciliationList = "/open/standard/report/payment/reconciliation/v4/list"
	// UriPaymethodStatistics 支付方式收款统计
)
