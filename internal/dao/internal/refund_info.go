// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RefundInfoDao is the data access object for the table refund_info.
type RefundInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RefundInfoColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RefundInfoColumns defines and stores column names for the table refund_info.
type RefundInfoColumns struct {
	Id        string // 售后退款表
	Number    string // 售后订单号
	OrderId   string // 订单id
	GoodsId   string // 要售后的商品id
	Reason    string // 退款原因
	Status    string // 状态 1待处理 2同意退款 3拒绝退款
	UserId    string // 用户id
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// refundInfoColumns holds the columns for the table refund_info.
var refundInfoColumns = RefundInfoColumns{
	Id:        "id",
	Number:    "number",
	OrderId:   "order_id",
	GoodsId:   "goods_id",
	Reason:    "reason",
	Status:    "status",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewRefundInfoDao creates and returns a new DAO object for table data access.
func NewRefundInfoDao(handlers ...gdb.ModelHandler) *RefundInfoDao {
	return &RefundInfoDao{
		group:    "default",
		table:    "refund_info",
		columns:  refundInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RefundInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RefundInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RefundInfoDao) Columns() RefundInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RefundInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RefundInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RefundInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
