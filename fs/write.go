package fs

import "os"

func WriteUpdateSh(res string) error {
	path, err := getNowPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path+"/update.sh", []byte(res), 0777)
}
