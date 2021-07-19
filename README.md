# starlight-wallpaper

##### 一个简易的随机壁纸服务端

主要功能点
1. 获取Bing每日壁纸，并缓存在内存中，每日定时更新
2. 服务器本地的壁纸随机请求
3. Bing历史壁纸收集和随机请求

接口定义：
|URI|Type|DataType|Function|
|:-|:-:|:-:|:-|
|/api/v1/bing/today|GET|JSON|获取Bing每日壁纸，Json格式，全8张|
|/api/v1/bing/today/:index|GET|JSON|获取Bing每日壁纸，Json格式，第1张|
|/api/v1/rand/:type|GET|IMAGE|根据类型获取服务器本地的壁纸随机请求|
|/api/v1/image/:id|GET|IMAGE|返回一个特定的图片|
|/api/v1/bing/history/rand|GET|IMAGE|随机请求一个Bing的历史壁纸|