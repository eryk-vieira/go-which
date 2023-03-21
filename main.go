package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("not enought arguments, please specify the file/program you want to discover the path")

		return
	}

	file := args[1]

	path := os.Getenv("PATH")

	splitedPath := filepath.SplitList(path)

	for _, directory := range splitedPath {
		fullPath := filepath.Join(directory, file)

		stat, err := os.Stat(fullPath)

		if err != nil {
			continue
		}

		mode := stat.Mode()

		if mode.IsRegular() {
			if mode&0111 != 0 {
				var size float64 = float64(stat.Size()) / 1e6
				var metric string = "MB"

				if size < 1 {
					size = size * 1000

					metric = "KB"
				}

				fmt.Println("Path:", fullPath)
				fmt.Printf("Size: %.2f %s\n", size, metric)
			}
		}

		return
	}

	fmt.Println("Executable not found")
	return
}
