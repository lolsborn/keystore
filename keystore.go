package keystore

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Keystore struct {
	db *sql.DB
}

const (
	create_table  string = "create table Settings (key text not null primary key, val text, ival number)"
	// Default database path
	default_db_path string = "./settings.db"
)

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func NewKeystore(path string) (ks *Keystore, err error) {
	ks = new(Keystore)
	var created bool = false
	exists, _ := exists(path)
	if exists == false {
		created = true
	}
	db, err := sql.Open("sqlite3", path)
	ks.db = db
	if created {
		log.Printf("New Keystore created at %s", path)
		ks.db.Exec(create_table)
	}
	return ks, err
}

func DefaultKeystore() (ks *Keystore, err error) {
	return NewKeystore(default_db_path)
}

func (ks *Keystore) Delete(key string) {
	ks.db.Exec("DELETE FROM Settings WHERE key = ?", key)
}

func (ks *Keystore) PutString(key, value string) {
	ks.db.Exec("REPLACE INTO Settings (key, val) VALUES (?, ?)", key, value)
}

func (ks *Keystore) PutObj(key string, obj interface{}) error {
	encoded, e := json.Marshal(obj)
	ks.PutString(key, string(encoded))
	return e
}

func (ks *Keystore) GetObj(key string, reciever interface{}) error {
	encoded, err := ks.GetString(key)
	if err != nil {
		return err
	}
	b := []byte(encoded)
	err = json.Unmarshal(b, reciever)
	return err
}

func (ks *Keystore) PutInt(key string, value int) {
	ks.db.Exec("REPLACE INTO Settings (key, ival) VALUES (?, ?)", key, value)
}

func (ks *Keystore) GetString(key string) (string, error) {
	res := ks.db.QueryRow("SELECT val FROM Settings WHERE key=?", key)
	var val string;
	err  := res.Scan(&val)
	return val, err
}

func (ks *Keystore) GetInt(key string) (int, error) {
	res := ks.db.QueryRow("SELECT ival FROM Settings WHERE key=?", key)
	var val int
	err := res.Scan(&val)
	return val, err
}

func (ks *Keystore) Close() {
	ks.db.Close()
}

