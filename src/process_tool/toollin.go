// +build linux

package process_tool

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)



var (
	Uptime int64 // 系统启动时间戳
	scClkTck = int64(C.sysconf(C._SC_CLK_TCK))
)


func updateUptime() (err error) {
	buf, err := ioutil.ReadFile("/proc/uptime")
	Uptime = time.Now().Unix()
	if err != nil {
		return err
	}
	if fields := strings.Fields(string(buf)); len(fields) == 2 {
		start, err := strconv.ParseFloat(fields[0], 10)
		if err == nil {
			Uptime = time.Now().Unix() - int64(start)//- sys.Uptime
		}
	}
	return err
}

func ProcessStartTime(pid int) (ts time.Time, err error)  {
	ts = time.Unix(0, 0)
	err = updateUptime()
	if err != nil{
		return ts,err
	}
	buf, err := ioutil.ReadFile(fmt.Sprintf("/proc/%v/stat", pid))
	if err != nil {
		return ts,err
	}
	if fields := strings.Fields(string(buf)); len(fields) > 22 {
		start, err := strconv.ParseInt(fields[21], 10, 0)
		if err == nil {
			if scClkTck > 0 {
				return time.Unix(Uptime+(start/scClkTck), 0),nil
			}
			return time.Unix(Uptime+(start/100), 0),nil
		}
	}
	return ts,errors.New("/proc/pid/stat len(fields) < 22")
}


