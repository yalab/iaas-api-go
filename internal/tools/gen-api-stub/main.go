// Copyright 2022 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/iaas-api-go/internal/define"
	"github.com/sacloud/iaas-api-go/internal/dsl"
	"github.com/sacloud/iaas-api-go/internal/tools"
)

const destination = "stub/zz_api_stubs.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-stub: ")
}

func main() {
	dsl.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", filepath.Join(tools.ProjectRootPath(), destination))
}

const tmpl = `// generated by 'github.com/sacloud/iaas-api-go/internal/tools/gen-api-stub'; DO NOT EDIT

package stub

import (
{{- range .ImportStatements "context" "log" }}
	{{ . }}
{{- end }}
)

{{ range . }} {{ $typeName := .TypeName }}{{ $resource := . }}

/************************************************* 
* {{ $typeName }}Stub
*************************************************/

{{ range .Operations }}
// {{ $typeName }}{{.MethodName}}StubResult is expected values of the {{ .MethodName }} operation
type {{ $typeName }}{{.MethodName}}StubResult struct {
	{{ range .StubFieldDefines -}}
	{{ . }}
	{{ end -}}	
	Err error
}
{{ end -}}

// {{ $typeName }}Stub is for trace {{ $typeName }}Op operations
type {{ $typeName }}Stub struct {
{{ range .Operations -}}
	{{.MethodName}}StubResult *{{ $typeName }}{{.MethodName}}StubResult 
{{ end -}}
}

// New{{ $typeName}}Stub creates new {{ $typeName}}Stub instance
func New{{ $typeName}}Stub(caller iaas.APICaller) iaas.{{$typeName}}API {
	return &{{ $typeName}}Stub{}
}

{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call with trace log
func (s *{{ $typeName }}Stub) {{ .MethodName }}(ctx context.Context{{if not $resource.IsGlobal}}, zone string{{end}}{{ range .Arguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	if s.{{$operationName}}StubResult == nil {
		log.Fatal("{{$typeName}}Stub.{{$operationName}}StubResult is not set")
	}
	{{.StubReturnStatement "s"}}
}
{{- end -}}

{{ end }}
`
