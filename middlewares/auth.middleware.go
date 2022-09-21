package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/tiyan-attirmidzi/go-gin-gorm/helpers"
	"github.com/tiyan-attirmidzi/go-gin-gorm/services"

	"github.com/gin-gonic/gin"
)

func Authorize(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helpers.ResponseError("Sorry, You're Unauthorized", "Token not found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helpers.ResponseError("Sorry, Token Is Not Valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

// TODO: add user access permisition
func HasAccess() {

}
