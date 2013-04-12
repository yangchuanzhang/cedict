package cedict

import (
  "github.com/yangchuanzhang/chinese"
)

func DetermineCharacterSet(text string) chinese.CharacterSet {
  if !dbLoaded {
    // FIXME deal with errors
    LoadDb()
    defer CloseDb()
  }

  for _,c := range text {
    isTradChar := false
    tradRecords, _ := FindRecords(string(c), chinese.Trad)
    if len(tradRecords) > 0 {
      isTradChar = true
    }

    if !isTradChar {
      simpRecords, _ := FindRecords(string(c), chinese.Simp)
      if len(simpRecords) > 0 {
        return chinese.Simp
      }
    }
  }

  return chinese.Trad
}



