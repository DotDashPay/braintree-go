// +build testing
//
// This file is compiled in for testing and allows you to redirect the braintree endpoint.
//
// go build -tags testing

package braintree

const Testing Environment = "testing"

var testUrl = ""

// SetTestUrl sets the URL used in lieu of a Braintree url
// when using the Testing Environment, i.e. all http requests
// will go to this URL instead of a Braintree URL.
func SetTestUrl(url string) {
	testUrl = url
}

func (e Environment) BaseURL() string {
	switch e {
	case Testing:
		return testUrl
	case Sandbox:
		return "https://sandbox.braintreegateway.com"
	case Production:
		return "https://www.braintreegateway.com"
	}
	panic(`invalid environment "` + e + `"`)
}
