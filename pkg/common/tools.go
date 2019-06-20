/**
 * Copyright 2017 ~ 2025 the original author or authors[983708408@qq.com].
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"umc-agent/pkg/config"
	"github.com/shirou/gopsutil/net"
)

var commandPath = "./pkg/resources/net.port.sh"
var command string

//var sumCommand = "ss -n sport == 22|awk '{sumup += $3};{sumdo += $4};END {print sumup,sumdo}'"

func GetNet(port string) string {
	if command == "" {
		command = ReadAll(commandPath)
	}
	command2 := strings.Replace(command, "#{port}", port, -1)
	s, _ := ExecShell(command2)
	fmt.Println(s)
	return s
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)
	return out.String(), err
}

//错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
}

func ReadAll(filePth string) string {
	f, err := os.Open(filePth)
	if err != nil {
		panic(err)
	}
	s, _ := ioutil.ReadAll(f)
	return string(s)
}

func ToJsonString(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("Marshal data error:%v\n", err)
	}
	return string(s)
}

func GetPhysicalId() string{
	var physicalId string
	nets, _ := net.Interfaces()
	var found bool = false
	for _, value := range nets {
		if strings.EqualFold(config.GlobalPropertiesObj.PhysicalPropertiesObj.Net, value.Name) {
			hardwareAddr := value.HardwareAddr
			fmt.Println("found net card:" + hardwareAddr)
			physicalId = hardwareAddr
			reg := regexp.MustCompile(`(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3}`)
			for _, addr := range value.Addrs {
				add := addr.Addr
				if len(reg.FindAllString(add, -1)) > 0 {
					fmt.Println("found ip " + add)
					//id = add+" "+id
					a := strings.Split(add,"/")
					if(len(a)>=2){
						physicalId = a[0]
					}else{
						physicalId = add
					}
					found = true
					break
				}
			}
		}
	}
	if !found {
		panic("net found ip,Please check the net conf")
	}
	return physicalId
}

