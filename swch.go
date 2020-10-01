// wert gen (sorce = file name String) - generates
// file based on /sorce/ file; that would look like
// "*name*.swch" (file = file name String) - makes
// changes in /file/ acording to "*fileName*.swch"
//  wert (^(-a|-n) \d :: .*)
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type simplifiedLine struct {
	content string
}

type simfileinfos interface {
	task() string
	line() string
}

func (sf simplifiedLine) task() string {
	check, _ := regexp.Compile("^-(a|n|r)")
	typeOfTask := check.FindStringSubmatch(sf.content)
	return string(typeOfTask[1])
}

func (sf simplifiedLine) line() string {
	check, _ := regexp.Compile(" :: (.*)")
	line := check.FindStringSubmatch(sf.content)
	return line[1]
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
checkargs:
	switch args[0] {
	case "b", "build":
		var files []string
		filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
			files = append(files, path)
			return nil
		})
		check, _ := regexp.Compile(".+[.]")
		fmt.Println(check.FindStringIndex(args[1]))
		executiveName := args[1][:check.FindStringIndex(args[1])[1]] + "swch"
		fmt.Println(executiveName, args[1])
		acc := 0
		fin := false
		for _, a := range files {
			if a == executiveName {
				fin = true
			}
			if a != args[1] {
				acc++
			}
		}
		if acc == len(files)-1 && fin {
			swch(args[1], true)
		}
		if acc == len(files) {
			fmt.Println("not file, sorry ::((")
		} else if !fin {
			u := ""
			fmt.Println("executive to this file doesn't exist")
			fmt.Println("should i generate it? (y/n)")
			fmt.Scanln(&u)
			if u == "y" {
				createSwchFile(args[1])
			}
		}
	case "d", "dev":
		var files []string
		filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
			files = append(files, path)
			return nil
		})
		check, _ := regexp.Compile(".+[.]")
		executiveName := args[1][:check.FindStringIndex(args[1])[1]] + "swch"
		acc := 0
		fin := false
		for _, a := range files {
			if a == executiveName {
				fin = true
			}
			if a != args[1] {
				acc++
			}
		}
		if acc == len(files)-1 && fin {
			swch(args[1], false)
		}
		if acc == len(files) {
			fmt.Println("not file, sorry ::((")
		} else if !fin {
			u := ""
			fmt.Println("executive to this file doesn't exist")
			fmt.Println("should i generate it? (y/n)")
			fmt.Scanln(&u)
			if u == "y" {
				createSwchFile(args[1])
			}
		}
	case "gen", "generate":
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
func swch(fileName string, build bool) {
	check, _ := regexp.Compile(".+[.]")
	executive := fileName[:check.FindStringIndex(fileName)[1]] + "swch"
	nlCheck, _ := regexp.Compile("\n")
	executiveFile, _ := os.OpenFile(executive, os.O_RDWR, 0666)
	exeFileInfo, _ := executiveFile.Stat()
	executiveContent := make([]byte, exeFileInfo.Size())
	executiveFile.Read(executiveContent)
	executiveFile.Close()
	listOfExecNls := nlCheck.FindAllStringSubmatchIndex(string(executiveContent), -1)
	linesOfExecutive := make(map[int]simplifiedLine)
	for i, a := range listOfExecNls {
		if i < 1 {
			linesOfExecutive[i] = simplifiedLine{string(executiveContent)[:a[0]]}
		} else {
			linesOfExecutive[i] = simplifiedLine{string(executiveContent)[listOfExecNls[i-1][1]:a[0]]}
		}
		if i == len(listOfExecNls)-1 {
			linesOfExecutive[i+1] = simplifiedLine{string(executiveContent)[a[1]:]}
		}
	}
	originFile, _ := os.OpenFile(fileName, os.O_RDWR, 0666)
	origFileInfo, _ := originFile.Stat()
	originContent := make([]byte, origFileInfo.Size())
	originFile.Read(originContent)
	listOfOrigNls := nlCheck.FindAllStringSubmatchIndex(string(originContent), -1)
	fmt.Println(listOfOrigNls)
	var linesOfOrigin []string
	fmt.Println(len(listOfOrigNls))
	for i, a := range listOfOrigNls {
		if i < 1 {
			linesOfOrigin = append(linesOfOrigin, string(originContent)[0:a[0]])
		} else {
			linesOfOrigin = append(linesOfOrigin, string(originContent)[listOfOrigNls[i-1][1]:a[0]])
		}
		if i == len(listOfOrigNls)-1 {
			linesOfOrigin = append(linesOfOrigin, string(originContent)[a[1]:])
		}
	}
	// for _, a := range linesOfOrigin {
	// 	fmt.Println(a)
	// }
	for i := 0; i < 5; i++ {
		fmt.Println(linesOfExecutive[i].content)
	}
	buffOfInputFile := ""
	var checkLater []string
	for i := len(linesOfExecutive) - 1; i >= 0; i-- {
		e, _ := regexp.MatchString(`^-(a|n|r) \d :: .*`, linesOfExecutive[i].content)
		if e {
			switch linesOfExecutive[i].task() {
			case "n":
				if i == len(linesOfExecutive)-1 {
					buffOfInputFile = linesOfExecutive[i].line()
				} else {
					buffOfInputFile = linesOfExecutive[i].line() + "\n" + buffOfInputFile
				}
			case "a":
				fmt.Println("a" + "a")
				if build {
					for _, a := range checkLater {
						if i == len(linesOfExecutive)-1 {
							buffOfInputFile = a
						} else {
							buffOfInputFile = a + "\n" + buffOfInputFile
						}
					}
				}
				if i == len(linesOfExecutive)-1 {
					buffOfInputFile = linesOfExecutive[i].line()
				} else {
					buffOfInputFile = linesOfExecutive[i].line() + "\n" + buffOfInputFile
				}
			case "r":
				if !build {
					if i == len(linesOfExecutive)-1 {
						buffOfInputFile = linesOfExecutive[i].line()
					} else {
						buffOfInputFile = linesOfExecutive[i].line() + "\n" + buffOfInputFile
					}
				}
			}
		} else if a, _ := regexp.MatchString(`^\s.*`, linesOfExecutive[i].content); a {
			checkLater = append([]string{linesOfExecutive[i].content}, checkLater...)
		}
	}
	fmt.Println([]byte(buffOfInputFile))
	fmt.Println(originContent)
	originFile.Truncate(0)
	originFile.WriteAt([]byte(buffOfInputFile), 0)
	originFile.Close()
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
	executive.Close()
	entrydata.Close()
}
func plainSwchGenerator(fileContent []byte) []byte {
	check, _ := regexp.Compile("\n")
	for c := 0; c < len(check.FindAllStringSubmatchIndex(string(fileContent), -1)); c++ {
		nlArr := check.FindAllStringSubmatchIndex(string(fileContent), -1)
		var cv []byte
		if c == 0 {
			cv = append([]byte(fmt.Sprint("-n ", c, " :: ")))
		}
		cv = append(cv, fileContent[:nlArr[c][0]]...)
		a := append(cv, []byte(fmt.Sprint("\n", "-n ", c+1, " :: "))...)
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
	genfile, _ := os.Create(filename)
	genfile.Write(byt)
	entrydata.Close()
}
