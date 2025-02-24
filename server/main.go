package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	// Импортируем сгенерированный пакет (поменяйте путь на актуальный для вашего проекта)
	pb "github.com/girolle/grpc-lesson/proto"
)

// server реализует интерфейсы для обоих сервисов.
// Встраивание Unimplemented...Server гарантирует, что если какой-либо метод не реализован,
// сервер вернёт корректную ошибку.
type server struct {
	pb.UnimplementedPartnerServiceServer
	pb.UnimplementedLinkShorteningServiceServer
}

// CreatePartner реализует метод создания партнёра.
func (s *server) CreatePartner(ctx context.Context, req *pb.CreatePartnerRequest) (*pb.CreatePartnerResponse, error) {
	// Для простоты создаем партнёра с фиксированным id.
	partner := &pb.Partner{
		Id:    1,
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
	log.Printf("Создан партнер: %+v", partner)
	return &pb.CreatePartnerResponse{Partner: partner}, nil
}

// GetPartnerStats реализует метод получения статистики партнёра.
func (s *server) GetPartnerStats(ctx context.Context, req *pb.GetPartnerStatsRequest) (*pb.PartnerStats, error) {
	// Возвращаем фиктивные данные статистики.
	stats := &pb.PartnerStats{
		PartnerId:   req.GetPartnerId(),
		Clicks:      100,
		Conversions: 10,
	}
	log.Printf("Получена статистика партнера (ID %d): %+v", req.GetPartnerId(), stats)
	return stats, nil
}

// ShortenLink реализует метод сокращения ссылки.
func (s *server) ShortenLink(ctx context.Context, req *pb.ShortenLinkRequest) (*pb.ShortenLinkResponse, error) {
	// Фиктивное сокращение: возвращаем статическую короткую ссылку.
	shortURL := "http://short.url/abc123"
	log.Printf("Сокращена ссылка: %s -> %s", req.GetOriginalUrl(), shortURL)
	return &pb.ShortenLinkResponse{ShortUrl: shortURL}, nil
}

// GetLinkStats реализует метод получения статистики по сокращённой ссылке.
func (s *server) GetLinkStats(ctx context.Context, req *pb.GetLinkStatsRequest) (*pb.LinkStats, error) {
	// Возвращаем фиктивные данные статистики для ссылки.
	stats := &pb.LinkStats{
		ShortUrl: req.GetShortUrl(),
		Clicks:   250,
	}
	log.Printf("Получена статистика для ссылки %s: %+v", req.GetShortUrl(), stats)
	return stats, nil
}

func main() {
	// Создаем TCP-слушатель на порту 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось начать прослушивание: %v", err)
	}

	// Создаем новый gRPC-сервер.
	grpcServer := grpc.NewServer()

	// Регистрируем сервисы.
	pb.RegisterPartnerServiceServer(grpcServer, &server{})
	pb.RegisterLinkShorteningServiceServer(grpcServer, &server{})

	log.Println("gRPC сервер запущен на порту :50051")
	// Запускаем сервер. Если возникнет ошибка – логируем её.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}