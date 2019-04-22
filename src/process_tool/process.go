package process_tool

import "time"

type ProcessInter interface {
  ProcessStartTime(pid int) (ts time.Time, err error)
}
