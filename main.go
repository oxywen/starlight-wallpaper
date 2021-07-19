package main

import (
	"flag"
	"fmt"
	"net/http"
	"starlight-wallpaper/cache"
	"starlight-wallpaper/router"
	"starlight-wallpaper/service"
	"starlight-wallpaper/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommandOption struct {
	//服务端口 -p
	ServerPort int
	//图片存放路径（随机壁纸和存储bing历史图片用） -d
	OSSPath string
}

func main() {
	//初始化运行参数
	option := ParseCommandOptions()
	//初始化缓存
	InitializationCache(option)
	//构建路由
	r := gin.Default()
	//注册静态资源
	r.Static("/api/v1/image", option.OSSPath)
	//注册router
	InitializationRouter(r)
	//启用Web服务
	r.Run(fmt.Sprintf(":%d", option.ServerPort))
}

func InitializationRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	v1 := r.Group("/api/v1")
	router.InitializationBingRouter(v1.Group("/bing"))
	router.InitializationStarlightRouter(v1.Group("/rand"))
}

func InitializationCache(option *CommandOption) {
	archive, err := service.GetTodayBingImageArchive()
	if err != nil {
		panic("initialization bing today image error," + err.Error())
	}
	cache.SetBingTodayImages(archive)
	indexs := make(map[string]*[]string, 10)
	service.IndexResourcePath(&indexs, option.OSSPath, ".jpg|.jpeg|.png", option.OSSPath)
	cache.SetImagesPathIndex(&indexs)
	//更新图片类别索引,暂无使用场景
	// types := make([]string, 0, len(indexs))
	// for t := range indexs {
	// 	types = append(types, t)
	// }
	// cache.SetImagesTypes(&types)
}

// 解析启动时命令行传入的参数
func ParseCommandOptions() *CommandOption {
	var option CommandOption
	flag.IntVar(&option.ServerPort, "p", 8806, "服务端口,默认为8806")
	flag.StringVar(&option.OSSPath, "d", "E:\\starlight-wallpaper\\", "图片存放路径,默认当前目录下的resources文件夹")
	flag.Parse()
	option.OSSPath = strings.TrimSuffix(option.OSSPath, utils.GetSystemFileSeparator())
	return &option
}
