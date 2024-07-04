package main

import (
	"fmt"
	"testing"
)


func Test_GetCommandLinePath(t *testing.T){
    paths, err := getFilePath()
    if err != nil {
        t.Errorf("err from the path:%s", err)
    }
    fmt.Println(paths)
}


