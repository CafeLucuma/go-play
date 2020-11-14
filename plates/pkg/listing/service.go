package listing

type Service interface {
	GetPlate(int) (Plate, error)
	GetPlates() ([]Plate, error)
}

type Repository interface {
	GetPlate(int) (Plate, error)
	GetPlates() ([]Plate, error)
}

type service struct {
	pR Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetPlates() ([]Plate, error) {
	return s.pR.GetPlates()
}
func (s *service) GetPlate(ID int) (Plate, error) {
	return s.pR.GetPlate(ID)
}
