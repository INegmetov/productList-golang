package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	productApp "github.com/inegmetov/productList-golang"
)

// @Summary Create a new product
// @Tags product
// @Description create a new product
// @Accept  json
// @Produce  json
// @Param input body productApp.ProductItem true "product info"
// @Success 200 {integer} integer 1
// @Router /api/lists/{id}/products [post]
func (h *Handler) createItem(c *gin.Context) {

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	var input productApp.ProductItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.ProductItem.Create(listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all products from list
// @Tags product
// @Description Get all products from list
// @Accept  json
// @Produce  json
// @Param id   path      int  true  "List ID"
// @Success 200 {object} []productApp.ProductItem
// @Router /api/lists/{id}/products [get]
func (h *Handler) getAll(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	item, err := h.service.ProductItem.GetAll(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Get all products
// @Tags product
// @Description Get all products from db
// @Accept  json
// @Produce  json
// @Success 200 {object} []productApp.ProductItem
// @Router /api/products [get]
func (h *Handler) getAllItems(c *gin.Context) {

	item, err := h.service.ProductItem.GetAllItems()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) addItemToList(c *gin.Context) {

}
