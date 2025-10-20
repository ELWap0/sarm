package archive

import (
	"os"
	"fmt"
	"io"
	"compress/gzip"
)

func Store(path, deletionTime, hash string) error {
	src, err := os.Open(path)
	if err != nil {
		return err
	}
	
	dst, err := os.Create(fmt.Sprintf("%s-%s",deletionTime,hash))
	gzw := gzip.NewWriter(dst)
	io.Copy(gzw,src)
	return nil
}

func Restore(path, deletionTime, hash string) error {
	dst, err := os.Open(path)
	if err != nil {
		return err
	}
	cPath := fmt.Sprintf("%s-%s",deletionTime,hash)
	src, err := os.Open(cPath)
	if err != nil {
		return err
	}
	gzr, err  := gzip.NewReader(src)
	if err != nil {
		return err
	}
	io.Copy(dst,gzr)
	return nil
}
