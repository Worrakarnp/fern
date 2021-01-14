package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/instructorroom"
)

// InstructorRoomController defines the struct for the instructorroom controller
type InstructorRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// InstructorRoom defines the struct for the instructorroom
type InstructorRoom struct {
	ROOM     string
	BUILDING string
}

// CreateInstructorRoom handles POST requests for adding instructorroom entities
// @Summary Create instructorroom
// @Description Create instructorroom
// @ID create-instructorroom
// @Accept   json
// @Produce  json
// @Param instructorroom body ent.InstructorRoom true "InstructorRoom entity"
// @Success 200 {object} ent.InstructorRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorrooms [post]
func (ctl *InstructorRoomController) CreateInstructorRoom(c *gin.Context) {
	obj := InstructorRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "instructorroom binding failed",
		})
		return
	}

	d, err := ctl.client.InstructorRoom.
		Create().
		SetROOM(obj.ROOM).
		SetBUILDING(obj.BUILDING).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetInstructorRoom handles GET requests to retrieve a instructorroom entity
// @Summary Get a instructorroom entity by ID
// @Description get instructorroom by ID
// @ID get-instructorroom
// @Produce  json
// @Param id path int true "InstructorRoom ID"
// @Success 200 {object} ent.InstructorRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorrooms/{id} [get]
func (ctl *InstructorRoomController) GetInstructorRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cl, err := ctl.client.InstructorRoom.
		Query().
		Where(instructorroom.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListInstructorRoom handles request to get a list of instructorroom entities
// @Summary List instructorroom entities
// @Description list instructorroom entities
// @ID list-instructorroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.InstructorRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorrooms [get]
func (ctl *InstructorRoomController) ListInstructorRoom(c *gin.Context) {
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

	instructorrooms, err := ctl.client.InstructorRoom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, instructorrooms)
}

// NewInstructorRoomController creates and registers handles for the instructorroom controller
func NewInstructorRoomController(router gin.IRouter, client *ent.Client) *InstructorRoomController {
	rc := &InstructorRoomController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *InstructorRoomController) register() {
	instructorrooms := ctl.router.Group("/instructorrooms")

	instructorrooms.POST("", ctl.CreateInstructorRoom)
	instructorrooms.GET(":id", ctl.GetInstructorRoom)
	instructorrooms.GET("", ctl.ListInstructorRoom)

}
