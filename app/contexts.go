//************************************************************************//
// API "goa Swagger service": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/swagger-service/design
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// HealthAeContext provides the ae health action context.
type HealthAeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHealthAeContext parses the incoming request URL and body, performs validations and creates the
// context used by the ae controller health action.
func NewHealthAeContext(ctx context.Context, service *goa.Service) (*HealthAeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := HealthAeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *HealthAeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// StartAeContext provides the ae start action context.
type StartAeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewStartAeContext parses the incoming request URL and body, performs validations and creates the
// context used by the ae controller start action.
func NewStartAeContext(ctx context.Context, service *goa.Service) (*StartAeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := StartAeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *StartAeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// ShowSpecContext provides the spec show action context.
type ShowSpecContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Pkg string
}

// NewShowSpecContext parses the incoming request URL and body, performs validations and creates the
// context used by the spec controller show action.
func NewShowSpecContext(ctx context.Context, service *goa.Service) (*ShowSpecContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowSpecContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPkg := req.Params["pkg"]
	if len(paramPkg) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("pkg"))
	} else {
		rawPkg := paramPkg[0]
		rctx.Pkg = rawPkg
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSpecContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/swagger+json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// UnprocessableEntity sends a HTTP response with status code 422.
func (ctx *ShowSpecContext) UnprocessableEntity(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(422)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
