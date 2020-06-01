package main

import (
	"fmt"
	"os"
	"encoding/json"
	"flag"

	// https://github.com/hashicorp/terraform/tree/v0.11.4
	// go mod edit -require github.com/hashicorp/terraform@v0.11.4
	terraformV11 "github.com/hashicorp/terraform/terraform"
)

type output map[string]interface{}

func read(tfplan string) (string, error) {
	f, err := os.Open(tfplan)
	if err != nil {
		//
		return "", err
	}
	defer f.Close()

	// https://github.com/hashicorp/terraform/blob/v0.11.4/terraform/plan.go
	plan, err := terraformV11.ReadPlan(f)
	if err != nil {
		//
		return "", err
	}
	diff := output{}
	for _, v := range plan.Diff.Modules {
		convertModuleDiff(diff, v)
	}

	j, err := json.MarshalIndent(diff, "", "    ")
	if err != nil {
		//
		return "", err
	}

	return string(j), nil
}

func insert(out output, path []string, key string, value interface{}) {
	if len(path) > 0 && path[0] == "root" {
		path = path[1:]
	}
	for _, elem := range path {
		switch nested := out[elem].(type) {
		case output:
			out = nested
		default:
			new := output{}
			out[elem] = new
			out = new
		}
	}
	out[key] = value
}

func convertModuleDiff(out output, diff *terraformV11.ModuleDiff) {
	insert(out, diff.Path, "destroy", diff.Destroy)
	for k, v := range diff.Resources {
		convertInstanceDiff(out, append(diff.Path, k), v)
	}
}

func convertInstanceDiff(out output, path []string, diff *terraformV11.InstanceDiff) {
	insert(out, path, "destroy", diff.Destroy)
	insert(out, path, "destroy_tainted", diff.DestroyTainted)
	for k, v := range diff.Attributes {
		insert(out, path, k, v.New)
	}
}

func main() {

	var plan string
	var out string
	flag.StringVar(&plan, "tfplan", "terraform.tfplan", "a string var")
	flag.StringVar(&out, "output", "console", "console or json")

	flag.Parse()

	json, err := read(plan)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if out == "console"{
		fmt.Println(json)
	} else {
		jsonFile, err := os.Create("tfplan.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		jsonFile.Write([]byte(json))
	}

}