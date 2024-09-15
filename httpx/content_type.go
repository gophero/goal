package httpx

// http 支持的 ContentType 类型定义
const (
	ContentTypeAll                        ContentType = "*/*"
	ContentTypeApplicationAtomXml         ContentType = "application/atom+xml"
	ContentTypeApplicationCbor            ContentType = "application/cbor"
	ContentTypeApplicationFormUrlencoded  ContentType = "application/x-www-form-urlencoded"
	ContentTypeApplicationJson            ContentType = "application/json"
	ContentTypeApplicationJsonUtf8        ContentType = "application/json;charset=UTF-8"
	ContentTypeApplicationOctetStream     ContentType = "application/octet-stream"
	ContentTypeApplicationPdf             ContentType = "application/pdf"
	ContentTypeApplicationProblemJson     ContentType = "application/problem+json"
	ContentTypeApplicationProblemJsonUtf8 ContentType = "application/problem+json;charset=UTF-8"
	ContentTypeApplicationProblemXml      ContentType = "application/problem+xml"
	ContentTypeApplicationRssXml          ContentType = "application/rss+xml"
	ContentTypeApplicationStreamJson      ContentType = "application/stream+json"
	ContentTypeApplicationXhtmlXml        ContentType = "application/xhtml+xml"
	ContentTypeApplicationXml             ContentType = "application/xml"
	ContentTypeImageGif                   ContentType = "image/gif"
	ContentTypeImageJpeg                  ContentType = "image/jpeg"
	ContentTypeImagePng                   ContentType = "image/png"
	ContentTypeMultipartFormData          ContentType = "multipart/form-data"
	ContentTypeMultipartMixed             ContentType = "multipart/mixed"
	ContentTypeMultipartRelated           ContentType = "multipart/related"
	ContentTypeTextEventStream            ContentType = "text/event-stream"
	ContentTypeTextHtml                   ContentType = "text/html"
	ContentTypeTextMarkdown               ContentType = "text/markdown"
	ContentTypeTextPlain                  ContentType = "text/plain"
	ContentTypeTextXml                    ContentType = "text/xml"
)
