---
layout: "cloudfoundry"
page_title: "Cloud Foundry: cloudfoundry_space_quota"
sidebar_current: "docs-cf-datasource-space-quota"
description: |-
  Get information on a Cloud Foundry space Quota.
---

# cloudfoundry\_space\_quota

Gets information on a Cloud Foundry space quota.

## Example Usage

The following example looks up a space quota named 'myquota'

```hcl
data "cloudfoundry_space_quota" "q" {
  name = "myquota"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the space quota to look up
* `org` - (Optional) The organization within which the quota is defined

## Attributes Reference

The following attributes are exported:

* `id` - The GUID of the space quota
