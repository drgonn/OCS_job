package public

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"ocs-app/model"
	"ocs-app/response"
)

func (t *Table[MODEL]) Create(c *gin.Context) {
	entity := t.Model
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.InputArgsError, err.Error()))

		return
	}
	fmt.Println("update3:", entity)

	err = t.VerifyPostArg(entity)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.InputArgsError, err.Error()))

		return
	}

	fmt.Println(entity)

	err = t.Struct.Create(&entity)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.InsertSqlError, err.Error()))

		return
	}

	c.JSON(200, response.NewResponseData(response.Success, entity))

}

func (t *Table[MODEL]) List(c *gin.Context) {
	// 提取path参数
	urlArgs := c.Request.URL.Query()
	var entity []MODEL
	var count int64
	var totalPages int64
	pageQuery := response.QueryPages(urlArgs)
	keys := map[string]interface{}{}
	keyOpts := map[string]interface{}{}

	fmt.Println("keys", pageQuery.Keys)
	dao := model.DAOOption{
		Order: pageQuery.Sort,
		Where: pageQuery.Keys,
	}

	err := t.Struct.CountWithKeys(&entity, &count, keys, keyOpts, dao)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.SelectError, err.Error()))
		return
	}

	err = t.Struct.SearchByPagesWithKeys(&entity, keys, keyOpts, pageQuery.Current, pageQuery.PerPage, dao)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.SelectError, err.Error()))
		return
	}

	var data interface{}
	if len(entity) <= 0 {
		data = []struct{}{}
	} else {
		data = entity
	}

	if count%int64(pageQuery.PerPage) != 0 {
		totalPages = count/int64(pageQuery.PerPage) + 1
	} else {
		totalPages = count / int64(pageQuery.PerPage)
	}

	pageResponse := response.PageResponse{
		PerPage:    pageQuery.PerPage,
		PageSize:   pageQuery.PerPage,
		Current:    pageQuery.Current,
		TotalRows:  count,
		TotalPages: totalPages,
	}
	c.JSON(200, response.PageResponseData(pageResponse, data))
}

func (t *Table[MODEL]) Update(c *gin.Context) {
	var err error
	entity := t.Model
	id := c.Param("id")
	if len(id) <= 0 {
		c.JSON(200, response.ErrorResponse(response.PathPramError, err.Error()))
		return
	}
	id_unit64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.PathPramError, err.Error()))
		return
	}

	err = c.ShouldBindJSON(&entity)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.ParamBindError, err.Error()))
		return
	}

	err = t.VerifyPatchArg(entity, id)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.InputArgsError, err.Error()))

		return
	}

	affectNum, err := t.Struct.UpdatePatch(&entity, id_unit64)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.UpdateSqlError, err.Error()))
		return
	}

	c.JSON(200, response.NewResponseData(response.Success, affectNum))

}

func (t *Table[MODEL]) Get(c *gin.Context) {
	id := c.Param("id")
	entity := t.Model
	keys := map[string]interface{}{t.Index: id}

	num, err := t.Struct.FindByKeys(&entity, keys)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.SelectError, err.Error()))
		return
	}
	if num < 1 {
		c.JSON(200, response.ErrorResponse(response.NotFoundDataError, err.Error()))
		return
	}
	c.JSON(200, response.NewResponseData(response.Success, entity))
}

func (t *Table[MODEL]) Delete(c *gin.Context) {
	id := c.Param("id")
	var entity MODEL

	affectNum, err := t.Struct.Delete(&entity, "id", id)
	if err != nil {
		c.JSON(200, response.ErrorResponse(response.SqlError, err.Error()))
		return
	}

	c.JSON(200, response.NewResponseData(response.Success, affectNum))
}
