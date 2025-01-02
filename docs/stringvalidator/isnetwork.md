---
hide:
    - navigation
---
# `IsNetwork`

!!! quote inline end "Released in v1.8.0"

This validator is a generic validator for checking if the string is a valid network format.

Some network formats are :

## How to use it

The validator takes a list of NetworkValidatorType and a boolean as argument.

The list can contain one or more of the following values:

* `IPV4` - Check if the string is a valid IPV4 address.
* `IPV4WithCIDR` - Check if the string is a valid IPV4 address with CIDR.
* `IPV4WithNetmask`- Check if the string is a valid IPV4 address with netmask.
* `IsRFC1918` - Check if the string is a valid [RFC1918](https://en.wikipedia.org/wiki/Private_network) address.

The boolean is used to define if the value must be at least one of the network types.

### Example OR

The following example will check if the string is a valid IPV4 address with CIDR or a valid IPV4 address with netmask.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "ip_address": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "IPV4 for ...",
                Validators: []validator.String{
                    fstringvalidator.IsNetwork([]fstringvalidator.NetworkValidatorType{
                        fstringvalidator.IPV4WithCIDR,
                        fstringvalidator.IPV4WithNetmask,
                    }, true)
                },
            },
```

### Example AND

The following example will check if the string is a valid IPV4 and a valid RFC1918 address.

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "ip_address": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "IPV4 for ...",
                Validators: []validator.String{
                    fstringvalidator.IsNetwork([]fstringvalidator.NetworkValidatorType{
                        fstringvalidator.IPV4,
                        fstringvalidator.IsRFC1918,
                    }, false)
                },
            },
```
