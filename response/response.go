package response

// ApiResponse
type ApiResponse struct {
	ID      string       `json:"id,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
	Status  bool         `json:"success"`
	Code    int          `json:"error_code"`       // 业务编码
	Data    interface{}  `json:"data"`             // 成功时返回的数据
	Message string       `json:"errmsg,omitempty"` // 错误描述
	Page    PageResponse `json:"page,omitempty"`
}

type Playload struct {
	Total   *int64 `json:"total,omitempty"`
	Current int64  `json:"current,omitempty"`
	PerPage int    `json:"pageSize,omitempty"`
	Size    *int   `json:"pagecount,omitempty"`
}

func ResponseSuccess() *ApiResponse {
	res := &ApiResponse{
		Status: true,
		Code:   0,
	}
	return res
}

// @description	返回错误信息并包含data的JSON信息,一般用于展示列表和详情
func NewResponseData(err *Error, data interface{}) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}
	return &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Data:    data,
		Message: err.Msg(),
	}
}

// @description	返回错误的JSON信息，自带错误提示msg包含具体错误信息,
func ErrorResponse(err *Error, message string) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}
	return &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Message: err.Msg() + " : " + message,
	}
}

func PageResponseData(payload PageResponse, data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: true,
		Code:   0,
		Data:   data,
		Page:   payload,
	}
}
