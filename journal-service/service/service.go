package service

import (
	"context"
	"journal-service/domain/journal"
	"log"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo           JournalRepository
	journalFactory journal.Factory
}

func New(jouRepo JournalRepository, journalFactory journal.Factory) Service {
	return Service{
		repo:           jouRepo,
		journalFactory: journalFactory,
	}
}

func (s Service) RegisterJournal(ctx context.Context, journal journal.Journal, studentIDs []uuid.UUID) error {
	if err := s.repo.CreateJournal(ctx, journal); err != nil {
		return err
	}
	if err := s.repo.CreateJournalStats(ctx, journal.ID(), studentIDs); err != nil {
		if err = s.repo.DeleteJournal(ctx, journal.ID()); err != nil {
			log.Println("[ERROR] could not delete journal after failed insertion of journal statuses")
		}
		return err
	}

	return nil
}

func (s Service) GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error) {
	jour, err := s.repo.GetJournal(ctx, id)
	if err != nil {
		return journal.Journal{}, nil
	}
	return jour, nil
}

func (s Service) UpdateJournal(ctx context.Context, jour journal.Journal) error {
	return s.repo.UpdateJournal(ctx, jour)
}

func (s Service) DeleteJournal(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteJournal(ctx, id)
}

func (s Service) MarkStudent(ctx context.Context, st journal.Stat) error {
	return s.repo.MarkStudent(ctx, st)
}

func (s Service) SetStudentAttendance(ctx context.Context, st journal.Stat) error {
	return s.repo.SetStudentAttendance(ctx, st)
}

func (s Service) GetStudentJournalEntries(ctx context.Context, studentID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	return s.repo.GetStudentJournalEntries(ctx, studentID, start, end)
}

func (s Service) GetTeacherJournalEntries(ctx context.Context, teacherID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	return s.repo.GetTeacherJournalEntries(ctx, teacherID, start, end)
}