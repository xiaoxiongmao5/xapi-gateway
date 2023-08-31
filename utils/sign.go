package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/** 加密密码（bcrypt 哈希算法）
 */
func HashPasswordByBcrypt(password string) (string, error) {
	// bcrypt.DefaultCost: 这是 bcrypt 哈希算法的工作因子（cost factor），表示计算哈希时使用的迭代次数。工作因子越高，计算哈希所需的时间和资源就越多，因此更难受到暴力破解。bcrypt.DefaultCost 是库中预定义的默认工作因子值。
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

/** 校验加密密码
 */
func CheckHashPasswordByBcrypt(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

/** 生成随机字符串
 */
func GenerateRandomKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	// 生成随机的字节序列
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// 转为base64编码的字符串
	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

/** 生成带盐的哈希值（SHA-256 哈希算法）
 */
func HashBySHA256WithSalt(data, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

/** 生成API签名
 */
func GetAPISign(body string, secretKey string) (sign string) {
	hasher := sha256.New()
	hasher.Write([]byte(body + "." + secretKey))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// func GetHeaderMap(accessKey string) (string, bool) {
// 	hashMap := make(map[string]string, 0)
// 	hashMap["accessKey"] = accessKey
// 	nonce, err := GenerateRandomKey(4)
// 	if err != nil {
// 		return "", false
// 	}
// 	hashMap["nonce"] = nonce
// 	hashMap["timestamp"] = string(time.Now().Unix())
// 	hashMap["sign"] = GetAPISign()
// 	return hashMap, true
// }
