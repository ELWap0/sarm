package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/ELWap0/sarm/internal/trash"
)

func main() {
	var action string
	var path string
	flag.StringVar(&path, "file", "", "file to perform action against")
	flag.StringVar(&action, "action", "", "command to perform")
	flag.Usage = func() {
		fmt.Printf("Usage: %s -action <action> -file <path>\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Printf("actions:\n\tclean\n\tdelete\n\tfind\n\tlist\n\trestore\n")
	}
	flag.Parse()
	tl, err  := trash.NewTrashMan()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch action {
	case "delete":
		tl.Remove(path)
		tl.Save()
	case "restore":
		tl.Restore(path)	
		tl.Save()
	case "find":
		tl.FuzzyFind(path)
	case "list":
		tl.List()
	case "clean":
		tl.Clean(path)
	default:
		flag.Usage()	
	}
}
