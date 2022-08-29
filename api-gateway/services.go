package main
import (
	"context"
	"github.com/Golibbek0414/crmprotos/teacherpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServices(TeacherServiceURL string) Services {
	conn, err:!=nil {
		panic(err)
	}
	return Services{
		TeacherService: teacherServiceAdapter{
			client: teacherpb.NewTeacherServiceClient(conn),
		},
	}
}
type Services struct {
	TeacherService TeacherServiceClient
}
type TeacherServiceClient interface {
	RegisterTeacher(context.Context, RegisterTeacherRequest) (Teacher, error)
	GetTeacher(context.Context, string) (Teacher, error)
	CreateSubject(context.Context, CreateSubjectRequest) (Subject, error)
	GetSubject(context.Context, string) (Subject, error)

}
type TeacherServiceAdapter struct {
	client teacherpb.TeacherServiceClient
}

func (a teacherServiceAdapter) RegisterTeacher(ctx context.Context, request RegisterTeacherRequest) (Teacher, error){

	grpcRequest :=&teacherpb.RegisterTeacherRequest{
		FirstName: request.FirstName,
		LastName: request.LastName,
		Email: request.Email,
		PhoneNumber: request.PhoneNumber,
		SubjectId: request.SubjectID,
	}
	response, err:=a.client.RegisterTeacher(ctx,grpcRequest)
	if err!=nil{
		return Teacher{}, err
	}
	return Teacher{
		ID: response.Id,
		FirstName: response.FirstName,
		LastName: response.LastName,
		Email: response.Email,
		PhoneNumber: response.PhoneNumber,
		SubjectID: response.SubjectId,
	}, nil
}

func (a teacherServiceAdapter) GetTeacher(ctx context.Context, id string) (Teacher, error) {
	response, err := a.client.GetTeacher(ctx, &teacherpb.GetTeacherRequest{TeacherId: id})
	if err != nil {
		return Teacher{}, err
	}

	return Teacher{
		ID:          response.Id,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		Email:       response.Email,
		PhoneNumber: response.PhoneNumber,
		SubjectID:   response.SubjectId,
	}, nil
}

func (a teacherServiceAdapter) CreateSubject(ctx context.Context, request CreateSubjectRequest) (Subject, error) {
	response, err := a.client.CreateSubject(
		ctx,
		&teacherpb.CreateSubjectRequest{
			Name:        request.Name,
			Description: request.Description,
		},
	)
	if err != nil {
		return Subject{}, err
	}

	return Subject{
		ID:          response.Id,
		Name:        response.Name,
		Description: response.Description,
	}, nil
}

func (a teacherServiceAdapter) GetSubject(ctx context.Context, id string) (Subject, error) {
	response, err := a.client.GetSubject(ctx, &teacherpb.GetSubjectRequest{
		SubjectId: id,
	})
	if err != nil {
		return Subject{}, err
	}

	return Subject{
		ID:          response.Id,
		Name:        response.Name,
		Description: response.Description,
	}, nil
}