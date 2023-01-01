package tools

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsDir 判断某个地址是否为一个文件夹
func IsDir(f os.FileInfo) bool {
	return f.IsDir()
}

// CopyFile 复制文件
func CopyFile(dstFile string, src string) (bool, string) {
	//打开源文件
	reader, err := os.Open(src)

	if err == nil {
		defer reader.Close()

		//创建拷贝文件
		dst, errCopyOpen := os.OpenFile(dstFile, os.O_CREATE|os.O_RDWR, 0777)

		if errCopyOpen == nil {
			defer dst.Close()

			_, errCopy := io.Copy(dst, reader)
			if errCopy == nil {
				return true, ""
			} else {
				return false, errCopy.Error()
			}
		}
		return false, errCopyOpen.Error()
	}
	return false, err.Error()
}

// Copy 核心复制函数
func Copy(dst string, src string) (bool, string) {
	f, err := os.Stat(src)
	if err != nil {
		return false, err.Error()
	}

	//非文件夹直接复制
	if !IsDir(f) {
		return CopyFile(dst, src)
	}

	//文件夹类型
	//创建目标文件夹
	os.Mkdir(dst, 0777)
	list, err := ioutil.ReadDir(src)
	if err == nil {
		for _, file := range list {
			Copy(filepath.Join(dst, file.Name()), filepath.Join(src, file.Name()))
		}
		return true, "ok"
	}
	return false, err.Error()
}