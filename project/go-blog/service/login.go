package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/untils"
)

func LogIn(name, pwd string) (*models.LoginResponse, error) {
	//pwd = untils.Md5Crypt(pwd, "mszl")
	u, err := dao.CheckUser(name, pwd)
	if err != nil {
		return nil, err
	}
	if u == nil || u.Uid == 0 {
		return nil, errors.New("账号或密码不正确")
	}
	var loginRes = &models.LoginResponse{}
	loginRes.UserInfo = models.UserInfo{
		Uid:      u.Uid,
		UserName: u.UserName,
		Avatar:   u.Avatar,
	}
	token, _ := untils.GenerateToken(&u.Uid)
	loginRes.Token = token
	return loginRes, nil
}
