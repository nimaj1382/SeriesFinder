package archive

import "TheCollector/app/internal/movie"

type Service struct {
	repo         *Repository
	movieService *movie.Service
}

func NewService(repo *Repository, movieService *movie.Service) *Service {
	return &Service{repo: repo, movieService: movieService}
}

func (s *Service) Create(archive *Archive) error {
	existence, err := s.repo.Exists(archive)
	if err != nil {
		return err
	}

	if !existence {
		err = s.repo.Create(archive)
		if err != nil {
			return err
		}
	}

	return s.Update(archive)
}

func (s *Service) Update(archive *Archive) error {
	return Crawl(archive.URL, archive.Depth, s.movieService)
}

func (s *Service) Remove(archive *Archive) error {
	return s.repo.Remove(archive)
}

func (s *Service) GetAll() ([]*Archive, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateAll() error {
	archives, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	for _, archive := range archives {
		err = s.Update(archive)
		if err != nil {
			return err
		}
	}

	return nil
}
