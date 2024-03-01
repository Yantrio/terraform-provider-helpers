package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = EchoFunction{}
)

func NewEchoFunction() function.Function {
	return EchoFunction{}
}

type EchoFunction struct{}

func (r EchoFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "echo"
}

func (r EchoFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Echo function",
		MarkdownDescription: "Echoes given argument as result",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "String to echo",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r EchoFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string
	err := req.Arguments.Get(ctx, &data)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	err = resp.Result.Set(ctx, data)
	resp.Error = function.ConcatFuncErrors(err)
}
