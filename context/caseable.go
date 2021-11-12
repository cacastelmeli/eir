package context

import "github.com/iancoleman/strcase"

type Caseable string

func (c Caseable) Snake() Caseable {
	return Caseable(strcase.ToSnake(string(c)))
}

func (c Caseable) LowerCamel() Caseable {
	return Caseable(strcase.ToLowerCamel(string(c)))
}

func (c Caseable) UpperCamel() Caseable {
	return Caseable(strcase.ToCamel(string(c)))
}
