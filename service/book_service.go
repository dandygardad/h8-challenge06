package service

import (
	"challenge06/entity"
	"challenge06/repository"
	"errors"
)

func CreateBook(newBook entity.Book) error {
	err := repository.Create(newBook)
	if err != nil {
		return errors.New("Error menambahkan buku")
	}
	return nil
}

func GetAllBooks() error {
	books, err := repository.GetAll()
	if err != nil {
		return errors.New("Terjadi kesalahan ambil buku")
	}

	// Cek jika tidak ada data buku dalam database
	if len(books) <= 0 {
		return errors.New("Tidak ada data buku")
	}

	return nil
}

func GetBook(id int) (entity.Book, error) {
	book, err := repository.GetOne(id)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func UpdateBook(id int, newBook entity.Book) error {
	oldBook, errGet := repository.GetOne(id)
	if errGet != nil {
		return errGet
	}

	newBook.ID = oldBook.ID // Isi id lama ke data buku baru
	errUpdate := repository.UpdateOne(id, newBook)
	if errUpdate != nil {
		return errors.New("Terjadi kesalahan update buku")
	}

	return nil
}

func DeleteBook(id int) error {
	book, errGet := repository.GetOne(id)
	if errGet != nil {
		return errGet
	}

	err := repository.DeleteOne(book.ID)
	if err != nil {
		return errors.New("Terjadi kesalahan hapus buku")
	}

	return nil
}
