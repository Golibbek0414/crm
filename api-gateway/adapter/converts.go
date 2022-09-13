package adapter

import (
	"api-gateway/response"
	"github.com/Golibbek0414/crmprotos/journalpb"
	"time"
	"github.com/Golibbek0414/crmprotos/schedulepb"
	"github.com/Golibbek0414/crmprotos/studentpb"
	"github.com/GOlibbek0414/crmprotos/teacherpb"
)

func fromProtoToResponseEntries(protoEntries *journalpb.Entries) []response.JournalEntry {
	entries := make([]response.JournalEntry, 0, len(protoEntries.ENtries))
	for _, e:=range protoEntries.Entries {
		entries = append(entries, response.JournalEntry{
			JournalID: e.JournalID,
			Date: e.Date.AsTime(),
			ScheduleID: e.ScheduleId,
			Mark: e.Mark,
			Attended: e.Attended,
		})
	}
	return entries
}

func fromProtoToTeacherList(teachers *teacherpb.ListTeachersResponse) ([]response.Teacher, error) {
	tchs :=make([]response.Teacher,0, len(teachers.Teachers))
	for _, item :=range teachers.Teachers {
		tchs = append(tchs,fromProtoToResponseTeacher(item))

	}
	return tchs,nil
}

func fromProtoToResponseTeacher(teacher *teacherpb.Teacher) response.Teacher{
	return response.Teacher{
		ID: teacher.Id,
		FirstName: teacher.FirstName,
		LastName: teacher.LastName,
		Email: teacher.Email,
		PhoneNumber: teacher.PhoneNumber,
		SubjectID: teacher.SubjectId,
	}
}

func fromProtoToSubjectList(subjects *teacherpb.ListSubjectsResponse) ([]response.Subject,error){
	sbjs := make([]response.Subject, 0, len(subjects.Subjects))
	for _, item := range subjects.Subjects {
		sbjs = append(sbjs, fromProtoToResponseSubject(item))
	}
	return sbjs,nil
}

func fromProtoToResponseSubject(subject *teacherpb.Subject) response.Subject {
	return response.Subject{
		ID:          subject.Id,
		Name:        subject.Name,
		Description: subject.Description,
	}
}

func fromProtoToResponseGroups(groups *studentpb.GroupList) ([]response.Group, error) {
	grs := make([]response.Group, 0, len(groups.Groups))
	for _, item := range groups.GetGroups() {
		grs = append(grs, fromProtoToResponseGroup(item))
	}
	return grs, nil
}

func fromProtoToResponseGroup(group *studentpb.Group) response.Group {
	return response.Group{
		ID:            group.Id,
		Name:          group.Name,
		MainTeacherID: group.MainTeacherId,
	}
}

func fromProtoToResponseSubject(subject *teacherpb.Subject) response.Subject {
	return response.Subject{
		ID:          subject.Id,
		Name:        subject.Name,
		Description: subject.Description,
	}
}

func fromProtoToResponseGroups(groups *studentpb.GroupList) ([]response.Group, error) {
	grs := make([]response.Group, 0, len(groups.Groups))
	for _, item := range groups.GetGroups() {
		grs = append(grs, fromProtoToResponseGroup(item))
	}
	return grs, nil
}

func fromProtoToResponseGroup(group *studentpb.Group) response.Group {
	return response.Group{
		ID:            group.Id,
		Name:          group.Name,
		MainTeacherID: group.MainTeacherId,
	}
}

func fromProtoScheduleListToResponseScheduleSlice(list *schedulepb.ScheduleList) []response.Schedule {
	schedules := make([]response.Schedule, 0, len(list.Schedules))
	for _, item := range list.Schedules {
		schedules = append(schedules, response.Schedule{
			ID:           item.Id,
			GroupId:      item.GroupId,
			SubjectID:    item.SubjectId,
			TeacherID:    item.TeacherId,
			WeekDay:      time.Weekday(item.Weekday),
			LessonNumber: item.LessonNumber,
		})
	}

	return schedules
}