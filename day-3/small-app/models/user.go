package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type CacheStore map[string]User

type Conn struct {
	//db *sql.DB
	cache CacheStore
	mu    *sync.Mutex
}

func NewConn() Conn {
	return Conn{cache: make(map[string]User, 100), mu: &sync.Mutex{}}
}

func (c *Conn) CreateUser(n NewUser) (User, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	us := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(passHash),
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// put the new user in the map
	c.cache = CacheStore{n.Email: us}
	return us, nil
}
