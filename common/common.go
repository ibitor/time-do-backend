package common

import "time"

//type Item struct {
//	Name        string    `json:"name"`
//	Count       int       `json:"count"`
//	ProduceDate time.Time `json:"produce_date"`
//}
type Item struct {
	Name        string    `form:"name"`
	Count       int       `form:"count"`
	ProduceDate time.Time `form:"produce_date" time_format:"2006-01-02" time_utc:"8"`
	SafeDay     int       `form:"safe_day"`
}
