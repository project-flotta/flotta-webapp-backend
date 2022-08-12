package helpers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

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

// SplitFilenameAndDir returns the filename and directory (in order) of a given path
func SplitFilenameAndDir(path string) (string, string) {
	return path[strings.LastIndex(path, "/")+1:], path[:strings.LastIndex(path, "/")]
}
