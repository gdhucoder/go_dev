package model

import "errors"

const (
	OFFLINE = iota
	ONLINE
)

var (
	ErrUserNameOrPasswordWrong = errors.New("user name or password is wrong!")
	ErrUserNotFound            = errors.New("user not found!")
)

type User struct {
	UserName    string
	Password    string
	StudentInfo *Student // 保存学生的ptr
	Status      int
}

// 用户注册
func Register(username, password string, name string, grade int, id string, sex int) *User {
	user := &User{
		UserName:    username,
		Password:    password,
		StudentInfo: CreateStudent(name, grade, id, sex),
	}
	return user
}

// 用户登录
func Login(user []*User, uesrname, password string) (err error) {
	for _, u := range user {

		if u.UserName == uesrname && u.Password == password {
			u.Status = ONLINE
			return
		}
		if u.UserName == uesrname && u.Password != password {
			err = ErrUserNameOrPasswordWrong
			return
		}
	}
	err = ErrUserNotFound
	return
}

// 用户下线，登出
func (u *User) LogOff() {
	u.Status = OFFLINE
}
