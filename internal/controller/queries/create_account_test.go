package queries

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMutationCreateAccountQuery(t *testing.T) {
	queries, err := ToGraphqlQuery([]byte(MutationCreateAccount))

	if err != nil {
		t.Error(err)
	}

	assert.Len(t, queries, 1)
	assert.Equal(t, MutationCreateAccountName, queries[0].Type)
	r := reflect.ValueOf(MutationCreateAccountVariables{})
	for varName, _ := range queries[0].Variables {
		assert.True(t, r.FieldByName(varName).IsValid(), "Cannot find required variable for query!")

	}
}
