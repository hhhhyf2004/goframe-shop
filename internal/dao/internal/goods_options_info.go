// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsOptionsInfoDao is the data access object for the table goods_options_info.
type GoodsOptionsInfoDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  GoodsOptionsInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// GoodsOptionsInfoColumns defines and stores column names for the table goods_options_info.
type GoodsOptionsInfoColumns struct {
	Id        string //
	GoodsId   string // 商品id
	PicUrl    string // 图片
	Name      string // 商品名称
	Price     string // 价格 单位分
	Stock     string // 库存
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// goodsOptionsInfoColumns holds the columns for the table goods_options_info.
var goodsOptionsInfoColumns = GoodsOptionsInfoColumns{
	Id:        "id",
	GoodsId:   "goods_id",
	PicUrl:    "pic_url",
	Name:      "name",
	Price:     "price",
	Stock:     "stock",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewGoodsOptionsInfoDao creates and returns a new DAO object for table data access.
func NewGoodsOptionsInfoDao(handlers ...gdb.ModelHandler) *GoodsOptionsInfoDao {
	return &GoodsOptionsInfoDao{
		group:    "default",
		table:    "goods_options_info",
		columns:  goodsOptionsInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GoodsOptionsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GoodsOptionsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GoodsOptionsInfoDao) Columns() GoodsOptionsInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GoodsOptionsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GoodsOptionsInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GoodsOptionsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
