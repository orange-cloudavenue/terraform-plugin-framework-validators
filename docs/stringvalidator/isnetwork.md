---
hide:
    - navigation
---
# `IsNetwork`

!!! quote inline end "Released in v1.8.0"

This validator is a generic validator for checking if the string respects a case.

## How to use it

The validator takes a list of NetworkValidatorType and a boolean as argument.

The list can contain one or more of the following values:

**IPV4**

* `IPV4` - Check if the string is a valid IPV4 address (Ex: 192.168.0.1).
* `IPV4WithCIDR` - Check if the string is a valid IPV4 address with CIDR (Ex: 192.168.0.0/24).
* `IPV4WithNetmask`- Check if the string is a valid IPV4 address with netmask (Ex: 192.168.0.0/255.255.255.0).
* `IPV4Range` - Check if the string is a valid IPV4 address range (Ex: 192.168.0.1-192.168.0.10).
* `RFC1918` - Check if the string is a valid [RFC1918](https://en.wikipedia.org/wiki/Private_network) address.

**TCP/UDP**

* `TCPUDPPortRange` - Check if the string is a valid TCP/UDP port range (Ex: 80-90).

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
                        fstringvalidator.RFC1918,
                    }, false)
                },
            },
```
