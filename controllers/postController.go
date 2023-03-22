package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"simpleblog/loggerService"
	"simpleblog/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogController struct {
	DB *gorm.DB
}

func NewBlogController(DB *gorm.DB) BlogController {
	return BlogController{DB}
}

// [...] Create Blog Handler
func (pc *BlogController) CreateBlog(ctx *gin.Context) {
	var payload models.Blog
	if err := ctx.BindJSON(&payload); err != nil {
		return
	}
	result := pc.DB.Create(&payload)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			loggerService.Log.Print("Trying to post Duplicate key")
			ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Blog with that title already exists", "data": nil})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error(), "data": nil})
		return
	}
	loggerService.Log.Print("Blog created")
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "Success", "data": payload})
}

// [...] Get Single Blog Handler
func (pc *BlogController) FindBlogById(ctx *gin.Context) {
	BlogId := ctx.Param("id")
	var blog models.Blog
	result := pc.DB.First(&blog, "id = ?", BlogId)
	if result.Error != nil {
		loggerService.Log.Print("Post id not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Fail", "data": nil})
		return
	}
	loggerService.Log.Print("Blog id is found and send as a response")
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": blog})
}

// [...] Get All Blogs Handler
func (pc *BlogController) FindBlogs(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Blogs []models.Blog
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&Blogs)
	if results.Error != nil {
		loggerService.Log.Printf("Trying to get blogs but found error %v", results.Error)

		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": results.Error, "data": nil})
		return
	}
	loggerService.Log.Print("All blogs are send as a respose ")
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": Blogs})
}
