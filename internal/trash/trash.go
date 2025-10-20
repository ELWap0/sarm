package trash

import (
	"path/filepath"
	"time"
	"errors"
	"sarm/internal/archive"
	"sarm/internal/common"
)

type Trash struct {
	Origin string `json:"origin"`
	Hash string `json:"hash"`
	DeletedAt string  `json:"time"`
}


func (t *Trash) Store() error {
	t.DeletedAt = time.Now().Format(time.RFC3339) 
	if err := archive.Store(t.Origin, t.DeletedAt, t.Hash); err != nil{
		return err
	}
	return nil
}

func (t Trash) restore() error {
	if err 	:= archive.Restore(t.Origin, t.DeletedAt, t.Hash); err != nil {
		return err
	}
	return nil
}

func newTrash(path string) (Trash,error) {
	origin, err := filepath.Abs(path)
	if err != nil {
		return Trash{}, nil
	}

	if !common.FileExits(origin) {
		return Trash{},errors.New("file does not exist") 
	}
	hash, err := common.GenHash(origin)
	if err != nil {
		return Trash{}, err
	}
	t := Trash{Origin:origin, Hash:hash}
	return t, nil
}
