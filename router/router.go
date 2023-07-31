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
	//Router.GET("/ws", controller.LinkChatRoom)
	//Router.POST("/api/chatgpt/chat1", middleware.Authorization, controller.ChatgptChatgPRC)
	//Router.POST("/api/chatgpt/chat", middleware.Authorization, controller.ChatgptChat)
	//Router.POST("/api/user/history_message", middleware.Authorization, controller.GetHistoryMessage)
	//Router.POST("/api/user/FirstMessage", middleware.Authorization, controller.FirstMessage)
	////命名规范 URL的path部分使用 系统/模块/操作 的格式，如 api/user/register
	//Router.POST("/api/user/register", controller.Register)
	//Router.GET("/api/user/is_register_name", controller.NameIsRegister)
	//Router.GET("/api/user/is_register_email", controller.EmailIsRegister)
	Router.POST("/api/user/login", user.Login)
	Router.POST("/questionnaires/create", questionnaireController.CreateQuestionnaire)
	Router.GET("/questionnaires/getAll", questionnaireController.GetAllQuestionnaires)
	Router.GET("/questionnaires/getQ", questionnaireController.GetQuestionnaireByID)
	Router.POST("/questionnaires/addSingle", questionnaireController.AddSingleChoiceQuestion)
	Router.POST("/questionnaires/addMultiple", questionnaireController.AddMultipleQuestion)
	Router.POST("/questionnaires/addSubjective", questionnaireController.AddSubjectiveQuestion)
}
