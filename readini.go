package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getValue0(file string, expectpara string, expectkey string) string{
	result0 := ""
	reader, err := os.Open(file)
	if err != nil{
		return "error"
	}
	defer reader.Close()

	r := bufio.NewReader(reader)

	var paratemp string
	for{
		stringtemp,err := r.ReadString('\n')

		if(err != nil){
			break
		}

		strs := strings.TrimSpace(stringtemp)

		if strs == "" {
			continue
		}

		if strs[0] == '[' && strs[len(strs)-1] == ']'{
			paratemp = strs[1:len(strs)-1]
		}else if paratemp == expectpara{
			pair := strings.Split(stringtemp, "=")
			if len(pair)==2{
				key := pair[0]

				key = strings.TrimSpace(key)
				if key == expectkey{
					result0 = strings.TrimSpace(pair[1])
				}
			}
		}


	}
	return result0
}

func main(){
	var expectpara string
	var expectkey string
	flag.StringVar(&expectpara, "p", "", "choose expect para")
	flag.StringVar(&expectkey, "k", "", "choose expect key")

	flag.Parse()

	fmt.Println(getValue0("example.ini", expectpara, expectkey));
}