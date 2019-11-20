kind = "service-splitter",
name = "payments"

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
