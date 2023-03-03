package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/mumingv/gin-blog/dao"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	LoginCount int
	LastTime   time.Time
	LastIp     string
	State      int8
	Created    time.Time
	Updated    time.Time
}

const secret = "wenjie.blog.csdn.net"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

// 登录
func Login(userName string, password string) (user []*User, err error) {

	// 生成加密密码
	db := dao.DB
	db = db.Where("username = ?", userName)
	db = db.Where("password = ?", encryptPassword([]byte(password)))
	if err = db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
