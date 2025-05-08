package jwt

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

const UID = "uid"

func CreateToken(uid uint, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		UID:   uid,
		"exp": gtime.Now().AddDate(0, 1, 0).Timestamp(),
	})

	return token.SignedString([]byte(key))
}

func ExtractToken(bearerToken string) string {
	// bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ParseToken(tokenString, key string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return gconv.Uint(claims[UID]), nil
	}

	return 0, fmt.Errorf("invalid token")
}
