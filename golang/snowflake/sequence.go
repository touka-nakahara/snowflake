package snowflake

import "sync"

type Sequence struct {
	Count uint16
	Time  uint64
	sync.Mutex
}

func (s *Sequence) GetSequenceValue(currentTime uint64) uint16 {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Millisecondがお大きいなら初期化して返す
	if currentTime > s.Time {
		s.Time = currentTime
		s.Count = 0
		return s.Count
	}

	// Millisecondが同じならカウントアップして返す
	if currentTime == s.Time {
		return MaxSequence & (s.Count + 1)
	}

	return s.Count
}
