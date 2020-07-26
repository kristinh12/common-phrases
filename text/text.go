package text

import (
	"fmt"
	"regexp"
	"strings"
)

type TextData struct {
	Text          string   `json:"text"`
	CommonPhrases []string `json:"common_phrases"`
}

func (t *TextData) GetCommonPhrases() error {
	caseInsensitive := strings.ToLower(t.Text)
	noNewLines := strings.Replace(caseInsensitive, "\\n", " ", -1)
	fmt.Println(noNewLines)

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
	if err != nil {
		return fmt.Errorf("Regex did not compile: %w", err)
	}
	processedString := reg.ReplaceAllString(noNewLines, "")

	wordArray := strings.Split(processedString, " ")
	numPhrases := len(wordArray)

	var phrases map[string]int

	for i := 1; i <= numPhrases; i++ {
		currentPhrase := strings.Join(wordArray[i-1:i+2], " ")
		if _, present := phrases[currentPhrase]; present {
			phrases[currentPhrase] += 1
		} else {
			phrases[currentPhrase] = 1
		}
	}

	for key, value := range phrases {
		fmt.Println("inside phrases")
		fmt.Printf("The phrase %v occurs %v times", key, value)
	}

	// Creates a sorted array of the most common phrases.
	mostCommon := make([]string, len(phrases))
	for key, value := range phrases {
		for i, currentPhrase := range mostCommon {
			if mostCommon[i] == "" {
				mostCommon[i] = key
				break
			} else if value > phrases[currentPhrase] {
				if i == 0 {
					mostCommon = append([]string{key}, mostCommon...)
				} else {
					firstPart := append(mostCommon[:i], key)
					mostCommon = append(firstPart, mostCommon[i+1:]...)
				}
				break
			}
		}
	}

	if len(mostCommon) > 100 {
		mostCommon = mostCommon[:100]
	} else if len(mostCommon) > len(phrases) {
		mostCommon = mostCommon[:len(phrases)]
	}

	t.CommonPhrases = mostCommon
	return nil
}
