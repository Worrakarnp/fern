package controllers

import (
	"context"
	"strconv"

	"github.com/B5871803/app/ent"
	"github.com/B5871803/app/ent/petition"
	"github.com/gin-gonic/gin"
)

// PetitionController defines the struct for the petition controller
type PetitionController struct {
	client *ent.Client
	router gin.IRouter
}

// Petition defines the struct for the petition controller
type Petition struct {
	PetitionName string
	Request      string
	Academic     string
	SubjectID    int
}

// CreatePetition handles POST requests for adding petition entities
// @Summary Create petition
// @Description Create petition
// @ID create-petition
// @Accept   json
// @Produce  json
// @Param petition body ent.Petition true "Petition entity"
// @Success 200 {object} ent.Petition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petitions [post]
func (ctl *PetitionController) CreatePetition(c *gin.Context) {
	obj := ent.Petition{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "petition binding failed",
		})
		return
	}
	dr, err := ctl.client.Petition.
		Query().
		Where(petition.IDEQ(int(obj.Petition))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Petition not found",
		})
		return
	}
	deg, err := ctl.client.Petition.
		Create().
		SetPetitionName(obj.PetitionName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, deg)
}

// GetPetition handles GET requests to retrieve a petition entity
// @Summary Get a petition entity by ID
// @Description get petition by ID
// @ID get-petition
// @Produce  json
// @Param id path int true "Petition ID"
// @Success 200 {object} ent.Petition
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petitions/{id} [get]
func (ctl *PetitionController) GetPetition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dep, err := ctl.client.Petition.
		Query().
		Where(petition.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dep)
}

// ListPetition handles request to get a list of petition entities
// @Summary List petition entities
// @Description list petition entities
// @ID list-petition
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Petition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petitions [get]
func (ctl *PetitionController) ListPetition(c *gin.Context) {
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

	petitions, err := ctl.client.Degree.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, petitions)
}

// NewPetitionController creates and registers handles for the petition controller
func NewPetitionController(router gin.IRouter, client *ent.Client) *PetitionController {
	depd := &PetitionController{
		client: client,
		router: router,
	}

	depd.register()

	return depd

}

func (ctl *PetitionController) register() {
	petitions := ctl.router.Group("/petitions")

	//crud
	petitions.POST("", ctl.CreatePetition)
	petitions.GET(":id", ctl.GetPetition)
	petitions.GET("", ctl.ListPetition)

}
