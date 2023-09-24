package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"questionnaire/config"
	"questionnaire/controller/questionnaireController"
	"questionnaire/controller/user"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()                                                     //创建gin框架使用的路由实例
	corsConfig := cors.DefaultConfig()                                         //使用默认的cors配置
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization") //添加自定义头Authorization
	if config.Config.Dev {
		corsConfig.AllowAllOrigins = true //如果是开发模式，接受所有源的请求
	} else {
		corsConfig.AllowOrigins = config.Config.Server.AllowOrigins //如果是生产模式，仅接受特定源的请求
	}
	Router.Use(cors.New(corsConfig)) //将cors中间件挂载到路由上

	SetRouter() //为路由添加处理函数
}
func SetRouter() {
	// 处理 WebSocket 连接请求

	Router.POST("/api/user/login", user.Login)
	Router.POST("/questionnaires/create", questionnaireController.CreateQuestionnaire)
	Router.GET("/questionnaires/getAll", questionnaireController.GetAllQuestionnaires)
	Router.GET("/questionnaires/getQ", questionnaireController.GetQuestionnaireByID)
	Router.POST("/questionnaires/addSingle", questionnaireController.AddSingleChoiceQuestion)
	Router.POST("/questionnaires/addMultiple", questionnaireController.AddMultipleQuestion)
	Router.POST("/questionnaires/addSubjective", questionnaireController.AddSubjectiveQuestion)
}
