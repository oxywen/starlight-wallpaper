package cache

import "starlight-wallpaper/model"

var bingTodayImages model.BingHPImageArchive

var imagesPathIndex map[string]*[]string

var imageTypes []string

func GetImagesPathIndex() *map[string]*[]string {
	return &imagesPathIndex
}

func SetImagesPathIndex(data *map[string]*[]string) {
	imagesPathIndex = *data
}

func GetBingTodayImages() *model.BingHPImageArchive {
	return &bingTodayImages
}

func SetBingTodayImages(data *model.BingHPImageArchive) {
	bingTodayImages = *data
}

func GetImageTypes() *[]string {
	return &imageTypes
}

func SetImagesTypes(data *[]string) {
	imageTypes = *data
}
