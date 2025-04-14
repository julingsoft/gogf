package guidx

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sony/sonyflake"
)

func New() (uint64, error) {
	st := sonyflake.Settings{
		StartTime: gtime.New("2025-01-01").Time,
		MachineID: func() (uint16, error) {
			return MachineID()
		},
	}

	s, err := sonyflake.New(st)
	if err != nil {
		return 0, err
	}

	return s.NextID()
}
