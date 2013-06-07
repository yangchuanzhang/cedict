package cedict

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
func SplitChineseTextIntoWords(text string) ([]ChineseTextWord, error) {
  output := make([]ChineseTextWord,0)

  charSet := DetermineCharSet(text)

  index := 0

  for index < len([]rune(text)) {
    // try to find the next word by going from the longest word length down to zero
    for substringLength := maxRunecount; substringLength >= 0; substringLength-=1 {

      // there's no word in the dictionary at the current index
      if substringLength == 0 {
        output = append(output, ChineseTextWord{T: WordTypeString, S: string([]rune(text)[index]), R: nil})
        index +=1
        break
      }

      // get next string of length substringLength
      var substring string
      if index+substringLength > len([]rune(text))-1 {
        substring = string([]rune(text)[index:])
      } else {
        substring = string([]rune(text)[index:index+substringLength])
      }

      records,err := FindRecords(substring, charSet)
      if err != nil {
        return nil, err
      }
      if len(records) > 0 {
        output = append(output, ChineseTextWord{T: WordTypeRecords, S: "", R: records})
        index += substringLength
        break
      }
    }
  }



  return output, nil
}

