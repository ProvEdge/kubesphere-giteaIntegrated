// Code generated from specification version 5.6.15: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newMsearchFunc(t Transport) Msearch {
	return func(body io.Reader, o ...func(*MsearchRequest)) (*Response, error) {
		var r = MsearchRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// Msearch allows to execute several search operations in one request.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-multi-search.html.
//
type Msearch func(body io.Reader, o ...func(*MsearchRequest)) (*Response, error)

// MsearchRequest configures the Msearch API request.
//
type MsearchRequest struct {
	Index        []string
	DocumentType []string

	Body io.Reader

	MaxConcurrentSearches *int
	PreFilterShardSize    *int
	SearchType            string
	TypedKeys             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MsearchRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len(strings.Join(r.Index, ",")) + 1 + len(strings.Join(r.DocumentType, ",")) + 1 + len("_msearch"))
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
	}
	if len(r.DocumentType) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.DocumentType, ","))
	}
	path.WriteString("/")
	path.WriteString("_msearch")

	params = make(map[string]string)

	if r.MaxConcurrentSearches != nil {
		params["max_concurrent_searches"] = strconv.FormatInt(int64(*r.MaxConcurrentSearches), 10)
	}

	if r.PreFilterShardSize != nil {
		params["pre_filter_shard_size"] = strconv.FormatInt(int64(*r.PreFilterShardSize), 10)
	}

	if r.SearchType != "" {
		params["search_type"] = r.SearchType
	}

	if r.TypedKeys != nil {
		params["typed_keys"] = strconv.FormatBool(*r.TypedKeys)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f Msearch) WithContext(v context.Context) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.ctx = v
	}
}

// WithIndex - a list of index names to use as default.
//
func (f Msearch) WithIndex(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Index = v
	}
}

// WithDocumentType - a list of document types to use as default.
//
func (f Msearch) WithDocumentType(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.DocumentType = v
	}
}

// WithMaxConcurrentSearches - controls the maximum number of concurrent searches the multi search api will execute.
//
func (f Msearch) WithMaxConcurrentSearches(v int) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.MaxConcurrentSearches = &v
	}
}

// WithPreFilterShardSize - a threshold that enforces a pre-filter roundtrip to prefilter search shards based on query rewriting if the??number of shards the search request expands to exceeds the threshold. this filter roundtrip can limit the number of shards significantly if for instance a shard can not match any documents based on it's rewrite method ie. if date filters are mandatory to match but the shard bounds and the query are disjoint..
//
func (f Msearch) WithPreFilterShardSize(v int) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.PreFilterShardSize = &v
	}
}

// WithSearchType - search operation type.
//
func (f Msearch) WithSearchType(v string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.SearchType = v
	}
}

// WithTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
//
func (f Msearch) WithTypedKeys(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.TypedKeys = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f Msearch) WithPretty() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f Msearch) WithHuman() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f Msearch) WithErrorTrace() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f Msearch) WithFilterPath(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f Msearch) WithHeader(h map[string]string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}
