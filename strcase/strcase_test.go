package strcase_test

import (
	"testing"

	"github.com/kaloseia/go-util/strcase"
	"github.com/stretchr/testify/suite"
)

type StringCaseTestSuite struct {
	suite.Suite
}

func TestStringCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StringCaseTestSuite))
}

func (suite *StringCaseTestSuite) TestToSnakeCase() {
	suite.Equal(strcase.ToSnakeCase("UpperCamelCase"), "Upper_Camel_Case")
	suite.Equal(strcase.ToSnakeCase("lowerCamelCase"), "lower_Camel_Case")

	suite.Equal(strcase.ToSnakeCase("Upper-Kebab-Case"), "Upper_Kebab_Case")
	suite.Equal(strcase.ToSnakeCase("lower-kebab-case"), "lower_kebab_case")

	suite.Equal(strcase.ToSnakeCase("Upper_Snake_Case"), "Upper_Snake_Case")
	suite.Equal(strcase.ToSnakeCase("lower_snake_case"), "lower_snake_case")

	suite.Equal(strcase.ToSnakeCase("CaseIDThing"), "Case_ID_Thing")
	suite.Equal(strcase.ToSnakeCase("CaseID"), "Case_ID")
	suite.Equal(strcase.ToSnakeCase("IDCase"), "ID_Case")

	suite.Equal(strcase.ToSnakeCase("CaseSQLThing"), "Case_SQL_Thing")
	suite.Equal(strcase.ToSnakeCase("CaseSQL"), "Case_SQL")
	suite.Equal(strcase.ToSnakeCase("SQLCase"), "SQL_Case")

	suite.Equal(strcase.ToSnakeCase("Single"), "Single")
	suite.Equal(strcase.ToSnakeCase("SINGLE"), "SINGLE")
	suite.Equal(strcase.ToSnakeCase("single"), "single")

	suite.Equal(strcase.ToSnakeCase("num01Word"), "num01_Word")
	suite.Equal(strcase.ToSnakeCase("01Word"), "01_Word")
	suite.Equal(strcase.ToSnakeCase("Word01"), "Word01")
}

func (suite *StringCaseTestSuite) TestToSnakeCaseLower() {
	suite.Equal(strcase.ToSnakeCaseLower("UpperCamelCase"), "upper_camel_case")
	suite.Equal(strcase.ToSnakeCaseLower("lowerCamelCase"), "lower_camel_case")

	suite.Equal(strcase.ToSnakeCaseLower("Upper-Kebab-Case"), "upper_kebab_case")
	suite.Equal(strcase.ToSnakeCaseLower("lower-kebab-case"), "lower_kebab_case")

	suite.Equal(strcase.ToSnakeCaseLower("Upper_Snake_Case"), "upper_snake_case")
	suite.Equal(strcase.ToSnakeCaseLower("lower_snake_case"), "lower_snake_case")

	suite.Equal(strcase.ToSnakeCaseLower("CaseIDThing"), "case_id_thing")
	suite.Equal(strcase.ToSnakeCaseLower("CaseID"), "case_id")
	suite.Equal(strcase.ToSnakeCaseLower("IDCase"), "id_case")

	suite.Equal(strcase.ToSnakeCaseLower("CaseSQLThing"), "case_sql_thing")
	suite.Equal(strcase.ToSnakeCaseLower("CaseSQL"), "case_sql")
	suite.Equal(strcase.ToSnakeCaseLower("SQLCase"), "sql_case")

	suite.Equal(strcase.ToSnakeCaseLower("Single"), "single")
	suite.Equal(strcase.ToSnakeCaseLower("SINGLE"), "single")
	suite.Equal(strcase.ToSnakeCaseLower("single"), "single")

	suite.Equal(strcase.ToSnakeCaseLower("num01Word"), "num01_word")
	suite.Equal(strcase.ToSnakeCaseLower("01Word"), "01_word")
	suite.Equal(strcase.ToSnakeCaseLower("Word01"), "word01")
}

func (suite *StringCaseTestSuite) TestToCamelCase() {
	suite.Equal(strcase.ToCamelCase("UpperCamelCase"), "upperCamelCase")
	suite.Equal(strcase.ToCamelCase("lowerCamelCase"), "lowerCamelCase")

	suite.Equal(strcase.ToCamelCase("Upper-Kebab-Case"), "upperKebabCase")
	suite.Equal(strcase.ToCamelCase("lower-kebab-case"), "lowerKebabCase")

	suite.Equal(strcase.ToCamelCase("Upper_Snake_Case"), "upperSnakeCase")
	suite.Equal(strcase.ToCamelCase("lower_snake_case"), "lowerSnakeCase")

	suite.Equal(strcase.ToCamelCase("Case_ID_Thing"), "caseIDThing")
	suite.Equal(strcase.ToCamelCase("Case_ID"), "caseID")
	suite.Equal(strcase.ToCamelCase("ID_Case"), "idCase")

	suite.Equal(strcase.ToCamelCase("Case_SQL_Thing"), "caseSQLThing")
	suite.Equal(strcase.ToCamelCase("Case_SQL"), "caseSQL")
	suite.Equal(strcase.ToCamelCase("SQL_Case"), "sqlCase")

	suite.Equal(strcase.ToCamelCase("Single"), "single")
	suite.Equal(strcase.ToCamelCase("SINGLE"), "single")
	suite.Equal(strcase.ToCamelCase("single"), "single")

	suite.Equal(strcase.ToCamelCase("num01_Word"), "num01Word")
	suite.Equal(strcase.ToCamelCase("01_Word"), "01Word")
	suite.Equal(strcase.ToCamelCase("Word01"), "word01")
}

func (suite *StringCaseTestSuite) TestToPascalCase() {
	suite.Equal(strcase.ToPascalCase("UpperCamelCase"), "UpperCamelCase")
	suite.Equal(strcase.ToPascalCase("lowerCamelCase"), "LowerCamelCase")

	suite.Equal(strcase.ToPascalCase("Upper-Kebab-Case"), "UpperKebabCase")
	suite.Equal(strcase.ToPascalCase("lower-kebab-case"), "LowerKebabCase")

	suite.Equal(strcase.ToPascalCase("Upper_Snake_Case"), "UpperSnakeCase")
	suite.Equal(strcase.ToPascalCase("lower_snake_case"), "LowerSnakeCase")

	suite.Equal(strcase.ToPascalCase("Case_ID_Thing"), "CaseIDThing")
	suite.Equal(strcase.ToPascalCase("Case_ID"), "CaseID")
	suite.Equal(strcase.ToPascalCase("ID_Case"), "IDCase")

	suite.Equal(strcase.ToPascalCase("Case_SQL_Thing"), "CaseSQLThing")
	suite.Equal(strcase.ToPascalCase("Case_SQL"), "CaseSQL")
	suite.Equal(strcase.ToPascalCase("SQL_Case"), "SQLCase")

	suite.Equal(strcase.ToPascalCase("Single"), "Single")
	suite.Equal(strcase.ToPascalCase("SINGLE"), "SINGLE")
	suite.Equal(strcase.ToPascalCase("single"), "Single")

	suite.Equal(strcase.ToPascalCase("num01_Word"), "Num01Word")
	suite.Equal(strcase.ToPascalCase("01_Word"), "01Word")
	suite.Equal(strcase.ToPascalCase("Word01"), "Word01")
}
