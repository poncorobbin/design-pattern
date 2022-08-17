package main

import "fmt"

type component interface {
	search(keyword string)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in file '%s'\n", keyword, f.name)
}

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword '%s' in folder '%s'\n", keyword, f.name)
	for _, c := range f.components {
		c.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

func main() {
	file1 := &file{"file 1"}
	file2 := &file{"file 2"}
	file3 := &file{"file 3"}

	folder1 := &folder{name: "folder 1"}
	folder1.add(file1)

	folder2 := &folder{name: "folder 2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	folder2.search("asdf")

}
