package src

import (
	"flag"
	"fmt"
	"process_tool"
)

func main() {
	var pid int
	flag.IntVar(&pid,"pid",0,"pid")
	flag.Parse()

	var ts,err = process_tool.ProcessStartTime(pid)
	if err==nil{
		 fmt.Println("start time:", ts)
	}else{
		fmt.Println(err.Error())
	}
}