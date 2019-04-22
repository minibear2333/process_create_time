package process_tool

/*
#include <stdio.h>
#include <Windows.h>
*/
import "C"

import (
	"syscall"
	"time"
)

type Process struct {
}

func (p *Process) ProcessStartTime(pid int) (ts time.Time, err error) {
	ts = time.Unix(0, 0)
	var creationTime,exitTime,kernelTime,userTime  syscall.Filetime
	handle,err := syscall.OpenProcess(C.PROCESS_ALL_ACCESS,false,uint32(pid))
	defer syscall.FreeLibrary(handle)
	if err!=nil{
		return  ts, err
	}
	err = syscall.GetProcessTimes(handle,&creationTime,&exitTime,&kernelTime,&userTime)
	if err!=nil {
		return  ts, err
	}
	ts = time.Unix(0,creationTime.Nanoseconds())
	return ts,err
}

