package training

import (
	"io"
	"os"
)

func FilesTrainig() {
	// file, _ := os.Create("hello.txt")
	data := []byte("Hello Bold!")
	// var dataFile []byte
	// buffer := make([]byte, 64)
	// file,_ := os.OpenFile("hello.txt",os.O_WRONLY, 0666) // для чтения, а O_WRONLY для записи
	file, _ := os.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	file.Write(data)
	file.Close()  
	file, _ = os.Open("hello.txt")
	// defer file.Close()
	// for {
	//   fmt.Println("Cirle")
	//   n, err := file.Read(buffer)
    //   if err == io.EOF {
	//   	 break
	//   }
	//   dataFile = append(dataFile, buffer[:n]...)
	// }
	io.Copy(os.Stdout, file)
}