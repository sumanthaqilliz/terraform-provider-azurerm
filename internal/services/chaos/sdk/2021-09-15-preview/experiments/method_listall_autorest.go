package experiments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ListAllResponse struct {
	HttpResponse *http.Response
	Model        *[]Experiment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAllResponse, error)
}

type ListAllCompleteResult struct {
	Items []Experiment
}

func (r ListAllResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAllResponse) LoadMore(ctx context.Context) (resp ListAllResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAllOptions struct {
	ContinuationToken *string
	Running           *bool
}

func DefaultListAllOptions() ListAllOptions {
	return ListAllOptions{}
}

func (o ListAllOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ContinuationToken != nil {
		out["continuationToken"] = *o.ContinuationToken
	}

	if o.Running != nil {
		out["running"] = *o.Running
	}

	return out
}

// ListAll ...
func (c ExperimentsClient) ListAll(ctx context.Context, id SubscriptionId, options ListAllOptions) (resp ListAllResponse, err error) {
	req, err := c.preparerForListAll(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAll(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListAllComplete retrieves all of the results into a single object
func (c ExperimentsClient) ListAllComplete(ctx context.Context, id SubscriptionId, options ListAllOptions) (ListAllCompleteResult, error) {
	return c.ListAllCompleteMatchingPredicate(ctx, id, options, ExperimentPredicate{})
}

// ListAllCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ExperimentsClient) ListAllCompleteMatchingPredicate(ctx context.Context, id SubscriptionId, options ListAllOptions, predicate ExperimentPredicate) (resp ListAllCompleteResult, err error) {
	items := make([]Experiment, 0)

	page, err := c.ListAll(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListAllCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListAll prepares the ListAll request.
func (c ExperimentsClient) preparerForListAll(ctx context.Context, id SubscriptionId, options ListAllOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Chaos/experiments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAllWithNextLink prepares the ListAll request with the given nextLink token.
func (c ExperimentsClient) preparerForListAllWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListAll handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (c ExperimentsClient) responderForListAll(resp *http.Response) (result ListAllResponse, err error) {
	type page struct {
		Values   []Experiment `json:"value"`
		NextLink *string      `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAllResponse, err error) {
			req, err := c.preparerForListAllWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAll(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "ListAll", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
