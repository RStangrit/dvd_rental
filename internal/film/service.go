package film

import (
	"errors"
	"time"
)

func (f *Film) Validate() error {
	if f.Title == "" || len(f.Title) > 255 {
		return errors.New("title is required and must be less than 255 characters")
	}
	if f.ReleaseYear < 1900 || f.ReleaseYear > time.Now().Year()+5 {
		return errors.New("release year is invalid")
	}
	if f.LanguageID <= 0 {
		return errors.New("language ID must be a positive number")
	}
	if f.RentalDuration <= 0 {
		return errors.New("rental duration must be a positive number")
	}
	if f.RentalRate <= 0 || f.RentalRate > 99.99 {
		return errors.New("rental rate is invalid")
	}
	if f.Length <= 0 {
		return errors.New("length must be a positive number")
	}
	if f.ReplacementCost <= 0 || f.ReplacementCost > 999.99 {
		return errors.New("replacement cost is invalid")
	}
	if !isValidRating(f.Rating) {
		return errors.New("invalid rating")
	}
	if len(f.SpecialFeatures) > 0 {
		for _, feature := range f.SpecialFeatures {
			if !isValidFeature(feature) {
				return errors.New("invalid special feature")
			}
		}
	}
	return nil
}

func isValidRating(rating string) bool {
	validRatings := map[string]bool{"G": true, "PG": true, "PG-13": true, "R": true, "NC-17": true}
	_, exists := validRatings[rating]
	return exists
}

func isValidFeature(feature string) bool {
	validFeatures := map[string]bool{"Trailers": true, "Commentaries": true, "Deleted Scenes": true, "Behind the Scenes": true}
	_, exists := validFeatures[feature]
	return exists
}
