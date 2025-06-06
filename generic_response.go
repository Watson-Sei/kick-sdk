package kicksdk

import "net/http"

type (
	apiResponse[Payload any] struct {
		Payload Payload `json:"data,omitempty"`
		Message string  `json:"message,omitempty"`
	}

	authErrorResponse struct {
		Error            string `json:"error,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
)

// EmptyResponse is a response that is used as a stub in case endpoint is not returning any
// data in response Body.
type EmptyResponse any

// Response is a response that will be returned to the user as a result of a call to any
// Kick API endpoint.
type Response[Payload any] struct {
	Payload          Payload
	ResponseMetadata ResponseMetadata
}

// ResponseMetadata is a metadata of the Kick API response.
type ResponseMetadata struct {
	StatusCode int
	Header     http.Header

	// KickMessage is a message that Kick sends along with the optional data in response to the API requests.
	// In case of an unsuccessful Request it will contain error message as to why the Request failed.
	KickMessage          string
	KickError            string
	KickErrorDescription string
}
