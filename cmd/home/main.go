package main

import "log"

func main(){
    getFilePath()
    defaultPath, err := setDefaultPath()
    if err != nil {
        log.Fatal(err)
    }
    searchDirs(defaultPath)
}
