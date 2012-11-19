package keystore

import (
	"testing"
	"strconv"
)

type TestObj struct {
	Bar string
	Foo string
	Baz int
}

const (
	test_db_path string = "./test_keystore_db.sqlite3"
)

func TestPutString(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	ks.PutString("foo", "bar")
	ks.PutString("foo", "bar2")

	ret, err := ks.GetString("foo")
	if err != nil {
		t.Errorf(err.Error())
	}
	if ret != "bar2" {
		t.Errorf("Expeccted bar2, but got %s", ret)
	}
}

func TestPutInt(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	ks.PutInt("putint", 1)
	ks.PutInt("putint", 55)	

	ret, err := ks.GetInt("putint")
	if err != nil {
		t.Errorf(err.Error())
	}
	if ret != 55 {
		t.Errorf("Expeccted 55, but got %d", ret)
	}
}

func TestDelete(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	ks.PutInt("foey", 1)

	ks.Delete("foey")
	ret, err := ks.GetInt("foey")

	if err == nil { // Expect an error here
		t.Errorf("After delete, expected error, but got %d", ret)
	}
}

func TestNotFoundString(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	_, err := ks.GetString("1ac62cc610fb4495920445ffc6ebea4a")

	if err == nil { // Expect an error here
		t.Errorf("expected error fetching 1ac62cc610fb4495920445ffc6ebea4a")
	}
}

func TestNotFoundInt(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	_, err := ks.GetInt("1ac62cc610fb4495920445ffc6ebea4a")

	if err == nil { // Expect an error here
		t.Errorf("expected error fetching 1ac62cc610fb4495920445ffc6ebea4a")
	}	
}

func TestNotFoundObj(t *testing.T) {
	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	obj := &TestObj{}
	err := ks.GetObj("1ac62cc610fb4495920445ffc6ebea4a", obj)

	if err == nil { // Expect an error here
		t.Errorf("expected error fetching 1ac62cc610fb4495920445ffc6ebea4a")
	}
}

func TestPutObj(t *testing.T) {

	obj := &TestObj{}
	obj.Bar = "asdfadfs"
	obj.Foo = "Hello, 世界"
	obj.Baz = 42

	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()
	ks.PutObj("test.obj", obj)


	obj2 := &TestObj{}
	ks.GetObj("test.obj", obj2)
	if obj2.Baz != obj.Baz {
		t.Errorf("Expeccted %d, but got %d", obj.Baz, obj2.Baz)
	}

	if obj2.Foo != obj.Foo {
		t.Errorf("Expeccted %s, but got %s", obj.Foo, obj2.Foo)
	}
}

// Benchmarks 10,000 replace operations
func Benchmark10KUpdateObj(b *testing.B) {

	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()

	for i := 0; i < 10000; i++ {
		obj := &TestObj{}
		obj.Bar = "asdfadfs"
		obj.Foo = "Hello, 世界"
		obj.Baz = 49
		ks.PutObj("benchmark.obj", obj)
	}

}

// Benchmarks 10,000 insert operations
func Benchmark10kRawInsertObj(b *testing.B) {

	ks, _ := NewKeystore(test_db_path)
	defer ks.Close()

	for i := 0; i < 10000; i++ {
		obj := &TestObj{}
		obj.Bar = "asdfadfs"
		obj.Foo = "Hello, 世界"
		obj.Baz = 55
		ks.PutObj("benchmark_insert" + strconv.Itoa(55) , obj)
	}

}
