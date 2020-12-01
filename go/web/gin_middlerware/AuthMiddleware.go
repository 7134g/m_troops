package middlerware

import (
	"github.com/gin-gonic/gin"
	"mblog/common"
	"mblog/model"
	"mblog/response"
)

func AuthMiddeware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			response.Fail(ctx, nil, "user Insufficient permissions")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, clamis, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Fail(ctx, nil, "user Insufficient permissions")
			ctx.Abort()
			return
		}

		userId := clamis.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			response.Fail(ctx, nil, "user Insufficient permissions")
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}
