package trash

import (
	"path/filepath"
	"time"
	"errors"
	"fmt"
	"os"
	"github.com/ELWap0/sarm/internal/archive"
	"github.com/ELWap0/sarm/internal/common"
)

const TrashRoute = "tm/sarm/"

type Trash struct {
	Origin string `json:"origin"`
	Hash string `json:"hash"`
	DeletedAt string  `json:"time"`
}

func (t Trash) getPath() string{
	fileName := fmt.Sprintf("%s-%s",t.DeletedAt,t.Hash)
	return filepath.Join(TrashRoute, fileName)
}


func (t *Trash) Store() error {
	t.DeletedAt = time.Now().Format(time.RFC3339) 
	if err := archive.Store(t.Origin, t.getPath()); err != nil{
		return err
	}
	return nil
}

func (t Trash) Restore() error {
	if err 	:= archive.Restore(t.getPath(),t.Origin); err != nil {
		return err
	}
	return nil
}

func (t Trash) Delete() error {
	return os.Remove(t.getPath())
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
