package webtest

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"

import (
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


func Init() {
	buf, err := ioutil.ReadFile("/proc/uptime")
	Uptime = time.Now().Unix()
	if err != nil {
		fmt.Println("read file /proc/uptime faile 1")
	}
	if fields := strings.Fields(string(buf)); len(fields) == 2 {
		start, err := strconv.ParseFloat(fields[0], 10)
		if err == nil {
			Uptime = time.Now().Unix() - int64(start)//- sys.Uptime
		}else{
			fmt.Println("read file /proc/uptime faile 2")
		}
	}
}

func ProcessStartTime(pid int) (ts time.Time) {
	buf, err := ioutil.ReadFile(fmt.Sprintf("/proc/%v/stat", pid))
	if err != nil {
		return time.Unix(0, 0)
	}
	if fields := strings.Fields(string(buf)); len(fields) > 22 {
		start, err := strconv.ParseInt(fields[21], 10, 0)
		if err == nil {
			if scClkTck > 0 {
				return time.Unix(Uptime+(start/scClkTck), 0)
			}
			return time.Unix(Uptime+(start/100), 0)
		}
	}
	return time.Unix(0, 0)
}