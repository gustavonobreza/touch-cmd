package main

import (
	"os"
)

func main() {
	fileName := os.Args[1]

	_file, err := os.Stat(fileName)

	if _file != nil {
		println("errror: this file alredy exists!")
		os.Exit(1)
	} else if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		defer file.Close()
	} else {
		println("errror: ", err.Error())
		os.Exit(1)
	}

}
