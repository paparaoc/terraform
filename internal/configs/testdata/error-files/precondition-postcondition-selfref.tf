resource "test" "test" {
  lifecycle {
    precondition {
      condition     = test.test.foo # ERROR: Invalid reference in precondition
      error_message = "Cannot refer to self."
    }
    postcondition {
      condition     = test.test.foo # ERROR: Invalid reference in postcondition
      error_message = "Cannot refer to self."
    }
  }
}

data "test" "test" {
  lifecycle {
    precondition {
      condition     = data.test.test.foo # ERROR: Invalid reference in precondition
      error_message = "Cannot refer to self."
    }
    postcondition {
      condition     = data.test.test.foo # ERROR: Invalid reference in postcondition
      error_message = "Cannot refer to self."
    }
  }
}
