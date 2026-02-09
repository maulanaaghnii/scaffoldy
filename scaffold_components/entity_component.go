package scaffoldcomponents

import (
	"fmt"
	"scaffoldy/pkg/utils"
)

func EntityContent(fields []utils.FieldInfo) string {
	res := "type CreateBarangRequest struct {\n"

	for _, field := range fields {
		res += fmt.Sprintf("\t%s %s\n", field.Name, field.Type)
	}

	res += "}"

	return res
}
