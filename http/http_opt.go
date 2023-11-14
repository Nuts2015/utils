package http

import (
	"github.com/go-resty/resty/v2"
)

type Option struct {
	headers    map[string]string // header
	params     map[string]string // params
	formData   map[string]string // formdata
	files      map[string]string // files
	outputFile string
}

type option func(*Option)

func WithHeader(k, v string) option {
	return func(o *Option) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[k] = v
	}
}

func WithParam(k, v string) option {
	return func(o *Option) {
		if o.params == nil {
			o.params = make(map[string]string)
		}
		o.params[k] = v
	}
}

func WithFormData(k, v string) option {
	return func(o *Option) {
		if o.formData == nil {
			o.formData = make(map[string]string)
		}
		o.formData[k] = v
	}
}

func WithFile(param, filepath string) option {
	return func(o *Option) {
		if o.files == nil {
			o.files = make(map[string]string)
		}
		// o.files = append(o.files, file)
		o.files[param] = filepath
	}
}

func WithOutput(file string) option {
	return func(o *Option) {
		o.outputFile = file
	}
}

func (c *Client) handleOpt(request *resty.Request, opts ...option) {

	o := &Option{}

	for _, opt := range opts {
		opt(o)
	}

	// headers
	if o.headers != nil {

		request.SetHeaders(o.headers)
	}

	// params
	if o.params != nil {
		request.SetQueryParams(o.params)
	}

	// formdata
	if o.formData != nil {
		request.SetFormData(o.formData)
	}

	// files
	if o.files != nil {
		request.SetFiles(o.files)
	}

	if o.outputFile != "" {
		request.SetOutput(o.outputFile)
	}

}
