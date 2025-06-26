package store

type Service interface {
	GetAll() ([]Store, error)
	GetByID(id int) (*Store, error)
	Create(s Store) error
	Update(s Store) error
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Store, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Store, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(store Store) error {
	return s.repo.Create(store)
}

func (s *service) Update(store Store) error {
	return s.repo.Update(store)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
