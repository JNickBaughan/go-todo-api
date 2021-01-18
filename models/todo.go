package models

// Todo is the basic type to hold a todo item
type Todo struct {
	ID string `json:"ID"`
    ParentID string `json:"ParentID"`
	Desc string `json:"Desc"`
	Complete bool `json:"Complete"`
}