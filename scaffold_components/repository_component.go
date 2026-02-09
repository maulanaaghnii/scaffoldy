package scaffoldcomponents

import (
	"scaffoldy/pkg/utils"
)

func RepositoryContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {

	if tableName == "" {
		tableName = "tbl" + utils.LowerFirst(domainName)
	}

	res := "package " + utils.LowerFirst(domainName) + "\n\n"
	res += "import ("
	res += "\n\t\"database/sql\""
	res += "\n\t\"errors\""
	res += "\n\t\"fmt\""
	res += "\n\t\"github.com/go-sql-driver/mysql\""
	res += ")\n\n"

	res += "// Save\n"
	res += "// Update\n"
	res += "// FindAll\n"
	res += "// FindById\n"
	res += "// FindByCode\n"
	res += "// SoftDelete\n"
	res += "// Delete\n"

	res += "var ("
	res += "\n\tErr" + domainName + "NotFound = errors.New(\"" + domainName + " not found\")\n"
	res += "\n\tErr" + domainName + "CodeDuplicate = errors.New(\"" + domainName + " code already exists\")\n"
	res += ")\n\n"

	res += "type Repository struct {\n"
	res += "\tdb *sql.DB\n"
	res += "}\n\n"

	res += "func NewRepository(db *sql.DB) *Repository {\n"
	res += "\treturn &Repository{db: db}\n"
	res += "}\n\n"

	res += SaveContent(domainName, fieldsQuery, tableName)
	res += UpdateContent(domainName, fieldsQuery, tableName)
	res += FindAllContent(domainName, fieldsQuery, tableName)
	res += FindByIdContent(domainName, fieldsQuery, tableName)
	res += FindByCodeContent(domainName, fieldsQuery, tableName)
	res += SoftDeleteContent(domainName, fieldsQuery, tableName)
	res += DeleteContent(domainName, fieldsQuery, tableName)

	return res
}

func SaveContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) Save(" + utils.LowerFirst(domainName) + " " + domainName + ") error {\n"
	res += "\tquery := `\n"
	res += "\t\tINSERT INTO " + tableName + " (\n"
	for _, field := range fieldsQuery {
		res += "\t\t\t" + field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ",\n"
		}
	}
	res += "\n\t\t)\n"
	res += "\t\tVALUES ("
	for _, field := range fieldsQuery {
		res += "?"
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ","
		}
	}
	res += "\n\t\t)\n"
	res += "\t`"
	// res += "\t`\n"
	res += "\n"
	res += "\t_, err := r.db.Exec(query,\n"

	for _, field := range fieldsQuery {
		res += "\t\t" + utils.LowerFirst(domainName) + "." + field.Name
		res += ",\n"
	}

	res += "\t)\n"

	res += "\n"
	res += "\tif err != nil {\n"
	res += "\t\tif mysqlErr, ok := err.(*mysql.MySQLError); ok {\n"
	res += "\t\t\tif mysqlErr.Number == 1062 {\n"
	res += "\t\t\t	return Err" + domainName + "CodeDuplicate\n"
	res += "\t\t\t}\n"
	res += "\t\t}\n\t\t\n"
	res += "\t\treturn fmt.Errorf(\"failed to save " + domainName + ": %w\", err)\n"
	res += "\t}\n\n"
	res += "\treturn nil\n"
	res += "}\n"

	return res
}

func UpdateContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) Update(" + utils.LowerFirst(domainName) + " " + domainName + ") error {\n"
	res += "\tquery := `\n"
	res += "\t\tUPDATE " + tableName + " SET "
	res += "\n"
	for _, field := range fieldsQuery {
		res += "\t\t\t" + field.Name + " = ?"
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ",\n"
		}
	}
	res += "\n\t\tWHERE ID = ?\n"
	res += "\t`\n\n"
	res += "\tresult, err := r.db.Exec(query,\n"

	for _, field := range fieldsQuery {
		res += "\t\t" + utils.LowerFirst(domainName) + "." + field.Name
		res += ",\n"
	}

	res += "\t)\n"

	res += "\n"
	res += "\tif err != nil {\n"
	res += "\t\tif mysqlErr, ok := err.(*mysql.MySQLError); ok {\n"
	res += "\t\t	if mysqlErr.Number == 1062 {\n"
	res += "\t\t		return Err" + domainName + "CodeDuplicate\n"
	res += "\t\t	}\n"
	res += "\t\t}\n"
	res += "\t\t\n"
	res += "\t\treturn fmt.Errorf(\"failed to update " + domainName + ": %w\", err)\n"
	res += "\t}\n\n"

	res += "\trowsAffected, err := result.RowsAffected()\n"
	res += "\tif err != nil {\n"
	res += "\t	return fmt.Errorf(\"failed to get rows affected: %w\", err)\n"
	res += "\t}\n\n"

	res += "\tif rowsAffected == 0 {\n"
	res += "\t	return Err" + domainName + "NotFound\n"
	res += "\t}\n\n"

	res += "\treturn nil\n"
	res += "}\n"

	res += "\n"

	return res
}

func FindAllContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) FindAll() ([]" + domainName + ", error) {\n"
	res += "\tquery := `\n"
	res += "\t	SELECT "
	for _, field := range fieldsQuery {
		res += field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ", "
		}
	}
	res += "\n\t\tFROM " + tableName + "\n"
	res += "\t`\n\n"
	res += "\trows, err := r.db.Query(query)\n"
	res += "\tif err != nil {\n"
	res += "\t	return nil, fmt.Errorf(\"failed to query " + domainName + ": %w\", err)\n"
	res += "\t}\n"
	res += "\tdefer rows.Close()\n\n"
	res += "\t " + utils.LowerFirst(domainName) + "List := make([]" + domainName + ", 0)\n"
	res += "\tfor rows.Next() {\n"
	res += "\t	var " + utils.LowerFirst(domainName) + " " + domainName + "\n"
	res += "\t	err := rows.Scan(\n"
	for _, field := range fieldsQuery {
		res += "\t		&" + utils.LowerFirst(domainName) + "." + field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ",\n"
		}
	}

	res += "\t	)\n"
	res += "\t	if err != nil {\n"
	res += "\t		return nil, fmt.Errorf(\"failed to scan " + domainName + ": %w\", err)\n"
	res += "\t	}\n"
	res += "\t	" + utils.LowerFirst(domainName) + "List = append(" + utils.LowerFirst(domainName) + "List, " + utils.LowerFirst(domainName) + ")\n"
	res += "\t}\n\n"
	res += "\treturn " + utils.LowerFirst(domainName) + "List , nil\n"
	res += "}\n"
	res += "\n"
	return res
}

func FindByIdContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) FindById(id string) (" + domainName + ", error) {\n"
	res += "\tquery := `\n"
	res += "\t	SELECT "
	for _, field := range fieldsQuery {
		res += field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ", "
		}
	}
	res += "\n\t\tFROM " + tableName + "\n"
	res += "\t\tWHERE ID = ?\n"
	res += "\t`\n\n"
	res += "\trow := r.db.QueryRow(query, id)\n"
	res += "\tvar " + utils.LowerFirst(domainName) + " " + domainName + "\n"
	res += "\terr := row.Scan(\n"
	for _, field := range fieldsQuery {
		res += "\t\t&" + utils.LowerFirst(domainName) + "." + field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ",\n"
		}
	}
	res += "\t)\n"
	res += "\tif err != nil {\n"
	res += "\t	if err == sql.ErrNoRows {\n"
	res += "\t		return " + domainName + "{}, Err" + domainName + "NotFound\n"
	res += "\t	}\n"
	res += "\t	return " + domainName + "{}, fmt.Errorf(\"failed to scan " + domainName + ": %w\", err)\n"
	res += "\t}\n"
	res += "\treturn " + utils.LowerFirst(domainName) + ", nil\n"
	res += "}\n"
	res += "\n"
	return res
}

func FindByCodeContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) FindByCode(code string) (" + domainName + ", error) {\n"
	res += "\tquery := `\n"
	res += "\t	SELECT "
	for _, field := range fieldsQuery {
		res += field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ", "
		}
	}
	res += "\n\t\tFROM " + tableName + "\n"
	res += "\t\tWHERE Code = ?\n"
	res += "\t`\n\n"
	res += "\trow := r.db.QueryRow(query, code)\n"
	res += "\tvar " + utils.LowerFirst(domainName) + " " + domainName + "\n"
	res += "\terr := row.Scan(\n"
	for _, field := range fieldsQuery {
		res += "\t\t&" + utils.LowerFirst(domainName) + "." + field.Name
		if field.Name != fieldsQuery[len(fieldsQuery)-1].Name {
			res += ",\n"
		}
	}
	res += "\t)\n"
	res += "\tif err != nil {\n"
	res += "\t	if err == sql.ErrNoRows {\n"
	res += "\t		return " + domainName + "{}, Err" + domainName + "NotFound\n"
	res += "\t	}\n"
	res += "\t	return " + domainName + "{}, fmt.Errorf(\"failed to scan " + domainName + ": %w\", err)\n"
	res += "\t}\n"
	res += "\treturn " + utils.LowerFirst(domainName) + ", nil\n"
	res += "}\n"
	res += "\n"
	return res
}

func SoftDeleteContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) SoftDelete(id string) error {\n"
	res += "\tquery := `\n"
	res += "\t	UPDATE " + tableName + " SET IsActive = false\n"
	res += "\t	WHERE ID = ?\n"
	res += "\t`\n\n"
	res += "\t_, err := r.db.Exec(query, id)\n"
	res += "\tif err != nil {\n"
	res += "\t	return fmt.Errorf(\"failed to delete " + domainName + ": %w\", err)\n"
	res += "\t}\n"
	res += "\treturn nil\n"
	res += "}\n"
	res += "\n"
	return res
}

func DeleteContent(domainName string, fieldsQuery []utils.FieldInfo, tableName string) string {
	res := "func (r *Repository) Delete(id string) error {\n"
	res += "\tquery := `\n"
	res += "\t	DELETE FROM " + tableName + "\n"
	res += "\t	WHERE ID = ?\n"
	res += "\t`\n\n"
	res += "\t_, err := r.db.Exec(query, id)\n"
	res += "\tif err != nil {\n"
	res += "\t	return fmt.Errorf(\"failed to delete " + domainName + ": %w\", err)\n"
	res += "\t}\n"
	res += "\treturn nil\n"
	res += "}\n"
	res += "\n"
	return res
}
