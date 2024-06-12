package main

import (
	"fmt"
	"snowflake/snowflake"
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
}
