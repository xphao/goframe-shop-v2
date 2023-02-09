// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-shop-v2/internal/dao/internal"
)

// internalOrderGoodsCommentsInfoDao is internal type for wrapping internal DAO implements.
type internalOrderGoodsCommentsInfoDao = *internal.OrderGoodsCommentsInfoDao

// orderGoodsCommentsInfoDao is the data access object for table order_goods_comments_info.
// You can define custom methods on it to extend its functionality as you wish.
type orderGoodsCommentsInfoDao struct {
	internalOrderGoodsCommentsInfoDao
}

var (
	// OrderGoodsCommentsInfo is globally public accessible object for table order_goods_comments_info operations.
	OrderGoodsCommentsInfo = orderGoodsCommentsInfoDao{
		internal.NewOrderGoodsCommentsInfoDao(),
	}
)

// Fill with you ideas below.