package serializer

import (
	"github.com/YouthCampProj/douyin/model"
)

const (
	CodeUserNotFound = 1000 + iota
	CodeUserLoginFailed
)

var CodeUserMessages = map[int]string{
	CodeSuccess:         "",
	CodeUserNotFound:    "用户不存在",
	CodeUserLoginFailed: "用户名或密码错误",
}

// User 用户信息
type User struct {
	ID            int    `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int    `json:"follow_count"`   // 关注总数
	FollowerCount int    `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// UserRegisterResponse 用户注册响应
type UserRegisterResponse struct {
	Response
	UserID int    `json:"user_id"` // 用户ID
	Token  string `json:"token"`   // 用户鉴权token
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	Response
	UserID int    `json:"user_id,omitempty"` // 用户ID
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

// UserResponse 获取用户信息响应
type UserResponse struct {
	Response
	User *User `json:"user"` // 用户信息
}

func BuildUserResponse(user *model.User, isFollow bool) *User {
	res := &User{
		ID:            user.UserID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	return res
}

func BuildUserLoginResponse(code int, user *model.User, token string) *UserLoginResponse {
	res := &UserLoginResponse{}
	if code != CodeSuccess {
		res.Response = NewResponse(code, CodeUserMessages[code])
		return res
	} else {
		res.Response = NewResponse(CodeSuccess, CodeUserMessages[CodeSuccess])
	}
	res.UserID = user.UserID
	res.Token = token
	return res
}

func BuildUserRegisterResponse(code int, user *model.User, token string) *UserRegisterResponse {
	res := &UserRegisterResponse{}
	if code != CodeSuccess {
		res.Response = NewResponse(code, CodeUserMessages[code])
		return res
	} else {
		res.Response = NewResponse(CodeSuccess, CodeUserMessages[CodeSuccess])
	}
	res.UserID = user.UserID
	res.Token = token
	return res
}
