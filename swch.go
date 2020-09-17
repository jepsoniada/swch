// wert
// gen (sorce = file name String) - generates file based on /sorce/ file; that would look like "*name*.swch"
// (file = file name String) - makes changes in /file/ acording to "*fileName*.swch"
//  wert (^(-a|-n) \d :: .*)
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type SimplifyFile struct {
	lines map[int]string
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	checkargs:
	switch args[0] {
		default:
			var files []string
			filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
				files = append(files, path)
				return nil
			})
			check, _ := regexp.Compile(".+[.]")
			executiveName := args[0][:check.FindStringIndex(args[0])[1]] + "swch"
			fmt.Println(executiveName)
			acc := 0
			fin := false
			for _, a := range files {
				if a == executiveName {
					fin = true
				}
				if a != args[0] {
					acc++
				}
			}
			if acc == len(files) -1 && fin {
				swch(args[0])
			}
			if acc == len(files) {
				fmt.Println("not file, sorry ::((")
			} else if !fin {
				u := ""
				fmt.Println("executive to this file doesn't exist")
				fmt.Println("should i generate it? (y/n)")
				fmt.Scanln(&u)
				if u == "y" {
					createSwchFile(args[0])
				}
			}
		case "generate", "gen":
			if len(args) > 1 {
				check, _ := regexp.Compile(".+[.]")
				filename := args[1][:check.FindStringIndex(args[1])[1]] + "swch"
				var files []string
				filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
					files = append(files, path)
					return nil
				})
				acc := 0
				for _, a := range files {
					if filename == a {
						u := ""
						fmt.Println("this file already exists")
						fmt.Println("should i overwrite this? (y/n): ")
						fmt.Scanln(&u)
						if u == "y" {
							updateFile(args[1])
						}
						goto skip
					}
				}
				for _, a := range files {
					if a != args[1] {
						acc++
					}
				}
				if acc < len(files) {
					createSwchFile(args[1])
				} else {
					args[1] = enterArg()
					goto checkargs
				}
				skip:
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

func swch (fileName string) {
	check, _ := regexp.Compile(".+[.]")
	executiveName := fileName[:check.FindStringIndex(fileName)[1]] + "swch"
	qwe, _ := os.OpenFile(executiveName, os.O_RDWR, 0666)
    qwei, _ := qwe.Stat()
	byt := make([]byte, qwei.Size())
	qwe.Read(byt)
	findLines, _ := regexp.Compile("\n")
	listofnewline := findLines.FindAllIndex(byt, -1)
	filelines := make(map[int]string)
	fmt.Println(filelines)
	if len(listofnewline) < 1 {
		filelines[0] = string(byt[:])
	} else {
		for i, _ := range listofnewline {
			if i < 1 {
				filelines[i] = string(byt[:listofnewline[i][0]])
			} else {
				filelines[i] = string(byt[listofnewline[i-1][1]:listofnewline[i][0]])
			}
			if i >= len(listofnewline)-1 {
				filelines[i+1] = string(byt[listofnewline[i][1]:])
				break
			}
		}
	}
	fmt.Println(filelines[0])
}

func enterArg() string {
	u := ""
	fmt.Println("problem with base arg")
	fmt.Print("on which document shoud it be based: ")
	fmt.Scanln(&u)
	if len(u) < 1 {
		enterArg()
	} else {
		return u
	}
	return u
}

// only one content arg per call
func updateFile(fileName string, content ...[]byte) {
	check, _ := regexp.Compile(".+[.]")
	executiveName := fileName[:check.FindStringIndex(fileName)[1]] + "swch"

	executive, _ := os.OpenFile(executiveName, os.O_RDWR, 0666)
	entrydata, _ := os.OpenFile(fileName, os.O_RDWR, 0666)
	entrydataInfo, _ := entrydata.Stat()
	if len(content) < 1 {
		contentFromEntry := make([]byte, entrydataInfo.Size())
		entrydata.Read(contentFromEntry)
		fmt.Println(contentFromEntry)
		a := plainSwchGenerator(contentFromEntry)
		executive.Write(a)
		fmt.Println(a)
		return
	}
	executive.Truncate(0)
	executive.Write(content[0])
}

func plainSwchGenerator(fileContent []byte) []byte {
	check, _ := regexp.Compile("\n")
	// nlArr := check.FindAllStringSubmatchIndex(string(fileContent), -1)
	for c := 0; c < len(check.FindAllStringSubmatchIndex(string(fileContent), -1)); c++ {
		nlArr := check.FindAllStringSubmatchIndex(string(fileContent), -1)
		var cv []byte
		if c == 0 {
			cv = append([]byte(fmt.Sprint(c, " :: ")))
		}
		cv = append(cv, fileContent[:nlArr[c][0]]...)
		a := append(cv, []byte(fmt.Sprint("\n", c + 1, " :: "))...)
		b := append(a, fileContent[nlArr[c][1]:]...)
		fileContent = b
	}
	return fileContent
}

func createSwchFile(fileName string) {
    entrydata, _ := os.OpenFile(fileName, os.O_RDWR, 0666)
	entrydataInfo, _ := entrydata.Stat()
	fmt.Println("here", entrydataInfo)
    byt := make([]byte, entrydataInfo.Size())
    entrydata.Read(byt)
	entrydata.Close()
	byt = plainSwchGenerator(byt)
	check, _ := regexp.Compile(".+[.]")
	filename := entrydataInfo.Name()[:check.FindStringIndex(entrydataInfo.Name())[1]] + "swch"
	// koe = koe[:dotindex]
	genfile, _ := os.Create(filename)
	// os.Create(filename)
	genfile.Write(byt)
}