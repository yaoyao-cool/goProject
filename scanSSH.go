package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func scan(ip string,port string,timeout time.Duration,username,password string) (result bool,err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return err
		},
		Timeout: timeout,
	}
	result,err=conn(ip,port,config)
	return result, err
}
func conn(ip string,port string,config *ssh.ClientConfig) (bool,error) {
	client,err:=ssh.Dial("tcp",fmt.Sprintf("%v:%v",ip,port),config)
	if err==nil{
		defer client.Close()
		session,err:=client.NewSession()
		errRet:=session.Run("echo 1")
		if err==nil && errRet==nil {
			defer session.Close()
			return true,nil
		}
	}
	return false,err
}
func main()  {
	ip:="152.136.124.150"
	port:="22"
	timeout:=time.Second
	username:="ubuntu"
	password:="Nimda123"
	r,e:=scan(ip,port,timeout,username,password)
	fmt.Printf("scan:%v,结果:%v,错误原因:%v",ip,r,e)
}