package common

import (
	"archive/zip"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			// log.L.Info(fmt.Sprintf("get dir error![%v]\n", err))
			return err
		}
		if exist {
			// log.L.Info(fmt.Sprintf("has dir![%v]\n"+_dir))
		} else {
			// log.L.Info(fmt.Sprintf("no dir![%v]\n"+_dir))
			// 创建文件夹
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				// log.L.Error(fmt.Sprintf("mkdir error![%v]\n",err))
			} else {
				// log.L.Info("mkdir success!\n")
			}
		}
	}
	return err
}

func CWD() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

func GetCurrentProcessPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}


func DoZlibCompress(src []byte) ([]byte,error) {
	var header zip.FileHeader
	buffer := new(bytes.Buffer)
	zipwriter := zip.NewWriter(buffer)
	header.Method = zip.Deflate
	fw, err := zipwriter.CreateHeader(&header)
	if err != nil {
		return nil,err
	}
	fw.Write(src)
	zipwriter.Close()
	data := buffer.Bytes()
	return data,nil
}