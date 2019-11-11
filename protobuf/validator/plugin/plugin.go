package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init()  {
	generator.RegisterPlugin(new(validator))
}

type validator struct {
	gen *generator.Generator
}

func (v *validator) Name() string {
	return "go-validator"
}

func (v *validator) Init(g *generator.Generator) {
	v.gen = g
}
func (v *validator) Generate(file *generator.FileDescriptor) {

}
func (v *validator) GenerateImports(file *generator.FileDescriptor) {}

func GetValidatorMsg(desc []*generator.Descriptor) []string {

}

func HasValidatorOption(msg *generator.Descriptor) bool {

}