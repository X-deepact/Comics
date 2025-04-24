package database

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type DBInterface interface {
	Ctx(ctx context.Context) DBInterface
	Create(value interface{}) error
	Save(value interface{}) error
	Delete(value interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Model(value interface{}) DBInterface
	Where(query interface{}, args ...interface{}) DBInterface

	Updates(value interface{}) error
	UpdateColumn(column string, value interface{}) error
	UpdateColumns(values map[string]interface{}) error
	UpdateSimple(column string, value interface{}) error

	Exec(sql string, values ...interface{}) error
	Raw(sql string, values ...interface{}) DBInterface
	Scan(dest interface{}) error

	Select(query interface{}, args ...interface{}) DBInterface
	Joins(query string, args ...interface{}) DBInterface

	Or(query interface{}, args ...interface{}) DBInterface
	Not(query interface{}, args ...interface{}) DBInterface

	Group(name string) DBInterface
	Having(query interface{}, args ...interface{}) DBInterface
	Order(value interface{}) DBInterface
	Limit(limit int) DBInterface
	Offset(offset int) DBInterface

	Count(count *int64) error
	Distinct(args ...interface{}) DBInterface

	Preload(string, ...any) DBInterface
	Table(string) DBInterface
	Debug() DBInterface
	ToSQL(value interface{}) (string, error)

	Transaction(ctx context.Context, fn TransactionFunc) error

	Close()
	BeginTx() (*gorm.DB, error)
	RollbackTx(tx *gorm.DB) error
	CommitTx(tx *gorm.DB) error
}

type TransactionFunc func(ctx context.Context, tx DBInterface) error

func (d *DatabaseClient) Ctx(ctx context.Context) DBInterface {
	return &DatabaseClient{db: d.db.WithContext(ctx)}
}

func (d *DatabaseClient) Create(value interface{}) error {
	return d.db.Create(value).Error
}

func (d *DatabaseClient) Save(value interface{}) error {
	return d.db.Save(value).Error
}

func (d *DatabaseClient) Delete(value interface{}) error {
	return d.db.Delete(value).Error
}

func (d *DatabaseClient) Find(dest interface{}, conds ...interface{}) error {
	return d.db.Find(dest, conds...).Error
}

func (d *DatabaseClient) First(dest interface{}, conds ...interface{}) error {
	return d.db.First(dest, conds...).Error
}

func (d *DatabaseClient) Model(value interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Model(value)}
}

func (d *DatabaseClient) Where(query interface{}, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Where(query, args...)}
}

func (d *DatabaseClient) Updates(value interface{}) error {
	return d.db.Updates(value).Error
}

func (d *DatabaseClient) UpdateColumn(column string, value interface{}) error {
	return d.db.UpdateColumn(column, value).Error
}

func (d *DatabaseClient) UpdateColumns(values map[string]interface{}) error {
	return d.db.Updates(values).Error
}

func (d *DatabaseClient) UpdateSimple(column string, value interface{}) error {
	return d.db.Update(column, value).Error
}

func (d *DatabaseClient) Exec(sql string, values ...interface{}) error {
	return d.db.Exec(sql, values...).Error
}

func (d *DatabaseClient) Raw(sql string, values ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Raw(sql, values...)}
}

func (d *DatabaseClient) Select(query interface{}, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Select(query, args...)}
}

func (d *DatabaseClient) Joins(query string, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Joins(query, args...)}
}

func (d *DatabaseClient) Or(query interface{}, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Or(query, args...)}
}

func (d *DatabaseClient) Not(query interface{}, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Not(query, args...)}
}

func (d *DatabaseClient) Group(name string) DBInterface {
	return &DatabaseClient{db: d.db.Group(name)}
}

func (d *DatabaseClient) Having(query interface{}, args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Having(query, args...)}
}

func (d *DatabaseClient) Order(value interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Order(value)}
}

func (d *DatabaseClient) Limit(limit int) DBInterface {
	return &DatabaseClient{db: d.db.Limit(limit)}
}

func (d *DatabaseClient) Offset(offset int) DBInterface {
	return &DatabaseClient{db: d.db.Offset(offset)}
}

func (d *DatabaseClient) Count(count *int64) error {
	return d.db.Count(count).Error
}

func (d *DatabaseClient) Distinct(args ...interface{}) DBInterface {
	return &DatabaseClient{db: d.db.Distinct(args...)}
}

func (d *DatabaseClient) Preload(column string, conditions ...any) DBInterface {
	return &DatabaseClient{db: d.db.Preload(column, conditions...)}
}

func (d *DatabaseClient) Table(name string) DBInterface {
	return &DatabaseClient{db: d.db.Table(name)}
}

func (d *DatabaseClient) Debug() DBInterface {
	return &DatabaseClient{db: d.db.Debug()}
}

func (d *DatabaseClient) Transaction(ctx context.Context, fn TransactionFunc) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		return fn(ctx, &DatabaseClient{db: tx})
	})
}

func (db *DatabaseClient) Close() {
	slog.Info("closing database connection")
	instance, err := db.db.DB()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to close database connection, %v", err))
		panic(err)
	}
	instance.Close()
}

func (db *DatabaseClient) BeginTx() (*gorm.DB, error) {
	tx := db.db.Begin()
	return tx, nil
}

func (db *DatabaseClient) RollbackTx(tx *gorm.DB) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("failed to rollback transaction, %v", err)
	}
	return nil
}

func (db *DatabaseClient) CommitTx(tx *gorm.DB) error {
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction, %v", err)
	}
	return nil
}

func (d *DatabaseClient) ToSQL(model interface{}) (string, error) {
	if d.db == nil {
		return "", fmt.Errorf("database connection is nil")
	}

	dryRunDB := d.db.Session(&gorm.Session{DryRun: true, QueryFields: true})

	stmt := dryRunDB.Model(model).Find(model).Statement

	sql := stmt.SQL.String()
	if sql == "" {
		return "", fmt.Errorf("generated SQL query is empty")
	}

	sql = dryRunDB.Dialector.Explain(sql, stmt.Vars...)

	return sql, nil
}

func (d *DatabaseClient) Scan(dest interface{}) error {
	return d.db.Scan(dest).Error
}
