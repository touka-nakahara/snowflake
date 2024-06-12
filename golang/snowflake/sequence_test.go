package snowflake

import (
	"testing"
	"time"
)

func TestSequence_GetSequenceValue(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		time     uint64
	}{
		{name: "Sequence", sequence: Sequence{}, time: uint64(time.Now().UTC().UnixMilli())},
	}
	ch := make(chan uint16, 5000)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 5000; i++ {
				// seq := tt.sequence.GetSequenceValue(tt.time)
				// fmt.Println(seq)
				go func() {
					t := uint64(time.Now().UTC().UnixMilli())
					seq := tt.sequence.GetSequenceValue(t)
					ch <- seq
				}()
			}

			for i := 0; i < 100; i++ {
				<-ch
				// fmt.Println(id)
			}

		})
	}
}
