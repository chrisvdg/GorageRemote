package db

import (
	"io/ioutil"
	"testing"

	e "github.com/chrisvdg/GorageRemote/entities"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	assert := assert.New(t)

	dbFile, err := ioutil.TempFile("", "testdb")
	assert.NoError(err)

	db, err := NewDB(dbFile.Name())
	assert.NoError(err)
	assert.NotNil(db)

	// fetch admin and check admin bool
	u, err := GetUser(db, "admin")
	assert.NoError(err)
	assert.Equal("admin", u.Name)
	assert.True(u.Admin)

	// change password and admin rights
	u = &e.User{
		Name:     "admin",
		Password: "Garage123",
		Admin:    false,
	}
	err = UpdatePassword(db, u)
	assert.NoError(err)
	err = UpdateAdmin(db, u)
	assert.NoError(err)

	// close first db conn
	db.Close()

	// use same db again (should not set default admin account)
	newDB, err := NewDB(dbFile.Name())
	assert.NoError(err)
	assert.NotNil(newDB)

	// check if password and admin rights are persisted
	err = CheckPassword(newDB, &e.User{
		Name:     "admin",
		Password: "Garage123",
	})
	assert.NoError(err)
	err = CheckPassword(newDB, &e.User{
		Name:     "admin",
		Password: "Gorage123",
	})
	assert.EqualError(err, ErrFailedAuth.Error())
	u, err = GetUser(newDB, "admin")
	assert.NoError(err)
	assert.False(u.Admin)
}
