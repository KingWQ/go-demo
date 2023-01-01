package main

import (
	"bufio"
	"fmt"
	"go-demo/tools"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//fileRead1()
	//fileRead2()
	//fileRead3()
	//fileWrite1()
	//fileWrite2()
	//fileWrite3()
	//testCopy1()
	//testCopy2()
	//deleteFile()
	//renameFile()
	testCopyFile()
}
func testCopyFile() {
	res, err := tools.Copy("./static/song1", "./static/song")
	fmt.Println(res, err)
}

func deleteFile() {
	err := os.Remove("./static/dst.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func renameFile() {
	os.Mkdir("./static/song", 0777)
	err := os.Rename("./static/dst2.txt", "./static/song/song.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func testCopy1() {
	dst, err := os.OpenFile("./static/dst.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err == nil {
		reader, err := os.Open("./static/test.txt")
		if err == nil {
			res, err := io.Copy(dst, reader)
			if err == nil {
				fmt.Println(res)
			} else {
				fmt.Println(err.Error())
			}
		}
		reader.Close()
	}
	dst.Close()
}
func testCopy2() {
	res, err := copyFile("./static/dst2.txt", "./static/test.txt")
	if res {
		fmt.Println("复制成功")
	} else {
		fmt.Println("复制失败：", err)
	}
}
func copyFile(dstFile string, src string) (bool, string) {
	//打开源文件
	reader, err := os.Open(src)
	defer reader.Close()

	if err == nil {
		//创建拷贝文件
		dst, errCopyOpen := os.OpenFile(dstFile, os.O_CREATE|os.O_RDWR, 0777)
		defer dst.Close()

		if errCopyOpen == nil {
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

func fileRead1() {
	file, err := os.Open("./static/test.txt")
	defer file.Close()

	if err == nil {
		var content []byte
		for {
			//创建切片保存字节
			var tempStr = make([]byte, 128)
			len, err := file.Read(tempStr)
			if err == io.EOF {
				break
			} else {
				content = append(content, tempStr[:len]...)
			}
		}

		fmt.Println(string(content))
	} else {
		fmt.Println(err)
	}
}
func fileRead2() {
	file, err := os.Open("./static/test.txt")
	defer file.Close()

	if err == nil {
		reader := bufio.NewReader(file)
		var content string
		for {
			//一次读取一行
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				//最后一行也要拼接
				content += str
				break
			} else {
				content += str
			}
		}
		fmt.Println(content)
	} else {
		fmt.Println(err)
	}
}
func fileRead3() {
	content, err := ioutil.ReadFile("./static/test.txt")
	if err == nil {
		fmt.Println(string(content))
	}
}

func fileWrite1() {
	file, err := os.OpenFile("./static/song1.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()
	if err == nil {
		fmt.Println("开始写入song1")
		file.WriteString("讲不出再见\r\n")
		file.WriteString("偏爱\r\n")
	}
}
func fileWrite2() {
	file, err := os.OpenFile("./static/song2.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()
	if err == nil {
		fmt.Println("开始写入song2")
		writer := bufio.NewWriter(file)
		writer.WriteString("天地孤影任我行\r\n")
		writer.Flush()
	}
}
func fileWrite3() {
	var content string = "笑话说太多，该说的没说出口"
	ioutil.WriteFile("./static/song3.txt", []byte(content), 0777)
}
