package migration

import (
	"github.com/alvinfebriando/gin-gorm-skeleton/entity"
	"github.com/alvinfebriando/gin-gorm-skeleton/hash"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	users := []*entity.User{
		{Name: "Alice", Email: "alice@example.com", Password: hashPassword("alice123")},
		{Name: "Bob", Email: "bob@example.com", Password: hashPassword("bob123")},
		{Name: "Charlie", Email: "charlie@example.com", Password: hashPassword("charlie123")},
	}

	db.Create(users)
}

func hashPassword(text string) string {
	hasher := hash.NewHasher()
	hashedText, err := hasher.Hash(text)
	if err != nil {
		return ""
	}
	return string(hashedText)
}
