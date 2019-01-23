package main

import (
	"os"
	f "fmt"
	io "io/ioutil"
	s "strings"
)

func read_dic(file string) []string {
	b, err := io.ReadFile(file)
	if err != nil {
		f.Println(err)
	}
	return s.Fields(string(b))
}

func distance(word_a string, word_b string) uint {
	word_min := word_a
	word_max := word_b
	if( len(word_a) > len(word_b) ){
		word_min = word_b
		word_max = word_a
	}
	var dist uint = 0
	for i := 0; i < len(word_min); i++ {
		if word_min[i] != word_max[i] {
			dist += 1
		}
	}
	dist += uint(len(word_max) - len(word_min))
	return dist
}

func inArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func compare(word string, dic []string) []string {
	sug := []string{"", "", ""}
	min_distances := [3]uint{^uint(0), ^uint(0), ^uint(0)}
	if !(inArray(word,dic)) {
		for i := 0; i < len(dic); i++ {
			d := distance(word,dic[i])
			for j := 0; j < 3; j++ {
				if d < min_distances[j] {
					min_distances[j] = d
					sug[j] = dic[i]
					break
				}
			}
		}
	}
	return sug
}

func main(){
	
	if len(os.Args) < 2 {
		f.Println("Erro")
		os.Exit(1)
	}

	word := string( os.Args[1] )
	// dictionary from https://github.com/dwyl/english-words/blob/master/words.txt
	dic := read_dic("words.txt")
	result := compare(word,dic)

	if len(result[0]) > 0 {
		f.Println( s.Join(result, " ") )
	} else {
		f.Println("no errors")
	}

	
}
