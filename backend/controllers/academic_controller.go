package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B5871803/ent"
	"github.com/gin-gonic/gin"
)

// AcademicController defines the struct for the academic controller
type AcademicController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateAcademic handles POST requests for adding academic entities
// @Summary Create academic
// @Description Create academic
// @ID create-academic
// @Accept   json
// @Produce  json
// @Param academic body ent.Academic true "Academic entity"
// @Success 200 {object} ent.Academic
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /academics [post]
func (ctl *AcademicController) CreateAcademic(c *gin.Context) {
	obj := ent.Academic{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Academic binding failed",
		})
		return
	}

	f, err := ctl.client.Academic.
		Create().
		SetAcademicName(obj.AcademicName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, f)
}

// GetAcademic handles GET requests to retrieve a academic entity
// @Summary Get a academic entity by ID
// @Description get academic by ID
// @ID get-academic
// @Produce  json
// @Param id path int true "Academic ID"
// @Success 200 {object} ent.Academic
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /academics/{id} [get]
func (ctl *AcademicController) GetAcademic(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Academic.
		Query().
		Where(academic.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListAcademic handles request to get a list of academic entities
// @Summary List academic entities
// @Description list academic entities
// @ID list-academic
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Academic
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /academics [get]
func (ctl *AcademicController) ListAcademic(c *gin.Context) {
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

	academics, err := ctl.client.Academic.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, academics)
}

// DeleteAcademic handles DELETE requests to delete a academic entity
// @Summary Delete a academic entity by ID
// @Description get academic by ID
// @ID delete-academic
// @Produce  json
// @Param id path int true "academic ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /academics/{id} [delete]
func (ctl *AcademicController) DeleteAcademic(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Academic.
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

// UpdateAcademic handles PUT requests to update a academic entity
// @Summary Update a academic entity by ID
// @Description update academic by ID
// @ID update-academic
// @Accept   json
// @Produce  json
// @Param id path int true "academic ID"
// @Param academic body ent.Academic true "Academic entity"
// @Success 200 {object} ent.Academic
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /academics/{id} [put]
func (ctl *AcademicController) UpdateAcademic(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Academic{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Academic binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	f, err := ctl.client.Academic.
		UpdateOne(&obj).
		SetAcademicName(obj.AcademicName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, f)
}

// NewAcademicController creates and registers handles for the fund controller
func NewAcademicController(router gin.IRouter, client *ent.Client) *AcademicController {
	uc := &AcademicController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitAcademicController registers routes to the main engine
func (ctl *AcademicController) register() {
	academics := ctl.router.Group("/academics")

	academics.GET("", ctl.ListAcademic)
	// CRUD
	academics.POST("", ctl.CreateAcademic)
	academics.GET(":id", ctl.GetAcademic)
	academics.PUT(":id", ctl.UpdateAcademic)
	academics.DELETE(":id", ctl.DeleteAcademic)
}
