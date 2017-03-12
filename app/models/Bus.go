package models

type BusDetails struct {
	Id          uint64
	NumberPlate string
	DriverName  string
	MaxSeats    uint32
	IsECab      bool
}
