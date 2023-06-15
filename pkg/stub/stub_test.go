package stub

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubValidate(t *testing.T) {
	testCases := []struct {
		name string
		in   *Stub
		want error
	}{
		{
			name: "missing service",
			in:   &Stub{},
			want: errors.New("missing service"),
		}, {
			name: "missing method",
			in:   &Stub{Service: "srv.fake.tt"},
			want: errors.New("missing method"),
		}, {
			name: "valid no input",
			in: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				Out:     &Output{},
			},
		}, {
			name: "invalid input",
			in: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In:      &Input{},
			},
			want: errors.New("require at least one of equals, contains or matches"),
		}, {
			name: "missing output",
			in: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name": "fake",
					},
				},
			},
			want: errors.New("missing output"),
		}, {
			name: "valid stub",
			in: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name": "fake",
					},
				},
				Out: &Output{},
			},
		},
	}

	for _, tc := range testCases {
		err := tc.in.Validate()
		assert.Equal(t, tc.want, err, tc.name)
	}
}

func TestStubMatch(t *testing.T) {
	testCases := []struct {
		name string
		st   *Stub
		in   map[string]interface{}
		want bool
	}{
		{
			name: "equals-match",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name": "foo",
						"info": map[string]interface{}{
							"no":   100,
							"city": "beijing",
						},
						"hobbies": []string{"climbing", "running"},
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
				"info": map[string]interface{}{
					"city": "beijing",
					"no":   100,
				},
				"hobbies": []string{"climbing", "running"},
			},
			want: true,
		}, {
			name: "equals-mismatch-slice",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name":    "foo",
						"age":     10,
						"hobbies": []string{"climbing", "running"},
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name":    "foo",
				"age":     10,
				"hobbies": []string{"running", "climbing"},
			},
			want: false,
		}, {
			name: "equals-mismatch",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name":    "foo",
						"age":     10,
						"hobbies": []string{"climbing", "running"},
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name":    "foo",
				"age":     20,
				"hobbies": []string{"climbing", "running"},
			},
			want: false,
		}, {
			name: "equals-mismatch",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Equals: map[string]interface{}{
						"name": "foo",
						"age":  10,
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
			},
			want: false,
		}, {
			name: "contains-match",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Contains: map[string]interface{}{
						"name": "foo",
						"age":  10,
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
			},
			want: true,
		}, {
			name: "contains-match",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Contains: map[string]interface{}{
						"name": "foo",
						"age":  10,
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
				"age":  10,
			},
			want: true,
		}, {
			name: "contains-mismatch",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Contains: map[string]interface{}{
						"name": "foo",
						"age":  10,
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "bar",
				"age":  10,
			},
			want: false,
		}, {
			name: "contains-mismatch",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Contains: map[string]interface{}{
						"name": "foo",
						"age":  10,
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name1": "foo",
				"age":   10,
			},
			want: false,
		}, {
			name: "matches-match",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				In: &Input{
					Matches: map[string]interface{}{
						"name": "^foo$",
					},
				},
				Out: &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
			},
			want: true,
		}, {
			name: "no-input-match",
			st: &Stub{
				Service: "srv.fake.tt",
				Method:  "Hello",
				Out:     &Output{},
			},
			in: map[string]interface{}{
				"name": "foo",
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		b := tc.st.Match(tc.in)
		assert.Equal(t, tc.want, b, tc.name)
	}
}
