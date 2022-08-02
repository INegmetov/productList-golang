package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	productApp "github.com/inegmetov/productList-golang"
)

// @Summary Create a new list
// @Tags List
// @Description create a new list
// @Accept  json
// @Produce  json
// @Param input body productApp.ProductList true "product info"
// @Success 200 {integer} integer 1
// @Router /api/lists/ [post]
func (h *Handler) createList(c *gin.Context) {
	var input productApp.ProductList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.ProductList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get list by ID
// @Tags List
// @Description get list by ID
// @Accept  json
// @Produce  json
// @Param id   path      int  true  "List ID"
// @Success 200 {object} productApp.ProductList
// @Router /api/lists/{id} [post]
func (h *Handler) getListById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	list, err := h.service.ProductList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
