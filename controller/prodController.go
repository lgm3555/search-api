package controller

import (
	"net/http"

	"github.com/search-api/handler"

	"github.com/search-api/response"

	"github.com/gin-gonic/gin"
)

// ProdSearch prod index search
// @Tags prod INDEX API
// @Summary prod INDEX SEARCH API
// @Description prod INDEX SEARCH
// @Accept json
// @Produce json
// @Param keyword query string true "상품명 키워드 삼성 노트북"
// @Param minPrice query string true "최저 가격 5000"
// @Param maxPrice query string true "최대 가격 6000"
// @Param sort query string true "정렬 1|1"
// @Param start query string true "시작 번호"
// @Param limit query string true "결과 개수"
// @Success 200 {object} response.Response
// @Router /prod/search [get]
func (c *Controller) ProdSearch(ctx *gin.Context) {
	resp := response.Response{}
	query := make(map[string]string)

	query["from"] = ctx.Query("start")
	query["size"] = ctx.Query("limit")
	query["keyword"] = ctx.Query("keyword")
	query["sort"] = ctx.Query("sort")
	query["limit"] = ctx.Query("limit")
	query["minPrice"] = ctx.Query("minPrice")
	query["maxPrice"] = ctx.Query("maxPrice")

	if rst, err := handler.EsProdSearch(query); err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Desc = err.Error()
	} else {
		resp.Status = http.StatusOK
		resp.Result = rst
	}

	ctx.JSON(http.StatusOK, resp)
}
