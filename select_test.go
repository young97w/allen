package allen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
}

func TestSelect(t *testing.T) {
	testCases := []struct {
		name      string
		selector  *Selector[User]
		wantQuery *Query
		wantErr   error
	}{
		{
			name:      "user_struct",
			selector:  &Selector[User]{},
			wantQuery: &Query{SQL: "SELECT * FROM User;"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.selector.Builder()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}
