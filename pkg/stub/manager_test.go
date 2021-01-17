package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubCRUD(t *testing.T) {
	type serviceMethod struct {
		service, method string
	}
	type findStubsArgs struct {
		service, method string
		in              map[string]interface{}
		out             []*Stub
	}

	testCases := []struct {
		added   []*Stub
		deleted []serviceMethod
		wants   []findStubsArgs
	}{
		{
			added: []*Stub{
				{
					Service: "greeter",
					Method:  "SayHello",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "foo",
						},
					},
				}, {
					Service: "greeter",
					Method:  "SayHello",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "bar",
						},
					},
				}, {
					Service: "greeter1",
					Method:  "SayHello1",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "foo",
						},
					},
				}, {
					Service: "greeter1",
					Method:  "SayHello1",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "bar",
						},
					},
				},
			},
			wants: []findStubsArgs{
				{
					service: "greeter",
					method:  "SayHello",
					out: []*Stub{
						{
							Service: "greeter",
							Method:  "SayHello",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "foo",
								},
							},
						}, {
							Service: "greeter",
							Method:  "SayHello",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "bar",
								},
							},
						},
					},
				}, {
					service: "greeter",
					method:  "SayHello",
					in: map[string]interface{}{
						"name": "bar",
					},
					out: []*Stub{
						{
							Service: "greeter",
							Method:  "SayHello",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "bar",
								},
							},
						},
					},
				}, {
					service: "greeter1",
					method:  "SayHello1",
					out: []*Stub{
						{
							Service: "greeter1",
							Method:  "SayHello1",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "foo",
								},
							},
						}, {
							Service: "greeter1",
							Method:  "SayHello1",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "bar",
								},
							},
						},
					},
				},
			},
		}, {
			added: []*Stub{
				{
					Service: "greeter",
					Method:  "SayHello",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "foo",
						},
					},
				}, {
					Service: "greeter",
					Method:  "SayHello",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "bar",
						},
					},
				}, {
					Service: "greeter1",
					Method:  "SayHello1",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "foo",
						},
					},
				}, {
					Service: "greeter1",
					Method:  "SayHello1",
					In: &Input{
						Equals: map[string]interface{}{
							"name": "bar",
						},
					},
				},
			},
			deleted: []serviceMethod{
				{
					service: "greeter",
					method:  "SayHello",
				},
			},
			wants: []findStubsArgs{
				{
					service: "greeter",
					method:  "SayHello",
				}, {
					service: "greeter1",
					method:  "SayHello1",
					out: []*Stub{
						{
							Service: "greeter1",
							Method:  "SayHello1",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "foo",
								},
							},
						}, {
							Service: "greeter1",
							Method:  "SayHello1",
							In: &Input{
								Equals: map[string]interface{}{
									"name": "bar",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		mgr := NewManager()
		for _, add := range tc.added {
			err := mgr.AddStub(add)
			assert.NoError(t, err)
		}
		for _, del := range tc.deleted {
			err := mgr.DeleteStub(del.service, del.method)
			assert.NoError(t, err)
		}
		for _, want := range tc.wants {
			stubs := mgr.FindStubs(want.service, want.method, want.in)
			assert.Equal(t, want.out, stubs)
		}
	}
}
