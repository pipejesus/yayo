package Lando

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"yayo/Helpers"
	"yayo/InputsCollector"
)

func UpdateLandoFile(inputsCollector *InputsCollector.Collector) {
	projectName, errProjectName := inputsCollector.GetAnswer("projectName")
	phpVersion, errPhpVersion := inputsCollector.GetAnswer("phpVersion")
	dbVersion, errDbVersion := inputsCollector.GetAnswer("dbVersion")

	if errProjectName != nil || errPhpVersion != nil || errDbVersion != nil {
		Helpers.CustomErrorPrefix().Printfln("Brakuje odpowiedzi na wszystkie pytania!")
		os.Exit(0)
	}

	landoFilePath := fmt.Sprintf("%s/.lando.yml", projectName)
	landoConfig, err := ReadLandoYaml(landoFilePath)

	if err != nil {
		Helpers.CustomErrorPrefix().Printfln(err.Error())
		os.Exit(0)
	}

	var landoStruct LandoYamlStructure
	err = yaml.Unmarshal(landoConfig, &landoStruct)

	landoStruct, err = SetYamlArrayItem(landoStruct, projectName, "name")
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln(err.Error())
		os.Exit(0)
	}

	landoStruct, err = SetYamlArrayItem(landoStruct, phpVersion, "config", "php")
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln(err.Error())
		os.Exit(0)
	}
	landoStruct, err = SetYamlArrayItem(landoStruct, dbVersion, "config", "database")
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln(err.Error())
		os.Exit(0)
	}

	landoStruct, err = SetYamlArrayItem(landoStruct, "projectName="+projectName, "services", "appserver", "overrides", "environment", "PHP_IDE_CONFIG")
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln(err.Error())
		os.Exit(0)
	}

	newLandoContents, err := yaml.Marshal(landoStruct)
	if err != nil {
		Helpers.CustomErrorPrefix().Printfln("Nie mogę zamienić na YAML!")
		os.Exit(0)
	}

	if err = WriteLandoYaml(landoFilePath, newLandoContents); err != nil {
		Helpers.CustomErrorPrefix().Printfln("Nie mogę zapisać do pliku .lando.yaml!")
		os.Exit(0)
	}

}
