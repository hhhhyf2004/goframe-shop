package backend

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:500#分页数量最大50条" dc:"分页数量，最大500"`
}

type CommonPaginationRes struct {
	List  any `dc:"列表数据"`
	Total int `dc:"总数"`
}
