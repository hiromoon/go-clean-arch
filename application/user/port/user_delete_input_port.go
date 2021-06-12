package port

type UserDeleteInputData struct {
	UserID string
}

type UserDeleteOutputData struct {}

type UserDeleteInputPort interface {
	Handle(input *UserDeleteInputData) (*UserDeleteOutputData, error)
}
