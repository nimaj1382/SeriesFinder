package movie

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Search(title string) ([]*Movie, error) {
	return s.Repo.Search(title)
}

func (s *Service) Add(movie *Movie) error {
	existence, err := s.Repo.Exists(movie)
	if err != nil {
		return err
	}
	if !existence {
		return s.Repo.Add(movie)
	}
	return nil
}
