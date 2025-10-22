package main

import (
	"flag"
	"fmt"
	"github.com/ELWap0/sarm/internal/trash"
)

func main(){
	var del, restore, find, clean string
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
			tm.Remove(del)
		case restore != "":
			tm.Restore(restore)
		case find != "":
			tm.FuzzyFind(find)
		case clean != "":
			tm.Clean(find)
		case list:
			tm.List()
		case purge:
			tm.Purge()
		default:
			flag.PrintDefaults()
	}
}
