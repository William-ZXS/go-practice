package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"encoding/json"
	"fmt"
	"sync"
	"github.com/garyburd/redigo/redis"
)





func init() {

	passWd := redis.DialPassword("william")
	db := redis.DialDatabase(1)
	conn, err := redis.Dial("tcp", "127.0.0.1:6379",passWd,db)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	//defer c.Close()
	memStore := &MemStore{
		conn:conn,
	}

	factory.Register("mem", memStore)


}


type MemStore struct {
	sync.RWMutex
	conn redis.Conn
}

// Create creates a new Book in the store.
func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()


	bookDataRedis, err := redis.Bytes(ms.conn.Do("GET", book.Id))
	if bookDataRedis != nil && len(bookDataRedis) >0 {
		return mystore.ErrExist
	}
	data, _ := json.Marshal(book)
	_, err = ms.conn.Do("SET", book.Id, data)
	if err !=nil {
		return err
	}

	return nil
}

// Update updates the existed Book in the store.
func (ms *MemStore) Update(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	bookDataRedis, err := redis.Bytes(ms.conn.Do("GET", book.Id))
	if err != nil {
		return err
	}
	if bookDataRedis == nil && len(bookDataRedis) <=0 {
		return mystore.ErrNotFound
	}

	oldBook := mystore.Book{}
	err = json.Unmarshal(bookDataRedis, &oldBook)
	if err != nil {
		fmt.Println(err)
	}
	nBook := oldBook

	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Authors != nil {
		nBook.Authors = book.Authors
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	data, _ := json.Marshal(nBook)
	_, err = ms.conn.Do("SET", book.Id, data)
	if err !=nil {
		return err
	}

	return nil
}

// Get retrieves a book from the store, by id. If no such id exists. an
// error is returned.
func (ms *MemStore) Get(id string) (mystore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	bookDataRedis, err := redis.Bytes(ms.conn.Do("GET", id))
	if err != nil {
		return mystore.Book{},err
	}
	if bookDataRedis == nil && len(bookDataRedis) <=0 {
		return mystore.Book{},mystore.ErrNotFound
	}

	book := mystore.Book{}
	err = json.Unmarshal(bookDataRedis, &book)
	if err != nil {
		fmt.Println(err)
	}
	return book,nil

}

// Delete deletes the book with the given id. If no such id exist. an error
// is returned.
func (ms *MemStore) Delete(id string) error {
	ms.Lock()
	defer ms.Unlock()


	bookDataRedis, err := redis.Bytes(ms.conn.Do("GET", id))
	if err != nil {
		return err
	}
	if bookDataRedis == nil && len(bookDataRedis) <=0 {
		return mystore.ErrNotFound
	}

	_, err = ms.conn.Do("DEL", id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll returns all the books in the store, in arbitrary order.
func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	//ms.RLock()
	//defer ms.RUnlock()
	//
	//allBooks := make([]mystore.Book, 0, len(ms.books))
	//for _, book := range ms.books {
	//	allBooks = append(allBooks, *book)
	//}
	return make([]mystore.Book,0), nil
}
