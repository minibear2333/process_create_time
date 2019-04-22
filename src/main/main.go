package main

import (
	"fmt"
	"process_tool"
	"flag"
)

func main() {
	var pid int
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