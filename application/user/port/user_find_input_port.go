package port

type UserFindInputData struct {
	UserID string
}

type UserFindOutputData struct {
	User *UserData
}

type UserFindInputPort interface {
	Handle(input *UserFindInputData) (*UserFindOutputData, error)
}
