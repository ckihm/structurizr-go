package main

import (
	"os"
	"strings"
	"text/template"
	"unicode"
)

type SoftwareSystem struct {
	Name        string
	Description string
	Tags        []string
}

type DeploymentEnvironment struct {
	Name            string
	DeploymentNodes []DeploymentNode
}

type DeploymentNode struct {
	Name                string
	Description         string
	Technology          string
	SoftwareSystemNames []string
}

func main() {
	softwareSystems := map[string]SoftwareSystem{
		"System 1": {
			Name:        "System 1",
			Description: "This is system 1",
			Tags:        []string{"web", "backend"},
		},
		"System 2": {
			Name:        "System 2",
			Description: "This is system 2",
			Tags:        []string{"worker", "backend"},
		},
	}

	deploymentEnvironments := []DeploymentEnvironment{
		{
			Name: "Production",
			DeploymentNodes: []DeploymentNode{
				{
					Name:                "Web Server",
					Description:         "Serves the system 1 web application",
					Technology:          "Apache Tomcat",
					SoftwareSystemNames: []string{"System1"},
				},
				{
					Name:                "Worker",
					Description:         "Processes background jobs",
					Technology:          "Celery",
					SoftwareSystemNames: []string{"System2", "System1"},
				},
			},
		},
	}

	// Define a function to get a software system by name from the softwareSystems map
	getSoftwareSystem := func(name string) SoftwareSystem {
		softwareSystem := softwareSystems[name]
		return softwareSystem
	}

	replaceAllWhitespace := func(str string) string {
		firstChar := []rune(str)[0]                  // get the first character as a rune
		lowerFirstChar := unicode.ToLower(firstChar) // convert the first character to lowercase
		rest := str[1:]                              // get the rest of the string
		result := string(lowerFirstChar) + rest      // combine the modified first character and the rest of the string
		return strings.ReplaceAll(result, " ", "")   // remove all whitespaces
	}

	// Define the template
	tpl := `workspace {
	model {
		{{range $name, $softwareSystem := .SoftwareSystems}}
		{{$name | replaceAllWhitespace}} = softwareSystem "{{$name}}" {
			description "{{$softwareSystem.Description}}"
			tags {{range $softwareSystem.Tags}} "{{.}}" {{end}}
		}
		{{end}}
		{{range .DeploymentEnvironments}}
		{{.Name | replaceAllWhitespace}} = deploymentEnvironment "{{.Name}}" {
			{{range .DeploymentNodes}}
			deploymentNode "{{.Name}}" {
				description "{{.Description}}"
				technology "{{.Technology}}"
				{{range .SoftwareSystemNames}}
				softwareSystemInstance "{{.}}" {
				}
				{{end}}
			}
			{{end}}
		}
	  	{{end}}
	}

	views {
        deployment * production {
            include *
            autoLayout lr
        }
    }
}`

	// Create a new template and parse the template string
	t := template.New("structurizr").Funcs(template.FuncMap{
		"getSoftwareSystem":    getSoftwareSystem,
		"replaceAllWhitespace": replaceAllWhitespace,
	})

	t, err := t.Parse(tpl)
	if err != nil {
		panic(err)
	}

	// Render the template to a file
	f, err := os.Create("workspace.dsl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Execute the template with the data
	err = t.Execute(f, struct {
		SoftwareSystems        map[string]SoftwareSystem
		DeploymentEnvironments []DeploymentEnvironment
	}{
		SoftwareSystems:        softwareSystems,
		DeploymentEnvironments: deploymentEnvironments,
	})

	if err != nil {
		panic(err)
	}
}
