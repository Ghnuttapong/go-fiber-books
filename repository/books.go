package repository

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const BooksCollection = "books"

type BooksRepository interface{
	Save(book *models.Book) error
	Update(book *models.Book) error
	GetById(id string) (book *models.Book, err error) 
	GetAll() (books []*models.Book, err error) 
	Delete(id string) error
	GetName(name string) (book *models.Book, err error)
}

type booksRepository struct {
	c *mgo.Collection
}

func NewBooksRepository(conn configs.Connection) BooksRepository {
	return &booksRepository{conn.DB().C(BooksCollection)}
}

func (r *booksRepository) Save(book *models.Book) error {
	return r.c.Insert(book)
}

func (r *booksRepository) Update(book *models.Book) error {
	return r.c.UpdateId(book.Id, book)
}
func (r *booksRepository) GetById(id string) (book *models.Book, err error) {
	err = r.c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&book)
	return	book, err
}
func (r *booksRepository) GetAll() (books []*models.Book, err error) {
	err = r.c.Find(bson.M{}).All(&books)
	return	books, err
}
func (r *booksRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
func (r *booksRepository) GetName(name string) (book *models.Book, err error){
	err = r.c.Find(bson.M{"name": name}).One(&book)
	return book, err
}