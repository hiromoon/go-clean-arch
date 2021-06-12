package port

type UserListInputData struct {}

type UserListOutputData struct {
	Users []*UserData
}


type UserListInputPort interface {
	Handle(input *UserListInputData) (*UserListOutputData, error)
}
