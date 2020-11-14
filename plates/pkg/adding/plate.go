package adding

import "errors"

type Plate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (p Plate) Validate() error {

	if p.Name == "" {
		return errors.New("'name' field empty")
	}

	if p.Description == "" {
		return errors.New("'description' field empty")
	}

	if p.Type == "" {
		return errors.New("'type' field empty")
	}

	if p.Type != "Pre-olimpica" && p.Type != "Olimpica" {
		return errors.New("Invalid plate type. Valid types: [Pre-olimpica, Olimpica]")
	}

	return nil
}
