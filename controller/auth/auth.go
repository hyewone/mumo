package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"mumogo/service"
)

type AuthController struct {
	Service *service.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Service: service.NewUserService(),
	}
}

func (con *AuthController) Login(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	userEmail := c.PostForm("user_email")
	providerType := c.PostForm("provider_type")

	if userEmail == "" || providerType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_email and provider_type must be provided"})
		return
	}

	user, err := con.Service.GetUserByEmailAndProvider(userEmail, providerType)
	// if err != nil {
	// 	print(user)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while checking user"})
	// }

	if user == nil {
		err := con.Service.CreateUser(userEmail, providerType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
			return
		}
		user, err = con.Service.GetUserByEmailAndProvider(userEmail, providerType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving created user"})
			return
		}
	}

	mumoAccessToken, err := CreateAndGetAccessToken(userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating or getting access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_info": user, "token": mumoAccessToken})
}

func (con *AuthController) GetUsers(c *gin.Context) {
	users, err := con.Service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while Select user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_list": users})
}

// func Logout(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, nil)
// }

func Signup(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (con *AuthController) CheckAccessToken(c *gin.Context) {

	// 보호되지 않는 경로를 미들웨어에서 바로 처리
	if c.FullPath() == "/auth/login" || c.FullPath() == "/open" {
		c.Next()
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("jdnfksdmfksd"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func CreateAndGetAccessToken(email string) (string, error) {

	var (
		key      []byte
		jwtToken *jwt.Token
	)

	key = []byte("jdnfksdmfksd")

	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ue":   email,
		"auth": "user",
		"iss":  "mumo",
		"sub":  "access_token",
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
		"iat":  time.Now().Unix(),
	})

	return jwtToken.SignedString(key)
}
