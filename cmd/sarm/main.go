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
		fmt.Printf("actions:\n\tdelete - deletes the file\n\trestore - restores deleted file\n\tfind - dumpter dive for a file\n\tlist - list all deleted files\n\tclean - purges all deleted files\n")
	}
	flag.Parse()
	tm, err  := trash.NewTrashMan()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch action {
	case "delete":
		tm.Remove(path)
		tm.List()
 		if err := tm.Save(); err != nil {
			fmt.Println(err.Error())
		}
	case "restore":
		tm.Restore(path)	
		tm.Save()
	case "find":
		tm.FuzzyFind(path)
	case "list":
		tm.List()
	case "clean":
		tm.Clean(path)
	default:
		flag.Usage()	
	}
}
