package v1

import (
	"net/http"
	"starlight-wallpaper/cache"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BingToday(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": cache.GetBingTodayImages(),
	})
}

func BingTodayByIndex(c *gin.Context) {
	i := c.Param("index")
	index, err := strconv.Atoi(i)
	if err != nil || index < 0 {
		index = 0
	}
	images := cache.GetBingTodayImages()
	length := len(images.Images)
	if index >= length {
		index = length - 1
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"Images":   images.Images[index],
			"UpdateAt": images.UpdateAt,
		},
	})
}

func BingHistoryRand(c *gin.Context) {

}
