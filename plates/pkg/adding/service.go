package adding

// Service provides beer adding operations.
type Service interface {
	AddPlate(Plate) error
}

// Repository provides access to beer repository.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddPlate(Plate) error
}

type service struct {
	pR Repository
}

func NewService(r Repository) Service {
	return &service{pR: r}
}

func (s *service) AddPlate(p Plate) error {
	if err := s.pR.AddPlate(p); err != nil {
		return err
	}

	return nil
}
