package utils

import (
	"snowflake/snowflake"
)

func CompareFlake(f1, f2 snowflake.SID) *snowflake.SID {
	// flakeを比較して新しい方を返す？
	if f1.ID > f2.ID {
		return &f1
	} else if f1.ID < f2.ID {
		return &f2
	} else {
		return nil
	}
}
