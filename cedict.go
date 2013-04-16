/*
TODO: add package description
*/
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

// LoadDb opens the database file. It's not necessary to call this function
// explicitly, as all functions that access the db open and close it themselves,
// if the database is not open, but it can improve performance by avoiding constant
// calls to sql.Open and Close.
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

// CloseDb closes the database connection if it's open. Otherwise, it does nothing.
func CloseDb() {
  if dbLoaded {
    db.Close()
    dbLoaded = false
  }
}

func isDbLoaded() bool {
  return dbLoaded
}

// FindRecords searches the database of cedict records and returns a slice of type
// []Record and an error. It returns an empty slice if no matches could be found.
func FindRecords(word string, charSet chinese.CharSet) ([]Record, error) {
  if !dbLoaded {
    err := LoadDb()
    if err != nil {
      return nil, err
    }
    defer CloseDb()
  }

  // construct db query based on charSet
  sql := "SELECT * FROM dict "

  switch charSet {
    case chinese.Trad: 
    sql += "WHERE trad = '"+word+"'"
    case chinese.Simp: 
    sql += "WHERE simp = '"+word+"'"
  default:
    return nil, fmt.Errorf("cedict: unrecognized character set")
  }

  // execute the query and defer closing it
  rows, err := db.Query(sql)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  // create slice to hold records
  records := make([]Record, 0)

  // populate records with the data from the db query
  for rows.Next() {
    var id, runecount int
    var trad, simp, pinyin, english string
    rows.Scan(&id, &trad, &simp, &pinyin, &english, &runecount)
    records = append(records, Record{Trad: trad, Simp: simp, Pinyin: pinyin, English: english})
  }

  return records, nil
}

// This method implements the Stringer interface for Record
func (r Record) String() string {
  return fmt.Sprintf("[simp: %q  trad: %q  pinyin: %q  english: %q]", r.Simp, r.Trad, r.Pinyin, r.English)
}
