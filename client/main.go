package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	// Импорт сгенерированного пакета (поменяйте путь на актуальный)
	pb "github.com/girolle/grpc-lesson/proto"
)

func main() {
	// Устанавливаем соединение с gRPC-сервером.
	// Для упрощения используем insecure соединение.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	// Создаем клиентов для каждого сервиса.
	partnerClient := pb.NewPartnerServiceClient(conn)
	linkClient := pb.NewLinkShorteningServiceClient(conn)

	// Определяем таймаут для контекста (1 секунда).
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Пример вызова метода CreatePartner.
	partnerResp, err := partnerClient.CreatePartner(ctx, &pb.CreatePartnerRequest{
		Name:  "Test Partner",
		Email: "partner@example.com",
	})
	if err != nil {
		log.Fatalf("Ошибка при создании партнера: %v", err)
	}
	log.Printf("Партнер создан: %+v", partnerResp.Partner)

	// Пример вызова метода GetPartnerStats.
	statsResp, err := partnerClient.GetPartnerStats(ctx, &pb.GetPartnerStatsRequest{
		PartnerId: partnerResp.Partner.Id,
	})
	if err != nil {
		log.Fatalf("Ошибка при получении статистики партнера: %v", err)
	}
	log.Printf("Статистика партнера: %+v", statsResp)

	// Пример вызова метода ShortenLink.
	shortResp, err := linkClient.ShortenLink(ctx, &pb.ShortenLinkRequest{
		OriginalUrl: "http://example.com/some/very/long/url",
	})
	if err != nil {
		log.Fatalf("Ошибка при сокращении ссылки: %v", err)
	}
	log.Printf("Сокращенная ссылка: %s", shortResp.ShortUrl)

	// Пример вызова метода GetLinkStats.
	linkStatsResp, err := linkClient.GetLinkStats(ctx, &pb.GetLinkStatsRequest{
		ShortUrl: shortResp.ShortUrl,
	})
	if err != nil {
		log.Fatalf("Ошибка при получении статистики ссылки: %v", err)
	}
	log.Printf("Статистика ссылки: %+v", linkStatsResp)
}