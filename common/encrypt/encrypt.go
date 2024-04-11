package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/zeromicro/go-zero/core/codec"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	passwordEncryptSeed = "(akita)@#$"
	mobileAesKey        = "5A2E746B08D846502F37A6E2D85D583B"
)

func EncPassword(password string) string {
	return Md5Sum([]byte(strings.TrimSpace(password + passwordEncryptSeed)))
}

// ValidatePassword 验证用户密码是否正确
func ValidatePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// 密码不匹配
		return err
	}
	// 密码匹配
	return nil
}

// GenerateSalt 生成随机盐值
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}

func EncMobile(mobile string) (string, error) {
	data, err := codec.EcbEncrypt([]byte(mobileAesKey), []byte(mobile))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func DecMobile(mobile string) (string, error) {
	originalData, err := base64.StdEncoding.DecodeString(mobile)
	if err != nil {
		return "", err
	}
	data, err := codec.EcbDecrypt([]byte(mobileAesKey), originalData)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}
