package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const mapSym = string(1)

//checks to see if a line scan worked, and prints why it failed if it did indeed fail
func passCheck(worked bool, err error) bool {

	if !worked {
		if err == nil {
			fmt.Println("Reached EOF!")
			return false
		} else {
			fmt.Println(err)
			return false
		}
	}
	return true

}

func main() {
	wordArray := []string{"", "", "", ""}
	wordMap := make(map[string]int)

	fileR, err := os.Open("wordlist.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//wordC starts at 3 to keep first four values empty because the unicode symbols are useful
	wordC := 3
	fileScan := bufio.NewScanner(fileR)

	//reads through the dictionary on file; in this case it reads through wordlist.txt
	for passCheck(fileScan.Scan(), fileScan.Err()) {
		wordC++
		wordMap[fileScan.Text()] = wordC
		wordArray = append(wordArray, fileScan.Text())

	}
	fileR.Close()

	//sets up text scanner for input
	textGet := bufio.NewScanner(os.Stdin)
	textGet.Scan()
	testInput := textGet.Text()

	encodeString := ""

	for _, t := range strings.Fields(testInput) {
		t = strings.ToLower(t)

		switch {
		case len(t) > 3 && wordMap[t] != 0:
			encodeString += mapSym + string(wordMap[t]) + " "

		default:
			encodeString += t + " "
		}
	}
	//cuts off trailing space character, easier to always do it and assume it is there
	//than messing with if statements all for the sake of one character
	encodeString = encodeString[:len(encodeString)-1]

	//passing the encoded string along as if these were fuctions
	toDecode := encodeString

	//need this to split fields by white space - because strings.Fields splits by all kinds of white space, not just normal space.
	splitF := func(c rune) bool {
		return c == ' '
	}

	decodeString := ""
	for _, t := range strings.FieldsFunc(toDecode, splitF) {
		temp := 0

		switch {
		case strings.HasPrefix(t, mapSym) && len(t) >= 2 && len(t) <= 4:

			temp = int([]rune(t[1:])[0])

			if wordC >= temp {
				decodeString += wordArray[temp] + " "
				continue

			}
			decodeString += t + " "

		default:
			decodeString += t + " "

		}
	}
	decodeString = decodeString[:len(decodeString)-1]

	fmt.Println("encoded:", encodeString)
	fmt.Println("decoded:", decodeString)

	fmt.Println("original bytes:", len(testInput), "\nmodified bytes:", len(encodeString))

}
