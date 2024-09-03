package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ThunderGod77/dbam/internal/core"
)

type schemaElements struct {
	schemaName     string
	tableName      string
	columnName     string
	underlyingType string

	charMaxLen        sql.NullInt64
	numPrecision      sql.NullInt64
	numScale          sql.NullInt64
	dateTimePrecision sql.NullInt64
}

func (pr *postgresRepo) GetSchemaElements(ctx context.Context) ([]*core.SchemaData, error) {
	rawSchemaElements, err := pr.querySchemaElements(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	schemaData := getSchemaData(rawSchemaElements)

	return schemaData, nil
}

func getColumnsData(elem *schemaElements) *core.Column {
	typeSuffix := ""

	switch true {
	case elem.charMaxLen.Valid:
		typeSuffix = fmt.Sprintf("(%d)", elem.charMaxLen.Int64)
	case elem.numPrecision.Valid && elem.numScale.Valid:
		typeSuffix = fmt.Sprintf("(%d,%d)", elem.numPrecision.Int64, elem.numScale.Int64)
	case elem.dateTimePrecision.Valid:
		typeSuffix = fmt.Sprintf("(%d)", elem.dateTimePrecision.Int64)
	}

	return &core.Column{
		ColumnName: elem.columnName,
		ColumnType: elem.underlyingType + typeSuffix,
	}
}

func getSchemaData(elems []*schemaElements) []*core.SchemaData {
	if len(elems) == 0 {
		return []*core.SchemaData{}
	}

	schemaData := []*core.SchemaData{}
	currentSchemaData := &core.SchemaData{
		SchemaName: "",
		Tables:     []*core.TableData{},
	}
	currentTableData := &core.TableData{
		TableName: "",
		Columns:   []*core.Column{},
	}

	for _, val := range elems {
		if currentSchemaData.SchemaName != val.schemaName {
			if currentSchemaData.SchemaName != "" {
				currentSchemaData.Tables = append(currentSchemaData.Tables, currentTableData)
				schemaData = append(schemaData, currentSchemaData)

			}

			currentTableData = &core.TableData{
				TableName: val.tableName,
				Columns:   []*core.Column{getColumnsData(val)},
			}

			currentSchemaData = &core.SchemaData{
				SchemaName: val.schemaName,
				Tables:     []*core.TableData{},
			}

			continue

		}

		if val.tableName != currentTableData.TableName {
			currentSchemaData.Tables = append(currentSchemaData.Tables, currentTableData)

			currentTableData = &core.TableData{
				TableName: val.tableName,
				Columns:   []*core.Column{getColumnsData(val)},
			}
			continue
		}

		currentTableData.Columns = append(currentTableData.Columns, getColumnsData(val))
	}

	currentSchemaData.Tables = append(currentSchemaData.Tables, currentTableData)

	schemaData = append(schemaData, currentSchemaData)

	return schemaData
}

func (pr *postgresRepo) querySchemaElements(ctx context.Context) ([]*schemaElements, error) {
	query := `SELECT table_schema, table_name, column_name, udt_name,
                   character_maximum_length, numeric_precision, numeric_scale, datetime_precision 
                   FROM information_schema.columns ORDER BY table_schema, table_name`

	pr.RLock()
	defer pr.RUnlock()

	rows, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	result := []*schemaElements{}

	for rows.Next() {
		cd := &schemaElements{}

		err := rows.Scan(
			&cd.schemaName, &cd.tableName, &cd.columnName, &cd.underlyingType,
			&cd.charMaxLen, &cd.numPrecision, &cd.numScale, &cd.dateTimePrecision,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, cd)
	}

	return result, nil
}
