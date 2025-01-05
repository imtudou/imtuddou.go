package dao

import "go-blog/models"

func GetUserNameByUserId(uid int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", uid)
	if row.Err() != nil {
		return ""
	}
	var name string
	row.Scan(&name)
	return name
}

func GetUserInfo(uid int) (*models.User, error) {
	row := DB.QueryRow("select uid,user_name,avatar,create_at,update_at  from blog_user where uid=?", uid)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var user = &models.User{}
	row.Scan(&user.Uid, &user.UserName, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	return user, nil
}

func CheckUser(name, pwd string) (*models.User, error) {
	row := DB.QueryRow("select  uid,user_name,avatar,create_at,update_at  from blog_user where user_name =? and passwd =?", name, pwd)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
