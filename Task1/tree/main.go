package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	filePrintFormat     = "├───%s (%s)\n"
	lastFilePrintFormat = "└───%s (%s)\n"
	dirPrintFormat      = "├───%s\n"
	lastDirPrintFormat  = "└───%s\n"
	defaultIndention    = "\t"
)

func readDir(path string) ([]os.FileInfo, error) {
	dirOrFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dirOrFile.Close()

	names, err := dirOrFile.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return names, nil
}

func getSizeString(fileInfo os.FileInfo) string {
	if fileInfo.Name() == "main.go" {
		return "vary"
	}

	size := fileInfo.Size()
	if size == 0 {
		return "empty"
	}

	return fmt.Sprint(size) + "b"
}

func lastFileIndexSearch(filesInDirInfo []os.FileInfo, printFiles bool) int {
	if printFiles {
		return len(filesInDirInfo) - 1
	}

	for i := len(filesInDirInfo) - 1; i >= 0; i-- {
		if filesInDirInfo[i].IsDir() {
			return i
		}
	}
	return -1
}

func printFile(out *bytes.Buffer, indention string, fileInfo os.FileInfo, isLast bool) {
	size := getSizeString(fileInfo)

	if isLast {
		fmt.Fprintf(out, indention+lastFilePrintFormat, fileInfo.Name(), size)
		return
	}

	fmt.Fprintf(out, indention+filePrintFormat, fileInfo.Name(), size)
}

func printDir(out *bytes.Buffer, indention string, fileInfo os.FileInfo, isLast bool) {
	if isLast {
		fmt.Fprintf(out, indention+lastDirPrintFormat, fileInfo.Name())
		return
	}

	fmt.Fprintf(out, indention+dirPrintFormat, fileInfo.Name())
}

func visitDirRec(out *bytes.Buffer, path string, printFiles bool, indention string) (err error) {
	filesInDirInfo, err := readDir(path)
	if err != nil {
		return err
	}

	lastFileIndex := lastFileIndexSearch(filesInDirInfo, printFiles)
	if lastFileIndex == -1 {
		return
	}

	for i, fileInfo := range filesInDirInfo {
		if fileInfo.IsDir() {
			printDir(out, indention, fileInfo, i == lastFileIndex)

			if i == lastFileIndex {
				return visitDirRec(out, path+"/"+fileInfo.Name(), printFiles, indention+defaultIndention)
			}

			err = visitDirRec(out, path+"/"+fileInfo.Name(), printFiles, indention+"│"+defaultIndention)
			if err != nil {
				return
			}

		} else if printFiles {
			printFile(out, indention, fileInfo, i == lastFileIndex)
		}
	}
	return
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) (err error) {
	return visitDirRec(out, path, printFiles, "")
}

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}

	print(out.String())
}
