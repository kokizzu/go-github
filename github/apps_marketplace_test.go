// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMarketplaceService_ListPlans(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/plans", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = false
	ctx := context.Background()
	plans, _, err := client.Marketplace.ListPlans(ctx, opt)
	if err != nil {
		t.Errorf("Marketplace.ListPlans returned error: %v", err)
	}

	want := []*MarketplacePlan{{ID: Int64(1)}}
	if !reflect.DeepEqual(plans, want) {
		t.Errorf("Marketplace.ListPlans returned %+v, want %+v", plans, want)
	}

	const methodName = "ListPlans"
	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Marketplace.ListPlans(ctx, opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestMarketplaceService_Stubbed_ListPlans(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/stubbed/plans", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = true
	ctx := context.Background()
	plans, _, err := client.Marketplace.ListPlans(ctx, opt)
	if err != nil {
		t.Errorf("Marketplace.ListPlans (Stubbed) returned error: %v", err)
	}

	want := []*MarketplacePlan{{ID: Int64(1)}}
	if !reflect.DeepEqual(plans, want) {
		t.Errorf("Marketplace.ListPlans (Stubbed) returned %+v, want %+v", plans, want)
	}
}

func TestMarketplaceService_ListPlanAccountsForPlan(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/plans/1/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = false
	ctx := context.Background()
	accounts, _, err := client.Marketplace.ListPlanAccountsForPlan(ctx, 1, opt)
	if err != nil {
		t.Errorf("Marketplace.ListPlanAccountsForPlan returned error: %v", err)
	}

	want := []*MarketplacePlanAccount{{ID: Int64(1)}}
	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Marketplace.ListPlanAccountsForPlan returned %+v, want %+v", accounts, want)
	}

	const methodName = "ListPlanAccountsForPlan"
	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Marketplace.ListPlanAccountsForPlan(ctx, 1, opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestMarketplaceService_Stubbed_ListPlanAccountsForPlan(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/stubbed/plans/1/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = true
	ctx := context.Background()
	accounts, _, err := client.Marketplace.ListPlanAccountsForPlan(ctx, 1, opt)
	if err != nil {
		t.Errorf("Marketplace.ListPlanAccountsForPlan (Stubbed) returned error: %v", err)
	}

	want := []*MarketplacePlanAccount{{ID: Int64(1)}}
	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Marketplace.ListPlanAccountsForPlan (Stubbed) returned %+v, want %+v", accounts, want)
	}
}

func TestMarketplaceService_GetPlanAccountForAccount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1, "marketplace_pending_change": {"id": 77}}`)
	})

	client.Marketplace.Stubbed = false
	ctx := context.Background()
	account, _, err := client.Marketplace.GetPlanAccountForAccount(ctx, 1)
	if err != nil {
		t.Errorf("Marketplace.GetPlanAccountForAccount returned error: %v", err)
	}

	want := &MarketplacePlanAccount{ID: Int64(1), MarketplacePendingChange: &MarketplacePendingChange{ID: Int64(77)}}
	if !reflect.DeepEqual(account, want) {
		t.Errorf("Marketplace.GetPlanAccountForAccount returned %+v, want %+v", account, want)
	}

	const methodName = "GetPlanAccountForAccount"
	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Marketplace.GetPlanAccountForAccount(ctx, 1)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestMarketplaceService_Stubbed_GetPlanAccountForAccount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketplace_listing/stubbed/accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	client.Marketplace.Stubbed = true
	ctx := context.Background()
	account, _, err := client.Marketplace.GetPlanAccountForAccount(ctx, 1)
	if err != nil {
		t.Errorf("Marketplace.GetPlanAccountForAccount (Stubbed) returned error: %v", err)
	}

	want := &MarketplacePlanAccount{ID: Int64(1)}
	if !reflect.DeepEqual(account, want) {
		t.Errorf("Marketplace.GetPlanAccountForAccount (Stubbed) returned %+v, want %+v", account, want)
	}
}

func TestMarketplaceService_ListMarketplacePurchasesForUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/marketplace_purchases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"billing_cycle":"monthly"}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = false
	ctx := context.Background()
	purchases, _, err := client.Marketplace.ListMarketplacePurchasesForUser(ctx, opt)
	if err != nil {
		t.Errorf("Marketplace.ListMarketplacePurchasesForUser returned error: %v", err)
	}

	want := []*MarketplacePurchase{{BillingCycle: String("monthly")}}
	if !reflect.DeepEqual(purchases, want) {
		t.Errorf("Marketplace.ListMarketplacePurchasesForUser returned %+v, want %+v", purchases, want)
	}

	const methodName = "ListMarketplacePurchasesForUser"
	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Marketplace.ListMarketplacePurchasesForUser(ctx, opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestMarketplaceService_Stubbed_ListMarketplacePurchasesForUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/marketplace_purchases/stubbed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"billing_cycle":"monthly"}]`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	client.Marketplace.Stubbed = true
	ctx := context.Background()
	purchases, _, err := client.Marketplace.ListMarketplacePurchasesForUser(ctx, opt)
	if err != nil {
		t.Errorf("Marketplace.ListMarketplacePurchasesForUser returned error: %v", err)
	}

	want := []*MarketplacePurchase{{BillingCycle: String("monthly")}}
	if !reflect.DeepEqual(purchases, want) {
		t.Errorf("Marketplace.ListMarketplacePurchasesForUser returned %+v, want %+v", purchases, want)
	}
}
