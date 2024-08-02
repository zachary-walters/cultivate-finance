package services

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type LinkService interface{}

type linkService struct {
	PGClient *sqlx.DB
	Ctx      context.Context
}

func NewLinkService(pgClient *sqlx.DB, ctx context.Context) *linkService {
	return &linkService{
		PGClient: pgClient,
		Ctx:      ctx,
	}
}

func (svc *linkService) GetAvailableLink() (Link, error) {
	link := Link{}
	sql := `
		select
		  *
		from 
		  links
		where 
		  date is null 
			or date < current_date - interval '30' day
		limit 1
		;
	`

	err := svc.PGClient.QueryRow(sql).Scan(&link.Slug, &link.Date, &link.InUse, &link.Input, &link.Link)
	if err != nil {
		log.Println(err)
	}

	return link, err
}

func (svc *linkService) SaveLink(input Input) (Link, error) {
	link, err := svc.GetAvailableLink()
	if err != nil {
		log.Println(err)
		return link, err
	}

	updateSql := `update links set date = :date, in_use = :in_use, link = :link where slug = :slug ;`

	now := time.Now()

	link.Date = &now
	link.InUse = true
	link.Link = &input.Link

	_, err = svc.PGClient.NamedExec(updateSql, link)
	if err != nil {
		log.Println("Error saving link:", err)
	}

	return link, err
}

func (svc *linkService) GetLink(linkKey string) (Link, error) {
	link := Link{}

	if err := svc.PGClient.QueryRow("select * from links where slug = $1", linkKey).Scan(&link.Slug, &link.Date, &link.InUse, &link.Input, &link.Link); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return link, err
		}
		log.Println(err)
		return link, err
	}

	return link, nil
}

func (svc *linkService) GenerateAll(length int, characters string) error {
	chunkSize := 65535
	sql := `insert into links (slug, date, in_use, input, link) values (:slug, :date, :in_use, :input, :link);`

	m := []map[string]interface{}{}
	mChunks := [][]map[string]interface{}{}
	ns := svc.nextString(length, characters)
	for {
		str := ns()
		if len(str) == 0 {
			break
		}

		m = append(m, map[string]interface{}{
			"slug":   str,
			"date":   nil,
			"in_use": false,
			"input":  nil,
			"link":   nil,
		})

		if len(m) >= chunkSize/5 {
			mChunks = append(mChunks, m)
			m = []map[string]interface{}{}
		}
	}

	mChunks = append(mChunks, m)

	for _, m := range mChunks {
		_, err := svc.PGClient.NamedExec(sql, m)
		if err != nil {
			log.Println("Error generating all:", err)
			return err
		}
	}

	return nil
}

func (svc *linkService) GenerateSlice(length int, characters string) []string {
	s := []string{}

	ns := svc.nextString(length, characters)
	for {
		str := ns()
		if len(str) == 0 {
			break
		}
		s = append(s, str)
	}

	return s
}

func (svc *linkService) nextString(n int, c string) func() string {
	r := []rune(c)
	p := make([]rune, n)
	x := make([]int, len(p))
	return func() string {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = r[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(r) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return string(p)
	}
}
