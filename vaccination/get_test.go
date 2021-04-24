package vaccination

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
	"testing/iotest"
	"time"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
)

const sampleData = `[
  {
    "area":"ΑΙΤΩΛΟΑΚΑΡΝΑΝΙΑΣ",
    "areaid":701,
    "dailydose1":665,
    "dailydose2":171,
    "daydiff":118,
    "daytotal":836,
    "referencedate":"2021-04-22T00:00:00",
    "totaldistinctpersons":27193,
    "totaldose1":27193,
    "totaldose2":7111,
    "totalvaccinations":34304
  },
  {
    "area":"ΑΝΑΤΟΛΙΚΗΣ ΑΤΤΙΚΗΣ",
    "areaid":901,
    "dailydose1":1054,
    "dailydose2":107,
    "daydiff":-39,
    "daytotal":1161,
    "referencedate":"2021-04-22T00:00:00",
    "totaldistinctpersons":51872,
    "totaldose1":51872,
    "totaldose2":17531,
    "totalvaccinations":69403
  }
]`

func TestNewDefaultGetParams(t *testing.T) {
	tests := []struct {
		name      string
		expected  *GetParams
		apiClient *api.Client
		opts      []GetParamsOpts
	}{
		{
			name:      "should return default get params without opts",
			apiClient: api.NewClient("some-token"),
			expected: &GetParams{
				client: api.NewClient("some-token"),
				dateTo: time.Now(),
			},
		},
		{
			name:      "should return default get params with opts",
			apiClient: api.NewClient("some-token"),
			expected: &GetParams{
				client:   api.NewClient("some-token"),
				dateTo:   time.Now().Add(time.Hour * 24 * 7),
				dateFrom: time.Now().Add(-time.Hour * 24 * 7),
			},
			opts: []GetParamsOpts{
				SetDateFrom(time.Now().Add(-time.Hour * 24 * 7)),
				SetDateTo(time.Now().Add(time.Hour * 24 * 7)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewDefaultGetParams(tt.apiClient, tt.opts...)
			assert.That(t, actual.client).IsEqualTo(tt.expected.client)
			if len(tt.opts) == 0 {
				assert.ThatTime(t, actual.dateFrom).IsNotDefined()
			} else {
				assert.ThatTime(t, actual.dateFrom).IsAlmostSameAs(tt.expected.dateFrom)
			}
			assert.ThatTime(t, actual.dateTo).IsAlmostSameAs(tt.expected.dateTo)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		params      *GetParams
		expected    List
		expectedErr error
	}{
		{
			name: "should fail if API returns an error",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path:  "mdg_emvolio",
					Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					Err:   errors.New("cannot fetch data"),
				})),
			)),
			expectedErr: errors.New("get vaccination data: cannot fetch data"),
		},
		{
			name: "should fail if the API response cannot be un-marshaled",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path:  "mdg_emvolio",
					Query: fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
				})),
			)),
			expectedErr: errors.New("get vaccination data: failed to un-marshal response unexpected end of JSON input"),
		},
		{
			name: "should fail if the API returns an unexpected response",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					StatusCode:   401,
					Path:         "mdg_emvolio",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: api.NewMockBody(`{"details":"λανθασμένο token"}`),
				})),
			)),
			expectedErr: errors.New("get vaccination data: {\"details\":\"λανθασμένο token\"}"),
		},
		{
			name: "should fail if the API returns an unexpected response and body cannot be read",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					StatusCode:   401,
					Path:         "mdg_emvolio",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: ioutil.NopCloser(iotest.ErrReader(errors.New("cannot read response data"))),
				})),
			)),
			expectedErr: errors.New("get vaccination data: cannot read response data"),
		},
		{
			name: "should return the expected data",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path:         "mdg_emvolio",
					Query:        fmt.Sprintf("date_to=%s", time.Now().Format("2006-01-02")),
					ResponseBody: api.NewMockBody(sampleData),
				})),
			)),
			expected: expectedList(),
		},
		{
			name: "should return the expected with query params",
			params: NewDefaultGetParams(api.NewClient(api.MockAPIToken,
				api.SetHTTPClient(api.NewMockHTTPClient(api.MockRequest{
					Path: "mdg_emvolio",
					Query: fmt.Sprintf("date_from=%s&date_to=%s",
						time.Now().Add(-7*time.Hour*24).Format("2006-01-02"),
						time.Now().Format("2006-01-02")),
					ResponseBody: api.NewMockBody(sampleData),
				})),
			), SetDateFrom(time.Now().Add(-7*time.Hour*24))),
			expected: expectedList(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Get(tt.params)
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
			assert.That(t, actual).IsEqualTo(tt.expected)
		})
	}
}

func expectedList() List {
	return List{
		{
			Area:                 "ΑΙΤΩΛΟΑΚΑΡΝΑΝΙΑΣ",
			AreaID:               701,
			DailyShot1:           665,
			DailyShot2:           171,
			DayDiff:              118,
			DayTotal:             836,
			ReferenceDate:        "2021-04-22T00:00:00",
			TotalDistinctPersons: 27193,
			TotalShot1:           27193,
			TotalShot2:           7111,
			TotalVaccinations:    34304,
		},
		{
			Area:                 "ΑΝΑΤΟΛΙΚΗΣ ΑΤΤΙΚΗΣ",
			AreaID:               901,
			DailyShot1:           1054,
			DailyShot2:           107,
			DayDiff:              -39,
			DayTotal:             1161,
			ReferenceDate:        "2021-04-22T00:00:00",
			TotalDistinctPersons: 51872,
			TotalShot1:           51872,
			TotalShot2:           17531,
			TotalVaccinations:    69403,
		},
	}
}
