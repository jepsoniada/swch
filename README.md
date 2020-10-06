# SWCH

## Indroducion

swch is tool for switching between dev and build file by makeing changes in selected file

## Instalation

download swch.go file and compile it with [go cli](https://golang.org/dl/)

## Basic Usage

### swch files

swch files are used for editing original file by simple commands

### swch syntax
**basic example for swch file would look like this:**
```log
-n 0 :: foo
-r 1 :: qwe
```
"-n" and "-r" are tasks; these defines what should be done to specific lines; digit, :: and "foo" in first line are identificator that tells which line should be mutated by task and what this line stores

**arguments:**

args is passed only to specific tasks (like -a); they are defined by single widespace and string input, for example:
```log
-a 0 :: foo
	baz
 jd
```

**tasks:**

- -n - does nothing
- -r - removes line in build
- -a - adds line after this task (unlimitet arguments)

**swch cli options:**

- swch build \*filename\* - switches to build version of file based on \*filename\*.swch tasks !don't edit in build 
- swch dev \*filename\* - switches to basic file mode
- swch gen \*filename\* - generates basic \*filename\*.swch file filled with -n tasks
