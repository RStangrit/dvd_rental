package models

import (
	"reflect"
	"sort"
)

var ModelRegistry []any

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
	"user.User":                  16,
}

func RegisterModel(model any) {
	ModelRegistry = append(ModelRegistry, model)
}

func ReorderModels() {
	sort.SliceStable(ModelRegistry, func(i, j int) bool {
		return getOrder(ModelRegistry[i]) < getOrder(ModelRegistry[j])
	})
}

func getOrder(model any) int {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if order, ok := modelOrder[modelType.String()]; ok {
		return order
	}
	return 999
}
