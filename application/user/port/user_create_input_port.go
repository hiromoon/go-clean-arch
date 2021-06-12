package port

type UserCreateInputData struct {
	User *UserData
}

type UserCreateOutputData struct {}

type UserCreateInputPort interface {
	Handle(input *UserCreateInputData) (*UserCreateOutputData, error)
}
