package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/degree"
)

// DegreeController defines the struct for the degree controller
type DegreeController struct {
	client *ent.Client
	router gin.IRouter
}

// Degree defines the struct for the degree controller
type Degree struct {
	DegreeName  string
	
}

// CreateDegree handles POST requests for adding degree entities
// @Summary Create degree
// @Description Create degree
// @ID create-degree
// @Accept   json
// @Produce  json
// @Param degree body ent.Degree true "Degree entity"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [post]
func (ctl *DegreeController) CreateDegree(c *gin.Context) {
	obj := ent.Degree{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "degree binding failed",
		})
		return
	}

	deg, err := ctl.client.Degree.
		Create().
		SetDegreeName(obj.DegreeName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, deg)
}

// GetDegree handles GET requests to retrieve a degree entity
// @Summary Get a degree entity by ID
// @Description get degree by ID
// @ID get-degree
// @Produce  json
// @Param id path int true "Degree ID"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees/{id} [get]
func (ctl *DegreeController) GetDegree(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dep, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dep)
}

// ListDegree handles request to get a list of degree entities
// @Summary List degree entities
// @Description list degree entities
// @ID list-degree
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [get]
func (ctl *DegreeController) ListDegree(c *gin.Context) {
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

	degrees, err := ctl.client.Degree.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, degrees)
}

// NewDegreeController creates and registers handles for the degree controller
func NewDegreeController(router gin.IRouter, client *ent.Client) *DegreeController {
	depd := &DegreeController{
		client: client,
		router: router,
	}

	depd.register()

	return depd

}

func (ctl *DegreeController) register() {
	degrees := ctl.router.Group("/degrees")

	//crud
	degrees.POST("", ctl.CreateDegree)
	degrees.GET(":id", ctl.GetDegree)
	degrees.GET("", ctl.ListDegree)

}
