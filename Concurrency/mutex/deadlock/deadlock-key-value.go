package main

import (
	"fmt"
	"sync"
)

//DataStore là một key-value in memory caching
type DataStore struct {
	sync.Mutex // ← Mutex để đồng bộ truy xuất vào biến cache
	cache      map[string]string
}

//Khởi tạo DataStore
func New() *DataStore {
	return &DataStore{
		cache: make(map[string]string),
	}
}

//Gán một cặp key-value vào DataStore
func (ds *DataStore) set(key string, value string) {
	ds.Lock()
	defer ds.Unlock()
	ds.cache[key] = value
}

//Lấy giá trị từ caching
func (ds *DataStore) get(key string) string {
	ds.Lock() //Lock Mutex lần thứ nhất
	defer ds.Unlock()
	if ds.count() > 0 { //Hàm ds.count() tiếp tục lock Mutex lần nữa
		item := ds.cache[key]
		return item
	}
	return ""
}
func (ds *DataStore) count() int {
	ds.Lock()
	defer ds.Unlock()
	return len(ds.cache)
}
func main() {
	/*
		Running this below will deadlock because the get() method will take a lock and will call the count() method which will also take a lock before the set() method unlocks()
	*/
	store := New()
	store.set("Go", "Lang")
	result := store.get("Go")
	fmt.Println(result)
}
