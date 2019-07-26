package queries

import (
	"common-go-example/internal/controller"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type MutationCreateAccountVariables struct {
	ID string `json:"id"`
}

type MutationCreateAccountResponse struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (m MutationCreateAccountResponse) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
		validation.Field(&m.UserName, validation.Required),
		validation.Field(&m.Email, validation.Required, is.Email),
		validation.Field(&m.CreatedAt, validation.Required))
}

func (m MutationCreateAccountResponse) ToCreateAccountResponse() *controller.CreateAccountRespone {
	return &controller.CreateAccountRespone{
		ID:        m.ID,
		UserName:  m.UserName,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
}

const MutationCreateAccountName = "insert_accounts"
const MutationCreateAccount = `
mutation($id: uuid){
    insert_accounts(
        where: { id: { _eq: $id } }
    ){
        id
        username
        email
        created_at
        updated_at
    }
}
`
