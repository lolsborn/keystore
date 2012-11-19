
# Simple Example

	// Get a database struct
	ks, _ := DefaultKeystore()

	// Save a string 
	ks.PutString("mykey", "foo") 

	// Fetch the string "foo" for key "mykeyy"
	bar := ks.GetString("mykey")
