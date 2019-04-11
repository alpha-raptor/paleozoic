package main

import (
	"fmt"
	"os"
	str "strings"
)

type worker struct {
	id int
}

func occur(w worker, word string, extr []byte) int {
	occurr := 0
	aux := len(extr)
	for i := 0; i < (aux); i++ {
		if str.Compare(word, string(extr[i])) == 0 {
			occurr += 1
		}
	}
	return occurr
}

//IDEIA: DIVIDIR O TAMANHO DO ARQUIVO EM UMA QUANTIDADE X DE TRABALHADORES

func readBookByTXT(path string) []byte {
	book, err := os.Open(path)
	check(err)
	//pega a informação de um file
	lenght, err := book.Stat()
	check(err)
	eof := make([]byte, lenght.Size())
	//content, err := book.Read(eof)
	check(err)
	//fmt.Printf("%d bytes:%s\n", content, string(eof))
	//fmt.Printf("%d bytes \n", content)
	return eof
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//var s string = "/home/raptor/Documentos/go/sd/trabalho/dom_casmurro.txt"
	var s string = "/home/raptor/Documentos/go/sd/trabalho/test.txt"

	var book = readBookByTXT(s)
	var w1 worker
	w1.id = 1
	fmt.Println(occur(w1, "a", book))
}
