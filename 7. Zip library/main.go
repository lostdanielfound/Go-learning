package main

import (
	"archive/zip" //used to read zip files
	"fmt"         //used for stdio output
	"io"
	"log" //User to log error output
	"os"
)

func main() {
	//open a zip archive for reading.
	filehandlerR, err := zip.OpenReader("SecureFolder.zip")
	if err != nil { //error checking to see if we were able to get a filehandler
		log.Fatal(err)
	}

	defer filehandlerR.Close() //close the filehandler as soon as the scope of this function ends

	//Iterate through the files in the archive,
	//printing some of their contents
	for _, f := range filehandlerR.File {
		fmt.Printf("Contents of %s:\n", f.Name)

		fileContent, err := f.Open() //opening the file
		if err != nil {              //error checking
			log.Fatal(err)
		}

		_, err = io.CopyN(os.Stdout, fileContent, 68) //Copy 68 bytes of each file into stdout
		if err != nil {
			log.Fatal(err)
		}

		fileContent.Close()
		fmt.Println()
	}
}
