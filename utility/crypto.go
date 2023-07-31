package utility

import "golang.org/x/crypto/bcrypt"

// PasswordHash 为输入的密码进行哈希加密
func PasswordHash(pwd string) (string, error) {
	// bcrypt.GenerateFromPassword 函数将输入的密码进行哈希加密
	// 需要传入两个参数，第一个参数是 byte 类型的密码，第二个是哈希加密的后续参数（默认为 10）
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 将加密后的 byte 数组转换为 string 类型，然后返回
	return string(bytes), nil
}

// PasswordVerify 比较已加密密码和未加密密码匹配程度
func PasswordVerify(pwd, hash string) bool {
	// bcrypt.CompareHashAndPassword 函数比较加密后的字符串和输入密码的匹配程度
	// 两个参数都是 byte 类型，一个是加密成功返回的 byte 数组，一个是未加密的密码 byte 数组
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil // 函数返回值是 bool 类型，表示两个参数完全匹配结果
}
