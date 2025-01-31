package db

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

var modelRegistry []interface{}

var modelOrder = map[string]int{
	"language.Language":          1,
	"actor.Actor":                2,
	"film.Film":                  3,
	"category.Category":          4,
	"film_actor.FilmActor":       5,
	"inventory.Inventory":        6,
	"film_category.FilmCategory": 7,
	"country.Country":            8,
	"city.City":                  9,
	"address.Address":            10,
	"customer.Customer":          11,
	"staff.Staff":                12,
	"store.Store":                13,
	"rental.Rental":              14,
	"payment.Payment":            15,
}

func reorderModels() {
	sort.SliceStable(modelRegistry, func(i, j int) bool {
		return getOrder(modelRegistry[i]) < getOrder(modelRegistry[j])
	})
}

func getOrder(model interface{}) int {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if order, ok := modelOrder[modelType.String()]; ok {
		return order
	}
	return 999
}

func RegisterModel(model interface{}) {
	modelRegistry = append(modelRegistry, model)
}

func createTables() error {
	reorderModels()

	if len(modelRegistry) == 0 {
		return fmt.Errorf("no models registered for migration")
	}

	for _, model := range modelRegistry {
		log.Printf("Starting migration for table: %T\n", model)
	}

	return GORM.AutoMigrate(modelRegistry...)
}
