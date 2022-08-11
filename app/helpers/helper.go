package helpers

import "github.com/gin-gonic/gin"

func FormatErrorMessage(c *gin.Context, status int, title, detail string) {
	c.JSON(status, gin.H{
		"errors": []map[string]interface{}{
			{
				"status": status,
				"title":  title,
				"detail": detail,
			},
		},
	})
}
