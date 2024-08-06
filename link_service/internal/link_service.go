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
	DB  *sqlx.DB
	Ctx context.Context
}

func NewLinkService(pgClient *sqlx.DB, ctx context.Context) *linkService {
	return &linkService{
		DB:  pgClient,
		Ctx: ctx,
	}
}

func (svc *linkService) UpdateExpiredLinks() error {
	sql := `
		update links set
			date = null,
			input = null,
			link = null
		where 
		  date < current_date - interval '30' day
		;
	`

	_, err := svc.DB.Exec(sql)
	if err != nil {
		log.Println("Error updating expired links:", err)
	}

	return err
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

	err := svc.DB.QueryRow(sql).Scan(&link.Slug, &link.Date, &link.Metadata, &link.Link)
	if err != nil {
		log.Println(err)
	}

	return link, err
}

func (svc *linkService) SaveLink(link Link) (Link, error) {
	newLink, err := svc.GetAvailableLink()
	if err != nil {
		log.Println(err)
		return link, err
	}

	updateSql := `update links set date = :date, link = :link where slug = :slug ;`

	now := time.Now()

	newLink.Date = &now
	newLink.Link = link.Link

	_, err = svc.DB.NamedExec(updateSql, newLink)
	if err != nil {
		log.Println("Error saving link:", err)
	}

	return newLink, err
}

func (svc *linkService) GetLink(linkKey string) (Link, error) {
	link := Link{
		Slug: linkKey,
	}

	if err := svc.DB.QueryRow("select * from links where slug = $1", linkKey).Scan(&link.Slug, &link.Date, &link.Metadata, &link.Link); err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return link, err
		}
		log.Println(err)
		return link, err
	}

	if link.Date == nil {
		return link, nil
	} else if link.Date.Before(time.Now().AddDate(0, -1, 0)) {
		err := svc.ResetLink(linkKey)
		if err != nil {
			return link, err
		}
	} else {
		link, err := svc.SaveLink(link)
		if err != nil {
			return link, err
		}
	}

	return link, nil
}

func (svc *linkService) ResetLink(linkKey string) error {
	sql := `update links set date = null, input = null, link = null where slug = $1`

	_, err := svc.DB.Exec(sql, linkKey)
	if err != nil {
		log.Println("Error resetting link:", err)
	}

	return err
}

func (svc *linkService) GenerateAll(length int, characters string) error {
	chunkSize := 65535
	sql := `insert into links (slug, date, metadata, link) values (:slug, :date, :metadata, :link);`

	m := []map[string]interface{}{}
	mChunks := [][]map[string]interface{}{}
	ns := svc.nextString(length, characters)
	for {
		str := ns()
		if _, ok := badWords[str]; ok {
			continue
		}

		if len(str) == 0 {
			break
		}

		m = append(m, map[string]interface{}{
			"slug":     str,
			"date":     nil,
			"metadata": nil,
			"link":     nil,
		})

		if len(m) >= chunkSize/5 {
			mChunks = append(mChunks, m)
			m = []map[string]interface{}{}
		}
	}

	mChunks = append(mChunks, m)

	for _, m := range mChunks {
		_, err := svc.DB.NamedExec(sql, m)
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
		if _, ok := badWords[str]; ok {
			continue
		}

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

var badWords = map[string]bool{
	"anal": true,
	"anus": true,
	"asss": true,
	"boob": true,
	"cock": true,
	"cumm": true,
	"coon": true,
	"dick": true,
	"fagg": true,
	"fagt": true,
	"fags": true,
	"fuck": true,
	"jizz": true,
	"milf": true,
	"nazi": true,
	"niga": true,
	"nigg": true,
	"nigr": true,
	"piss": true,
	"poon": true,
	"rape": true,
	"scat": true,
	"shit": true,
	"smut": true,
	"suck": true,
	"tits": true,
	"titt": true,
	"twat": true,
	"wank": true,
}
