package codewars

import (
	"log"
	"slices"
	"sort"
	"strings"
)

var messages = []string{"dance on the table", "hide my beers", "scouts rocks"}
var secrets = []string{"egncd pn thd tgbud", "hked mr bddys", "scplts ypcis"}
 //=> agdeikluopry
func Codewars() {
	result := FindTheKey(messages, secrets)
	log.Println(result)
}


func FindTheKey(messages []string, secrets []string) string {
	cryptResultWordArray := []string{}
	for index ,message := range messages {
        cryptWordArray := strings.SplitAfter(message, "")
        secretsWordArray := strings.SplitAfter(secrets[index], "")
		for indexI, messageSecond := range cryptWordArray {
          if messageSecond != secretsWordArray[indexI] {
			valueAppended := messageSecond + secretsWordArray[indexI]
			if idx := slices.Index(cryptResultWordArray, valueAppended); idx != -1 { 
			  continue
			}
			valueAppendedSortedArr := strings.SplitAfter(valueAppended, "")
			sort.Strings(valueAppendedSortedArr)
			valueAppendedSorted := strings.Join(valueAppendedSortedArr,"")
			if idx := slices.Index(cryptResultWordArray, valueAppendedSorted); idx != -1 { 
			  continue
			}
			cryptResultWordArray = append(cryptResultWordArray, valueAppendedSorted)
		  }
		}
	}
	slices.Sort(cryptResultWordArray)
	resultString := strings.Join(cryptResultWordArray,"")
	return resultString
}

func GetCount(str string)  int {
	// Enter solution here
	var strArr = strings.Split(str,"")
	var count int = 0
	var vowelsArray = []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}

	for _ , value := range strArr {
		isVowels := slices.Contains(vowelsArray, value)
		if (isVowels) {
			count++
		}
		
	}
	return count
}