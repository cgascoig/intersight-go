package main

import (
	"fmt"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func stripEnvelope(result interface{}) interface{} {
	vr := reflect.ValueOf(result)
	if vr.Kind() == reflect.Ptr {
		vr = vr.Elem()
	}

	if vr.Kind() == reflect.Struct {
		payloadValue := vr.FieldByName("Payload")
		if payloadValue.IsValid() {
			resultsValue := reflect.Indirect(payloadValue).FieldByName("Results")
			if resultsValue.IsValid() {
				return resultsValue.Interface()
			}
			return payloadValue.Interface()
		}
	}

	return result
}

func indent(lines []string, slice bool) []string {
	ret := []string{}

	first := true
	for _, line := range lines {
		if slice && first {
			first = false
			ret = append(ret, fmt.Sprintf("- %s", line))
		} else {
			ret = append(ret, fmt.Sprintf("  %s	", line))
		}
	}

	return ret
}

func printResultHumanValue(r reflect.Value) []string {
	vr := reflect.Indirect(r)
	ret := []string{}

	switch vr.Kind() {
	case reflect.Struct:
		for i := 0; i < vr.NumField(); i++ {
			f := vr.Field(i)
			t := vr.Type().Field(i)
			if t.PkgPath == "" { // empty PkgPath indicates exported field
				newLines := printResultHumanValue(f)
				if !t.Anonymous {
					if len(newLines) <= 1 {
						ret = append(ret, fmt.Sprintf("%s: %s", t.Name, strings.Join(newLines, "")))
					} else {
						ret = append(ret, fmt.Sprintf("%s: ", t.Name))
						ret = append(ret, indent(newLines, false)...)
					}
				} else {
					ret = append(ret, newLines...)
				}
			}
		}
	case reflect.Slice:
		for i := 0; i < vr.Len(); i++ {
			newLines := printResultHumanValue(vr.Index(i))

			ret = append(ret, indent(newLines, true)...)
		}
	default:
		ret = []string{fmt.Sprintf("%s", vr.String())}
	}

	return ret
}

func printResultHuman(result interface{}) {

	for _, line := range printResultHumanValue(reflect.ValueOf(result)) {
		fmt.Println(line)
	}
}

func printResultYAML(result interface{}) {
	buf, err := yaml.Marshal(result)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Println(string(buf))
}
