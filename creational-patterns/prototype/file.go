package main

import "fmt"

type File struct{
	name string
}

func (f *File) print(indentation string) {
    fmt.Println(indentation + f.name)
}

func (f *File) clone() IItem {
	return &File{
		name: "Copy of " + f.name,
	}
}