package memory

import (
	"errors"
	"time"

	"github.com/CafeLucuma/go-play/domain-hex/pkg/adding"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/listing"
)

var (
	plates []Plate
)

type Storage struct {
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	return new(Storage), nil
}

func (s *Storage) AddPlate(p adding.Plate) error {
	newPlate := Plate{
		ID:          len(plates) + 1,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   time.Now(),
		Type:        p.Type,
	}

	plates = append(plates, newPlate)

	return nil
}

func (s *Storage) GetPlates() ([]listing.Plate, error) {

	var listingPlates []listing.Plate
	for _, plate := range plates {
		listingPlates = append(listingPlates, listing.Plate{
			ID:          plate.ID,
			Name:        plate.Name,
			Description: plate.Description,
			CreatedAt:   plate.CreatedAt,
			Type:        plate.Type,
		})
	}
	//error when accesing database goes here
	return listingPlates, nil
}

func (s *Storage) GetPlate(ID int) (listing.Plate, error) {
	for _, plate := range plates {
		if plate.ID == ID {
			return listing.Plate{
				ID:          plate.ID,
				Name:        plate.Name,
				Description: plate.Description,
				CreatedAt:   plate.CreatedAt,
				Type:        plate.Type,
			}, nil
		}
	}

	return listing.Plate{}, errors.New("Couldnt find plate with id " + string(ID))
}
