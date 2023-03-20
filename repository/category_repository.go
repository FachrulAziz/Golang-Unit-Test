package repository

import "golang-unit-test/entity"

type CategoryRepository interface {
	FindbyId(id string) *entity.Category
}
