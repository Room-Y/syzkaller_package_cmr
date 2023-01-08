package main

import (
	"fmt"
	"syzkaller_package_cmr/StringHash"
)

func main() {
	file := "logSysPerFuncBB"
	SysFunc, FuncSys := StringHash.GetFileToMap(file)
	fmt.Println(len(SysFunc), len(FuncSys))
}
