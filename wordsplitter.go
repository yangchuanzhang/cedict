package cedict

import (
  "github.com/yangchuanzhang/chinese"
)

type WordType int

cont (
  WordTypeString = iota
  WordTypeRecords
)

type ChineseTextWord struct {
  T WordType
  S string
  R []records
}

func SplitChineseTextIntoWords(text string) []ChineseTextWord {
  
  return nil
}
