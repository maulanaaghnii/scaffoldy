package utils

import (
	"reflect"
	"strings"
)

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
