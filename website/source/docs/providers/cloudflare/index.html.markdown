---
layout: "cloudflare"
page_title: "Provider: Cloudflare"
sidebar_current: "docs-cloudflare-index"
---

# CloudFlare Provider

The CloudFlare provider is used to interact with the
DNS resources supported by CloudFlare. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Configure the CloudFlare provider
provider "cloudflare" {
    email = "${var.cloudflare_email}"
    token = "${var.cloudflare_token}"
}

# Create a record
resource "cloudflare_record" "www" {
    ...
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The email associated with the account
* `token` - (Required) The Cloudflare API token


