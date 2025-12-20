package main

import "fmt"

type Folder struct{
	name string
	children []IItem
}

func (f *Folder) print(indentation string){
	fmt.Println(indentation + f.name)

	for _, child := range f.children{
		child.print(indentation + " ")
	}
}

func (f *Folder) clone() IItem {
	childrenCopies := make([]IItem, 0, len(f.children))

	for _, child := range(f.children){
		childrenCopies = append(childrenCopies, child.clone())
	}

	return &Folder{
		name: "Copy of " + f.name,
		children: childrenCopies,
	}
}