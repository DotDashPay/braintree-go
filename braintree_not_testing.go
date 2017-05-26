// +build !testing

// This file is compiled in for production and disallows redirecting the braintree endpoints.

package braintree

func (e Environment) BaseURL() string {
	switch e {
	case Sandbox:
		return "https://sandbox.braintreegateway.com"
	case Production:
		return "https://www.braintreegateway.com"
	}
	panic(`invalid environment "` + e + `"`)
}
