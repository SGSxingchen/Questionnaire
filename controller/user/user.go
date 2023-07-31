package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
	"questionnaire/utility"
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	var user *models.User
	user, err = GetUserByEmail(data.Email) //利用邮箱登录获取对应USER
	if !IsValidEmail(data.Email) {
		utility.Response(http.StatusBadRequest, " 邮箱格式不规范", nil, c)
		return
	}
	if user == nil {
		utility.Response(http.StatusNotFound, "未找到用户", nil, c)
		return
	}
	if !utility.PasswordVerify(data.Password, user.Password) {
		utility.Response(http.StatusBadRequest, "密码错误", nil, c)
		return
	}
	fmt.Println("Login: ", user.UserID)
	token := utility.GenerateStandardJwt(&utility.JwtData{
		ID: strconv.Itoa(int(user.UserID)),
	})
	_id := user.Name
	utility.Response(http.StatusOK, "OK", gin.H{"token": token, "id": _id}, c)
}
func IsValidEmail(email string) bool {
	// 正则表达式来匹配RFC 2822规范的电子邮件地址
	pattern := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	match, _ := regexp.MatchString(pattern, email)
	// 如果match返回true，则表示匹配成功，反之则表示匹配失败
	return match
}
func isValidUsername(username string) bool {
	// 检查用户名长度是否在 4-20 个字符之间
	if utf8.RuneCountInString(username) < 4 || utf8.RuneCountInString(username) > 20 {
		return false
	}
	// 检查用户名是否只包含字母、数字、下划线或中文字符
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '_' && !unicode.Is(unicode.Scripts["Han"], char) {
			return false
		}
	}
	return true
}
func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := db.DB.Where("email = ?", email).First(&user).Error // 根据 email 查询用户数据
	if err != nil {                                           // 如果查询出错，则返回错误信息
		return nil, err
	} else { // 否则返回查找到的用户信息
		return user, nil
	}
}
