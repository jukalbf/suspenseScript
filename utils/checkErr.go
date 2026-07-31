package utils

import "log"

// Verifica se houve erro
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
