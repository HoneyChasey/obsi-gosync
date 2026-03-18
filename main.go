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
func read_file_in_folder(pathfolder string) ([]*os.File, error) {
	entries, err := os.ReadDir(pathfolder)
	if err != nil {return nil, fmt.Errorf("Please provide correct path to existing folder or cannot access")}

	var pointers []*os.File

	for _, entry := range(entries){
		if entry.IsDir(){
			newPathFolder := filepath.Join(pathfolder, entry.Name())
			new_pointer, err := read_file_in_folder(newPathFolder)	
			if err != nil {return nil, fmt.Errorf("Error, cannot read the folder recursive")}
			pointers = append(pointers, new_pointer...) // Wtf is ...
			continue
		}

		fullpath := filepath.Join(pathfolder, entry.Name())
		file, err := os.Open(fullpath)
		if err != nil {return nil, fmt.Errorf("Error cannot read or not exist")}
		pointers = append(pointers, file)
	}
	return pointers, nil
}

func write_into_archive(files []*os.File, zipFile *os.File, rootFolder string) (error) {
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

	relativePath, err := filepath.Rel(rootFolder, file.Name())
	if err != nil {return fmt.Errorf("Canno't calculate the relative path")}

	header.Name = relativePath
	header.Method = zip.Deflate

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
	var folderName = "test"
	rootPath := filepath.Join("/Users/honeychasey/Documents/projetPerso/go/obsi-go-sync/obsi-gosync", folderName)
	archive, _ := create_zip("test.zip")
	files, err := read_file_in_folder(rootPath)
	if err != nil {fmt.Println(err)}
	fmt.Print(files)
	write_into_archive(files, archive, rootPath)
}
