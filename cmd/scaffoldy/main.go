package main

import (
	"flag"
	"os"
	"scaffoldy/internal/category"
	"scaffoldy/pkg/utils"
	_scaffold_components "scaffoldy/scaffold_components"
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
	}{
		DomainName:      domainName,
		DomainNameLower: utils.LowerFirst(domainName),
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func ScaffoldRepositoryComponents(domainName string, v any, tableName string) {
	fieldsQuery := utils.GetFieldsInfo(v, "query")
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

func ScaffoldRequestComponents(domainName string, v any) {
	fieldsCreate := utils.GetFieldsInfo(v, "create")
	fieldsUpdate := utils.GetFieldsInfo(v, "update")
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

func ScaffoldServiceComponents(domainName string, v any) {
	fieldsCreate := utils.GetFieldsInfo(v, "serviceCreate")
	fieldsUpdate := utils.GetFieldsInfo(v, "serviceUpdate")
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

func main() {
	// Define flags with empty defaults to make them "mandatory" in logic
	domainName := flag.String("domain-name", "", "The name of the domain")
	tableName := flag.String("table-name", "", "The name of the database table")

	// Parse the flags
	flag.Parse()

	// Validation: Check if mandatory flags are provided
	if *domainName == "" || *tableName == "" {
		println("\n[WARNING] Eksekusi dibatalkan!")
		println("Anda harus menyertakan --domain-name dan --table-name")
		println("Contoh: go run .\\cmd\\scaffoldy\\. --domain-name Category --table-name tblcategory\n")
		os.Exit(1)
	}

	// Use the values from flags (dereferencing the pointers)
	ScaffoldHanlderComponents(*domainName)
	ScaffoldRepositoryComponents(*domainName, category.Category{}, *tableName)
	ScaffoldRequestComponents(*domainName, category.Category{})
	ScaffoldServiceComponents(*domainName, category.Category{})

	println("Successfully scaffolded domain:", *domainName)
}
