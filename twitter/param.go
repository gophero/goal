package twitter

import "github.com/gophero/goal/stringx"

type FormParam struct {
	builder *stringx.Builder
}

func NewFormParam() *FormParam {
	return &FormParam{builder: stringx.NewBuilder()}
}

func (sp *FormParam) Param() string {
	return stringx.Trimright(sp.builder.String(), stringx.Ampersand)
}

func (sp *FormParam) Append(k string, v ...string) *FormParam {
	for _, s := range v {
		sp.builder.WriteString(k).WriteString("=").WriteString(s).WriteString(stringx.Ampersand)
	}
	return sp
}

func (sp *FormParam) FilterFields(ff *FieldFilter) *FormParam {
	if ff != nil {
		if len(ff.Expansions) > 0 {
			sp.Append("expansions", formatExpansion(ff.Expansions...))
		}
		if len(ff.UserFields) > 0 {
			sp.Append("user.fields", formatUserFields(ff.UserFields...))
		}
		if len(ff.TwitterFields) > 0 {
			sp.Append("tweet.fields", formatTweetFields(ff.TwitterFields...))
		}
	}
	return sp
}

type FormParamOption func(p *FormParam)
