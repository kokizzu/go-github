// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestActionsService_ListWorkflowJobs(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/actions/runs/29679449/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"jobs":[{"id":399444496,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"},{"id":399444497,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}]}`)
	})

	opts := &ListWorkflowJobsOptions{ListOptions: ListOptions{Page: 2, PerPage: 2}}
	ctx := context.Background()
	jobs, _, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflowJobs returned error: %v", err)
	}

	want := &Jobs{
		TotalCount: Ptr(4),
		Jobs: []*WorkflowJob{
			{ID: Ptr(int64(399444496)), RunID: Ptr(int64(29679449)), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
			{ID: Ptr(int64(399444497)), RunID: Ptr(int64(29679449)), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
		},
	}
	if !cmp.Equal(jobs, want) {
		t.Errorf("Actions.ListWorkflowJobs returned %+v, want %+v", jobs, want)
	}

	const methodName = "ListWorkflowJobs"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.ListWorkflowJobs(ctx, "\n", "\n", 29679449, opts)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_ListWorkflowJobs_Filter(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/actions/runs/29679449/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"filter": "all", "per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"jobs":[{"id":399444496,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"},{"id":399444497,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}]}`)
	})

	opts := &ListWorkflowJobsOptions{Filter: "all", ListOptions: ListOptions{Page: 2, PerPage: 2}}
	ctx := context.Background()
	jobs, _, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflowJobs returned error: %v", err)
	}

	want := &Jobs{
		TotalCount: Ptr(4),
		Jobs: []*WorkflowJob{
			{ID: Ptr(int64(399444496)), RunID: Ptr(int64(29679449)), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
			{ID: Ptr(int64(399444497)), RunID: Ptr(int64(29679449)), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
		},
	}
	if !cmp.Equal(jobs, want) {
		t.Errorf("Actions.ListWorkflowJobs returned %+v, want %+v", jobs, want)
	}
}

func TestActionsService_ListWorkflowJobsAttempt(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/actions/runs/29679449/attempts/1/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"jobs":[{"id":399444496,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z","run_attempt":2},{"id":399444497,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z","run_attempt":2}]}`)
	})
	opts := &ListOptions{Page: 2, PerPage: 2}
	ctx := context.Background()
	jobs, _, err := client.Actions.ListWorkflowJobsAttempt(ctx, "o", "r", 29679449, 1, opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflowJobsAttempt returned error: %v", err)
	}

	want := &Jobs{
		TotalCount: Ptr(4),
		Jobs: []*WorkflowJob{
			{
				ID:          Ptr(int64(399444496)),
				RunID:       Ptr(int64(29679449)),
				StartedAt:   &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
				CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
				RunAttempt:  Ptr(int64(2)),
			},
			{
				ID:          Ptr(int64(399444497)),
				RunID:       Ptr(int64(29679449)),
				StartedAt:   &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
				CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
				RunAttempt:  Ptr(int64(2)),
			},
		},
	}
	if !cmp.Equal(jobs, want) {
		t.Errorf("Actions.ListWorkflowJobsAttempt returned %+v, want %+v", jobs, want)
	}

	const methodName = "ListWorkflowJobsAttempt"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.ListWorkflowJobsAttempt(ctx, "\n", "\n", 29679449, 1, opts)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.ListWorkflowJobsAttempt(ctx, "o", "r", 29679449, 1, opts)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_GetWorkflowJobByID(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/actions/jobs/399444496", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":399444496,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}`)
	})

	ctx := context.Background()
	job, _, err := client.Actions.GetWorkflowJobByID(ctx, "o", "r", 399444496)
	if err != nil {
		t.Errorf("Actions.GetWorkflowJobByID returned error: %v", err)
	}

	want := &WorkflowJob{
		ID:          Ptr(int64(399444496)),
		StartedAt:   &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
		CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
	}
	if !cmp.Equal(job, want) {
		t.Errorf("Actions.GetWorkflowJobByID returned %+v, want %+v", job, want)
	}

	const methodName = "GetWorkflowJobByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.GetWorkflowJobByID(ctx, "\n", "\n", 399444496)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.GetWorkflowJobByID(ctx, "o", "r", 399444496)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_GetWorkflowJobLogs(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name              string
		respectRateLimits bool
	}{
		{
			name:              "withoutRateLimits",
			respectRateLimits: false,
		},
		{
			name:              "withRateLimits",
			respectRateLimits: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, mux, _ := setup(t)
			client.RateLimitRedirectionalEndpoints = tc.respectRateLimits

			mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				http.Redirect(w, r, "https://github.com/a", http.StatusFound)
			})

			ctx := context.Background()
			url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, 1)
			if err != nil {
				t.Errorf("Actions.GetWorkflowJobLogs returned error: %v", err)
			}
			if resp.StatusCode != http.StatusFound {
				t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusFound)
			}
			want := "https://github.com/a"
			if url.String() != want {
				t.Errorf("Actions.GetWorkflowJobLogs returned %+v, want %+v", url.String(), want)
			}

			const methodName = "GetWorkflowJobLogs"
			testBadOptions(t, methodName, func() (err error) {
				_, _, err = client.Actions.GetWorkflowJobLogs(ctx, "\n", "\n", 399444496, 1)
				return err
			})

			// Add custom round tripper
			client.client.Transport = roundTripperFunc(func(*http.Request) (*http.Response, error) {
				return nil, errors.New("failed to get workflow logs")
			})
			// propagate custom round tripper to client without CheckRedirect
			client.initialize()
			testBadOptions(t, methodName, func() (err error) {
				_, _, err = client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, 1)
				return err
			})
		})
	}
}

func TestActionsService_GetWorkflowJobLogs_StatusMovedPermanently_dontFollowRedirects(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name              string
		respectRateLimits bool
	}{
		{
			name:              "withoutRateLimits",
			respectRateLimits: false,
		},
		{
			name:              "withRateLimits",
			respectRateLimits: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, mux, _ := setup(t)
			client.RateLimitRedirectionalEndpoints = tc.respectRateLimits

			mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				http.Redirect(w, r, "https://github.com/a", http.StatusMovedPermanently)
			})

			ctx := context.Background()
			_, resp, _ := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, 0)
			if resp.StatusCode != http.StatusMovedPermanently {
				t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusMovedPermanently)
			}
		})
	}
}

func TestActionsService_GetWorkflowJobLogs_StatusMovedPermanently_followRedirects(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name              string
		respectRateLimits bool
	}{
		{
			name:              "withoutRateLimits",
			respectRateLimits: false,
		},
		{
			name:              "withRateLimits",
			respectRateLimits: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, mux, serverURL := setup(t)
			client.RateLimitRedirectionalEndpoints = tc.respectRateLimits

			// Mock a redirect link, which leads to an archive link
			mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				redirectURL, _ := url.Parse(serverURL + baseURLPath + "/redirect")
				http.Redirect(w, r, redirectURL.String(), http.StatusMovedPermanently)
			})

			mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				http.Redirect(w, r, "https://github.com/a", http.StatusFound)
			})

			ctx := context.Background()
			url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, 1)
			if err != nil {
				t.Errorf("Actions.GetWorkflowJobLogs returned error: %v", err)
			}

			if resp.StatusCode != http.StatusFound {
				t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusFound)
			}

			want := "https://github.com/a"
			if url.String() != want {
				t.Errorf("Actions.GetWorkflowJobLogs returned %+v, want %+v", url.String(), want)
			}
		})
	}
}

func TestActionsService_GetWorkflowJobLogs_unexpectedCode(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name              string
		respectRateLimits bool
	}{
		{
			name:              "withoutRateLimits",
			respectRateLimits: false,
		},
		{
			name:              "withRateLimits",
			respectRateLimits: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client, mux, serverURL := setup(t)
			client.RateLimitRedirectionalEndpoints = tc.respectRateLimits

			// Mock a redirect link, which leads to an archive link
			mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				redirectURL, _ := url.Parse(serverURL + baseURLPath + "/redirect")
				http.Redirect(w, r, redirectURL.String(), http.StatusMovedPermanently)
			})

			mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				w.WriteHeader(http.StatusNoContent)
			})

			ctx := context.Background()
			url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, 1)
			if err == nil {
				t.Fatal("Actions.GetWorkflowJobLogs should return error on unexpected code")
			}
			if !strings.Contains(err.Error(), "unexpected status code") {
				t.Error("Actions.GetWorkflowJobLogs should return unexpected status code")
			}
			if got, want := resp.Response.StatusCode, http.StatusNoContent; got != want {
				t.Errorf("Actions.GetWorkflowJobLogs return status %d, want %d", got, want)
			}
			if url != nil {
				t.Errorf("Actions.GetWorkflowJobLogs return %+v, want nil", url)
			}
		})
	}
}

func TestTaskStep_Marshal(t *testing.T) {
	t.Parallel()
	testJSONMarshal(t, &TaskStep{}, "{}")

	u := &TaskStep{
		Name:        Ptr("n"),
		Status:      Ptr("s"),
		Conclusion:  Ptr("c"),
		Number:      Ptr(int64(1)),
		StartedAt:   &Timestamp{referenceTime},
		CompletedAt: &Timestamp{referenceTime},
	}

	want := `{
		"name": "n",
		"status": "s",
		"conclusion": "c",
		"number": 1,
		"started_at": ` + referenceTimeStr + `,
		"completed_at": ` + referenceTimeStr + `
	}`

	testJSONMarshal(t, u, want)
}

func TestWorkflowJob_Marshal(t *testing.T) {
	t.Parallel()
	testJSONMarshal(t, &WorkflowJob{}, "{}")

	u := &WorkflowJob{
		ID:          Ptr(int64(1)),
		RunID:       Ptr(int64(1)),
		RunURL:      Ptr("r"),
		NodeID:      Ptr("n"),
		HeadBranch:  Ptr("b"),
		HeadSHA:     Ptr("h"),
		URL:         Ptr("u"),
		HTMLURL:     Ptr("h"),
		Status:      Ptr("s"),
		Conclusion:  Ptr("c"),
		CreatedAt:   &Timestamp{referenceTime},
		StartedAt:   &Timestamp{referenceTime},
		CompletedAt: &Timestamp{referenceTime},
		Name:        Ptr("n"),
		Steps: []*TaskStep{
			{
				Name:        Ptr("n"),
				Status:      Ptr("s"),
				Conclusion:  Ptr("c"),
				Number:      Ptr(int64(1)),
				StartedAt:   &Timestamp{referenceTime},
				CompletedAt: &Timestamp{referenceTime},
			},
		},
		CheckRunURL:  Ptr("c"),
		WorkflowName: Ptr("w"),
	}

	want := `{
		"id": 1,
		"run_id": 1,
		"run_url": "r",
		"node_id": "n",
		"head_branch": "b",
		"head_sha": "h",
		"url": "u",
		"html_url": "h",
		"status": "s",
		"conclusion": "c",
		"created_at": ` + referenceTimeStr + `,
		"started_at": ` + referenceTimeStr + `,
		"completed_at": ` + referenceTimeStr + `,
		"name": "n",
		"steps": [{
			"name": "n",
			"status": "s",
			"conclusion": "c",
			"number": 1,
			"started_at": ` + referenceTimeStr + `,
			"completed_at": ` + referenceTimeStr + `
		}],
		"check_run_url": "c",
		"workflow_name": "w"
	}`

	testJSONMarshal(t, u, want)
}

func TestJobs_Marshal(t *testing.T) {
	t.Parallel()
	testJSONMarshal(t, &Jobs{}, "{}")

	u := &Jobs{
		TotalCount: Ptr(1),
		Jobs: []*WorkflowJob{
			{
				ID:          Ptr(int64(1)),
				RunID:       Ptr(int64(1)),
				RunURL:      Ptr("r"),
				NodeID:      Ptr("n"),
				HeadBranch:  Ptr("b"),
				HeadSHA:     Ptr("h"),
				URL:         Ptr("u"),
				HTMLURL:     Ptr("h"),
				Status:      Ptr("s"),
				Conclusion:  Ptr("c"),
				CreatedAt:   &Timestamp{referenceTime},
				StartedAt:   &Timestamp{referenceTime},
				CompletedAt: &Timestamp{referenceTime},
				Name:        Ptr("n"),
				Steps: []*TaskStep{
					{
						Name:        Ptr("n"),
						Status:      Ptr("s"),
						Conclusion:  Ptr("c"),
						Number:      Ptr(int64(1)),
						StartedAt:   &Timestamp{referenceTime},
						CompletedAt: &Timestamp{referenceTime},
					},
				},
				CheckRunURL:  Ptr("c"),
				RunAttempt:   Ptr(int64(2)),
				WorkflowName: Ptr("w"),
			},
		},
	}

	want := `{
		"total_count": 1,
		"jobs": [{
			"id": 1,
			"run_id": 1,
			"run_url": "r",
			"node_id": "n",
			"head_branch": "b",
			"head_sha": "h",
			"url": "u",
			"html_url": "h",
			"status": "s",
			"conclusion": "c",
			"created_at": ` + referenceTimeStr + `,
			"started_at": ` + referenceTimeStr + `,
			"completed_at": ` + referenceTimeStr + `,
			"name": "n",
			"steps": [{
				"name": "n",
				"status": "s",
				"conclusion": "c",
				"number": 1,
				"started_at": ` + referenceTimeStr + `,
				"completed_at": ` + referenceTimeStr + `
			}],
			"check_run_url": "c",
			"run_attempt": 2,
			"workflow_name": "w"
		}]
	}`

	testJSONMarshal(t, u, want)
}
