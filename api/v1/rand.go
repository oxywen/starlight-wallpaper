package v1

import (
	"math/rand"
	"net/http"
	"starlight-wallpaper/cache"
	"time"

	"github.com/gin-gonic/gin"
)

func RandByType(c *gin.Context) {
	t := c.Param("type")
	if t == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Type cannot be empty",
		})
		return
	}
	m := cache.GetImagesPathIndex()
	list := (*m)[t]
	if list == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "This type was not found",
		})
		return
	}
	length := len(*list)
	if length <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "There are no files under this type",
		})
		return
	}
	rand.Seed(time.Now().UnixNano())
	uri := (*list)[rand.Intn(length)]
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "/api/v1/image" + uri,
	})
}
