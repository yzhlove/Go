package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//文件备份工具

var source = flag.String("source", "", "来源:")
var final = flag.String("final", "", "最终:")

func main() {
	var (
		sourceBytes, backupBytes int64
		err                      error
	)
	if *source == "" || *final == "" {
		fmt.Printf("请输入文件路径")
		return
	}
	//来源库
	if sourceBytes, err = checkFilePath(*source); err != nil {
		panic(err)
	}
	//最终写入的库
	if backupBytes, err = checkFilePath(*final); err != nil {
		panic(err)
	}

	//备份来源
	if err = backup(*source, sourceBytes); err != nil {
		panic(err)
	}

	//备份最终写入的库
	if err = backup(*final, backupBytes); err != nil {
		panic(err)
	}

}

func checkFilePath(path string) (bytes int64, err error) {
	var info os.FileInfo
	if info, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return 0, fmt.Errorf("文件不存在")
		}
		if info.IsDir() {
			return 0, fmt.Errorf("不能是文件夹")
		}
	}
	return info.Size(), nil
}

func backup(path string, num int64) error {
	var (
		directory = filepath.Dir(path)
		base      = filepath.Base(path)
		//ext        = filepath.Ext(path)
		sourceFile *os.File
		backupFile *os.File
		bytes      int64
		err        error
	)

	if sourceFile, err = os.Open(path); err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	if backupFile, err = os.Create(directory + base + ".backup"); err != nil {
		panic(err)
	}

	defer backupFile.Close()

	//备份
	if bytes, err = io.Copy(backupFile, sourceFile); err != nil {
		panic(err)
	}

	fmt.Printf("bytes = %d num = %d \n", bytes, num)

	if err = backupFile.Sync(); err != nil {
		panic(err)
	}

	return nil
}
