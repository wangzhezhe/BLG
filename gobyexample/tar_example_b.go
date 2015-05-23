//参考的是go语言程序设计中的内容
package main

import (
	"archive/tar"
	//	"fmt"
	"compress/gzip"
	"io"
	"os"
	"s"
	"strings"
)

func createTar(filename string, files []string) error {
	//默认是在当前的目录下创建
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileWriter io.WriterCloser = file
	if strings.HasSuffix(filename, ".gz") {
		fileWriter = gzip.NewWriter(file)
		defer fileWriter.Close()
	}

	writer := tar.NewWriter(fileWriter)
	defer writer.Close()
	for _, name := range files {
		if err := writeFileToTar(writer, name); err != nil {
			return err
		}
	}
	return nil
}

func writeFileToTar(writer *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	//实例化header结构体
	header := &tar.Header{
		Name:    sanitizedName(filename),
		Mode:    int64(stat.Mode()),
		Uid:     os.Geteuid(),
		Gid:     os.Geteuid(),
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
	}
	if err = writer.WriteHeader(header); err != nil {
		return err
	}
	_, err = io.Copy(writer, file)

	return err
}
