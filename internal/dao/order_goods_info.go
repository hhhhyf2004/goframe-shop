// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"goframe-shop/internal/dao/internal"
)

// orderGoodsInfoDao is the data access object for the table order_goods_info.
// You can define custom methods on it to extend its functionality as needed.
type orderGoodsInfoDao struct {
	*internal.OrderGoodsInfoDao
}

var (
	// OrderGoodsInfo is a globally accessible object for table order_goods_info operations.
	OrderGoodsInfo = orderGoodsInfoDao{internal.NewOrderGoodsInfoDao()}
)

// Add your custom methods and functionality below.
