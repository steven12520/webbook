package ziptest

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func TestCompress() {

	f1, err := os.Open(`E:\dada6\267.jpg`)
	if err != nil {

	}
	defer f1.Close()
	f2, err := os.Open(`E:\dada6\268.jpg`)
	if err != nil {

	}
	defer f2.Close()
	f3, err := os.Open(`E:\dada6\269.jpg`)
	if err != nil {

	}
	defer f3.Close()

	f4, err := os.Open(`E:\dada6\270.jpg`)
	if err != nil {

	}
	defer f4.Close()

	f5, err := os.Open(`E:\dada6\271.jpg`)
	if err != nil {

	}
	defer f5.Close()
	f6, err := os.Open(`E:\dada6\272.jpg`)
	if err != nil {

	}
	defer f6.Close()

	var files = []*os.File{f1, f2, f3, f4, f5, f6}
	dest := "E:/dada6/dada6.zip"
	err = Compress(files, dest)
	if err != nil {

	}
}
func TestDeCompress() {
	err := DeCompress("E:/dada6/dada6.zip", "E:/dada6/")
	if err != nil {

	}
}

func GetFile() {
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(`E:/dada6`)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		println(file.Name())
	}
}
