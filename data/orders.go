package data

import "time"

type Orders struct {
	Order_ID    int
	Client_ID   int
	Des_emp_ID  int
	Order_TYPE  int
	Description string
	Order_date  time.Time
	Close_date  time.Time
	Master_date time.Time
}
