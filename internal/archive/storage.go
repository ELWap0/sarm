package archive

import (
	"os"
	"io"
	"compress/gzip"
)

func Store(src, dst string) error {
	fdSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	fdDst, err := os.Create(dst)
	gzw := gzip.NewWriter(fdDst)
	io.Copy(gzw, fdSrc)
	fdSrc.Close()
	fdDst.Close()
	os.Remove(src)
	return nil
}

func Restore(src, dst string) error {
	fdDst, err := os.Open(src)
	if err != nil {
		return err
	}
	fdSrc, err := os.Open(dst)
	if err != nil {
		return err
	}
	gzr, err := gzip.NewReader(fdSrc)
	if err != nil {
		return err
	}
	io.Copy(fdDst,gzr)
	fdSrc.Close()
	fdDst.Close()
	os.Remove(src)
	return nil
}
