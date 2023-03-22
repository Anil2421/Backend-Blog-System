package routes

import (
	"simpleblog/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRouteController struct {
	BlogController controllers.BlogController
}

func NewRouteBlogController(BlogController controllers.BlogController) BlogRouteController {
	return BlogRouteController{BlogController}
}

func (pc *BlogRouteController) BlogRoute(rg *gin.Engine) {

	rg.POST("/articles", pc.BlogController.CreateBlog)
	rg.GET("/articles", pc.BlogController.FindBlogs)
	rg.GET("/article/:id", pc.BlogController.FindBlogById)
}
