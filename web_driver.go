package aaronweb

var DefaultWebDriver WebEngine

func SetWebDriver(driver WebEngine) {
	if driver == nil {
		panic("web driver cannot be nil")
	}
	DefaultWebDriver = driver
}
