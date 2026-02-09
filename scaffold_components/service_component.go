package scaffoldcomponents

import "scaffoldy/pkg/utils"

func GetAllServiceContent(domainName string) string {
	res := ""
	res += "func (s *Service) GetAll" + domainName + "() ([]" + domainName + ", error) {\n"
	res += "\treturn s.repository.FindAll()\n"
	res += "}\n\n"
	return res
}

func GetByIDServiceContent(domainName string) string {
	res := ""
	res += "func (s *Service) Get" + domainName + "ByID(id string) (" + domainName + ", error) {\n"
	res += "\treturn s.repository.FindById(id)\n"
	res += "}\n\n"
	return res
}

func GetByCodeServiceContent(domainName string) string {
	res := ""
	res += "func (s *Service) Get" + domainName + "ByCode(code string) (" + domainName + ", error) {\n"
	res += "\treturn s.repository.FindByCode(code)\n"
	res += "}\n\n"
	return res
}

func CreateServiceContent(domainName string, createFields []utils.FieldInfo) string {
	res := ""
	res += "func (s *Service) Create" + domainName + "(req Create" + domainName + "Request) (" + domainName + ", error) {\n"
	res += "\tif err := s.validateCreateRequest(req); err != nil {\n"
	res += "\t	return " + domainName + "{}, err\n"
	res += "\t}\n"

	res += "\t" + utils.LowerFirst(domainName) + " := " + domainName + "{\n"

	for _, field := range createFields {

		// ✅ Special case: AuditTrails
		if field.IsEmbedded && field.Name == "AuditTrails" {
			res += "\t\tAuditTrails: shared.AuditTrails{\n"
			res += "\t\t\tCreatedAt: time.Now(),\n"
			res += "\t\t\tCreatedBy: \"system\",\n"
			res += "\t\t},\n"
			continue
		}

		// Normal field
		res += "\t\t" + field.Name + ": " +
			utils.GetFieldValueExpression(field.Type, field.Name) + ",\n"
	}

	res += "\t}\n"

	res += "\terr := s.repository.Save(" + utils.LowerFirst(domainName) + ")\n"
	res += "\treturn " + utils.LowerFirst(domainName) + ", err\n"
	res += "}\n\n"
	return res
}

func UpdateServiceContent(domainName string, updateFields []utils.FieldInfo) string {
	res := ""
	res += "func (s *Service) Update" + domainName + "(id string, req Update" + domainName + "Request) (" + domainName + ", error) {\n"
	res += "\texisting, err := s.repository.FindById(id)\n"
	res += "\tif err != nil {\n"
	res += "\t	return " + domainName + "{}, err\n"
	res += "\t}\n"

	res += "\tif err := s.validateUpdateRequest(req); err != nil {\n"
	res += "\t	return " + domainName + "{}, err\n"
	res += "\t}\n"

	for _, field := range updateFields {
		// ✅ Special case: AuditTrails
		if field.IsEmbedded && field.Name == "AuditTrails" {
			res += "\texisting.UpdatedAt = time.Now()\n"
			res += "\texisting.UpdatedBy = \"system\"\n"
			continue
		}

		// Normal field

		res += "\texisting." + field.Name + " = " +
			utils.GetFieldValueExpression(field.Type, field.Name) + "\n"

	}

	res += "err = s.repository.Update(existing)\n"
	res += "return existing, err\n"
	res += "}\n\n"
	return res
}

func SoftDeleteServiceContent(domainName string) string {
	res := ""
	res += "func (s *Service) SoftDelete" + domainName + "(id string) error {\n"
	res += "\treturn s.repository.SoftDelete(id)\n"
	res += "}\n\n"
	return res
}

func DeleteServiceContent(domainName string) string {
	res := ""
	res += "func (s *Service) Delete" + domainName + "(id string) error {\n"
	res += "\treturn s.repository.Delete(id)\n"
	res += "}\n\n"
	return res
}

func ValidationServiceContent(domainName string) string {
	res := ""
	res += "\n\n"
	res += "func (s *Service) validateCreateRequest(req Create" + domainName + "Request) error { \n"
	res += "\t// if strings.TrimSpace(req.Code) == \"\" {\n"
	res += "\t// 	return fmt.Errorf(\"code is required\")\n"
	res += "\t// }\n"
	res += "\t// if strings.TrimSpace(req.Name) == \"\" {\n"
	res += "\t// 	return fmt.Errorf(\"name is required\")\n"
	res += "\t// }\n"
	res += "\t// if req.Price < 0 {\n"
	res += "\t// 	return fmt.Errorf(\"price cannot be negative\")\n"
	res += "\t// }\n"
	res += "\t// if req.Stock < 0 {\n"
	res += "\t// 	return fmt.Errorf(\"stock cannot be negative\")\n"
	res += "\t// }\n"
	res += "\treturn nil\n"
	res += "}\n\n"
	res += "func (s *Service) validateUpdateRequest(req Update" + domainName + "Request) error {\n"
	res += "\t// if strings.TrimSpace(req.Name) == \"\" {\n"
	res += "\t// 	return fmt.Errorf(\"name is required\")\n"
	res += "\t// }\n"
	res += "\t// if req.Price < 0 {\n"
	res += "\t// 	return fmt.Errorf(\"price cannot be negative\")\n"
	res += "\t// }\n"
	res += "\t// if req.Stock < 0 {\n"
	res += "\t// 	return fmt.Errorf(\"stock cannot be negative\")\n"
	res += "\t// }\n"
	res += "\treturn nil\n"
	res += "}\n\n"
	return res
}

func ServiceContent(domainName string, createFields []utils.FieldInfo, updateFields []utils.FieldInfo) string {
	res := "package " + utils.LowerFirst(domainName) + "\n\n"

	res += "import (\n"
	// res += "\"fmt\"\n"
	res += "\"strings\"\n"
	res += "\"time\"\n\n"
	// res += "\"learn-golang/internal/shared\"\n\n"
	res += "\"github.com/google/uuid\"\n"
	res += "\"scaffoldy/internal/shared\"\n"
	res += ")\n\n"

	res += "type Service struct {\n"
	res += "\trepository *Repository\n"
	res += "}\n\n"
	res += "func NewService(repository *Repository) *Service {\n"
	res += "\treturn &Service{repository: repository}\n"
	res += "}\n\n"

	res += "// GetAll" + domainName + "\n"
	res += "// Get" + domainName + "ByID\n"
	res += "// Get" + domainName + "ByCode\n"
	res += "// Create" + domainName + "\n"
	res += "// Update" + domainName + "\n"
	res += "// Delete" + domainName + "\n"

	res += GetAllServiceContent(domainName)
	res += GetByIDServiceContent(domainName)
	res += GetByCodeServiceContent(domainName)
	res += CreateServiceContent(domainName, createFields)
	res += UpdateServiceContent(domainName, updateFields)
	res += SoftDeleteServiceContent(domainName)
	res += DeleteServiceContent(domainName)
	res += ValidationServiceContent(domainName)

	return res
}
