package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"questionnaire/utility"
	"strconv"
)

// Authorization 中间件函数验证用户授权

func Authorization(c *gin.Context) {
	token := c.GetHeader("Authorization") // 获取请求头中的Authorization字段
	if token == "" {                      // 如果token为空，说明未授权，返回未授权状态码，并终止后续函数调用
		utility.Response(http.StatusUnauthorized, "未收到token", nil, c)
		c.Abort()
		return
	}
	id, err := utility.ParseToken(token) // 调用ParseToken函数解析token
	if err != nil {                      // 如果解析token时出错，返回未授权状态码，并终止后续函数调用
		utility.Response(http.StatusUnauthorized, "Token验证失败", nil, c)
		c.Abort()
		return
	}
	idInt, ok := strconv.Atoi(id) // 转换解析出的用户ID为整数型
	if ok != nil {                // 如果转换失败，说明用户ID获取失败，返回服务器错误状态码，并打印错误信息
		utility.Response(http.StatusInternalServerError, "用户ID获取失败", nil, c)
		log.Println(ok)
		c.Abort()
	}
	c.Set("user_id", uint64(idInt)) // 保存转换出来的ID，并作为中间件处理后的参数传递给后续函数
	c.Next()                        // 继续执行后续函数
}
