package main

import (
"net"
"fmt"
"os"
)

func main(){
	conn ,err := net.Dial("tcp","127.0.0.1:6666")
	if err !=nil{
		fmt.Println("dial err",err)
		return
	}
	defer conn.Close()
	go func(){
		buf :=make([]byte,1024)
		for{
			//持续接收放进缓冲区buf中
			n,err:=os.Stdin.Read(buf)
			if err!=nil{
				fmt.Println("stdin err",err)
				return
			}
			//接收的信息直接直接发送给服务器
			_,err=conn.Write(buf[:n])
			if err!=nil{
				fmt.Println("write err",err)
				return
			}
		}
	}()
	buf := make([]byte,4096)
	for{
		n,err:=conn.Read(buf)
		if n==0{
			fmt.Println("服务器断开连接")
			break
		}
		if err!=nil {
			fmt.Println("read err", err)
			return
		}
		//接收到的信息直接打印出来
		fmt.Println(string(buf[:n]))
	}
}