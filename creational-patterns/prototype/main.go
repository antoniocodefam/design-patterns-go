package main

import (
	"fmt"
	"strings"
)

func main(){
	folder1 := &Folder{
		name: "First folder",
		children:[]IItem{
			&File{
				name: "First file",
			},
			&File {
				name: "Second file",
			},
		} ,
	}

	folder2 := &Folder{
		name: "Second folder",
		children: []IItem{
			&File{
				name: "Third file",
			},
			folder1,
			&File{
				name: "Fourth file",
			},
		},
	}

	folder2.print(" ")

	fmt.Println(strings.Repeat("-",70))
    
	clonedFolder2 := folder2.clone()
	clonedFolder2.print(" ")
}