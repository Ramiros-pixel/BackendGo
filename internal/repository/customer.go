package repository

import (
	"context"
	"time"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"shellrean.id/Go-RestAPI/domain"
)

type customerRepository struct{
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository{
	return &customerRepository{
		db: goqu.New("mysql", con),
	}
}

func (cr customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error){
	dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr customerRepository) FindById(ctx context.Context, id string) (result domain.Customer, err error){
	dataset := cr.db.From("customers").Where(
		goqu.C("deleted_at").IsNull(),
		goqu.C("id").Eq(id),
	)
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (cr customerRepository) Save(ctx context.Context, c *domain.Customer) error{
	executor := cr.db.Insert("customers").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Update(ctx context.Context, c *domain.Customer) error{
	executor := cr.db.Update("customers").Set(c).Where(goqu.C("id").Eq(c.ID)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Delete(ctx context.Context, id string) error{
	executor := cr.db.Update("customers").
		Set(goqu.Record{"deleted_at": time.Now()}).
		Where(goqu.C("id").Eq(id)).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
