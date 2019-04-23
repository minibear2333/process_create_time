package main

import (
	"flag"
	"fmt"
	"process_tool"
)

func main() {
	flag.IntVar(&pid,"pid",0,"pid")
	flag.Parse()

	var processInter process_tool.ProcessInter
	var process  *process_tool.Process
	processInter = process
	var ts,err = processInter.ProcessStartTime(pid)
	if err==nil{
		 fmt.Println("start time:", ts)
	}else{
		fmt.Println(err.Error())
	}
}