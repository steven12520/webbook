package common

import (
	"archive/zip"
	"github.com/astaxie/beego/logs"
	"github.com/fogleman/gg"
	"io"
	"os"
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
		prefix = prefix + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + fi.Name())
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
		header.Name = prefix + header.Name
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

//创建特殊照片
func Createpic(filename string) string {

	rectangle := "E:\\pic\\test\\rectangle.png"

	filepath := "E:\\pic\\pic\\" + filename + ".jpg"

	_, err := os.Stat(filepath)
	if err == nil {
		return filepath
	}

	const S = 500
	im, err := gg.LoadImage(rectangle)
	if err != nil {
		logs.Error(err)
		return ""
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.DrawStringAnchored(filename, S/2, S/2, 1, 1)

	dc.DrawRoundedRectangle(0, 0, 100, 100, 0)
	dc.DrawImage(im, 0, 0)
	//dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.Clip()

	er := dc.SavePNG(filepath)

	if er == nil {
		return filepath
	} else {
		return ""
	}
}
