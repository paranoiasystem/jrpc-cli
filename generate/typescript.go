package generate

import (
	"encoding/json"
	"fmt"
	"jrpc-cli/model"
	"jrpc-cli/util"
	"strings"
)

// typeMap is a mapping from JSON types to TypeScript types.
var typeMap = map[string]string{
	"string":    "string",
	"number":    "number",
	"boolean":   "boolean",
	"null":      "null",
	"date-time": "Date",
}

// Typescript takes a JSON schema as input and generates corresponding TypeScript interfaces and types.
func Typescript(byteSchema []byte) (string, error) {
	var jrpcSchema model.JRPCSchema

	fmt.Println("Generating TypeScript code")

	// Unmarshal the JSON schema.
	if err := json.Unmarshal(byteSchema, &jrpcSchema); err != nil {
		return "", fmt.Errorf("error parsing JSON-RPC schema: %w", err)
	}

	interfacesTypes := ""
	interfaces := "export interface " + util.ToCamelCase(jrpcSchema.Info.Name, false) + " {\r\n"

	for _, method := range jrpcSchema.Methods {
		methodInterface, additionalTypes, err := generateMethodInterface(method)
		if err != nil {
			return "", err
		}
		interfaces += methodInterface
		interfacesTypes += additionalTypes
	}

	interfaces += "}"

	return interfacesTypes + "\r\n" + interfaces, nil
}

// generateMethodInterface generates the TypeScript interface for a method.
func generateMethodInterface(method model.JRPCSchemaMethod) (string, string, error) {
	methodName := "\t" + util.ToCamelCase(method.Name, false)
	var params []string
	var additionalTypes string

	for _, param := range method.Params {
		paramInterface, additionalParamTypes, err := generateParamInterface(param)
		if err != nil {
			return "", "", err
		}
		params = append(params, paramInterface)
		additionalTypes += additionalParamTypes
	}

	tempGeneratedResponseTypes := generateTypesFromSchema(method.Result.Name+"Response", method.Result.Schema)
	additionalTypes += strings.Join(tempGeneratedResponseTypes, "\r\n")

	methodSignature := "(" + strings.Join(params, ", ") + "): Promise<" + util.ToCamelCase(method.Result.Name+"Response", true) + ">;\r\n"
	methodDescription := "\r\n\t/**\r\n\t* " + method.Description + "\r\n\t*/\r\n"
	methodInterface := methodDescription + methodName + methodSignature

	return methodInterface, additionalTypes, nil
}

func generateParamInterface(param model.JRPCSchemaMethodParam) (string, string, error) {
	singleType, _ := param.Schema.Type.(string)
	if singleType == "" {
		return "", "", fmt.Errorf("seems the schema is not properly defined for param %s", param.Name)
	}
	if singleType == "object" {
		tempGeneratedInputTypes := generateTypesFromSchema(param.Name, param.Schema)
		interfacesTypes := strings.Join(tempGeneratedInputTypes, "")
		return param.Name + ": " + util.ToCamelCase(param.Name, true), interfacesTypes, nil
	} else if singleType == "array" {
		// Get the type of items in the array.
		itemType, ok := param.Schema.Items.Type.(string)
		if !ok {
			return "", "", fmt.Errorf("unsupported item type in array: %v", param.Schema.Items.Type)
		}

		if itemType == "object" {
			// If items are objects, generate types for them.
			tempGeneratedInputTypes := generateTypesFromSchema(param.Name, *param.Schema.Items)
			interfacesTypes := strings.Join(tempGeneratedInputTypes, "")
			return param.Name + ": " + util.ToCamelCase(param.Name, true) + "[]", interfacesTypes, nil
		} else {
			// Otherwise, use the basic TypeScript type.
			return param.Name + ": " + typeMap[itemType] + "[]", "", nil
		}
	} else {
		return param.Name + ": " + typeMap[singleType], "", nil
	}
}

// generateTypesFromSchema generates TypeScript types from a schema.
func generateTypesFromSchema(name string, schema model.JsonSchema) []string {
	var typesToReturn []string
	singleType, _ := schema.Type.(string)
	if singleType == "object" {
		tempType := "export type " + util.ToCamelCase(name, true) + " = {\r\n"
		for key, value := range schema.Properties {
			innerSingleType, _ := value.Type.(string)
			if innerSingleType == "object" {
				typesToReturn = append(typesToReturn, generateTypesFromSchema(key, value)...)
				tempType += "\t" + key + ": " + util.ToCamelCase(key, true) + "\r\n"
			} else if innerSingleType == "array" {
				if value.Items.Type.(string) == "object" {
					typesToReturn = append(typesToReturn, generateTypesFromSchema(key, *value.Items)...)
				}
				tempType += "\t" + key + ": Array<" + util.ToCamelCase(key, true) + ">\r\n"
			} else if innerSingleType != "" {
				tempType += "\t" + key + ": " + typeMap[innerSingleType] + "\r\n"
			} else {
				var tempTypes []string
				for _, innerType := range value.Type.([]interface{}) {
					tempTypes = append(tempTypes, typeMap[innerType.(string)])
				}
				tempType += "\t" + key + ": " + strings.Join(tempTypes, " | ") + "\r\n"
			}
		}
		tempType += "}\r\n"
		typesToReturn = append(typesToReturn, tempType)
	} else if schema.Type == "array" {
		if schema.Items.Type.(string) == "object" {
			typesToReturn = append(typesToReturn, generateTypesFromSchema(name, *schema.Items)...)

		} else {
			typesToReturn = append(typesToReturn, "export type "+util.ToCamelCase(name, true)+" = "+typeMap[singleType]+"[]\r\n")
		}
	} else {
		typesToReturn = append(typesToReturn, "export type "+util.ToCamelCase(name, true)+" = "+typeMap[singleType]+"\r\n")
	}
	return typesToReturn
}
