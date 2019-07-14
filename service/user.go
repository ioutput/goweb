package service

import (
	"strconv"
	"demo/models"
	"golang.org/x/crypto/bcrypt"
)



//查询列表
func ListUser(params map[string]string) (result []models.User) {
	//转string为int
	page,_ := strconv.Atoi(params["page"])
	page_size,_ := strconv.Atoi(params["page_size"])
	offset := (page - 1) * page_size
	err :=models.DB.Offset(offset).Limit(page_size).Find(&result).Error
	if err != nil {
		panic(err)
	}
	return 
}

func Login(params map[string]string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	var user models.User
	err := models.DB.Where(&models.User{Username:params["username"]}).First(&user).Error
	if err != nil{
		result["code"] = 400
		result["msg"] = "账号或密码错误"
	}else{
		errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params["password"]))
		if errs != nil {
			result["code"] = 400
			result["msg"] = "账号或密码错误1"
		}else{
			result["code"] = 200
			result["data"] = user
		}
		
	}
	return 
}

func Register(user models.User) (result map[string]interface{}) {	
	result = make(map[string]interface{})
	//user :=models.User{Username:params["username"]}
	hash,_ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	err := models.DB.Create(&user).Error
	
	if err != nil{
		result["code"] = 400
		result["msg"] = err
	}else{
		result["code"] = 200
		result["data"] = user
	}

	return 
}