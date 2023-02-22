package StringHash

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintTest() {
	fmt.Println(111)
}

func GetFileToMap(file string) (map[string][]string, map[string][]string) {
	FuncToSys := make(map[string][]string, 0)
	SysToFunc := make(map[string][]string, 0)

	fi, _ := os.Open(file)
	defer fi.Close()

	syscallBegin := "Now We do Func: "
	syscallLen := len(syscallBegin)

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		str := string(a)
		if strings.HasPrefix(str, syscallBegin) {
			str = str[syscallLen:]
			// fmt.Println(str)
			SysToFunc[str] = make([]string, 0)
			br.ReadLine()
			for {
				a, _, _ := br.ReadLine()

				funcStr := string(a)
				if funcStr == "" {
					break
				}
				funcStr = strings.Split(funcStr, ":")[0]
				SysToFunc[str] = append(SysToFunc[str], funcStr)
				FuncToSys[funcStr] = append(FuncToSys[funcStr], str)
			}
		}
	}

	// rangeMapStringSplice(FuncToSys)
	return FuncToSys, SysToFunc
}

func rangeMapStringSplice(a map[string][]string) {
	for k, v := range a {
		fmt.Println(k)
		for _, o := range v {
			fmt.Print(o, "    ")
		}
		fmt.Print("\n\n")
	}
}
