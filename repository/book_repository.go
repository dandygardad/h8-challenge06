package repository

import (
	"challenge06/entity"
	"errors"
)

var DatabaseBook []entity.Book

func Create(newBook entity.Book) error {
	newBook.ID = len(DatabaseBook) + 1
	DatabaseBook = append(DatabaseBook, newBook)

	return nil
}

func GetAll() ([]entity.Book, error) {
	return DatabaseBook, nil
}

func GetOne(id int) (entity.Book, error) {
	var book entity.Book
	var isFound = false
	for _, v := range DatabaseBook {
		if v.ID == id {
			isFound = true
			book = v
			break
		}
	}

	// Jika ditemukan kembalikan error
	if !isFound {
		return book, errors.New("Buku tidak ditemukan")
	}

	return book, nil
}

func UpdateOne(id int, newBook entity.Book) error {
	DatabaseBook[id-1] = newBook
	return nil
}

func DeleteOne(id int) error {
	id = id - 1
	copy(DatabaseBook[id:], DatabaseBook[id+1:])      // Pindahkan data setelah id ke id yang mau dihapus
	DatabaseBook = DatabaseBook[:len(DatabaseBook)-1] // Timpa database yang lama

	return nil
}
