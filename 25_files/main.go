package main

import (
	"os"
)

func main() {

	// Example of getting file info
	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// println(fileInfo.Name())
	// println(fileInfo.Size())

	// Example of reading from a file
	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// buff := make([]byte, 10)
	// d, err := f.Read(buff)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(d, string(buff))

	// data, err := os.ReadFile("file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// f, err := os.Create("file2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// _, err = f.WriteString("Hello World")
	// if err != nil {
	// 	panic(err)
	// }

	// bytes := []byte("Hello World")
	// _, err = f.Write(bytes)
	// if err != nil {
	// 	panic(err)
	// }

	//Read and Write to a file

	// sourceFile, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("file2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("File copied successfully")

	// Deleting a file
	err := os.Remove("file2.txt")
	if err != nil {
		panic(err)
	}

}
