package middleware

import (
	"net/http"

	"github.com/JJDoneAway/addressbook/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary      Show Open API Doc
// @Description  Provide the OpenAPI json
// @ID	         swagger
// @Tags         swagger
// @Produce      plain
// @Success      200  {string}  string "OpenAPI json"
// @Router       /swagger/doc.json [get]
func RegisterSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html#/")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
