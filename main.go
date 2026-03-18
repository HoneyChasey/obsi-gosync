package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// thx to https://golang.cafe/blog/golang-zip-file-example

// is the arry of all file of the obsidian archive.
var pointers []*os.File



// Create the zip archive
func create_zip(name string) (*os.File, error){
	if name == ""{
		return nil, fmt.Errorf("Please provide a path")
	}
	archive, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return archive, nil
}


// get a pointer to the file present in the folder
func read_file_in_folder(pathfolder string, pointers []*os.File) ([]*os.File, error) {
	files, err := os.ReadDir(pathfolder)
	if err != nil {return nil, fmt.Errorf("Please provide correct path to existing folder or cannot access")}
	for _, entry := range(files){
		if entry.IsDir(){continue;}
		fullpath := filepath.Join(pathfolder, entry.Name())
		file, err := os.Open(fullpath)
		if err != nil {return nil, fmt.Errorf("Error cannot read or not exist")}
		pointers = append(pointers, file)
	}
return pointers, nil
}


func write_into_archive(files []*os.File, zipFile *os.File) (error) {
	if files == nil || zipFile == nil {return fmt.Errorf("Internal Error")}

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	
	for _, file := range files {
	info, err := file.Stat()
  if err != nil {
  	return fmt.Errorf("impossible de lire %s : %w", file.Name(), err)
  }

	header, err := zip.FileInfoHeader(info)
  if err != nil {
  	return fmt.Errorf("impossible de créer le header : %w", err)
  }

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
  	return fmt.Errorf("impossible d'écrire le header : %w", err)
  }

	_, err = io.Copy(writer, file)
  if err != nil {
  	return fmt.Errorf("impossible de copier %s : %w", file.Name(), err)
  }
	}
	return nil
}


func main(){
	archive, _ := create_zip("test.zip")
	files, err := read_file_in_folder("test", pointers)
	if err != nil {fmt.Println(err)}
	fmt.Print(files)
	fmt.Println("Creation of zip archive")
	write_into_archive(files, archive)
}
