package response

import (
	"ocs-app/global"
	"strconv"
)

type PageQuery struct {
	PerPage int    `form:"pageSize" json:"pageSize" query:"pageSize"`
	Current int    `form:"current" json:"current" query:"current"`
	Sort    string `form:"sort" json:"sort" query:"sort"`
	Search  string `form:"search" json:"search" query:"search"`
	Keys    map[string]interface{}
}

type PageResponse struct {
	PerPage    int   `json:"pageSize"`
	Current    int   `json:"current"`
	PageSize   int   `json:"pagesize"`
	TotalRows  int64 `json:"totalrows"`
	TotalPages int64 `json:"totalpages,omitempty"`
}

func QueryPages(urlArgs map[string][]string) PageQuery {
	var perPage = global.AppSetting.DefaultPageSize
	var current = 1
	var sort = "id desc"
	var search = ""
	var keys = map[string]interface{}{}

	for k, v := range urlArgs {
		switch k {
		case "pageSize":
			perPage = defaultNum(v, perPage).(int)
		case "current":
			current = defaultNum(v, 1).(int)
		case "sort":
			if len(sort) == 0 {
				sort = "id desc"
				break
			}
			sort = v[0]
			if sort[0] == '-' {
				sort = sort[1:] + " desc"
			}
		case "search":
			if len(v) > 0 {
				search = v[0]
			}
		default:
			if len(v) > 0 {
				keys[k] = v[0]
			}
		}
	}

	return PageQuery{
		PerPage: perPage,
		Current: current,
		Sort:    sort,
		Search:  search,
		Keys:    keys,
	}
}

func defaultNum(arg []string, d interface{}) (r interface{}) {
	var err error
	if len(arg) == 0 {
		r = d
	} else {
		perS := arg[0]
		r, err = strconv.Atoi(perS)
		if err != nil {
			r = d
		}
	}
	return r
}
