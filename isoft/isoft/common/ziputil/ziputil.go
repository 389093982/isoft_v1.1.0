package ziputil

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CompressZip(dirpath, destZipName string) error {
	//获取源文件列表
	f, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return err
	}
	fzip, _ := os.Create(destZipName)
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(filepath.Join(dirpath, file.Name()))
		if err != nil {
			return err
		}
		_, err = fw.Write(filecontent)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeCompressZip(zipName, unzipDir string) error {
	os.Mkdir(unzipDir, 0777)           //创建一个目录
	cf, err := zip.OpenReader(zipName) //读取zip文件
	if err != nil {
		return err
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}

		f, err := os.Create(filepath.Join(unzipDir, file.Name))
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
