package database

import (
	"testing"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	// Criando entidade user
	user, _ := entity.NewUser("Felipe", "felipe.com", "123456")

	// Criando uma instancia do "model" user
	userDB := NewUser(db)

	// Criando o user no banco
	err = userDB.Create(user)
	assert.Nil(t, err)

	// Buscando o usuário criado
	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	// Criando entidade user
	user, _ := entity.NewUser("Felipe", "felipe.com", "123456")

	// Criando uma instancia do "model" user
	userDB := NewUser(db)

	// Criando o user no banco
	err = userDB.Create(user)
	assert.Nil(t, err)

	// Buscando o usuário criado a partir do email
	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}
