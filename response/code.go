package response

var (
	Success = NewError(0, "成功")

	InputArgsError    = NewError(10001, "输入参数错误")
	InsertSqlError    = NewError(10002, "插入数据错误")
	UpdateSqlError    = NewError(10003, "修改数据错误")
	DuplicateError    = NewError(10004, "插入数据禁止重复")
	PathPramError     = NewError(10005, "路径参数错误")
	SelectError       = NewError(10006, "查询数据错误")
	NotFoundDataError = NewError(10007, "表内没有数据")

	UnknowError              = NewError(10201, "未知错误")
	SqlError                 = NewError(10202, "数据库错误")
	UnauthorizedTokenTimeout = NewError(10203, "token失效")
	ParamBindError           = NewError(10204, "参数绑定有误")
	MiddleError              = NewError(10205, "中间件错误")
)
