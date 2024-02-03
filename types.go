package vegagoja

// Param is a bound input parameter.
type Param struct {
	Name  string      `json:"name,omitempty" mapstructure:"name"`
	Bind  Bind        `json:"bind,omitempty" mapstructure:"bind"`
	Value interface{} `json:"value,omitempty" mapstructure:"value"`
}

// Valid returns true when v conforms to the parameter's bind specification.
func (p Param) Valid(v interface{}) bool {
	return p.Bind.Valid(v)
}

// Bind is the bound input parameter type definition.
type Bind struct {
	Input BindType `json:"input,omitempty" mapstructure:"input"`
	Min   int      `json:"min,omitempty" mapstructure:"min"`
	Max   int      `json:"max,omitempty" mapstructure:"max"`
	Step  int      `json:"step,omitempty" mapstructure:"step"`
}

// Valid returns true when v conforms to the specification.
func (bind Bind) Valid(v interface{}) bool {
	return false
}

// BindType are input types.
//
// See: https://vega.github.io/vega/docs/signals/#bind
type BindType string

// Bind types.
const (
	BindCheckbox BindType = "checkbox"
	BindRadio    BindType = "radio"
	BindSelect   BindType = "select"
)

// String satisfies the [fmt.Stringe] interface.
func (typ BindType) String() string {
	return string(typ)
}
