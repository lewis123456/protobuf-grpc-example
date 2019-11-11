package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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
	v.gen.Reset()
	msgs := getValidatorMsg(file.MessageType)
	for _, msg := range file.MessageType {
		if file.GetSyntax() == "proto3" {
			v.generateValidator(file, msg, "", msgs)
		}
	}
}
func (v *validator) GenerateImports(file *generator.FileDescriptor) {}

func getValidatorMsg(desc []*descriptor.DescriptorProto) []string {
	var msgs []string
	for _, msg := range desc {
		if hasValidatorOption(msg) {
			msgs = append(msgs, *msg.Name)
		}
		if len(msg.NestedType) != 0 {
			msgs = append(msgs, getValidatorMsg(msg.NestedType)...)
		}
	}
	return msgs
}

func hasValidatorOption(msg *descriptor.DescriptorProto) bool {
	for _, field := range msg.Field {
		if field.Options != nil {
			return true
		}
	}
	return false
}

func (v *validator) P(str ...interface{}) {
	v.gen.P(str...)
}

func (v *validator) generateValidator(file *generator.FileDescriptor, msg *descriptor.DescriptorProto, prefix string, msgs []string) {
	if !hasValidatorOption(msg) {
		return
	}

}