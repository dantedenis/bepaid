package contracts

// RequestSetTest wraps SetTest method
//
// Implement this interface if you don't want to set Test field on every request creation
//
// Api implementation should let user set a mode. If mode is "test",
// then assert on every incoming request on RequestSetTest and call SetTest with true
type RequestSetTest interface {
	SetTest(test bool)
}
