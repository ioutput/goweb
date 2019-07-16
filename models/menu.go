package models



type Menu struct {
	Model
	Name		string	`json:"name" form:"name"`
	Url			string	`json:"url" form:"url"`
	Status 		uint8	`json:"status"  form:"status"`
	Pid			uint8	`json:"pid"  form:"pid"`
	IsMenu		uint8	`json:"is_menu"  form:"is_menu"`
	Sort		uint8	`json:"sort"  form:"sort"`
	Remark		string	`json:"remark" form:"remark"`
}


