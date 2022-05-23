package main

import (
	"deploy/utils"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
)

func main() {
	str, _ := os.Getwd()
	args := os.Args
	env := args[1]
	cmd := exec.Command("yarn14", "build:" + env)
	cmd.Dir = str
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "command", err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", "command", string(output))

	v := viper.New()
	v.SetConfigFile(path.Join(str, "config.yaml"))
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败", err)
	}
	host := v.Get(env + ".host").(string)
	port := v.Get(env + ".port").(int)
	username := v.Get(env + ".username").(string)
	password := v.Get(env + ".password").(string)
	localPath := str + v.Get(env + ".localPath").(string)
	remotePath := v.Get(env + ".remotePath").(string) + v.Get("name").(string)
	backupPath := v.Get(env + ".backupPath").(string) + v.Get("name").(string) + ".zip"

	//fmt.Println(host, port, username, password, localPath, remotePath, str)
	ssh_connector.Deploy(host, port, username, password, localPath, remotePath, backupPath)
	//session, err := ssh_connector.Connect("root", "Datayes#@!", "10.21.141.50", 22)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer session.Close()
	//
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr
	//_ = session.Run("ls") //sh 命令路径
}