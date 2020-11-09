package postgres

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/CafeLucuma/go-play/domain-hex/pkg/adding"
	"github.com/CafeLucuma/go-play/domain-hex/pkg/listing"
	_ "github.com/lib/pq"
)

var (
	plates []Plate
)

type Storage struct {
	db *sql.DB
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	storage := new(Storage)

	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Println("Cant load databse url from environment")
		return nil, errors.New("Cant load databse url from environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	storage.db = db

	if err := storage.db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return storage, nil
}

func (s *Storage) CloseDB() {
	log.Println("Closing db...")
	log.Fatal(s.db.Close())
}

func (s *Storage) AddPlate(p adding.Plate) error {

	newPlate := Plate{
		ID:          len(plates) + 1,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   time.Now(),
		Type:        p.Type,
	}

	sqlStatement := "INSERT INTO PLATES (name, description, plate_type, created_on) VALUES ($1, $2, $3, $4)"
	_, err := s.db.Exec(sqlStatement, newPlate.Name, newPlate.Description, newPlate.Type, newPlate.CreatedAt)
	if err != nil {
		log.Println("Cant insert new plate to database")
		return err
	}

	return nil
}

func (s *Storage) GetPlates() ([]listing.Plate, error) {

	sqlStatement := "SELECT * FROM plates"

	plateRows, err := s.db.Query(sqlStatement)
	if err != nil {
		log.Println("Cant get plates from db")
		return nil, err
	}

	var listingPlates []listing.Plate

	for plateRows.Next() {
		var plate listing.Plate

		if err := plateRows.Scan(&plate.ID, &plate.Name, &plate.Description, &plate.Type, &plate.CreatedAt); err != nil {
			log.Println("Cant parse row to plate")
			return nil, err
		}

		listingPlates = append(listingPlates, plate)
	}

	//error when accesing database goes here
	return listingPlates, nil
}

func (s *Storage) GetPlate(ID int) (listing.Plate, error) {

	sqlStatement := "SELECT * FROM plates WHERE plate_id = $1"
	plateRow := s.db.QueryRow(sqlStatement, ID)

	var plate listing.Plate
	if err := plateRow.Scan(&plate.ID, &plate.Name, &plate.Description, &plate.Type, &plate.CreatedAt); err != nil {
		log.Println("Cant parse row to plate")
		return listing.Plate{}, err
	}

	return plate, nil
}
