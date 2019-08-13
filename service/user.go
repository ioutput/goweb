package service

import (
	"fmt"
	//"log"
	"strconv"
	"net/http"
	"goweb/models"
	"golang.org/x/crypto/bcrypt"
)



//查询列表
func ListUser(params map[string]string) (result []models.User) {
	//转string为int
	page,_ := strconv.Atoi(params["page"])
	page_size,_ := strconv.Atoi(params["page_size"])
	offset := (page - 1) * page_size
	query := models.DB.Offset(offset).Limit(page_size)
	if params["username"] !="" {
		query = query.Where("Username LIKE ?" , fmt.Sprintf("%s%s%s","%",params["username"],"%"))
	}
	if params["status"] !="" {
		query = query.Where("Status =?",params["status"])
	}
	err :=query.Find(&result).Error
	if err != nil {
		panic(err)
	}
	return 
}

func Login(username,password string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	var user models.User
	err := models.DB.Where(&models.User{Username:username}).First(&user).Error
	if err != nil{
		result["code"] = http.StatusBadRequest
		result["msg"] = "账号或密码错误"
	}else{
		errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if errs != nil {
			
			result["code"] = http.StatusBadRequest
			result["msg"] = "账号或密码错误!"
		}else{
			result["code"] = http.StatusOK
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
		result["code"] = http.StatusBadRequest
		result["msg"] = err
	}else{
		result["code"] = http.StatusOK
		result["data"] = user
	}

	return 
}