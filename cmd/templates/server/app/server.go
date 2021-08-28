package app

import (
	"ago_grpc_1/pkg/templates"
	templatesV1Pb "ago_grpc_1/pkg/templates/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Server struct {
	templatesV1Pb.UnimplementedServiceServer
	templatesSvc *templates.Service
}

func NewServer(templatesSvc *templates.Service) *Server {
	return &Server{templatesSvc: templatesSvc}
}

func (s *Server) Create(ctx context.Context, request *templatesV1Pb.CreateRequest) (*templatesV1Pb.Response, error) {
	template, err := s.templatesSvc.Create(ctx, request.Title, request.Phone)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return serviceTemplateToGRPCResponse(template), nil
}

func (s *Server) GetList(ctx context.Context, empty *emptypb.Empty) (*templatesV1Pb.AllResponse, error) {
	allTemplates, err := s.templatesSvc.GetList(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	items := make([]*templatesV1Pb.Response, len(allTemplates))
	for i, template := range allTemplates {
		items[i] = serviceTemplateToGRPCResponse(template)
	}

	return &templatesV1Pb.AllResponse{Items: items}, nil
}

func (s *Server) GetById(ctx context.Context, request *templatesV1Pb.GetByIdRequest) (*templatesV1Pb.Response, error) {
	template, err := s.templatesSvc.GetById(ctx, request.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return serviceTemplateToGRPCResponse(template), nil
}

func (s *Server) UpdateById(ctx context.Context, request *templatesV1Pb.UpdateRequest) (*templatesV1Pb.Response, error) {
	template, err := s.templatesSvc.Update(ctx, request.Id, request.Title, request.Phone)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return serviceTemplateToGRPCResponse(template), nil
}

func (s *Server) DeleteById(ctx context.Context, request *templatesV1Pb.GetByIdRequest) (*templatesV1Pb.Response, error) {
	template, err := s.templatesSvc.Delete(ctx, request.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return serviceTemplateToGRPCResponse(template), nil
}

func serviceTemplateToGRPCResponse(template *templates.Template) *templatesV1Pb.Response {
	return &templatesV1Pb.Response{
		Id:    template.Id,
		Title: template.Title,
		Phone: template.Phone,
		Created: &timestamppb.Timestamp{
			Seconds: template.Created,
			Nanos:   0,
		},
		Updated: &timestamppb.Timestamp{
			Seconds: template.Updated,
			Nanos:   0,
		},
	}
}
