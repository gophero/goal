package twitter

import (
	"github.com/gophero/goal/conv"
	"github.com/gophero/goal/stringx"
)

var GetParamOptions getParamOptions

type GetParam struct {
	builder *stringx.Builder
}

func NewGetParam() *GetParam {
	return &GetParam{builder: stringx.NewBuilder()}
}

func (sp *GetParam) Param() string {
	return stringx.Trimright(sp.builder.String(), stringx.Ampersand)
}

func (sp *GetParam) Append(k string, v ...string) *GetParam {
	for _, s := range v {
		sp.builder.WriteString(k).WriteString("=").WriteString(s).WriteString(stringx.Ampersand)
	}
	return sp
}

func (sp *GetParam) FilterFields(ff *FieldFilter) *GetParam {
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

type GetParamOption func(p *GetParam)

type getParamOptions struct {
}

func (o getParamOptions) MaxResults(maxResults uint32) GetParamOption {
	if maxResults < 1 || maxResults > 1000 {
		maxResults = 100
	}
	return func(p *GetParam) {
		p.Append("max_results", conv.Uint32ToStr(maxResults))
	}
}

func (o getParamOptions) PaginationToken(pt string) GetParamOption {
	return func(p *GetParam) {
		p.Append("pagination_token", pt)
	}
}
