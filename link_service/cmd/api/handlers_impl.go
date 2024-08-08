package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	services "github.com/zachary-walters/cultivate-finance/link_service/internal"
)

func getLink(db *sqlx.DB, ctx context.Context, slug string) (services.Link, error) {
	linkService := services.NewLinkService(db, ctx)
	return linkService.GetLink(slug)
}

func generateAll(db *sqlx.DB, ctx context.Context) error {
	linkService := services.NewLinkService(db, ctx)
	return linkService.GenerateAll(4, "abcdefghijklmnopqrstuvwxyz")
}

func saveLink(db *sqlx.DB, ctx context.Context) (services.Link, error) {
	linkService := services.NewLinkService(db, ctx)
	input := services.Input{
		Link: "somelink",
	}

	link := services.Link{
		Link: &input.Link,
	}

	return linkService.SaveLink(link)
}

func updateExpiredLinks(db *sqlx.DB, ctx context.Context) error {
	linkService := services.NewLinkService(db, ctx)
	return linkService.UpdateExpiredLinks()
}
