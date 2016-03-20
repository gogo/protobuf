// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/nourish/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*

The validator plugin generates a Validate method for each message.
By default, if none of the message's fields are annotated with the gogo validator annotation, it returns a nil error.
In case some of the fields are annotated, the Validate function returns nil upon sucessful validation, or an error
describing why the validation failed.
The Validate method is called recursively for all submessage of the message.

TODO(michal): ADD COMMENTS.

Equal is enabled using the following extensions:

  - equal
  - equal_all

While VerboseEqual is enable dusing the following extensions:

  - verbose_equal
  - verbose_equal_all

The equal plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

Let us look at:

  github.com/nourish/protobuf/test/example/example.proto

Btw all the output can be seen at:

  github.com/nourish/protobuf/test/example/*

The following message:



given to the equal plugin, will generate the following code:



and the following test code:


*/
package binding

import (
	"strconv"
	"strings"

	"github.com/nourish/protobuf/gogoproto"
	descriptor "github.com/nourish/protobuf/protoc-gen-gogo/descriptor"
	"github.com/nourish/protobuf/protoc-gen-gogo/generator"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg   generator.Single
	protoPkg   generator.Single
	bindingPkg generator.Single
	fmtPkg     generator.Single
	timePkg    generator.Single
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "binding"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.bindingPkg = p.NewImport("nourish/http/binding")
	p.timePkg = p.NewImport("time")
	p.fmtPkg = p.NewImport("fmt")
	p.regexPkg = p.NewImport("regexp")

	if !gogoproto.IsProto3(file.FileDescriptorProto) {
		return
	}

	if !gogoproto.BindingEnabledAll(file.FileDescriptorProto) {
		return
	}

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		p.generateMessage(file, msg)
	}
}

func (p *plugin) generateMessage(file *generator.FileDescriptor, message *generator.Descriptor) {
	if !gogoproto.BindingEnabled(file.FileDescriptorProto, message.DescriptorProto) {
		return
	}
	p.P(``)
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	shouldValidate := false
	shouldGenerateCreatedAt := false
	p.P(`func (this *`, ccTypeName, `) FieldMap() `+p.bindingPkg.Use()+`.FieldMap {`)
	p.In()
	p.P(`return ` + p.bindingPkg.Use() + `.FieldMap{`)
	p.In()
	for _, field := range message.Field {
		binding := gogoproto.GetBinding(field)
		fieldName := p.GetFieldName(message, field)
		//formName := fieldName
		formName := *field.Name
		if fieldName == "CreatedAt" {
			shouldGenerateCreatedAt = true
		}
		if binding != nil && binding.Ignore != nil && *binding.Ignore {
			continue
		}
		requiredStr := ""
		if binding != nil && binding.Required != nil && *binding.Required {
			requiredStr = ", Required: true"
		}
		if shouldValidateField(binding) {
			shouldValidate = true
		}
		if binding != nil && binding.CustomBinder != nil {
			p.P(`&this.` + fieldName + `: ` + p.bindingPkg.Use() + `.Field{`)
			p.In()
			p.P(`Form: "`, formName, `",`)
			if requiredStr != "" {
				p.P(`Required: true,`)
			}
			p.P(`Binder: `, binding.GetCustomBinder(), `,`)
			p.Out()
			p.P(`},`)
		} else {
			p.P(`&this.` + fieldName + `: ` + p.bindingPkg.Use() + `.Field{Form: "` + formName + `"` + requiredStr + `},`)
		}
	}
	p.Out()
	p.P(`}`)
	p.Out()
	p.P(`}`)
	if shouldValidate && gogoproto.ValidationEnabled(file.FileDescriptorProto, message.DescriptorProto) {
		// for _, field := range message.Field {
		// 	if !shouldValidate {
		// 		continue
		// 	}

		// }
		p.P(``)
		p.generateValidationMessage(file, message)
	}

	if shouldGenerateCreatedAt {
		p.P(``)
		p.generateCreateHook(message, shouldGenerateCreatedAt)
	}

}

func (p *plugin) generateCreateHook(message *generator.Descriptor, createdAt bool) {
	// ccTypeName := generator.CamelCaseSlice(message.TypeName())
	// p.P(`func (this *`, ccTypeName, `) BeforeCreate() (err error) {`)
	// p.In()
	// if createdAt {
	// 	p.P(`this.CreatedAt = NewTimestamp(`, p.timePkg.Use(), `.Now())`)
	// 	p.P(`return nil`)
	// }
	// p.Out()
	// p.P(`}`)
}

