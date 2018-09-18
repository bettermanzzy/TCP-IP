package main

import (
	"net"
	"fmt"
	"strings"
	"io"
)
func handleconn(conn net.Conn){

	buf :=make([]byte,1024)
	conn.Write([]byte("welcome my server"))
	for{
		n,err := conn.Read(buf)
		//client 断开连接
		if n==0{
			fmt.Println("out to leave",conn.RemoteAddr().String())
			break
		}
		//接收数据错误
		if err!=nil&&err!=io.EOF{
			fmt.Println("read err",err)
			return
		}
		//处理数据
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main(){
	listener,err:= net.Listen("tcp","127.0.0.1:6666")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for{
		conn,err := listener.Accept()
		if err!=nil{
			fmt.Println(err)
			return
		}
		go handleconn(conn)
	}

}