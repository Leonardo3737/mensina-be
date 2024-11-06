package middlewares

import (
	"mensina-be/core/services"
	"mensina-be/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchemma string = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, BearerSchemma) {
			c.AbortWithStatus(401)
			return
		}

		token := header[len(BearerSchemma):]

		tokenId, err := services.NewJWRService().GetIdByToken(token)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Set("userId", tokenId)
	}
}

func AuthById() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchemma string = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" || !strings.HasPrefix(header, BearerSchemma) {
			c.AbortWithStatus(401)
			return
		}

		token := header[len(BearerSchemma):]

		tokenId, err := services.NewJWRService().GetIdByToken(token)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		_id := c.Param("id")

		id, err := strconv.Atoi(_id)

		if err != nil {
			c.JSON(400, utils.ErrorResponse{
				Error: "ID has to be integer",
			})
			return
		}

		if id != int(tokenId) {
			c.AbortWithStatus(401)
			return
		}
		c.Set("userId", tokenId)
	}
}
