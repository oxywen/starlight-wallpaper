package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"starlight-wallpaper/constant"
	"starlight-wallpaper/model"
	"time"
)

func GetTodayBingImageArchive() (*model.BingHPImageArchive, error) {
	resp, err := http.Get(constant.BING_IMAGE_ARCHIVE_URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response code error,the code should be 200,not %d", resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data model.BingHPImageArchive
	json.Unmarshal(bytes, &data)
	data.UpdateAt = time.Now()
	return &data, nil
}
