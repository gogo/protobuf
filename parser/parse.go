package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"os/exec"
	"strings"
)

func ParseFile(filename string, paths ...string) (*descriptor.FileDescriptorSet, error) {
	return parseFile(filename, false, true, paths...)
}

func parseFile(filename string, includeSourceInfo bool, includeImports bool, paths ...string) (*descriptor.FileDescriptorSet, error) {
	args := []string{"--proto_path=" + strings.Join(paths, ":")}
	if includeSourceInfo {
		args = append(args, "--include_source_info")
	}
	if includeImports {
		args = append(args, "--include_imports")
	}
	args = append(args, "--descriptor_set_out=/dev/stdout")
	args = append(args, filename)
	cmd := exec.Command("protoc", args...)
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	fileDesc := &descriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fileDesc); err != nil {
		return nil, err
	}
	return fileDesc, nil
}
