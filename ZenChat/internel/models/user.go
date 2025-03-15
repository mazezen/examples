package models

import "time"

type CUser struct {
	Model        `xorm:"extends"`
	Username     string    `xorm:"varchar(50) notnull unique 'username' comment('用户名')" json:"username"`
	Password     string    `xorm:"varchar(64) notnull 'password' comment('密码')" json:"password"`
	Avatar       string    `xorm:"varchar(255) 'avatar' comment('头像')" json:"avatar"`
	Gender       string    `xorm:"varchar(10) 'gender' comment('性别: male:男 famale:女')" json:"gender"`
	Phone        string    `xorm:"varchar(11) unique 'phone' comment('手机号')" json:"phone"`
	Email        string    `xorm:"varchar(50) unique 'email' comment('邮箱')" json:"email"`
	Identity     string    `xorm:"varchar(255) unique 'identity' comment('身份证号')" json:"identity"`
	ClientIp     string    `xorm:"varchar(100) 'client_ip' comment('客户端IP')" json:"client_ip"`
	ClientPort   string    `xorm:"varchar(10) 'client_port' comment('客户端端口')" json:"client_port"`
	LoginTime    time.Time `xorm:"datetime 'login_time' comment('登录时间')" json:"login_time"`
	HearBeatTime time.Time `xorm:"datetime 'hear_beat_time' comment('心跳时间')" json:"hear_beat_time"`
	LoginOutTime time.Time `xorm:"datetime 'login_out_time' comment('退出时间')" json:"login_out_time"`
	IsLoginOut   uint      `xorm:"int 'is_login_out' comment('是否登出')" json:"is_login_out"`
	DeviceInfo   string    `xorm:"varchar(150) 'device_info' comment('设备信息')" json:"device_info"`
}

func (u *CUser) TableName() string {
	return "c_user"
}
