package gobinaryfile

import (
  "os"
)

func LoadFromPath(path string, buf *[]byte)error{

  f, err := os.Open(path)
  if err != nil {
    return err
  }

  c := make([]byte, 1)
  size := 0
  for {
    len, r_err := f.Read(c)
    if len == 0 {
      break
    }
    size += len
    if r_err != nil {
      return err
    }
    *buf = append(*buf, c[0])
  }

	return nil
}
