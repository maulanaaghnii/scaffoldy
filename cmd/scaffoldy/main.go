package main

import (
	"flag"
	"fmt"
	"os"
	"scaffoldy/pkg/utils"
	_scaffold_components "scaffoldy/scaffold_components"
	"strings"
	"text/template"
)

func ScaffoldHanlderComponents(domainName string) {
	tpl := _scaffold_components.HandlerContent()
	t, err := template.New("code").Parse(tpl)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("internal/" + utils.LowerFirst(domainName) + "/handler" + ".go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := struct {
		DomainName      string
		DomainNameLower string
		DomainNameKebab string
	}{
		DomainName:      domainName,
		DomainNameLower: utils.LowerFirst(domainName),
		DomainNameKebab: utils.ToKebabCase(domainName),
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func ScaffoldRepositoryComponents(domainName string, fields []utils.FieldInfo, tableName string) {
	fieldsQuery := utils.FilterFields(fields, "query")
	tpl := _scaffold_components.RepositoryContent(domainName, fieldsQuery, tableName)

	t, err := template.New("code").Parse(tpl)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("internal/" + utils.LowerFirst(domainName) + "/repository" + ".go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := struct {
		DomainName      string
		DomainNameLower string
	}{
		DomainName:      domainName,
		DomainNameLower: utils.LowerFirst(domainName),
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func ScaffoldRequestComponents(domainName string, fields []utils.FieldInfo) {
	fieldsCreate := utils.FilterFields(fields, "create")
	fieldsUpdate := utils.FilterFields(fields, "update")
	tpl := _scaffold_components.RequestComponent(domainName, fieldsCreate, fieldsUpdate)

	t, err := template.New("code").Parse(tpl)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("internal/" + utils.LowerFirst(domainName) + "/request" + ".go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := struct {
		DomainName      string
		DomainNameLower string
	}{
		DomainName:      domainName,
		DomainNameLower: utils.LowerFirst(domainName),
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func ScaffoldServiceComponents(domainName string, fields []utils.FieldInfo) {
	fieldsCreate := utils.FilterFields(fields, "serviceCreate")
	fieldsUpdate := utils.FilterFields(fields, "serviceUpdate")
	tpl := _scaffold_components.ServiceContent(domainName, fieldsCreate, fieldsUpdate)

	t, err := template.New("code").Parse(tpl)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("internal/" + utils.LowerFirst(domainName) + "/service" + ".go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := struct {
		DomainName      string
		DomainNameLower string
	}{
		DomainName:      domainName,
		DomainNameLower: utils.LowerFirst(domainName),
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func InjectRegistration(domainName string) {
	apiMainPath := "cmd/api/main.go"
	content, err := os.ReadFile(apiMainPath)
	if err != nil {
		fmt.Printf("Warning: Gagal membaca %s untuk auto-injection\n", apiMainPath)
		return
	}

	body := string(content)
	domainLower := utils.LowerFirst(domainName)
	registerLine := fmt.Sprintf("\t\t%s.Register(api, db)", domainLower)
	importLine := fmt.Sprintf("\t\"scaffoldy/internal/%s\"", domainLower)

	// 1. Inject Import if not exists
	if !strings.Contains(body, importLine) {
		importMarker := "import ("
		body = strings.Replace(body, importMarker, importMarker+"\n"+importLine, 1)
	}

	// 2. Inject Register Call if not exists
	if !strings.Contains(body, registerLine) {
		marker := "// [SCAFFOLDY_INSERT_MARKER]"
		body = strings.Replace(body, marker, registerLine+"\n\t\t"+marker, 1)
	}

	err = os.WriteFile(apiMainPath, []byte(body), 0644)
	if err != nil {
		fmt.Printf("Warning: Gagal menulis ke %s\n", apiMainPath)
	} else {
		fmt.Printf("Successfully injected %s registration into api/main.go\n", domainName)
	}
}

func main() {
	domainName := flag.String("domain-name", "", "The name of the domain")
	tableName := flag.String("table-name", "", "The name of the database table")
	flag.Parse()

	if *domainName == "" || *tableName == "" {
		fmt.Println("\n[WARNING] Anda harus menyertakan --domain-name dan --table-name")
		os.Exit(1)
	}

	// CARI FILE MODEL/ENTITY SECARA OTOMATIS
	domainLower := utils.LowerFirst(*domainName)
	entityPaths := []string{
		fmt.Sprintf("internal/%s/entity.go", domainLower),
		fmt.Sprintf("internal/%s/%s.go", domainLower, domainLower),
	}

	var fields []utils.FieldInfo
	// var err error
	foundFile := ""

	for _, path := range entityPaths {
		if _, err := os.Stat(path); err == nil {
			fields, err = utils.GetFieldsFromAST(path, *domainName)
			if err == nil && len(fields) > 0 {
				foundFile = path
				break
			}
		}
	}

	if foundFile == "" {
		fmt.Printf("Error: File entity untuk domain '%s' tidak ditemukan di internal/%s/\n", *domainName, domainLower)
		fmt.Println("Pastikan ada file entity.go atau " + domainLower + ".go dengan struct " + *domainName)
		os.Exit(1)
	}

	fmt.Printf("Menggunakan model dari: %s\n", foundFile)

	// Eksekusi Scaffold
	ScaffoldHanlderComponents(*domainName)
	ScaffoldRepositoryComponents(*domainName, fields, *tableName)
	ScaffoldRequestComponents(*domainName, fields)
	ScaffoldServiceComponents(*domainName, fields)

	// Auto Inject
	InjectRegistration(*domainName)

	fmt.Printf("\n--- Scaffold Complete for Domain: %s ---\n", *domainName)
}
