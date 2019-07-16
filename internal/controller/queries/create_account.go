package queries

import validation "github.com/go-ozzo/ozzo-validation"

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
		validation.Field(&m.Id, validation.Required))
	// TODO: Rest
}

const MutationCreateAccountName = "insert_accounts"
const MutationCreateAccountQuery = `
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
