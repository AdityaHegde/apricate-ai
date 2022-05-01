package caravans

import "github.com/lib/pq"

type CaravanEntity struct {
	Uuid               string        `json:"uuid" gorm:"primaryKey"`
	Id                 int64         `json:"id"`
	Origin             string        `json:"origin"`
	Destination        string        `json:"destination"`
	Assistants         pq.Int32Array `json:"assistants" gorm:"type:int[]"`
	ArrivalTime        int           `json:"arrival_time"`
	SecondsTillArrival int           `json:"seconds_till_arrival"`
}
