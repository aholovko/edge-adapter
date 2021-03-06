/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"

	"github.com/hyperledger/aries-framework-go/pkg/client/outofband"
	"github.com/hyperledger/aries-framework-go/pkg/doc/presexch"
)

// GetPresentationRequestResponse API response of getPresentationRequest.
type GetPresentationRequestResponse struct {
	PD                   *presexch.PresentationDefinition `json:"pd"`
	Inv                  *outofband.Invitation            `json:"invitation"`
	Credentials          []json.RawMessage                `json:"credentials,omitempty"`
	CredentialGovernance json.RawMessage                  `json:"credentialGovernance,omitempty"`
}

// CreateRPTenantRequest API request body to register an RP tenant.
type CreateRPTenantRequest struct {
	Label                string   `json:"label"`
	Callback             string   `json:"callback"`
	Scopes               []string `json:"scopes"`
	RequiresBlindedRoute bool     `json:"requiresBlindedRoute"`
	SupportsWACI         bool     `json:"supportsWACI"`
}

// CreateRPTenantResponse API response body to register an RP tenant.
type CreateRPTenantResponse struct {
	ClientID             string   `json:"clientID"`
	ClientSecret         string   `json:"clientSecret"`
	PublicDID            string   `json:"publicDID"`
	Scopes               []string `json:"scopes"`
	RequiresBlindedRoute bool     `json:"requiresBlindedRoute"`
	SupportsWACI         bool     `json:"supportsWACI"`
}

// HandleCHAPIResponse is the input message to the chapiResponseHandler handler.
type HandleCHAPIResponse struct {
	InvitationID           string          `json:"invID"`
	VerifiablePresentation json.RawMessage `json:"vp"`
}

// HandleCHAPIResponseResult is the body of the response to a HandleCHAPIResponse request.
type HandleCHAPIResponseResult struct {
	RedirectURL string `json:"redirectURL"`
}

// DIDDocReq model.
type DIDDocReq struct {
	ID   string `json:"@id,omitempty"`
	Type string `json:"@type,omitempty"`
}

// DIDDocResp model.
type DIDDocResp struct {
	ID   string          `json:"@id,omitempty"`
	Type string          `json:"@type,omitempty"`
	Data *DIDDocRespData `json:"data,omitempty"`
}

// DIDDocRespData model for error data in DIDDocResp.
type DIDDocRespData struct {
	ErrorMsg string          `json:"errorMsg,omitempty"`
	DIDDoc   json.RawMessage `json:"didDoc,omitempty"`
}

// ErrorResp model.
type ErrorResp struct {
	ID   string         `json:"@id,omitempty"`
	Type string         `json:"@type,omitempty"`
	Data *ErrorRespData `json:"data,omitempty"`
}

// ErrorRespData model for error data in ErrorResp.
type ErrorRespData struct {
	ErrorMsg string `json:"errorMsg,omitempty"`
}
