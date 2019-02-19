package application

type Application struct {
	// Application unique identifier.
	ID string `json:"id"`

	// Name is the human-readable string name of the client to be presented to the
	// end-user during authorization.
	Name string `json:"name"`

	// Description of the application.
	Description string `json:"description"`

	// RedirectURIs is an array of allowed redirect urls for the client, for example:
	// http://mydomain/oauth/callback.
	RedirectURIs []string `json:"redirectURIs"`

	// GrantTypes is an array of grant types the client is allowed to use.
	//
	// Pattern: client_credentials|authorize_code|implicit|refresh_token|password
	GrantTypes []string `json:"grantTypes"`

	// ResponseTypes is an array of the OAuth 2.0 response type strings that the client can
	// use at the authorization endpoint.
	//
	// Pattern: id_token|code|token
	ResponseTypes []string `json:"responseTypes"`

	// Scope is a string containing a space-separated list of scope values (as
	// described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client
	// can use when requesting access tokens.
	//
	// Pattern: ([a-zA-Z0-9\.]+\s)+
	Scopes []string `json:"scopes"`

	// Public is a boolean that identifies this client as public, meaning that it
	// does not have a secret. It will disable the client_credentials grant type for this client if set.
	Public bool `json:"public"`
}

type GetAllCmd struct{}

var ValidApp = Application{
	ID:            "some-app",
	Name:          "Some App",
	Description:   "A simple app",
	RedirectURIs:  []string{"some-url"},
	GrantTypes:    []string{"implicit", "refresh_token"},
	ResponseTypes: []string{"token", "code"},
	Scopes:        []string{"permissions"},
	Public:        true,
}
