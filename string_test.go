package godash

import (
	"gotest.tools/assert"
	"testing"
)

func TestCamelCase(t *testing.T) {
	str := "Hello world"
	str, _ = CamelCase(str)

	assert.Equal(t, str, "helloWorld")

	str, _ = CamelCase("Hello World from Ginkgo")
	assert.Equal(t, str, "helloWorldFromGinkgo")

	str, _ = CamelCase("001hello from gin")
	assert.Equal(t, str, "001HelloFromGin")

	str, _ = CamelCase("0aa1bb2cc")
	assert.Equal(t, str, "0Aa1Bb2Cc")

	str, _ = CamelCase("aa-bb-cc")
	assert.Equal(t, str, "aa-Bb-Cc")
}

func TestCamelCaseWithInit(t *testing.T) {
	str := "Hello world"
	str, _ = CamelCaseWithInit(str, true)

	assert.Equal(t, str, "HelloWorld")

	str, _ = CamelCaseWithInit("Hello World from Ginkgo", true)
	assert.Equal(t, str, "HelloWorldFromGinkgo")

	str, _ = CamelCaseWithInit("001hello from gin", true)
	assert.Equal(t, str, "001HelloFromGin")

	str, _ = CamelCaseWithInit("0aa1bb2cc", true)
	assert.Equal(t, str, "0Aa1Bb2Cc")

	str, _ = CamelCaseWithInit("aa-bb-cc", true)
	assert.Equal(t, str, "Aa-Bb-Cc")
}

func TestCapitalize(t *testing.T) {
	str := "jerry"
	str = Capitalize(str)

	assert.Equal(t, str, "Jerry")
}

func TestEndsWith(t *testing.T) {
	str := "Hello World"
	result := EndsWith(str, "d")
	assert.Equal(t, result, true)

	result = EndsWith(str, "l")
	assert.Equal(t, result, false)

	result = EndsWithByLength(str, "d", 1)
	assert.Equal(t, result, true)

	result = EndsWithByLength(str, "l", 1)
	assert.Equal(t, result, false)

	result = EndsWithByLength(str, "l", 2)
	assert.Equal(t, result, true)
}

func TestEscape(t *testing.T) {
	str := "<>"
	str = Escape(str)
	assert.Equal(t, str, "&lt;&gt;")
}

func TestEscapeRegExp(t *testing.T) {
	str := `Escaping symbols like: .+*?()|[]{}^$`
	str = EscapeRegExp(str)

	assert.Equal(t, str, `Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$`)
}
