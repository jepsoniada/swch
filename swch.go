package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	checkargs:
	switch args[0] {
		case "generate", "gen":
			if len(args) > 1 {
				var files []string
				filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
					files = append(files, path)
					return nil
				})
				acc := 0
				for _, a := range files {
					if args[1] == a {
						printContent(a)
					} else {
						acc++
					}
				}
				if acc == len(files) {
					args[1] = enterArg()
					goto checkargs
				}
			} else {
				var files []string
				filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
					files = append(files, path)
					return nil
				})
				arg := enterArg()
				for _, a := range files {
					if arg == a {
						args = append(args, a)
					}
				}
				goto checkargs
			}
	}
}

func enterArg() string {
	u := ""
	fmt.Println("problem with base arg")
	fmt.Print("on which document shoud it be based: ")
	fmt.Scanln(&u)
	if len(u) < 1 {
		enterArg()
	} else { 	
		// fmt.Println(u)
		return u
	}
	return u
}

func printContent(fileName string) {
    data, _ := os.OpenFile(fileName, os.O_RDWR, 0666)
    datainfo, _ := data.Stat()
    byt := make([]byte, datainfo.Size())
    data.Read(byt)
    data.Close()
	fmt.Printf("\n%s\n\n", string(byt[:]))
}