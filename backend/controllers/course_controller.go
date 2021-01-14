package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team19/app/ent"
	"github.com/team19/app/ent/course"
	"github.com/team19/app/ent/degree"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/subject"
)

// CourseController defines the struct for the course controller
type CourseController struct {
	client *ent.Client
	router gin.IRouter
}

// Course defines the struct for the course
type Course struct {
	DepartmentID int
	SubjectID    int
	DegreeID     int
	CourseName   string
}

// CreateCourse handles POST requests for adding course entities
// @Summary Create course
// @Description Create course
// @ID create-course
// @Accept   json
// @Produce  json
// @Param course body ent.Course true "Course entity"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [post]
func (ctl *CourseController) CreateCourse(c *gin.Context) {
	obj := Course{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "course binding failed",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.DepartmentID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Department not found",
		})
		return
	}

	subj, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.SubjectID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Subject not found",
		})
		return
	}

	deg, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(obj.DegreeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Degree  not found",
		})
		return
	}

	cou, err := ctl.client.Course.
		Create().
		SetDepartmentID(d).
		SetSubjectID(subj).
		SetDegreeID(deg).
		SetCourseName(obj.CourseName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   cou,
	})

}

// GetCourse handles GET requests to retrieve a course entity
// @Summary Get a course entity by ID
// @Description get course by ID
// @ID get-course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [get]
func (ctl *CourseController) GetCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cou, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cou)
}

// ListCourse handles request to get a list of course entities
// @Summary List course entities
// @Description list course entities
// @ID list-course
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [get]
func (ctl *CourseController) ListCourse(c *gin.Context) {
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

	courses, err := ctl.client.Course.
		Query().
		WithDepartmentID().
		WithSubjectID().
		WithDegreeID().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, courses)
}

// DeleteCourse handles DELETE requests to delete a course entity
// @Summary Delete a course entity by ID
// @Description get course by ID
// @ID delete-course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [delete]
func (ctl *CourseController) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Course.
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

// UpdateCourse handles PUT requests to update a course entity
// @Summary Update a course entity by ID
// @Description update course by ID
// @ID update-course
// @Accept   json
// @Produce  json
// @Param id path int true "Course ID"
// @Param course body ent.Course true "Course entity"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [put]
func (ctl *CourseController) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Course{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "course binding failed",
		})
		return
	}
	obj.ID = int(id)
	cou, err := ctl.client.Course.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cou)

}

// NewCourseController creates and registers handles for the course controller
func NewCourseController(router gin.IRouter, client *ent.Client) *CourseController {
	couu := &CourseController{
		client: client,
		router: router,
	}

	couu.register()

	return couu

}

func (ctl *CourseController) register() {
	courses := ctl.router.Group("/courses")

	courses.GET("", ctl.ListCourse)

	// CRUD
	courses.POST("", ctl.CreateCourse)
	courses.GET(":id", ctl.GetCourse)
	courses.PUT(":id", ctl.UpdateCourse)
	courses.DELETE(":id", ctl.DeleteCourse)
}
