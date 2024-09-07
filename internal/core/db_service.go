package core

import "context"

type Column struct {
	ColumnName string
	ColumnType string
}

type TableData struct {
	TableName string
	Columns   []*Column
}

type SchemaData struct {
	SchemaName string
	Tables     []*TableData
}

type DbDataService interface {
	ChangeDb(dbName string) error

	GetSchemaElements(ctx context.Context) ([]*SchemaData, error)

	GetAllDbNames(ctx context.Context) ([]string, error)

	RunQuery(ctx context.Context, query string) ([][]string, error)
}
