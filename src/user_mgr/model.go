package user_mgr

type UserModel struct {
	Id               int64  `json:"id" gorm:"primaryKey"`
	Email            string `json:"email" gorm:"index;unique"`
	Password         string `json:"password"`
	Token            string `json:"token" gorm:"index"`
	Forbidden        bool   `json:"forbidden" gorm:"index"`
	Roles            string `json:"roles"`
	Permissions      string `json:"permissions"`
	Created_unixtime int64  `json:"created_unixtime" gorm:"autoCreateTime"`
}
