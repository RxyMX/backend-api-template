package queries

import (
	"common-go-example/internal/controller"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type MutationCreateAccountVariables struct {
	id string `json:"id"`
}

type MutationCreateAccountResponse struct {
	Id        string `json:"id"`
	UserName  string `json"username"`
	Email     string `json"email"`
	CreatedAt string `json:"created_at"`
}

func (m MutationCreateAccountResponse) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Id, validation.Required),
		validation.Field(&m.UserName, validation.Required),
		validation.Field(&m.Email, validation.Required, is.Email),
		validation.Field(&m.CreatedAt, validation.Required))
}

func (m MutationCreateAccountResponse) ToCreateAccountResponse() *controller.CreateAccountRespone {
	return &controller.CreateAccountRespone{
		Id:        m.Id,
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
