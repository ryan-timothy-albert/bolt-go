// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltgo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://{environment}.bolt.com/v3",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// BoltTypescriptSDK - Bolt API Reference: A comprehensive Bolt API reference for interacting with Transactions, Orders, Product Catalog, Configuration, Testing, and much more.
type BoltTypescriptSDK struct {
	// Account endpoints allow you to view and manage shoppers' accounts. For example,
	// you can add or remove addresses and payment information.
	//
	Account  *Account
	Payments *Payments
	// Use this endpoint to retrieve an OAuth token. Use the token to allow your ecommerce server to make calls to the Account
	// endpoint and create a one-click checkout experience for shoppers.
	//
	//
	// https://help.bolt.com/products/accounts/direct-api/oauth-guide/
	OAuth *OAuth
	// Use the Orders API to create and manage orders, including orders that have been placed outside the Bolt ecosystem.
	//
	Orders *Orders
	// Endpoints that allow you to generate and retrieve test data to verify certain
	// flows in non-production environments.
	//
	Testing *Testing

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*BoltTypescriptSDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

type ServerEnvironment string

const (
	ServerEnvironmentAPI        ServerEnvironment = "api"
	ServerEnvironmentAPISandbox ServerEnvironment = "api-sandbox"
)

func (e ServerEnvironment) ToPointer() *ServerEnvironment {
	return &e
}

func (e *ServerEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api":
		fallthrough
	case "api-sandbox":
		*e = ServerEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerEnvironment: %v", v)
	}
}

// WithEnvironment allows setting the environment variable for url substitution
func WithEnvironment(environment ServerEnvironment) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["environment"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["environment"] = fmt.Sprintf("%v", environment)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security components.Security) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *BoltTypescriptSDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *BoltTypescriptSDK {
	sdk := &BoltTypescriptSDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "3.0.1",
			SDKVersion:        "0.2.0",
			GenVersion:        "2.250.15",
			UserAgent:         "speakeasy-sdk/go 0.2.0 2.250.15 3.0.1 github.com/BoltApp/bolt-go",
			ServerDefaults: []map[string]string{
				{
					"environment": "api-sandbox",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Account = newAccount(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.OAuth = newOAuth(sdk.sdkConfiguration)

	sdk.Orders = newOrders(sdk.sdkConfiguration)

	sdk.Testing = newTesting(sdk.sdkConfiguration)

	return sdk
}
