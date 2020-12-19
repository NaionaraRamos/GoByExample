package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"os"
	//"strings"
)

//DIRECTORIES

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil
}

//DIRECTORIES

func main() {

	p := fmt.Println

	//FILE PATHS
	/* j := filepath.Join

	a := j("dir1", "dir2", "filename")
	p("a: ", a)

	p(j("dir1//", "filename"))
	p(j("dir1/../dir1", "filename"))

	p("Dir(a):", filepath.Dir(a))
	p("Base(a):", filepath.Base(a))

	p(filepath.IsAbs("dir/file"))
	p(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	p(ext)

	p(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	p(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}

	p(rel) */

	//FILE PATHS

	//DIRECTORIES

/* 	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	p("Listing subdir/parent")
	for _, entry := range c {
		p(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = ioutil.ReadDir(".")
	check(err)

	p("Listing subdir/parent/child")

	for _, entry := range c {
		p(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	p("visiting subdir")
	err = filepath.Walk("subdir", visit) */

	//DIRECTORIES


	//Temporary Files and Directories

	f, err := ioutil.TempFile("", "sample")
	check(err)

	p("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampleDir")
	check(err)
	p("Tem dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0655)
	check(err)

	//Temporary Files and Directories
}