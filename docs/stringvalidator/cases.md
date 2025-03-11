---
hide:
    - navigation
---
# `Cases`

!!! quote inline end "Released in v1.12.0"

This validator is a generic validator for checking if the string is a valid network format.

## How to use it

The validator takes a list of CasesValidatorType.

The list can contain one or more of the following values:

* `CasesDisallowUpper` - Check if the string does not contain any uppercase characters.
* `CasesDisallowLower` - Check if the string does not contain any lowercase characters.
* `CasesDisallowSpace`- Check if the string does not contain any space characters.
* `CasesDisallowNumber` - Check if the string does not contain any number characters.

### Example DisallowUpper and DisallowSpace

The following example will check if the string does not contain any uppercase characters and does not contain any space characters.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "user_name": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Username for ...",
                Validators: []validator.String{
                    fstringvalidator.Cases([]fstringvalidator.CasesValidatorType{
                        fstringvalidator.CasesDisallowUpper,
                        fstringvalidator.CasesDisallowSpace,
                    }, true)
                },
            },
```
