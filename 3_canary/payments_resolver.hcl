kind = "service-resolver"
name = "payment"

# https://www.consul.io/api/health.html#filtering-2
# Show Node.Meta demonstration showing performance testing a new instance type
default_subset = "blue"

subsets = {
  blue = {
    filter = "Service.Meta.version == 2"
  }
  green = {
    filter = "Service.Meta.version == 3"
  }
}

