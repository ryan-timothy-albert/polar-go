// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package polargo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/polarsource/polar-go/internal/hooks"
	"github.com/polarsource/polar-go/internal/utils"
	"github.com/polarsource/polar-go/models/apierrors"
	"github.com/polarsource/polar-go/models/components"
	"github.com/polarsource/polar-go/models/operations"
	"github.com/polarsource/polar-go/retry"
	"github.com/spyzhov/ajson"
	"net/http"
	"net/url"
)

type ExternalOrganizations struct {
	sdkConfiguration sdkConfiguration
}

func newExternalOrganizations(sdkConfig sdkConfiguration) *ExternalOrganizations {
	return &ExternalOrganizations{
		sdkConfiguration: sdkConfig,
	}
}

// List External Organizations
// List external organizations.
func (s *ExternalOrganizations) List(ctx context.Context, request operations.ExternalOrganizationsListRequest, opts ...operations.Option) (*operations.ExternalOrganizationsListResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "external_organizations:list",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := url.JoinPath(baseURL, "/v1/external_organizations/")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ExternalOrganizationsListResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}
	res.Next = func() (*operations.ExternalOrganizationsListResponse, error) {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}

		b, err := ajson.Unmarshal(rawBody)
		if err != nil {
			return nil, err
		}
		p := *request.Page
		nP := int64(p + 1)
		nPs, err := ajson.Eval(b, "$.pagination.max_page")
		if err != nil {
			return nil, err
		}
		if !nPs.IsNumeric() {
			return nil, nil
		}

		nPsVal, err := nPs.GetNumeric()
		if err != nil {
			return nil, err
		}
		// GetNumeric returns as float64
		if int(nPsVal) <= int(p) {
			return nil, nil
		}
		r, err := ajson.Eval(b, "$.items")
		if err != nil {
			return nil, err
		}
		if !r.IsArray() {
			return nil, nil
		}
		arr, err := r.GetArray()
		if err != nil {
			return nil, err
		}
		if len(arr) == 0 {
			return nil, nil
		}
		l := int(*request.Limit)
		if len(arr) < l {
			return nil, nil
		}

		return s.List(
			ctx,
			operations.ExternalOrganizationsListRequest{
				Platform:       request.Platform,
				Name:           request.Name,
				OrganizationID: request.OrganizationID,
				Page:           &nP,
				Limit:          request.Limit,
				Sorting:        request.Sorting,
			},
			opts...,
		)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.ListResourceExternalOrganization
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ListResourceExternalOrganization = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.HTTPValidationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
