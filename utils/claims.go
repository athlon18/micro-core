package utils

import (
	"fmt"
	"github.com/athlon18/micro-core/config/structs"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}

func (Util) GetClaims(token string, jwt structs.JWT) (*CustomClaims, error) {
	j := NewJWT(jwt)
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
