package common

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func GetCurrentProcessName(path string) string {
	baseName := filepath.Base(path)
	ext := filepath.Ext(baseName)
	name := strings.TrimSuffix(baseName, ext)
	return name
}

func GetCurrentProcessPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

func CompareTwoMapInterface(data1 map[string]interface{}, data2 map[string]interface{}) bool {
	keySlice := make([]string, 0)
	dataSlice1 := make([]interface{}, 0)
	dataSlice2 := make([]interface{}, 0)
	for key, value := range data1 {
		keySlice = append(keySlice, key)
		dataSlice1 = append(dataSlice1, value)
	}
	for _, key := range keySlice {
		if data, ok := data2[key]; ok {
			dataSlice2 = append(dataSlice2, data)
		} else {
			return false
		}
	}
	dataStr1, _ := json.Marshal(dataSlice1)
	dataStr2, _ := json.Marshal(dataSlice2)

	return string(dataStr1) == string(dataStr2)
}

func DoZlibCompress(src []byte) ([]byte, error) {
	var header zip.FileHeader
	buffer := new(bytes.Buffer)
	zipwriter := zip.NewWriter(buffer)
	header.Method = zip.Deflate
	fw, err := zipwriter.CreateHeader(&header)
	if err != nil {
		return nil, err
	}
	fw.Write(src)
	zipwriter.Close()
	data := buffer.Bytes()
	return data, nil
}
