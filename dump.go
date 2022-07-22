package gobinaryfile

import (
	"os"
)

func DumpToFile(data *[]byte, path string)error{
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	file.Write(*data)

	return nil
}
