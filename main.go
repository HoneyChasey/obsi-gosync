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


// Get a pointer to the file present folder and sub folder. Recursive
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

// Write into the archvie file the pointers who get with read_file_in_folder
func write_into_archive(files []*os.File, zipFile *os.File, rootFolder string) (error) {
	if files == nil || zipFile == nil {return fmt.Errorf("Internal Error")}

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	
	for _, file := range files {
	info, err := file.Stat()
  if err != nil {return fmt.Errorf("Cannot read%s : %w", file.Name(), err)}

	header, err := zip.FileInfoHeader(info)
  if err != nil {return fmt.Errorf("Cannot read header: %w", err)}

	relativePath, err := filepath.Rel(rootFolder, file.Name())
	if err != nil {return fmt.Errorf("Cannot calculate the relative path")}

	header.Name = relativePath
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {return fmt.Errorf("Cannot create the header: %w", err)}

	_, err = io.Copy(writer, file)
  if err != nil {return fmt.Errorf("Cannot coppy %s : %w", file.Name(), err)}

	}
	return nil
}

func unzip_archvie(src string, dest string) (error) {
	if src == "" || dest == "" {return fmt.Errorf("Internal Error")}
	reader, err := zip.OpenReader(src)
	if err != nil {return fmt.Errorf("Cannot read the archive or not exist")}

	for _, file := range reader.File {
		destPath := filepath.Join(dest, file.Name)
		os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		outFile, err := os.Create(destPath)
		if err != nil {return fmt.Errorf("Error while creating file, not exist of not have access to him")}

		defer outFile.Close()
		zipFile, err := file.Open()
		if err != nil {return fmt.Errorf("Cannot read the header file of 1 file")}
		defer zipFile.Close()
		_, err = io.Copy(outFile, zipFile)
		if err != nil {return fmt.Errorf("Error while copy and unzip file")}
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
	unzip_archvie("test.zip", "test2")
}
