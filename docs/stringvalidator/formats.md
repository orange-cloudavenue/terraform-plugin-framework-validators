---
hide:
    - navigation
---
# `Formats`

The Formats validator is a flexible utility designed to validate whether a given string adheres to specific predefined formats. It supports multiple validation types, enabling you to ensure strings meet requirements such as Base64 encoding, UUIDs, etc...

## How to use it

The validator accepts a list of FormatsValidatorType values, which specify the formats to validate. You can include one or more of the following options:

* `IsBase64` - Check if the string is a valid Base64 encoded string.
* `IsUUIDv4` - Check if the string is a valid (v4) UUID.
* `IsURN` - Check if the string is a valid URN.

### Example IsBase64

The following example demonstrates how to use the validator to check if a string is a valid Base64-encoded value.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "content_encoded": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Content for ...",
                Validators: []validator.String{
                    fstringvalidator.Formats(
                        []fstringvalidator.FormatsValidatorType{
                            fstringvalidator.IsBase64,
                        }, 
                        false,
                    ),
                },
            },
```

### Example IsUUIDv4

The following example demonstrates how to use the validator to check if a string is a valid version 4 (v4) UUID.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {  
    resp.Schema = schema.Schema{  
        (...)  
        "uuid": schema.StringAttribute{  
            Optional:            true,  
            MarkdownDescription: "Unique identifier (UUID v4) for the resource.",  
            Validators: []validator.String{  
                fstringvalidator.Formats(  
                    []fstringvalidator.FormatsValidatorType{  
                        fstringvalidator.IsUUIDv4,  
                    }, 
                    false,
                ),
            },
        },
        (...)
    }
}
```

### Example IsURN

The following example will check if the string is a valid URN.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "urn": schema.StringAttribute{
            Optional:            true,
            MarkdownDescription: "Uniform Resource Name (URN) for the resource.",
            Validators: []validator.String{
                fstringvalidator.Formats(
                    []fstringvalidator.FormatsValidatorType{
                        fstringvalidator.IsURN,
                    }, 
                    false,
                ),
            },
        },
        (...)
    }
}
```

### Example with multiple formats checks

The Formats validator can also be used to validate multiple formats at once. You can combine different formats in a single validator by passing them as a slice of FormatsValidatorType values. This allows you to check if a string is valid for any of the specified formats.

The following example demonstrates how to use the validator to check if a string is a valid URN value **OR** a valid version 4 (v4) UUID.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "id": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "ID can be an urn or uuid format ...",
                Validators: []validator.String{
                    fstringvalidator.Formats(
                        []fstringvalidator.FormatsValidatorType{
                            fstringvalidator.IsURN,
                            fstringvalidator.IsUUIDv4,
                        }, 
                        true,
                    ),
                },
            },
        (...)
    }
}
```
