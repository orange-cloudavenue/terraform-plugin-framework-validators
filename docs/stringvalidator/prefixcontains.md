---
hide:
    - navigation
---
# `PrefixContains`

!!! quote inline end "Released in v1.7.0"

This validator is used to check if the string contains prefix in the given value.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "interface_type": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Type of ...",
                Validators: []validator.String{
                    fstringvalidator.PrefixContains("urn:test:demo:")
                },
            },
```

## Description and Markdown description

* **Description:**
must start with "urn:test:demo:"
* **Markdown description:**
This value must start with `urn:test:demo:`.
