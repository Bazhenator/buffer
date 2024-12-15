package entities

import "time"

type CleaningType uint // CleaningType is a special type of cleaning clients can request in cleaning-service

const (
	Daily CleaningType = iota + 1
	Deep
	PostConstruction
	Express
	Office
	Industrial
	PostEvent
	AfterSnow
	Pool
)

type Priority uint // Priority is a special type represented client's priority

const (
	Low Priority = iota + 1
	Medium
	Hight
)

type Request struct {
	Id           uint64
	ClientId     uint64
	CleaningType CleaningType
	Priority     Priority
	GeneratorId  uint64
	Status       uint32
	AppendTime   time.Time
	TimeInBuffer time.Duration
}