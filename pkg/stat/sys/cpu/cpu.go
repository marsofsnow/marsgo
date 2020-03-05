package cpu

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	interval time.Duration = time.Millisecond * 500 // 500ms
)

var (
	stats CPU    //一个接口变量,相当于一个指针 ,在init里面初始化,然后在一个携程里调用其函数,赋值给usage
	usage uint64 //包级变量,
)

// CPU is cpu stat usage.
type CPU interface {
	Usage() (u uint64, e error)
	Info() Info
}

func init() {
	var (
		err error
	)
	stats, err = newCgroupCPU()
	if err != nil {
		// fmt.Printf("cgroup cpu init failed(%v),switch to psutil cpu\n", err)
		stats, err = newPsutilCPU(interval)
		if err != nil {
			panic(fmt.Sprintf("cgroup cpu init failed!err:=%v", err))
		}
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			<-ticker.C
			u, err := stats.Usage()
			if err == nil && u != 0 {
				atomic.StoreUint64(&usage, u)
			}
		}
	}()
}

// Stat cpu stat.
type Stat struct {
	Usage uint64 // cpu use ratio.
}

// Info cpu info.
type Info struct {
	Frequency uint64
	Quota     float64
}

// ReadStat read cpu stat.
func ReadStat(stat *Stat) {
	stat.Usage = atomic.LoadUint64(&usage)
}

// GetInfo get cpu info.
func GetInfo() Info {
	return stats.Info()
}
