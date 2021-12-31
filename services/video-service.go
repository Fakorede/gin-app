package services

import (
	"github.com/Fakorede/gin-app/entity"
	"github.com/Fakorede/gin-app/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video)
	Delete(entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func NewVideoService(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.videoRepository.Save(video)
	return video
}

func (s *videoService) Update(video entity.Video) {
	s.videoRepository.Update(video)
}

func (s *videoService) Delete(video entity.Video) {
	s.videoRepository.Delete(video)
}

func (s *videoService) FindAll() []entity.Video {
	return s.videoRepository.FindAll()
}
