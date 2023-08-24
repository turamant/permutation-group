package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func init(){
	buildDictionary()
}

var dictionary map[string][]string


func alphabetize(word string) string{
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func buildDictionary(){
	dictionary = make(map[string][]string)
	file,err := os.Open("words.txt")
	
	if err != nil {
		log.Printf("Error : %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtwords []string
    
	for scanner.Scan(){
		txtwords = append(txtwords, scanner.Text())
	}
	
	file.Close()
	for _, v := range txtwords{
		fmt.Println(v)
	}
	for _, word := range txtwords{
	 	alphabetized := alphabetize(word)
		var list []string
		if len(dictionary) > 0 && len(dictionary[alphabetized]) > 0{
			list = dictionary[alphabetized]
		}else{
			list = []string{}
		}
		list = append(list, word)
		dictionary[alphabetized] = list
		
	}
	
}

func output(word string){
	wd := alphabetize(word)
	fmt.Printf("Permutation group of %s is %s", word, dictionary[wd])

}

func main(){
	output("esprits")
	
	

}