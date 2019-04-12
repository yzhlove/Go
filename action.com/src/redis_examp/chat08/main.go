package main

import (
	"os"
)

//文件备份

func main() {

	var (
		sourceFile, newFile *os.File
		err                 error
		sourcePath, newPath string
	)
	if sourceFile, err = os.Open(sourcePath); err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	if newFile, err = os.Open(newPath); err != nil {
		panic(err)
	}

	defer newFile.Close()

	//

}
