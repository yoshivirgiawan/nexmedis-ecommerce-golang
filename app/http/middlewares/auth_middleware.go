package middlewares

import (
	"ecommerce/app/modules/auth"
	"ecommerce/app/modules/user"
	"ecommerce/helper"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {
		userService := user.NewService()
		authService := auth.NewService()
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			expirationTime := time.Unix(int64(claim["exp"].(float64)), 0)
			if time.Now().After(expirationTime) {
				response := helper.APIResponse("Token expired", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		} else {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if len(requiredRoles) > 0 {
			hasValidRole := false
			for _, role := range requiredRoles {
				if user.Role == role {
					hasValidRole = true
					break
				}
			}
			if !hasValidRole {
				response := helper.APIResponse("Forbidden: Access denied", http.StatusForbidden, "error", nil)
				c.AbortWithStatusJSON(http.StatusForbidden, response)
				return
			}
		}

		c.Set("currentUser", user)
	}
}
