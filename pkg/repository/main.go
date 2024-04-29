package repository

type Repos struct {
	DB DBRepositoryInterface
	Fs FsRepositoryInterface
	Log LogRepositoryInterface
}

func New() Repos {
	return Repos{
		DB: &DBRepository{},
		Fs: &FsRepository{},
		Log: &LogRepository{},
	}
}

func NewMock() Repos {
	return Repos{
		DB: &DBRepository{},
		Fs: &FsMockRepository{},
		Log: &LogRepository{},
	}
}
