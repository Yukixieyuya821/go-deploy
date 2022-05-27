package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func main() {
	// 执行ping baidu的命令, 命令不会结束
	Command("node -v")

}

func Command(cmd string) error {
	c := exec.Command("cmd", "/C", cmd) 	// windows
	c.Dir = "D:\\node\\node-v16.14.2-win-x64"
	//c := exec.Command("bash", "-c", cmd)  // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			byte2String := ConvertByte2String([]byte(readString), "GB18030")
			fmt.Print(byte2String)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

