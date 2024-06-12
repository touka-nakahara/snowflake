package main

import (
	"fmt"
	"snowflake/snowflake"
	"snowflake/utils"
)

func main() {

	sf := snowflake.NewSnowflake()
	id := sf.ID()
	fmt.Println(id)
	p := sf.ParseID(id)
	fmt.Println(p.MachineID)
	fmt.Println(p.Sequence)
	fmt.Println(p.Timestamp)
	fmt.Println(p.GenericTimeStamp)

	id2 := sf.ID()
	fmt.Println(id2)
	p2 := sf.ParseID(id2)

	ans := utils.CompareFlake(*p2, *p)
	fmt.Println(ans.ID)
}
