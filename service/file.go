package service

import (
	"io/ioutil"
	"os"
	"path"
	"starlight-wallpaper/utils"
	"strings"
)

func IndexResourcePath(result *map[string]*[]string, resourcePath, fileType, localPathPrefix string) error {
	f, err := os.Stat(resourcePath)
	if err != nil || !f.IsDir() {
		return nil
	}
	rd, _ := ioutil.ReadDir(resourcePath)
	separator := utils.GetSystemFileSeparator()
	for _, fi := range rd {
		if fi.IsDir() {
			IndexResourcePath(result, resourcePath+separator+fi.Name(), fileType, localPathPrefix)
		} else {
			ext := path.Ext(fi.Name())
			if strings.Contains(fileType, ext) {
				imagePath := resourcePath + separator + fi.Name()
				uri := utils.LocalPathToURI(strings.TrimPrefix(imagePath, localPathPrefix))
				arr := strings.Split(uri, "/")
				if len(arr) <= 1 && arr[1] != "" {
					continue
				}
				p := (*result)[arr[1]]
				if p == nil {
					list := make([]string, 0)
					p = &list
					(*result)[arr[1]] = p
				}
				*p = append(*p, uri)
			}
		}
	}
	return nil
}
