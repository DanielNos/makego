package main

import "os"

func makeDirs(paths []string, perm os.FileMode) {
	for _, path := range paths {
		os.MkdirAll(path, perm)
	}
}

func writeLine(file *os.File, line string) {
	file.WriteString(line + "\n")
}
