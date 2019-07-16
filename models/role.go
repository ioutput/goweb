package models



type Role struct {
	Model
	Name		string	`json:"name" form:"name"`
	Status 		uint8	`json:"status"  form:"status"`
	MenuIds		string	`json:"menu_ids"  form:"menu_ids"`
	Remark		string	`json:"remark" form:"remark"`
}


