package dssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"time"
)

// 转发
func sForward(serverAddr string, remoteAddr string, localConn net.Conn, config *ssh.ClientConfig) {
	// 设置sshClientConn
	sshClientConn, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		fmt.Printf("ssh.Dial failed: %s", err)
	}

	// 设置Connection
	sshConn, err := sshClientConn.Dial("tcp", remoteAddr)

	// 将localConn.Reader复制到sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			fmt.Printf("io.Copy failed: %v", err)
		}
	}()

	// 将sshConn.Reader复制到localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			fmt.Printf("io.Copy failed: %v", err)
		}
	}()
}
func Tunnel(username string, password string, serverAddr string, remoteAddr string, localAddr string) {
	// 设置SSH配置
	fmt.Printf("%s，服务器：%s；远程：%s；本地：%s\n", "设置SSH配置", serverAddr, remoteAddr, localAddr)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// 设置本地监听器
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		fmt.Printf("net.Listen failed: %v\n", err)
	}

	for {
		// 设置本地
		localConn, err := localListener.Accept()
		if err != nil {
			fmt.Printf("localListener.Accept failed: %v\n", err)
		}
		go sForward(serverAddr, remoteAddr, localConn, config)
	}
}
