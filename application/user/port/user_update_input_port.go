package port

type UserUpdateInputData struct {
	User *UserData
}

type UserUpdateOutputData struct {}


type UserUpdateInputPort interface {
	Handle(input *UserUpdateInputData) (*UserUpdateOutputData, error)
}
