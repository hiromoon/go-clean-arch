package port

type UserData struct {
	ID       string
	Name     string
	Password string
}

type UserListInputData struct {}

type UserListOutputData struct {
	Users []*UserData
}


type UserListInputPort interface {
	Handle(input *UserListInputData) (*UserListOutputData, error)
}
