package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

// GlobConfig 定义一个 config结构体变量
var GlobConfig Config

// Config 声明 配置结构体
type Config struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	ServerAddr string `yaml:"serverAddr"`
	ServerPort string `yaml:"serverPort"`
	RemoteAddr string `yaml:"remoteAddr"`
	RemotePort string `yaml:"remotePort"`
	LocalAddr  string `yaml:"localAddr"`
	LocalPort  string `yaml:"localPort"`
}

// 读取配置文件
func (c *Config) getConfig() *Config {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

// 初始化读取配置
func init() {
	// 直接赋值给结构体
	GlobConfig.getConfig()
}

// 转发
func forward(localConn net.Conn, config *ssh.ClientConfig) {
	// 设置sshClientConn
	sshClientConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", GlobConfig.ServerAddr, GlobConfig.ServerPort), config)
	if err != nil {
		fmt.Printf("ssh.Dial failed: %s", err)
	}

	// 设置Connection
	sshConn, err := sshClientConn.Dial("tcp", fmt.Sprintf("%s:%s", GlobConfig.RemoteAddr, GlobConfig.RemotePort))

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

func main() {
	// 设置SSH配置
	fmt.Printf("%s，服务器：%s:%s；远程：%s:%s；本地：%s:%s\n", "设置SSH配置", GlobConfig.ServerAddr, GlobConfig.ServerPort, GlobConfig.RemoteAddr, GlobConfig.RemotePort, GlobConfig.LocalAddr, GlobConfig.LocalPort)
	config := &ssh.ClientConfig{
		User: GlobConfig.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(GlobConfig.Password),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// 设置本地监听器
	localListener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", GlobConfig.LocalAddr, GlobConfig.LocalPort))
	if err != nil {
		fmt.Printf("net.Listen failed: %v\n", err)
	}

	for {
		// 设置本地
		localConn, err := localListener.Accept()
		if err != nil {
			fmt.Printf("localListener.Accept failed: %v\n", err)
		}
		go forward(localConn, config)
	}
}
