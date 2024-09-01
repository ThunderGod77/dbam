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
	GetSchemaElements(ctx context.Context) ([]*SchemaData, error)
	GetAllDatabases(ctx context.Context) ([]string, error)
}
