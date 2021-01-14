package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
	"github.com/team19/app/ent/title"
)

// InstructorInfoController defines the struct for the instructorinfo controller
type InstructorInfoController struct {
	client *ent.Client
	router gin.IRouter
}

// InstructorInfo defines the struct for the instructorinfo
type InstructorInfo struct {
	Department     int
	InstructorRoom int
	Title          int
	Name           string
	PhoneNumber    string
	Email          string
	Password       string
}

// CreateInstructorInfo handles POST requests for adding instructorinfo entities
// @Summary Create instructorinfo
// @Description Create instructorinfo
// @ID create-instructorinfo
// @Accept   json
// @Produce  json
// @Param instructorinfo body ent.InstructorInfo true "InstructorInfo entity"
// @Success 200 {object} ent.InstructorInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorinfos [post]
func (ctl *InstructorInfoController) CreateInstructorInfo(c *gin.Context) {
	obj := InstructorInfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "instructorinfo binding failed",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Department not found",
		})
		return
	}

	ir, err := ctl.client.InstructorRoom.
		Query().
		Where(instructorroom.IDEQ(int(obj.InstructorRoom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "InstructorRoom not found",
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Title not found",
		})
		return
	}

	u, err := ctl.client.InstructorInfo.
		Create().
		SetNAME(obj.Name).
		SetPHONENUMBER(obj.PhoneNumber).
		SetEMAIL(obj.Email).
		SetPASSWORD(obj.Password).
		SetDepartment(d).
		SetInstructorroom(ir).
		SetTitle(t).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetInstructorInfo handles GET requests to retrieve a instructorinfo entity
// @Summary Get a instructorinfo entity by ID
// @Description get instructorinfo by ID
// @ID get-instructorinfo
// @Produce  json
// @Param id path int true "InstructorInfo ID"
// @Success 200 {object} ent.InstructorInfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorinfos/{id} [get]
func (ctl *InstructorInfoController) GetInstructorInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.InstructorInfo.
		Query().
		Where(instructorinfo.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   u,
	})
}

// ListInstructorInfo handles request to get a list of instructorinfo entities
// @Summary List instructorinfo entities
// @Description list instructorinfo entities
// @ID list-instructorinfo
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.InstructorInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorinfos [get]
func (ctl *InstructorInfoController) ListInstructorInfo(c *gin.Context) {
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

	instructorinfos, err := ctl.client.InstructorInfo.
		Query().
		WithDepartment().
		WithInstructorroom().
		WithTitle().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, instructorinfos)
}

// DeleteInstructorInfo handles DELETE requests to delete a instructorinfo entity
// @Summary Delete a instructorinfo entity by ID
// @Description get instructorinfo by ID
// @ID delete-instructorinfo
// @Produce  json
// @Param id path int true "InstructorInfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorinfos/{id} [delete]
func (ctl *InstructorInfoController) DeleteInstructorInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.InstructorInfo.
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

// UpdateInstructorInfo handles PUT requests to update a instructorinfo entity
// @Summary Update a instructorinfo entity by ID
// @Description update instructorinfo by ID
// @ID update-instructorinfo
// @Accept   json
// @Produce  json
// @Param id path int true "InstructorInfo ID"
// @Param instructorinfo body ent.InstructorInfo true "InstructorInfo entity"
// @Success 200 {object} ent.InstructorInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /instructorinfos/{id} [put]
func (ctl *InstructorInfoController) UpdateInstructorInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.InstructorInfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "instructorinfo binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.InstructorInfo.
		UpdateOneID(int(id)).
		SetNAME(obj.NAME).
		SetPHONENUMBER(obj.PHONENUMBER).
		SetEMAIL(obj.EMAIL).
		SetPASSWORD(obj.PASSWORD).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewInstructorInfoController creates and registers handles for the instructorinfo controller
func NewInstructorInfoController(router gin.IRouter, client *ent.Client) *InstructorInfoController {
	uc := &InstructorInfoController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *InstructorInfoController) register() {
	instructorinfos := ctl.router.Group("/instructorinfos")

	instructorinfos.GET("", ctl.ListInstructorInfo)

	// CRUD
	instructorinfos.POST("", ctl.CreateInstructorInfo)
	instructorinfos.GET(":id", ctl.GetInstructorInfo)
	instructorinfos.PUT(":id", ctl.UpdateInstructorInfo)
	instructorinfos.DELETE(":id", ctl.DeleteInstructorInfo)
}
