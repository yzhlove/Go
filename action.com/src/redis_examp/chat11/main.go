package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var master = flag.String("Master", "", "主库")
var slave = flag.String("Slave", "", "从库")

type Info struct {
	length int64
	path   string
}

func main() {
	flag.Parse()
	if *master == "" || *slave == "" {
		log.Print("请输入文件路径 \n")
		return
	}

	infos := []*Info{&Info{path: *master}, &Info{path: *slave}}

	if err := Init(infos); err != nil {
		log.Fatal(err)
	}
}

func Init(infos []*Info) error {

	var (
		wg  sync.WaitGroup
		msg = make(chan error, len(infos))
	)

	//检查文件信息
	for _, v := range infos {
		length, err := check(v.path)
		if err != nil {
			return err
		}
		v.length = length
	}
	for _, v := range infos {
		wg.Add(1)
		go func(info *Info) {
			if err := backup(info.path, info.length, &wg); err != nil {
				msg <- err
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(msg)
	}()

	for v := range msg {
		if v != nil {
			fmt.Printf("出现错误:%v \n", v)
			return v
		}
	}
	return nil
}

func check(path string) (int64, error) {
	var (
		info os.FileInfo
		err  error
	)
	if info, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return 0, fmt.Errorf("文件不长存在:%v ", path)
		}
	}
	if info.IsDir() {
		return 0, fmt.Errorf("不能是文件夹:%v ", path)
	}
	return info.Size(), nil
}

func backup(path string, length int64, wg *sync.WaitGroup) error {
	defer wg.Done()
	source, err := os.Open(path)
	if err != nil {
		log.Printf("打开 %v 失败:%v", path, err)
		return err
	}
	defer source.Close()
	back, err := os.Create(path + ".backup")
	if err != nil {
		log.Printf("创建备份文件 %v 失败:%v \n", path+".backup", err)
		return err
	}
	defer back.Close()
	num, err := io.Copy(back, source)
	if err != nil {
		log.Printf("备份过程失败:%v \n", err)
		return err
	}
	if num != length {
		return fmt.Errorf("拷贝内容不一致,file:%v copy:%v \n", length, num)
	}
	if err = back.Sync(); err != nil {
		log.Printf("文件写入磁盘失败:%v \n", err)
		return err
	}
	return nil
}
