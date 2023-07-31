package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(code int, msg string, data gin.H, c *gin.Context) {
	//如果传入的data不为空，部分响应带有data字段
	if data != nil {
		c.JSON(code, gin.H{
			"msg":  msg,  // 该响应状态码的简要描述
			"data": data, // 带有数据的响应内容
		})
	} else { // 否则只返回状态码及简要描述
		c.JSON(code, gin.H{
			"msg": msg,
		})
	}
	//response := ""
	//if data != nil {
	//	response = fmt.Sprintf(`{"msg": "%s", "data": %v}`, msg, data)
	//} else {
	//	response = fmt.Sprintf(`{"msg": "%s"}`, msg)
	//}
	//
	//// 发送响应到客户端
	//c.Writer.WriteHeader(code)
	//c.Writer.Write([]byte(response))
}
func Responses(code int, msg string, dataSlice []gin.H, c *gin.Context) {
	if dataSlice != nil && len(dataSlice) > 0 {
		// 创建一个保存数据数组的切片
		data := make([]gin.H, len(dataSlice))
		copy(data, dataSlice)
		// 将数据切片添加到响应内容中
		c.JSON(code, gin.H{
			"msg":  msg,
			"data": data,
		})
	} else {
		// 否则返回表示成功的响应
		c.JSON(code, gin.H{
			"msg": msg,
		})
	}
}

// ResponseBadRequest 返回客户端请求错误
func ResponseBadRequest(c *gin.Context) {
	Response(http.StatusBadRequest, "Bad request", nil, c) // 由函数调用者传入相关参数构造响应信息
}

// ResponseInternalServerError 返回服务器内部错误
func ResponseInternalServerError(c *gin.Context) {
	Response(http.StatusInternalServerError, "Internal server error", nil, c) // 由函数调用者传入相关参数构造响应信息
}

// ResponseOK 返回成功信息及data字段
func ResponseOK(c *gin.Context, data gin.H) {
	Response(http.StatusOK, "OK", data, c) // 由函数调用者传入相关参数构造响应信息
}
