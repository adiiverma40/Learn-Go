package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to files manipulation in go lang")

	content := "This needs to be written in file using go lang, \nlets see will it overwrite itslef,\n it does overwrite"
	file, err := os.Create("./myfiles.txt")
	if err != nil{
		panic(err)
	}

	lenght, err := io.WriteString(file,content)
	checkNilError(err)

	//woudnt it be better if 
	// 
	// lenght, checkNilError(err) := io.WriteString(file,content)
	
	// if err != nil{
	// 	panic(err)
	// }
	fmt.Println("the text of lenght: ", lenght ,"is created")
	defer file.Close()
	readFile(file.Name())
}


func readFile(filename string){
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	fmt.Println("text raw data is \n", dataByte)
	fmt.Println("text  is \n", string(dataByte))
	
}


func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}