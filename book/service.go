package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {

	books, err := s.repository.FindAll()

	return books, err

	//return s.repository.FindAll() cara simple
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)

	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// karena parameter nya adalah book request maka harus di pindah dari struct book -> book request
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}