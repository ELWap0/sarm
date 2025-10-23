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
	if err != nil {
		return err
	}
	defer fdDst.Close()

	gzw := gzip.NewWriter(fdDst)
	defer gzw.Close()

	if _, err = io.Copy(gzw, fdSrc); err != nil {
		return err
	}
	fdSrc.Close()
	os.Remove(src)
	return nil
}

func Restore(src, dst string) error {
	fdSrc, err := os.Open(src)
	if err != nil {
		return err
	}

	gzr, err := gzip.NewReader(fdSrc)
	if err != nil {
		return err
	}

	fdDst, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fdDst.Close()

	io.Copy(fdDst,gzr)
	gzr.Close()
	fdSrc.Close()
	os.Remove(src)

	return nil
}
