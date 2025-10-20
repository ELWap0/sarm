package sarm

import (
	"flag"
	"sarm/internal/trash"
)

func main() {
	var action string
	var path string
	flag.StringVar(&path, "file", "", "file to perform action against")
	flag.StringVar(&action, "action", "", "command to perform")
	flag.Parse()
	tl := trash.NewTrashList()
	trash.Read()
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
	case "save":
	}
}
