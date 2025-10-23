package main

import (
	"flag"
	"fmt"
	"github.com/ELWap0/sarm/internal/trash"
)

func main(){
	var del, restore, find, clean, found string
	var purge, list bool
	flag.StringVar(&del, "delete", "", "delete following file")
	flag.StringVar(&restore, "restore", "", "restore following file")
	flag.StringVar(&find, "find", "", "fuzzy find string in trash can")
	flag.BoolVar(&list, "list", false, "list all files in trash can")
	flag.StringVar(&clean, "clean", "", "delete file from trash can")
	flag.BoolVar(&purge, "purge", false, "delete all of trash can")
	flag.Parse()
	tm, err  := trash.NewTrashMan()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch {
		case del != "":
			err  = tm.Remove(del)
		case restore != "":
			err = tm.Restore(restore)
		case find != "":
			found, err = tm.FuzzyFind(find)
			fmt.Println(found)
		case clean != "":
			err = tm.Clean(find)
		case list:
			tm.List()
		case purge:
			err = tm.Purge()
		default:
			flag.PrintDefaults()
	}
	if err != nil {
		fmt.Printf("Error: %v",err.Error())
	}
}
