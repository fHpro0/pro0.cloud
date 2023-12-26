package database

import (
	"context"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	metrics "pro0.cloud/v2/lib/metric"
	"pro0.cloud/v2/lib/secureString"
	"strings"
	"sync"
)

type Dsn struct {
	secureString.SecureString
}

func (d *Dsn) String() string {
	if dsn, err := d.Get(); err == nil {
		return dsn
	}
	return ""
}

type Db struct {
	Dbh     rel.Adapter
	Repo    rel.Repository
	Metrics *metrics.Metrics
	mu      sync.RWMutex
}

func NewDb(dsn *Dsn) (*Db, error) {
	dbh, repo, err := connect(dsn)
	if err != nil {
		return nil, err
	}

	db := &Db{
		Dbh:  dbh,
		Repo: repo,
	}
	db.Repo.Instrumentation(db.logger)

	return db, nil
}

func (db *Db) Close() error {
	return db.Dbh.Close()
}

func (db *Db) Reconnect(dsn *Dsn) error {
	// Open new connection
	dbh, repo, err := connect(dsn)
	if err != nil {
		return err
	}

	_ = db.Close()

	db.mu.Lock()
	defer db.mu.Unlock()

	db.Dbh = dbh
	db.Repo = repo
	db.Repo.Instrumentation(db.logger)

	return nil
}

func connect(dsn *Dsn) (rel.Adapter, rel.Repository, error) {
	dbh, err := mysql.Open(dsn.String())
	if err != nil {
		return nil, nil, err
	}
	// Force connection
	if err := dbh.Ping(context.Background()); err != nil {
		return nil, nil, err
	}

	repo := rel.New(dbh)

	return dbh, repo, nil
}

// logger instrumentation to log queries and rel operation.
func (db *Db) logger(ctx context.Context, op string, message string, args ...any) func(err error) {
	// no op for rel functions.
	if strings.HasPrefix(op, "rel-") {
		return func(error) {}
	}

	return func(err error) {
		if err != nil {
			go db.Metrics.AddDb(metrics.DbError)
		}
	}
}
