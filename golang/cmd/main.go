package main

import (
	"fmt"
	"log"
	"math/bits"
	"time"
)

const (
	TimeStampLength uint8 = 41
	MachineIDLength uint8 = 10
	SequenceLength  uint8 = 12

	TimeStampMask = 1<<TimeStampLength - 1
	MachineIDMask = 1<<MachineIDLength - 1
	SequenceMask  = 1<<SequenceLength - 1

	MoveTimeStampLength = MachineIDLength + SequenceLength
	MoveMachineIDLength = SequenceLength
)

type Snowflake struct {
	MachineID uint16 // 10bit
	Sequence  uint16 // 12bit
	EpochTime time.Time
}

type SID struct {
	ID               uint64 
	MachineID        uint16
	Sequence         uint16
	Timestamp        uint64
	GenericTimeStamp time.Time
}

func NewSnowflake() *Snowflake {

	year := 2020
	month := 1
	day := 1
	hour := 0
	minute := 0
	second := 0
	micorsecond := 0
	nanosecond := micorsecond * 1000
	epochTime := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC)

	machineID := 3

	if machineID>>MachineIDLength != 0 {
		log.Fatalf("MachineID is under 12 bits. but, (%d)", bits.Len64(uint64(machineID)))
	}

	return &Snowflake{
		MachineID: uint16(machineID),
		Sequence:  0,
		EpochTime: epochTime,
	}
}

func (s *Snowflake) GetTimestamp() int64 {
	ct := time.Now().UTC()
	fmt.Println(ct)
	d := ct.Sub(s.EpochTime).Milliseconds()
	return d
}

func (s *Snowflake) ID() uint64 {
	timestamp := s.GetTimestamp()

	id := (uint64(timestamp) << uint64(MoveTimeStampLength)) | (uint64(s.MachineID) << uint64(MoveMachineIDLength)) | uint64(s.Sequence)

	return id
}

func (s *Snowflake) ParseID(id uint64) *SID {
	sid := &SID{ID: id}
	sid.MachineID = uint16(id >> uint64(MoveMachineIDLength) & MachineIDMask)
	sid.Sequence = uint16(id & SequenceMask)
	sid.Timestamp = id >> uint64(MoveTimeStampLength) & TimeStampMask
	sid.GenericTimeStamp = s.EpochTime.Add(time.Duration(sid.Timestamp) * time.Millisecond)
	return sid
}

func main() {

	sf := NewSnowflake()
	id := sf.ID()
	fmt.Println(id)
	p := sf.ParseID(id)
	fmt.Println(p.MachineID)
	fmt.Println(p.Sequence)
	fmt.Println(p.Timestamp)
	fmt.Println(p.GenericTimeStamp)
}

