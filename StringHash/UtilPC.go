package StringHash

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func InitSysFuncAndPC(file string, pcToFuncs map[uint64]string, funcToPCS map[string][]uint64) (map[uint64]map[string]struct{}, map[uint64]map[uint64]struct{}, map[uint64]uint64) {
	SysToFuncs := make(map[uint64]map[string]struct{})
	SysToPCs := make(map[uint64]map[uint64]struct{})
	SysTotalPCs := make(map[uint64]uint64)
	fi, _ := os.Open(file)
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		str := string(a)

		if strings.HasPrefix(str, "----") {
			strSysId, _, _ := br.ReadLine()
			SysId, _ := strconv.ParseUint(strings.Split(string(strSysId), "  >>>  ")[0], 10, 64)

			if _, ok := SysToPCs[SysId]; !ok {
				SysToFuncs[SysId] = make(map[string]struct{})
				SysToPCs[SysId] = make(map[uint64]struct{})
				SysTotalPCs[SysId] = 0
			}

			for {
				strPC, _, _ := br.ReadLine()
				if len(strPC) == 0 {
					break
				}

				PC, _ := strconv.ParseUint(strings.Split(string(strPC), " :  ")[0], 10, 64)
				if funcname, ok1 := pcToFuncs[PC]; ok1 {
					if _, ok2 := SysToFuncs[SysId][funcname]; !ok2 {
						SysToFuncs[SysId][funcname] = struct{}{}
						SysTotalPCs[SysId] += uint64(len(funcToPCS[funcname]))
					}
				}
			}
		}
	}

	return SysToFuncs, SysToPCs, SysTotalPCs
}
