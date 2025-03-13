---
hide:
    - navigation
---
# `Cases`

!!! quote inline end "Released in v1.15.0"

This validator is used to check if the string contains a valid http status code.
A parameter can be passed to allow a category of status codes.
The following categories are available:

* `1xx` - Informational responses
* `2xx` - Successful responses
* `3xx` - Redirection messages
* `4xx` - Client error responses
* `5xx` - Server error responses

## How to use it

The following example will check if the string does not contain any uppercase characters and does not contain any space characters.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "status_code": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Allowed HTTP status code",
                Validators: []validator.String{
                    fstringvalidator.HTTPCode(stringvalidator.HTTPCodeParams{
                      Allow1xx: false,
                      Allow2xx: true,
                      Allow3xx: true,
                      Allow4xx: false,
                      Allow5xx: false,
           })
                },
            },
```

In this example, the validator will check if the string is a valid HTTP status code and will allow only 2xx and 3xx status codes.
