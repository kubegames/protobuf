func New{{$.Name}}Error(code int, format string, a ...interface{}) error {
	return status.Error(codes.Code(code), fmt.Sprintf(format, a...))
}

type {{ $.InterfaceName }} interface {
	{{range .MethodSet}}
		{{.Name}}(ctx context.Context,request *{{.Request}}) (response *{{.Reply}},err error)
	{{end}}
}

func Register{{ $.InterfaceName }}(r gin.IRouter, srv {{ $.InterfaceName }}) {
	//init router
	s := _{{.Name}}{
		server: srv,
		router:     r,
	}
	s._RegisterService()
}

type _{{$.Name}} struct{
	server {{ $.InterfaceName }}
	router gin.IRouter
}

{{range .Methods}}
func (s *_{{$.Name}}) {{ .HandlerName }} (ctx *gin.Context) {
	var in {{.Request}}
{{if .HasPathParams }}
	if err := ctx.ShouldBindUri(&in); err != nil {
		_{{$.Name}}ParamsError(ctx, err)
		return
	}
{{end}}
{{if eq .Method "GET" "DELETE" }}
	if err := ctx.ShouldBindQuery(&in); err != nil {
		_{{$.Name}}ParamsError(ctx, err)
		return
	}
{{else if eq .Method "POST" "PUT" }}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		_{{$.Name}}ParamsError(ctx, err)
		return
	}
{{else}}
	if err := ctx.ShouldBind(&in); err != nil {
		_{{$.Name}}ParamsError(ctx, err)
		return
	}
{{end}}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.({{ $.InterfaceName }}).{{.Name}}(newCtx, &in)
	if err != nil {
		_{{$.Name}}Error(ctx, err)
		return
	}

	_{{$.Name}}Success(ctx, out)
}
{{end}}

func (s *_{{$.Name}}) _RegisterService() {
{{range .Methods}}
		s.router.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}

func _{{$.Name}}Error(ctx *gin.Context, err error) {
	code := 500

	msg := "unknow error"
	if err == nil {
		msg += ", err is nil"
		ctx.JSON(code, map[string]interface{}{
			"code": code,
			"msg":  msg,
			"data": nil,
		})
		return
	}

	if c, ok := status.FromError(err); ok {
		code = int(c.Code())
		msg = c.Message()
	}

	//_ = ctx.Error(err)

	ctx.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

func _{{$.Name}}ParamsError(ctx *gin.Context, err error) {
	ctx.JSON(400, map[string]interface{}{
		"code": 400,
		"msg":  err.Error(),
		"data": nil,
	})
}

func _{{$.Name}}Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

