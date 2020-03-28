package main

import (
	"log"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("file")

	file, err := os.OpenFile("./sample2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Hello World!\r\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString("こんにちは世界\r\n")
	if err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
