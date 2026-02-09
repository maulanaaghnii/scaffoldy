package scaffoldcomponents

import (
	"fmt"
	"scaffoldy/pkg/utils"
)

func RequestComponent(domainName string, fieldsCreate []utils.FieldInfo, fieldsUpdate []utils.FieldInfo) string {

	res := "package " + utils.LowerFirst(domainName) + "\n\n"
	res += "import " + "\"time\"" + "\n"
	res += "\n"
	res += "type Create" + domainName + "Request struct {\n"

	for _, field := range fieldsCreate {
		// We use double quotes for the json tag to avoid backtick nesting issues
		// or concat them if we really want backticks in the output string.
		res += fmt.Sprintf("\t%s %s `json:\"%s\"` \n", field.Name, field.Type, utils.LowerFirst(field.Name))
	}

	res += "} \n\n"

	res += "type Update" + domainName + "Request struct {\n"

	for _, field := range fieldsUpdate {
		// We use double quotes for the json tag to avoid backtick nesting issues
		// or concat them if we really want backticks in the output string.
		res += fmt.Sprintf("\t%s %s `json:\"%s\"` \n", field.Name, field.Type, utils.LowerFirst(field.Name))
	}

	res += "}"

	return res
}
