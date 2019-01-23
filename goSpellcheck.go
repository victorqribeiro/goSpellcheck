package main

import (
	f "fmt"
	io "io/ioutil"
	"os"
	s "strings"
)

func readDic(file string) []string {
	b, err := io.ReadFile(file)
	if err != nil {
		f.Println(err)
	}
	return s.Fields(string(b))
}

func distance(wordA string, wordB string) uint {
	wordMin := wordA
	wordMax := wordB
	if len(wordA) > len(wordB) {
		wordMin = wordB
		wordMax = wordA
	}
	var dist uint = 0
	for i := 0; i < len(wordMin); i++ {
		if wordMin[i] != wordMax[i] {
			dist += 1
		}
	}
<<<<<<< HEAD
	dist += uint(len(word_max) - len(word_min))
=======
	if len(wordMin) != len(wordMax) {
		for i := 0; i < (len(wordMax) - len(wordMin)); i++ {
			dist += 1
		}
	}
>>>>>>> 2d0a3d75afb0aba9cf0613293dbeadfb00cca929
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
	minDistances := [3]uint{^uint(0), ^uint(0), ^uint(0)}
	if !(inArray(word,dic)) {
		for i := 0; i < len(dic); i++ {
			d := distance(word,dic[i])
			for j := 0; j < 3; j++ {
<<<<<<< HEAD
				if d < min_distances[j] {
					min_distances[j] = d
=======
				if d < minDistances[j] && !(inArray(dic[i],sug)) {
					minDistances[j] = d
>>>>>>> 2d0a3d75afb0aba9cf0613293dbeadfb00cca929
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
		f.Println("Error")
		os.Exit(1)
	}

	word := string( os.Args[1] )
	// dictionary from https://github.com/dwyl/english-words/blob/master/words.txt
	dic := readDic("words.txt")
	result := compare(word,dic)

	if len(result[0]) > 0 {
		f.Println( s.Join(result, " ") )
	} else {
		f.Println("no errors")
	}
}
