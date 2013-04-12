package cedict

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "fmt"
  "github.com/yangchuanzhang/chinese"
)

type Record struct {
  Simp string
  Trad string
  Pinyin string
  English string
}

var db *sql.DB
var dbLoaded = false

func LoadDb() (err error) {
  if !dbLoaded {
    //FIXME get path to db file from somewhere else
    db, err = sql.Open("sqlite3", "/Users/json/cedict.sqlite3")
    if err == nil {
      dbLoaded = true
    }
  }
  return
}

func CloseDb() {
  if dbLoaded {
    db.Close()
    dbLoaded = false
  }
}

func FindRecords(word string, charSet chinese.CharacterSet) ([]Record, error) {
  if !dbLoaded {
    return nil, fmt.Errorf("cedict: Database not loaded")
  }

  sql := "SELECT * FROM dict "

  switch charSet {
    case chinese.Trad: 
    sql += "WHERE trad = '"+word+"'"
    case chinese.Simp: 
    sql += "WHERE simp = '"+word+"'"
  default:
    return nil, fmt.Errorf("cedict: unrecognized character set")
  }

  rows, err := db.Query(sql)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  records := make([]Record, 0)

  for rows.Next() {
    var id, runecount int
    var trad, simp, pinyin, english string
    rows.Scan(&id, &trad, &simp, &pinyin, &english, &runecount)
    records = append(records, Record{Trad: trad, Simp: simp, Pinyin: pinyin, English: english})
  }

  return records, nil
}

func main() {
  LoadDb()

  fmt.Println(FindRecords("的", chinese.Trad))

}
