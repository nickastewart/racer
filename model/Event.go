package model 

type Event struct {
	Location    string
	RaceType    string
	Position    string
	DriverInfo  DriverInfo
	DriverTimes []DriverTime
}
