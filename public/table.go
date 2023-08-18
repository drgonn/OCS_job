package public

import (
	"ocs-app/model"
)

type AllM interface {
	model.Stock
}

// Arg
type Arg struct {
	NotNull  bool
	Post     []string
	Put      []string
	List     []string
	Sort     bool
	Length   int
	Name     string
	JsonName string
	Type     string
	Mean     string
	About    string
	Default  string
}

type Table[MODEL AllM] struct {
	Struct    model.BaseDAO
	Model     MODEL
	Name      string
	Mean      string
	About     string
	UrlPrefix string
	Index     string
	Args      []Arg

	// 列表查询时的字段，如" name, about, time ",当需要指定查询字段时使用
	ListArgs string

	// selfCreate func(c *gin.Context, db *sqlx.DB) (response.Error, string)

}

// 多对多数据结构体
type Many struct {
	TableName string
	Relation  string
	Index     string
}

var (
	StockC *Table[model.Stock]
)

func TableInit() {
	StockC = parseModel2Table(&model.Stock{}, model.Stock{}, "stock", "stock", "id")

}
