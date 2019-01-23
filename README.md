# goSpellcheck

A terrbile spell checker written in Go.

# About

Since I've never really done anything in Go, I’ve decided to create a simple spell check. So I went to [https://github.com/dwyl/english-words/](https://github.com/dwyl/english-words/) and Got a bunch of words as a dictionary (words.txt).

After that I wrote some Go code to check if a given word was present in the dictionary, and if it’s not, gather the 3 more similar words and return as suggestions. The first time I wrote a similar code, in python, I used the euclidean distance and edit distance. This time around, however, I only check letter for letter and giver a plus 1 if they are different.

After that, I went to [Lists of common misspellings/For machines](https://en.wikipedia.org/wiki/Wikipedia:Lists_of_common_misspellings/For_machines) and tested my spell check against it with a little bash script I wrote.

It was a nice little project I did just so I don't get too rust while I'm on college vacation.

# Hot to use

Just call the the *goSpellcheck.go* file passing a word as parameter.

```
go run goSpellcheck.go word
```

# Score

The spell checker got 1185 from 4278 words, around 27,7%.
