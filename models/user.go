package models



type User struct {
	Model
	Username	string	`json:"username" form:"username"`
	Password	string	`json:"-" form:"password"`
	Status 		uint8	`json:"status"  form:"status"`
	RoleId		uint8	`json:"role_id"  form:"role_id"`
	Token		string	`json:"token"  form:"token"`
	Remark		string	`json:"remark" form:"remark"`
}
//修改数据库名
func (User) TableName() string {
  return "users"
}

