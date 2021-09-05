package services

import (
	"restfulalta/part-3-code-structuring/config"
	"restfulalta/part-3-code-structuring/models"
)

func GetBooks() ([]models.BookAPI, error) {
	var books []models.BookAPI

	res := config.Db.Model(&models.Book{}).Find(&books)

	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func GetBookById(targetId int) (models.BookAPI, int, error) {
	var book models.BookAPI

	res := config.Db.Model(&models.Book{}).Find(&book, targetId)

	if res.Error != nil {
		return models.BookAPI{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.BookAPI{}, 0, nil
	}

	return book, 1, nil
}

func AddBook(newBook *models.Book) (models.Book, error) {
	res := config.Db.Create(newBook)
	if res.Error != nil {
		return models.Book{}, res.Error
	}
	return *newBook, nil
}

func EditBook(newData models.Book, targetId int) (models.Book, int, error) {
	targetBook := models.Book{}
	res := config.Db.Find(&targetBook, targetId)

	if res.Error != nil {
		return models.Book{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Book{}, 0, nil
	}

	res = config.Db.Model(&targetBook).Updates(newData)

	if res.Error != nil {
		return models.Book{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Book{}, 0, nil
	}

	return targetBook, 1, nil
}

func DeleteBook(targetId int) (models.Book, int, error) {	
	targetBook := models.Book{}
	res := config.Db.Find(&targetBook, targetId)

	if res.Error != nil {
		return models.Book{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Book{}, 0, nil
	}
	
	deleted := targetBook

	res = config.Db.Unscoped().Delete(&targetBook)

	if res.Error != nil {
		return models.Book{}, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Book{}, 0, nil
	}

	return deleted, 1, nil

}
