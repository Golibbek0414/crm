package main

import (
	"log"
	"net"
	"schedule-service/config"
	"schedule-service/domain/schedule"
	"schedule-service/repository"
	"schedule-service/server"
	"schedule-service/service"

	"github.com/Golibbek0414/crmcommon/id"
	"github.com/Golibbek0414/crmprotos/schedulepb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("error with loading config: ", err)
	}

	repo, err := repository.NewPostgresRepository(cfg.Config)
	if err != nil {
		panic(err)
	}

	scheduleFactory := schedule.NewFactory(id.Generator{})

	svc := service.New(repo)
	svr := server.New(svc, scheduleFactory)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	schedulepb.RegisterScheduleServiceServer(grpcServer, svr)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}