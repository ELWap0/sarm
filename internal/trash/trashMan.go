package trash

import (
	"os"
	"fmt"
	"maps"
	"slices"
	"errors"
	"encoding/json"
	fuzzyfinder"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/ELWap0/sarm/internal/common"
	
)
const lockfile = "/tmp/sarm/lockfile.json"
type TrashMan struct{
	trashCans map[string]Trash
}

func NewTrashMan() (tm TrashMan, err error){
	tm = TrashMan{}
	tm.trashCans = make(map[string]Trash)
	if(!common.FileExits(lockfile)){
		if err = os.MkdirAll(TrashRoute, 0777); err != nil{
			return TrashMan{}, err
		}
		fd, err := os.Create(lockfile)
		defer fd.Close()
		if err != nil {
			return TrashMan{}, err
		}
	}else {
		data, err := os.ReadFile(lockfile)
		if err != nil {
			return TrashMan{}, err
		}
		json.Unmarshal(data,&tm.trashCans)	
	}
	return 
}

func (tm *TrashMan) Remove(path string) error {
	t, err := newTrash(path)
	if err != nil {
		return err
	}
	if err = t.Store(); err != nil {
		return err
	}
	tm.trashCans[path] = t
	return nil
}

func (tl TrashMan) List() {
	for file, _ := range tl.trashCans {
		fmt.Println(file)
	}
}

func (tl *TrashMan) Restore(path string) error {
	t, ok := tl.trashCans[path]
	if !ok {
		return errors.New("sarm does not have a record of " + path)
	}
	t.Restore()
	delete(tl.trashCans, path)
	return nil
}

func (tm *TrashMan) Clean(path string) error {
	t, ok := tm.trashCans[path]
	if !ok {
		return errors.New("sarm does not have a record of " + path)
	}
	t.Delete()
	return nil
}

func (tm *TrashMan) Purge() error {//wipes cache
	os.Remove(lockfile)
	return nil
}


func (tm TrashMan) FuzzyFind(in string) (string, error){
	keys := slices.Collect(maps.Keys(tm.trashCans))
	finds := fuzzyfinder.Find(in,keys)
	if len(finds) == 0 { 
		return "", errors.New("no match found")
	}
	return finds[0], nil
}


func (tl TrashMan) Save() (err error){
	var fd *os.File 
	data, err  := json.Marshal(tl.trashCans)
	fd.Write(data)
	return nil
}
