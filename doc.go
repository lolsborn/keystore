/*
# Overview

Keystore is a dead simple key-value store built using SQLite for Go.  It is
designed only with simplicity in mind.

## Installation

	go get github.com/steve918/keystore


## Simple Example

	import "github.com/steve918/keystore"
	
	// Get a keystore struct
	ks, _ := DefaultKeystore()

	// Save a string 
	ks.PutString("mykey", "foo") 

	// Fetch the string "foo" for key "mykeyy"
	bar := ks.GetString("mykey")
*/
package keystore