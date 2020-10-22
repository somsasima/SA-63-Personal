package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssaatw/app/ent"
	"github.com/ssaatw/app/ent/jobtitle"
)

type JobtitleController struct {
	client *ent.Client
	router gin.IRouter
}

type Jobtitle struct {
	Jobtitlename string
}

// CreateJobtitle handles POST requests for adding jobtitle entities
// @Summary Create jobtitle
// @Description Create jobtitle
// @ID create-jobtitle
// @Accept   json
// @Produce  json
// @Param jobtitle body Jobtitle true "Jobtitle entity"
// @Success 200 {object} Jobtitle
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /jobtitles [post]
func (ctl *JobtitleController) CreateJobtitle(c *gin.Context) {
	obj := Jobtitle{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "jobtitle binding failed",
		})
		return
	}

	jt, err := ctl.client.Jobtitle.
		Create().
		SetJobtitlename(obj.Jobtitlename).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, jt)
}

// GetJobtitle handles GET requests to retrieve a jobtitle entity
// @Summary Get a jobtitle entity by ID
// @Description get jobtitle by ID
// @ID get-jobtitle
// @Produce  json
// @Param id path int true "Jobtitle ID"
// @Success 200 {object} ent.Jobtitle
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /jobtitles/{id} [get]
func (ctl *JobtitleController) GetJobtitle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	jt, err := ctl.client.Jobtitle.
		Query().
		Where(jobtitle.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, jt)
}

// ListJobtitle handles request to get a list of jobtitle entities
// @Summary List jobtitle entities
// @Description list jobtitle entities
// @ID list-jobtitle
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Jobtitle
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /jobtitles [get]
func (ctl *JobtitleController) ListJobtitle(c *gin.Context) {
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

	jobtitles, err := ctl.client.Jobtitle.
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

	c.JSON(200, jobtitles)
}

// NewJobtitleController creates and registers handles for the jobtitle controller
func NewJobtitleController(router gin.IRouter, client *ent.Client) *JobtitleController {
	jtc := &JobtitleController{
		client: client,
		router: router,
	}

	jtc.register()

	return jtc

}

func (ctl *JobtitleController) register() {
	jobtitles := ctl.router.Group("/jobtitles")

	jobtitles.POST("", ctl.CreateJobtitle)
	jobtitles.GET(":id", ctl.GetJobtitle)
	jobtitles.GET("", ctl.ListJobtitle)

}
