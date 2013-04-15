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
func SplitChineseTextIntoWords(text string) []ChineseTextWord {
  // TODO: implement this method (low priority)
  return nil
}
