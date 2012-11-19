/* (http://golang.org/doc/gopher/help.png)
# Overview

Keystore is a dead simple key-value store built using SQLite for Go.  It is
designed only with simplicity in mind.

## Installation

	go get github.com/steve918/keystore


## Simple Example

	import "github.com/steve918/keystore"
	
	// Get a keystore struct
	ks, _ := DefaultKeystore()
	// Close the keystore when we are finished
	defer ks.Close()

	// Save a string 
	ks.PutString("mykey", "foo") 

	// Fetch the string "foo" for key "mykey"
	bar := ks.GetString("mykey")

## Usage
### string
	ks.PutString("foo", "bar")
	var ret string = ks.GetString("foo")

### int
	ks.PutInt("meaning", 42)
	var meaning int = ks.GetInt("meaning")

### interface{}
	obj := &TestObj{}
	obj.Bar = "asdfadfs"
	obj.Foo = "Hello, 世界"
	obj.Baz = 42

	ks.PutObj("test.obj", obj)

	out_obj := &TestObj{}
	ks.GetObj("test.obj", out_obj)


## Requirements

depends on http://github.com/mattn/go-sqlite3

*/
package keystore
