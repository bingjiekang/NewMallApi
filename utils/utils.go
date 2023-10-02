package utils

// 图形验证码
type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 200 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}
