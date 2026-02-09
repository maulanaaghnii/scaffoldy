package utils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

func GetFieldsFromAST(filePath string, structName string) ([]FieldInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	var infos []FieldInfo

	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// Jika structName kosong, ambil struct pertama yang ketemu
		// Jika tidak kosong, cari yang namanya cocok
		if structName != "" && ts.Name.Name != structName {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range st.Fields.List {
			fieldType := ""
			switch t := field.Type.(type) {
			case *ast.Ident:
				fieldType = t.Name
			case *ast.SelectorExpr:
				fieldType = t.X.(*ast.Ident).Name + "." + t.Sel.Name
			case *ast.StarExpr:
				if ident, ok := t.X.(*ast.Ident); ok {
					fieldType = "*" + ident.Name
				}
			}

			// Handle embedded fields (audit trails etc)
			isEmbedded := len(field.Names) == 0

			if isEmbedded {
				// Spesifik handle AuditTrails untuk sistem ini
				if fieldType == "shared.AuditTrails" || fieldType == "AuditTrails" {
					infos = append(infos, FieldInfo{
						Name:       "CreatedAt",
						Type:       "time.Time",
						IsEmbedded: false,
					})
					infos = append(infos, FieldInfo{
						Name:       "CreatedBy",
						Type:       "string",
						IsEmbedded: false,
					})
					infos = append(infos, FieldInfo{
						Name:       "UpdatedAt",
						Type:       "time.Time",
						IsEmbedded: false,
					})
					infos = append(infos, FieldInfo{
						Name:       "UpdatedBy",
						Type:       "string",
						IsEmbedded: false,
					})
					continue
				}
			}

			for _, name := range field.Names {
				infos = append(infos, FieldInfo{
					Name:       name.Name,
					Type:       fieldType,
					IsEmbedded: false,
				})
			}
		}
		return false // Berhenti setelah ketemu struct yang dicari
	})

	return infos, nil
}

type FieldInfo struct {
	Name       string
	Type       string
	IsEmbedded bool
}

func getStructFieldNames(v any) []string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// kalau embedded struct
		if field.Anonymous {
			embeddedType := field.Type
			if embeddedType.Kind() == reflect.Pointer {
				embeddedType = embeddedType.Elem()
			}

			for j := 0; j < embeddedType.NumField(); j++ {
				embeddedField := embeddedType.Field(j)
				fields = append(fields, embeddedField.Name)
			}
			continue
		}

		fields = append(fields, field.Name)
	}

	return fields
}

func CategoryFieldsCreateSlice(v any) []string {
	exclude := map[string]bool{
		"ID":        true,
		"UpdatedAt": true,
		"UpdatedBy": true,
	}

	allFields := getStructFieldNames(v)

	var result []string
	for _, f := range allFields {
		if !exclude[f] {
			result = append(result, f)
		}
	}

	return result
}

func CategoryFieldsFullSlice(v any) []string {
	allFields := getStructFieldNames(v)
	var result []string
	for _, f := range allFields {
		result = append(result, f)
	}
	return result
}

func FilterFields(fields []FieldInfo, requestType string) []FieldInfo {
	var infos []FieldInfo

	for _, field := range fields {
		if requestType == "serviceUpdate" && (strings.ToLower(field.Name) == "id" || strings.ToLower(field.Name) == "code") {
			continue
		}

		// Handle flattening for query if it's already flattened in AST or handle here
		// In my AST parser, I already flattened AuditTrails into individual fields for simplicity
		// So we just need to handle the serviceCreate/Update case where we might want the embedded block

		if requestType == "serviceCreate" || requestType == "serviceUpdate" {
			// Jika field-field audit trails ada, kita kumpulkan jadi satu "AuditTrails" field jika itu yang diharapkan oleh template
			if field.Name == "CreatedAt" || field.Name == "CreatedBy" || field.Name == "UpdatedAt" || field.Name == "UpdatedBy" {
				// Cek apakah AuditTrails sudah ada di infos
				found := false
				for _, info := range infos {
					if info.Name == "AuditTrails" {
						found = true
						break
					}
				}
				if !found {
					infos = append(infos, FieldInfo{
						Name:       "AuditTrails",
						Type:       "shared.AuditTrails",
						IsEmbedded: true,
					})
				}
				continue
			}
		}

		infos = append(infos, field)
	}

	return infos
}

func GetFieldsInfo(v any, requestType string) []FieldInfo {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var infos []FieldInfo

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 1. Skip ID / Code logic
		// if requestType == "create" && field.Name == "ID" { // lowercase !!!!!!!!
		// 	continue
		// }
		if requestType == "serviceUpdate" && (strings.ToLower(field.Name) == "id" || strings.ToLower(field.Name) == "code") { // lowercase !!!!!!!!
			continue
		}

		// 2. Embedded struct handling
		if field.Anonymous {
			embeddedType := field.Type
			if embeddedType.Kind() == reflect.Pointer {
				embeddedType = embeddedType.Elem()
			}

			// Handle AuditTrails
			if embeddedType.Name() == "AuditTrails" {
				if requestType == "query" { // lowercase !!!!!!!!
					// ✅ flatten AuditTrails fields
					for j := 0; j < embeddedType.NumField(); j++ {
						f := embeddedType.Field(j)
						infos = append(infos, FieldInfo{
							Name: f.Name,
							Type: f.Type.String(),
						})
					}
				}
				// ❌ jangan pernah append "AuditTrails" sebagai field
				if requestType == "serviceCreate" || requestType == "serviceUpdate" {
					infos = append(infos, FieldInfo{
						Name:       "AuditTrails",
						Type:       embeddedType.String(),
						IsEmbedded: true,
					})
					continue
				}
				continue
			}

		}

		infos = append(infos, FieldInfo{
			Name: field.Name,
			Type: field.Type.String(),
		})
	}

	return infos
}

func GetFieldValueExpression(dataType string, fieldName string) string {
	res := ""
	switch dataType {
	case "string":
		res = "strings.TrimSpace(req." + fieldName + ")"
	case "int":
		res = "req." + fieldName
	case "float64":
		res = "req." + fieldName
	case "bool":
		res = "req.IsActive"
	default:
		res = "req." + fieldName
	}

	if fieldName == "ID" {
		res = "uuid.New().String()"
	}

	if fieldName == "Code" {
		res = "strings.ToUpper(strings.TrimSpace(req.Code))"
	}

	if fieldName == "IsActive" {
		res = "true"
	}

	if fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
		res = "time.Now()"
	}

	return res

}
