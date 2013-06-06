package cedict

//import "fmt"

type WordType int

const (
  WordTypeString = iota
  WordTypeRecords
)

type ChineseTextWord struct {
  T WordType
  S string
  R []Record
}

// Once implemented, this method will split a string of chinese text
// into a slice of words of type WordType.
// TODO comment this function
func SplitChineseTextIntoWords(text string) []ChineseTextWord {
  output := make([]ChineseTextWord,0)

  charSet := DetermineCharSet(text)

  index := 0

  for index < len([]rune(text)) {
    for substringLength := maxRunecount; substringLength >= 0; substringLength-=1 {
      if substringLength == 0 {
        output = append(output, ChineseTextWord{T: WordTypeString, S: string([]rune(text)[index]), R: nil})
        index +=1
        break
      }

      var substring string

      if index+substringLength > len([]rune(text))-1 {
        substring = string([]rune(text)[index:])
      } else {
        substring = string([]rune(text)[index:index+substringLength])
      }

      // TODO deal with error
      records,_ := FindRecords(substring, charSet)
      if len(records) > 0 {
        output = append(output, ChineseTextWord{T: WordTypeRecords, S: "", R: records})
        index += substringLength
        break
      }
    }
  }



  return output
}

