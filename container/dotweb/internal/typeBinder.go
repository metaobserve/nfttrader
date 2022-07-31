package internal

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	web "github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/reflects"
	"github.com/pkg/errors"
	"strings"
)

type typeBinder struct {
}

//Bind decode res.Body or form-value to struct
func (binder *typeBinder) Bind(i interface{}, ctx web.Context) (err error) {
	fmt.Println("UserBind.Bind")
	req := ctx.Request()
	ctype := req.Header.Get(web.HeaderContentType)
	if req.Body == nil {
		err = errors.New("request body can't be empty")
		return err
	}
	err = errors.New("request unsupported MediaType -> " + ctype)
	switch {
	case strings.HasPrefix(ctype, web.MIMEApplicationJSON):
		err = json.Unmarshal(ctx.Request().PostBody(), i)
	case strings.HasPrefix(ctype, web.MIMEApplicationXML):
		err = xml.Unmarshal(ctx.Request().PostBody(), i)
		//case strings.HasPrefix(ctype, MIMEApplicationForm), strings.HasPrefix(ctype, MIMEMultipartForm),
		//	strings.HasPrefix(ctype, MIMETextHTML):
		//	err = reflects.ConvertMapToStruct(defaultTagName, i, ctx.FormValues())
	default:
		//check is use json tag, fixed for issue #91
		tagName := "form"
		if ctx.HttpServer().ServerConfig().EnabledBindUseJsonTag {
			tagName = "json"
		}
		//no check content type for fixed issue #6
		err = reflects.ConvertMapToStruct(tagName, i, ctx.Request().FormValues())
	}
	return err
}

//BindJsonBody default use json decode res.Body to struct
func (binder *typeBinder) BindJsonBody(i interface{}, ctx web.Context) (err error) {
	fmt.Println("UserBind.BindJsonBody")
	if ctx.Request().PostBody() == nil {
		err = errors.New("request body can't be empty")
		return err
	}
	err = json.Unmarshal(ctx.Request().PostBody(), i)
	return err
}

func NewTypeBinder() *typeBinder {
	return &typeBinder{}
}
