package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "hoge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// main()
			machineID := 0b111111111111 & 0b0000001001
			fmt.Printf("12ビットに制限されたmachineID: %010b", machineID)
		})
	}
}
