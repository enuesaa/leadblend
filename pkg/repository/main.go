package repository

type Repos struct {
	DB    DBRepositoryInterface
	Fs    FsRepositoryInterface
	Log   LogRepositoryInterface
	Using string
}

func New() Repos {
	return Repos{
		DB:    &DBRepository{},
		Fs:    &FsRepository{},
		Log:   &LogRepository{},
		Using: "",
	}
}

func NewMock() Repos {
	return Repos{
		DB:    &DBRepository{},
		Fs:    &FsMockRepository{},
		Log:   &LogRepository{},
		Using: "",
	}
}
