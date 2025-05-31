package service

import (
	"stroy-svaya/internal/model"
	"stroy-svaya/internal/repository"
)

type Service struct {
	repo *repository.SQLiteRepository
}

func NewService(r *repository.SQLiteRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) InitPileDrivingRecordLine() *model.PileDrivingRecordLine {
	return &model.PileDrivingRecordLine{
		RecordedBy: "Сатья Надела",
	}
}

func (s *Service) InsertPileDrivingRecordLine(rec *model.PileDrivingRecordLine) error {
	if err := s.repo.InsertPileDrivingRecordLine(rec); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPileDrivingRecord(projectId int) ([]model.PileDrivingRecordLine, error) {

	lines, err := s.repo.GetPileDrivingRecord(projectId)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
