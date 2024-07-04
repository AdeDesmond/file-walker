package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func getFilePath() (string, error) {
    var path string
    paths := flag.String("paths", "Users", "enter the path for the directory separated by spaces if the path is longer")
    flag.Parse()
    if *paths == "" {
        return "", errors.New("path cannot be and empty string")
    }
    fmt.Println(flag.Args())
    if len(flag.Args()) > 0 {
        path = "/" + strings.Join(flag.Args(), "/")
    }else {
        path = "/" + flag.Args()[0]
    }
    return path, nil
}

//the idea is the user can start searching from a default paths
//so we can set the default path to be the home dir paths

func setDefaultPath()(string, error){
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Printf("%s", err)
        return "", err
    }

    return homeDir, nil
}
