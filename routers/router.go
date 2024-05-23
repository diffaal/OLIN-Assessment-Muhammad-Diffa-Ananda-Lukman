package routers

import (
	"net/http"
	"olin-assessment-muhammad-diffa/controllers"
	"olin-assessment-muhammad-diffa/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	assessmentService := services.NewAssessmentService()
	assessmentController := controllers.NewAssessmentController(assessmentService)

	assessmentApi := router.Group("assessments/")
	{
		assessmentApi.POST("/code-one", assessmentController.AssessmentCodeOne)
		assessmentApi.POST("/code-two", assessmentController.AssessmentCodeTwo)
	}

	return router
}
