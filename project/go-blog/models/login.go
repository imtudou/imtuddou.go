package models

type LoginResponse struct {
	Token    string   `json:"token"`    // token
	UserInfo UserInfo `json:"userInfo"` // 用户信息
}
