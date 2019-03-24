package compressutil

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"isoft/isoft/common/fileutil"
	"os"
	"strings"
)

func CompressDir(dirPath, dest string) error {
	filePaths, _, err := fileutil.GetAllFile(dirPath, false)
	if err != nil {
		return err
	}

	files := make([]*os.File, 0)
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		files = append(files, file)
	}

	return Compress(files, dest)
}

//压缩 使用gzip压缩成tar.gz
func Compress(files []*os.File, dest string) error {
	dest = fileutil.ChangeToLinuxSeparator(dest)
	// 文件夹不存在时需要创建文件夹
	os.MkdirAll(string([]rune(dest)[0:strings.LastIndex(dest, "/")]), 0755)

	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	defer file.Close()
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
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func DeCompress(tarFile, dest string) error {
	// 转成 Linux 路径分隔符 /
	tarFile = fileutil.ChangeToLinuxSeparator(tarFile)
	// dest 末尾补齐 /
	dest = fileutil.ChangeToLinuxSeparator(dest)
	if !strings.HasSuffix(dest, "/") {
		dest += "/"
	}
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
