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
type TrashList struct{
	list map[string]Trash
}

func NewTrashList() TrashList {
	tl := TrashList{}
	tl.list = make(map[string]Trash)
	return tl
}

func (tl *TrashList) Remove(path string) error {
	t, err := newTrash(path)
	if err != nil {
		return err
	}
	tl.list[path] = t
	return nil
}

func (tl TrashList) List() {
	for file, _ := range tl.list {
		fmt.Println(file)
	}
}

func (tl *TrashList) Restore(path string) error {
	if tl.find(path) {
		if err := tl.list[path].restore(); err 	!= nil {
			return err
		}
		delete(tl.list, path)
	}
	return nil
}

func (tl *TrashList) Clean(path string)  {
	if tl.find(path) {
		delete(tl.list, path)
	}
}

func (tl *TrashList) Purge() error {//wipes cache
	for name, _ := range tl.list {
		tl.Clean(name)
		delete(tl.list,name)
	}
	os.Remove(lockfile)
	return nil
}

func (tl TrashList) find(in string) bool{
	for key, _ := range tl.list {
		if key == in {
			return true
		}
	}
	return false
}

func (tl TrashList) FuzzyFind(in string) (string, error){
	keys := slices.Collect(maps.Keys(tl.list))
	finds := fuzzyfinder.Find(in,keys)
	if len(finds) == 0 { 
		return "", errors.New("no match found")
	}
	return finds[0], nil
}

func Read()  (TrashList,error) {
	data , err :=os.ReadFile(lockfile)
	if err != nil {
		return TrashList{},err
	}
	tl := TrashList{}
	json.Unmarshal(data,tl.list)
	return tl,nil
}

func (tl TrashList) Save() (err error){
	var fd *os.File 
	if !common.FileExits(lockfile) {
		fd, err = os.Create(lockfile)
	}
	data, err  := json.Marshal(tl.list)
	fd.Write(data)
	return nil
}
