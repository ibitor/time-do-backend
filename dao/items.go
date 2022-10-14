package dao

import "timedo/sql"

var DB = sql.Test()

func GetAllItem() []sql.Item {
	var items []sql.Item
	DB.Find(&items)
	return items
}
