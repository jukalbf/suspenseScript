package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) {
	fmt.Println("=> Tentando ler conteúdo")

	_, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	os.Lstat(fileName)
}
