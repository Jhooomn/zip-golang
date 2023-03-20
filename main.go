package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    // Create a new zip archive.
    zipFile, err := os.Create("example.zip")
    if err != nil {
        panic(err)
    }
    defer zipFile.Close()

    // Create a new writer for the zip archive.
    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    // Add files to the zip archive.
    files := []string{"file1.txt", "file2.txt", "file3.txt"}
    for _, file := range files {
        // Open the file to be added to the archive.
        fileToZip, err := os.Open(file)
        if err != nil {
            panic(err)
        }
        defer fileToZip.Close()

        // Create a new file header for the file.
        fileInfo, err := fileToZip.Stat()
        if err != nil {
            panic(err)
        }
        header, err := zip.FileInfoHeader(fileInfo)
        if err != nil {
            panic(err)
        }

        // Set the file header name to the name of the file.
        header.Name = file

        // Add the file header to the zip archive.
        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            panic(err)
        }

        // Write the file contents to the zip archive.
        _, err = io.Copy(writer, fileToZip)
        if err != nil {
            panic(err)
        }
    }

    // Print a message indicating that the zip archive was created successfully.
    fmt.Println("Zip archive created successfully!")
}

