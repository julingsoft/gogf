package dbx

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func GetTableComment(ctx context.Context, tableName string) string {
	var (
		db      = g.DB()
		result  gdb.Result
		comment = ""
	)

	result, err := db.GetAll(ctx,
		`SELECT TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`,
		db.GetConfig().Name, tableName)
	if err != nil {
		return ""
	}

	if len(result) > 0 {
		comment = result[0]["TABLE_COMMENT"].String()
	}

	if gstr.SubStrRune(comment, gstr.LenRune(comment)-1) == "表" {
		comment = gstr.SubStrRune(comment, 0, gstr.LenRune(comment)-1)
	}

	return comment
}

func GetTableColumns(ctx context.Context, tableName string) []string {
	var (
		db     = g.DB()
		result gdb.Result
	)

	result, err := db.GetAll(ctx,
		`SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`,
		db.GetConfig().Name, tableName)
	if err != nil {
		return nil
	}

	var columns []string
	for _, item := range result {
		columns = append(columns, item["COLUMN_NAME"].String())
	}

	return columns // sorted
}

func GetAllColumns(ctx context.Context, table string) []map[string]string {
	fields, err := g.DB().TableFields(ctx, table)
	if err != nil {
		panic(err)
	}

	var keys = GetTableColumns(ctx, table)
	var allColumns = make([]map[string]string, len(keys))
	for i, key := range keys {
		fields[key].Comment = gstr.ReplaceByMap(fields[key].Comment, map[string]string{
			"：":  ":",
			"，":  ":",
			",":  ":",
			" ":  ":",
			"\t": ":",
		})
		if gstr.Contains(fields[key].Comment, ":") {
			fields[key].Comment = gstr.Split(fields[key].Comment, ":")[0]
		}

		var comment = fields[key].Comment
		if g.IsEmpty(comment) && fields[key].Name == "id" {
			comment = "ID"
		}
		allColumns[i] = map[string]string{
			"Name":               fields[key].Name,
			"Comment":            comment,
			"CaseCamelName":      gstr.CaseCamel(fields[key].Name),
			"CaseCamelLowerName": gstr.CaseCamelLower(fields[key].Name),
			"Type":               GetFieldType(fields[key].Type),
			"Key":                fields[key].Key,
		}
	}

	return allColumns
}

func GetKeyColumns(ctx context.Context, table string) []map[string]string {
	var allColumns = GetAllColumns(ctx, table)

	var keyColumns = make([]map[string]string, 0)
	for _, fieldInfo := range allColumns {
		if !g.IsEmpty(fieldInfo["Key"]) {
			keyColumns = append(keyColumns, fieldInfo)
		}
	}

	return keyColumns
}

func GetFieldType(fieldType string) string {
	m, err := gregex.MatchString(`(\w+)\(`, fieldType)
	if err != nil {
		panic(err)
	}

	if len(m) > 1 {
		fieldType = m[1]
	}

	var unsigned = gstr.ContainsI(fieldType, "unsigned")
	fieldType = gstr.ReplaceByMap(fieldType, map[string]string{
		" unsigned": "",
	})

	if gstr.InArray([]string{"bigint"}, fieldType) {
		if unsigned {
			return "uint64"
		}
		return "int64"
	}

	if gstr.InArray([]string{"bit", "int", "mediumint", "smallint", "tinyint", "enum"}, fieldType) {
		if unsigned {
			return "uint"
		}
		return "int"
	}

	if gstr.InArray([]string{"decimal", "float", "double"}, fieldType) {
		return "float64"
	}

	if gstr.InArray([]string{"blob", "binary"}, fieldType) {
		return "[]byte"
	}

	if gstr.InArray([]string{"date", "datetime", "timestamp", "time"}, fieldType) {
		return "*gtime.Time"
	}

	return "string"
}
