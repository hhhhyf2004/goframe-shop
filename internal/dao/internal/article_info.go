// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleInfoDao is the data access object for the table article_info.
type ArticleInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ArticleInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ArticleInfoColumns defines and stores column names for the table article_info.
type ArticleInfoColumns struct {
	Id        string //
	UserId    string // 作者id
	Title     string // 标题
	Desc      string // 摘要
	PicUrl    string // 封面图
	IsAdmin   string // 1后台管理员发布 2前台用户发布
	Praise    string // 点赞数
	Detail    string // 文章详情
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// articleInfoColumns holds the columns for the table article_info.
var articleInfoColumns = ArticleInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	Title:     "title",
	Desc:      "desc",
	PicUrl:    "pic_url",
	IsAdmin:   "is_admin",
	Praise:    "praise",
	Detail:    "detail",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewArticleInfoDao creates and returns a new DAO object for table data access.
func NewArticleInfoDao(handlers ...gdb.ModelHandler) *ArticleInfoDao {
	return &ArticleInfoDao{
		group:    "default",
		table:    "article_info",
		columns:  articleInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ArticleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ArticleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ArticleInfoDao) Columns() ArticleInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ArticleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ArticleInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ArticleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
