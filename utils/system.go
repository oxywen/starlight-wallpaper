package utils

import (
	"runtime"
	"strings"
)

func GetSystemFileSeparator() string {
	sysType := runtime.GOOS
	switch sysType {
	case "linux":
		return "/"
	case "windows":
		return "\\"
	default:
		return "\\"
	}
}

func LocalPathToURI(srcPath string) string {
	sysType := runtime.GOOS
	if sysType == "linux" {
		return srcPath
	}
	uri := strings.ReplaceAll(srcPath, "\\", "/")
	return uri
}
