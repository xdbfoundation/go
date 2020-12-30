package digitalbitstoml

import "log"

// ExampleGetTOML gets the digitalbits.toml file for coins.asia
func ExampleClient_GetDigitalBitsToml() {
	_, err := DefaultClient.GetDigitalBitsToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
