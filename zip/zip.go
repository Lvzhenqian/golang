package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	//oldFileName可以是文件或者目录
	oldFileName := "root.log"

	currentTime := time.Now()

	//获取s
	mSecond := fmt.Sprintf("%03d", currentTime.Nanosecond() / 1e6)

	//zip文件名
	zipFileName := strings.Split(oldFileName, ".")[0] + "_" + currentTime.Format("20060102150405") + mSecond + ".zip"

	//压缩文件
	zipFile(oldFileName, zipFileName)
}

func zipFile(source, target string) error{

	//创建目标zip文件
	zipFile , err := os.Create(target)

	if err != nil {
		fmt.Println(err)
		return err
	}

	//关闭文件
	defer zipFile.Close()

	//创建一个写zip的writer
	archive := zip.NewWriter(zipFile)

	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		//将文件或者目录信息转换为zip格式的文件信息
		header, err := zip.FileInfoHeader(info)

		if err != nil{
			return err
		}

		if !info.IsDir() {
			// 确定采用的压缩算法（这个是内建注册的deflate）
			header.Method = zip.Deflate
		}

		//
		header.SetModTime(time.Unix(info.ModTime().Unix(), 0))

		//文件或者目录名
		header.Name = path

		//创建在zip内的文件或者目录
		writer, err := archive.CreateHeader(header)

		if err != nil{
			return err
		}

		//如果是目录，只需创建无需其他操作
		if info.IsDir() {
			return nil
		}

		//打开需要压缩的文件
		file, err := os.Open(path)

		if err != nil{
			return err
		}

		defer file.Close()

		//将待压缩文件拷贝给zip内文件
		_, err = io.Copy(writer, file)

		return err

	})
}
