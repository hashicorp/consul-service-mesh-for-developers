# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

kind = "service-splitter",
name = "payment"

splits = [
  {
    weight = 50,
    service_subset = "blue"
  },
  {
    weight = 50,
    service_subset = "green"
  }
]
