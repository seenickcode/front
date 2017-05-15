package front

import "fmt"

const (
	FrontHostname     = "https://api2.frontapp.com"
	FrontBaseEndpoint = ""
)

type Front struct {
	jwtToken string
}

func New(jwtToken string) (*Front, error) {
	if len(jwtToken) == 0 {
		return nil, fmt.Errorf("Front JWT Token is required.")
	}
	return &Front{
		jwtToken: jwtToken,
	}, nil
}
