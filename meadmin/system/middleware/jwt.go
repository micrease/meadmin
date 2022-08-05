package middleware

import (
	"errors"
	"meadmin/system/config"
	"meadmin/system/consts"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			token = c.Request.FormValue("token")
		}

		if len(token) == 0 {
			token = c.Request.Header.Get("Authorization")
			if len(token) == 0 {
				token = c.Request.FormValue("Authorization")
			}
		}

		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求Header中未携带token,无权限访问",
			})
			c.Abort()
			return
		}

		jwtSecret := config.GetConfig().JwtSecret
		if jwtSecret == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "缺少JWT Secret配置",
			})
			c.Abort()
			return
		}

		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:] //截取字符
		}

		j := NewJWT(jwtSecret)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": 402,
					"msg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 405,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set(consts.JwtClaimDataKey, claims.Data)
		c.Next()
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = ""
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Data consts.JwtClaimData
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT(jwt_secret string) *JWT {
	return &JWT{
		[]byte(jwt_secret),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// 生成令牌
func GenerateToken(jwtSecret string, data consts.JwtClaimData) (string, error) {
	j := JWT{
		[]byte(jwtSecret),
	}
	claims := CustomClaims{
		data,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 36000), // 过期时间 一小时
			Issuer:    jwtSecret,                        //签名的发行者
		},
	}
	return j.CreateToken(claims)
}
