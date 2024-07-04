package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//ability to search directories

func searchDirs(path string){
    entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
    for _, entry := range entries {
        if entry.IsDir() {
            fmt.Println(entry.Name())
            fmt.Println("FILE JOIN PATH",filepath.Join(entry.Name()))
        }
    }
    fmt.Println(entries)
}
