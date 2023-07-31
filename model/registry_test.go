package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	UserName string `column:"first_name"`
	Id       int64
	Address  []string
}

func TestRegistry_ParseModel(t *testing.T) {
	r := Registry{}
	m, err := r.ParseModel(new(User))
	assert.NoError(t, err)
	m, err = r.Get(new(User))
	assert.NoError(t, err)
	fmt.Println(m)
}
