package snowflake

import (
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

	MaxSequence = SequenceMask
)

type Snowflake struct {
	MachineID uint16 // 10bit
	EpochTime int64
	Sequence  Sequence // 12bit
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
	epochTime := time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC).UnixMilli()

	machineID := 3

	if machineID>>MachineIDLength != 0 {
		log.Fatalf("MachineID is under 12 bits. but, (%d)", bits.Len64(uint64(machineID)))
	}

	return &Snowflake{
		MachineID: uint16(machineID),
		EpochTime: epochTime,
	}
}

func (s *Snowflake) GetTimestamp() uint64 {
	ct := time.Now().UTC().UnixMilli()
	// fmt.Println(time.UnixMilli(ct).UTC())
	d := ct - s.EpochTime
	return uint64(d)
}

func (s *Snowflake) ID() uint64 {
	timestamp := s.GetTimestamp()
	sequence := s.Sequence.GetSequenceValue(timestamp)

	// sequenceがオーバーフローしたら待つ必要がある
	if sequence == MaxSequence {
		time.Sleep(1 * time.Millisecond)
		return s.ID()
	}

	id := (uint64(timestamp) << uint64(MoveTimeStampLength)) | (uint64(s.MachineID) << uint64(MoveMachineIDLength)) | uint64(sequence)

	return id
}

func (s *Snowflake) ParseID(id uint64) *SID {
	sid := &SID{ID: id}
	sid.MachineID = uint16(id >> uint64(MoveMachineIDLength) & MachineIDMask)
	sid.Sequence = uint16(id & SequenceMask)
	sid.Timestamp = id >> uint64(MoveTimeStampLength) & TimeStampMask
	sid.GenericTimeStamp = time.UnixMilli(s.EpochTime + int64(sid.Timestamp)).UTC()
	return sid
}
