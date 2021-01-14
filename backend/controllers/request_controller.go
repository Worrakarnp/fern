package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B5871803/ent"
	"github.com/gin-gonic/gin"
)

// RequestController defines the struct for the request controller
type RequestController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRequest handles POST requests for adding request entities
// @Summary Create request
// @Description Create request
// @ID create-request
// @Accept   json
// @Produce  json
// @Param request body ent.Request true "Request entity"
// @Success 200 {object} ent.Request
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requests [post]
func (ctl *RequestController) CreateRequest(c *gin.Context) {
	obj := ent.Request{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Request binding failed",
		})
		return
	}

	f, err := ctl.client.Request.
		Create().
		SetRequestName(obj.RequestName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, f)
}

// GetRequest handles GET requests to retrieve a request entity
// @Summary Get a request entity by ID
// @Description get request by ID
// @ID get-request
// @Produce  json
// @Param id path int true "Request ID"
// @Success 200 {object} ent.Request
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requests/{id} [get]
func (ctl *RequestController) GetRequest(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Request.
		Query().
		Where(request.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListRequest handles request to get a list of request entities
// @Summary List request entities
// @Description list request entities
// @ID list-request
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Request
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requests [get]
func (ctl *RequestController) ListRequest(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	requests, err := ctl.client.Request.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, requests)
}

// DeleteRequest handles DELETE requests to delete a request entity
// @Summary Delete a request entity by ID
// @Description get request by ID
// @ID delete-request
// @Produce  json
// @Param id path int true "request ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requests/{id} [delete]
func (ctl *RequestController) DeleteRequest(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Request.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateRequest handles PUT requests to update a request entity
// @Summary Update a request entity by ID
// @Description update request by ID
// @ID update-request
// @Accept   json
// @Produce  json
// @Param id path int true "request ID"
// @Param request body ent.Request true "Request entity"
// @Success 200 {object} ent.Request
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requests/{id} [put]
func (ctl *RequestController) UpdateRequest(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Request{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Request binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Request.
		UpdateOne(&obj).
		SetRequestName(obj.RequestName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, f)
}

// NewRequestController creates and registers handles for the fund controller
func NewRequestController(router gin.IRouter, client *ent.Client) *RequestController {
	uc := &RequestController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRequestController registers routes to the main engine
func (ctl *RequestController) register() {
	requests := ctl.router.Group("/requests")

	requests.GET("", ctl.ListRequest)
	// CRUD
	requests.POST("", ctl.CreateRequest)
	requests.GET(":id", ctl.GetRequest)
	requests.PUT(":id", ctl.UpdateRequest)
	requests.DELETE(":id", ctl.DeleteRequest)
}