func (p *plugin) generateValidationMessage(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Validate(errs `+p.bindingPkg.Use()+`.Errors) `+p.bindingPkg.Use()+`.Errors {`)
	p.In()
	for _, field := range message.Field {
		json := gogoproto.GetJsonTag(field)
		binding := gogoproto.GetBinding(field)
		if !shouldValidateField(binding) {
			continue
		}
		fieldName := p.GetFieldName(message, field)
		formName := *field.Name
		if json != nil {
			jsonName := strings.Split(*json, ",")[0]
			if jsonName != "" {
				formName = jsonName
			}
		}
		variableName := "this.Get" + fieldName + "()"
		repeated := field.IsRepeated()
		if repeated {
			p.P(`for _, item := range `, variableName, `{`)
			p.In()
			variableName = "item"
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, formName, binding, field)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, formName, binding)
		} else if field.IsMessage() {

			// if binding != nil && binding.Required != nil && *binding.Required {
			// 	if !repeated {
			// 		p.P(`if nil == `, variableName, `{`)
			// 		p.In()
			// 		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:RequiredField", "`, formName, ` is required")`)
			// 		p.Out()
			// 		p.P(`}`)
			// 	} else {
			// 		p.P(`if len(`, variableName, `) <= 0 {`)
			// 		p.In()
			// 		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:RequiredRepeatedField", "At least one `, formName, ` is required")`)
			// 		p.Out()
			// 		p.P(`}`)
			// 	}
			// }
			// if nullable {
			// 	p.P(`if `, variableName, ` != nil {`)
			// 	p.In()
			// } else {
			// 	// non-nullable fields in proto3 store actual structs, we need pointers to operate on interaces
			// 	variableName = "&(" + variableName + ")"
			// }
			// p.P(`if err := proto.CallValidatorIfExists(`, variableName, `); err != nil {`)
			// p.In()
			// p.P(`return err`)
			// p.Out()
			// p.P(`}`)
			// if nullable {
			// 	p.Out()
			// 	p.P(`}`)
			// }
		}
		if repeated {
			// end the repeated loop
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return errs`)
	p.Out()
	p.P(`}`)
}

func shouldValidateField(binding *gogoproto.FieldBinding) bool {
	shouldValidate := binding != nil && !binding.GetIgnore() && (binding.GetRequired() || binding.GetMaxLen() > 0 || binding.GetIntLt() > 0 || binding.GetValidRegex() != "")
	return shouldValidate
}

func (p *plugin) generateIntValidator(variableName string, ccTypeName string, fieldName string, formName string, binding *gogoproto.FieldBinding) {
	//fieldIdentifier := ccTypeName + "." + fieldName
	fieldIdentifier := fieldName
	customMsg := binding.GetInvalidMessage()
	if binding.IntGte != nil {
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must greater than or equal to ` + strconv.Itoa(int(binding.GetIntGte())) + `.`
		}
		p.P(`if !(`, variableName, ` >= `, binding.IntGte, `){`)
		p.In()
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:IntGte", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
	if binding.IntGt != nil {
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must greater than ` + strconv.Itoa(int(binding.GetIntGt())) + `.`
		}
		p.P(`if !(`, variableName, ` > `, binding.IntGt, `){`)
		p.In()
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:IntGte", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
	if binding.IntLt != nil {
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must less than ` + strconv.Itoa(int(binding.GetIntLt())) + `.`
		}
		p.P(`if !(`, variableName, ` < `, binding.IntLt, `){`)
		p.In()
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:IntGte", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
	if binding.IntLte != nil {
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must less than or equal to ` + strconv.Itoa(int(binding.GetIntLte())) + `.`
		}
		p.P(`if !(`, variableName, ` <= `, binding.IntLt, `){`)
		p.In()
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:IntGte", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) isSupportedInt(field *descriptor.FieldDescriptorProto) bool {
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_INT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_UINT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	}
	return false
}

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {

		binding := gogoproto.GetBinding(field)
		if binding != nil && binding.ValidRegex != nil {
			fieldName := p.GetFieldName(message, field)
			p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile("`, binding.ValidRegex, `")`)
		}
	}
}

func (p *plugin) generateStringValidator(variableName string, ccTypeName string, fieldName string, formName string, binding *gogoproto.FieldBinding, field *descriptor.FieldDescriptorProto) {
	//fieldIdentifier := ccTypeName + "." + formName
	fieldIdentifier := fieldName
	customMsg := binding.GetInvalidMessage()
	if binding.MinLen != nil {
		p.P(`if len(`, variableName, `) < `, strconv.Itoa(int(binding.GetMinLen())), ` {`)
		p.In()
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must be at least ` + strconv.Itoa(int(binding.GetMinLen())) + ` characters long.`
		}
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:MinLen", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
	if binding.MaxLen != nil {
		p.P(`if len(`, variableName, `) > `, strconv.Itoa(int(binding.GetMaxLen())), ` {`)
		p.In()
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must be no more than ` + strconv.Itoa(int(binding.GetMaxLen())) + ` characters long.`
		}
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:MaxLen", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
	if binding.ValidRegex != nil {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(this.Get`, fieldName, `()) {`)
		p.In()
		msg := customMsg
		if msg == "" {
			msg = fieldIdentifier + ` must conform to regex '` + binding.GetValidRegex() + `'`
		}
		p.P(`errs.Add([]string{"`, formName, `"}, "ValidationError:Regexp", "`, msg, `")`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
