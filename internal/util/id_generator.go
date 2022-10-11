package util

import (
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

// Init init Sonyflake id generator
func Init() {

	var f sonyflake.Settings
	f.StartTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	sf = sonyflake.NewSonyflake(f)
	if sf == nil {
		panic("id generator init error.")
	}
}

func GetNextId() (int64, error) {
	ret, err := sf.NextID()
	if err != nil {
		return 0, err
	}

	return int64(ret), nil
}
