package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/subject"
)

// SubjectController defines the struct for the subject controller
type SubjectController struct {
	client *ent.Client
	router gin.IRouter
}

// Subject defines the struct for the subject controller
type Subject struct {
	SubjectName  string
	
}

// CreateSubject handles POST requests for adding subject entities
// @Summary Create subject
// @Description Create subject
// @ID create-subject
// @Accept   json
// @Produce  json
// @Param subject body ent.Subject true "Subject entity"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects [post]
func (ctl *SubjectController) CreateSubject(c *gin.Context) {
	obj := ent.Subject{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "subject binding failed",
		})
		return
	}

	sub, err := ctl.client.Subject.
		Create().
		SetSubjectName(obj.SubjectName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sub)
}

// GetSubject handles GET requests to retrieve a subject entity
// @Summary Get a subject entity by ID
// @Description get subject by ID
// @ID get-subject
// @Produce  json
// @Param id path int true "Subject ID"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects/{id} [get]
func (ctl *SubjectController) GetSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sub, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sub)
}

// ListSubject handles request to get a list of subject entities
// @Summary List subject entities
// @Description list subject entities
// @ID list-subject
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subjects [get]
func (ctl *SubjectController) ListSubject(c *gin.Context) {
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

	subjects, err := ctl.client.Subject.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, subjects)
}

// NewSubjectController creates and registers handles for the subject controller
func NewSubjectController(router gin.IRouter, client *ent.Client) *SubjectController {
	subs := &SubjectController{
		client: client,
		router: router,
	}

	subs.register()

	return subs

}

func (ctl *SubjectController) register() {
	subjects := ctl.router.Group("/subjects")

	//crud
	subjects.POST("", ctl.CreateSubject)
	subjects.GET(":id", ctl.GetSubject)
	subjects.GET("", ctl.ListSubject)

}
