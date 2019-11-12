package pprof

import (
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func PPMem(suffix string) error {
	fp, err := os.Create(fmt.Sprintf("mem.pprof.%v", suffix))
	if err != nil {
		return errors.New(fmt.Sprintf("PPMem failed. suffix: [%v], err: [%v]", suffix, err))
	}
	defer fp.Close()

	fmt.Println("PPMem fp", fp.Name())
	if err := pprof.WriteHeapProfile(fp); err != nil {
		return errors.New(fmt.Sprintf("PPMem failed. suffix: [%v], err: [%v]", suffix, err))
	}

	return nil
}

func PPCpu(suffix string, duration time.Duration) error {
	fp, err := os.Create(fmt.Sprintf("cpu.pprof.%v", suffix))
	if err != nil {
		return errors.New(fmt.Sprintf("PPCpu failed. suffix: [%v], duration: [%v], err: [%v]", suffix, duration, err))
	}
	defer fp.Close()

	fmt.Println("PPCpu fp", fp.Name())
	if err := pprof.StartCPUProfile(fp); err != nil {
		return errors.New(fmt.Sprintf("PPCpu failed. suffix: [%v], duration: [%v], err: [%v]", suffix, duration, err))
	}
	defer pprof.StopCPUProfile()

	time.Sleep(duration)

	return nil
}

// "mem [suffix=yyyymmddhh]" 生成内存 pprof 文件
// "cpu [duration=60s] [suffix=yyyymmddhh]" 生成 cpu pprof 文件
func PPCmd(command string) error {
	fields := strings.Split(command, " ")
	cmdType := fields[0]

	switch cmdType {
	case "mem":
		suffix := time.Now().Format("200601021504")
		if len(fields) >= 2 {
			suffix = fields[1]
		}

		return PPMem(suffix)
	case "cpu":
		duration := 60 * time.Second
		if len(fields) >= 2 {
			var err error
			duration, err = time.ParseDuration(fields[1])
			if err != nil {
				return errors.New(fmt.Sprintf("PPCmd failed. command: [%v], err: [%v]", command, err))
			}
		}

		suffix := time.Now().Format("200601021504")
		if len(fields) >= 3 {
			suffix = fields[2]
		}

		return PPCpu(suffix, duration)
	default:
		return errors.New(fmt.Sprintf("PPCmd failed. invalid command: [%v]", command))
	}
}
