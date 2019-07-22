package queries

import (
	test "github.com/kintohub/common-go/utils/testutils"
	"reflect"
	"testing"
)

func TestMutationCreateAccountQuery(t *testing.T) {
	test.AssertValidHasuraQuery(t,
		MutationCreateAccountName,
		MutationCreateAccount,
		reflect.ValueOf(MutationCreateAccountVariables{}))
}
