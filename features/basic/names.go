package basic

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func init() {
	// results in "Api" using ToCamel("API")
	strcase.ConfigureAcronym("API", "api")

	// results in "Sql" using ToCamel("SQL")
	strcase.ConfigureAcronym("SQL", "sql")
}

func toEnumValueName(valueName string) string {
	return strcase.ToCamel(strings.ToLower(valueName))
}
