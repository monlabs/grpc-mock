package pb

import (
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
)

func TestToStruct(t *testing.T) {
	testCases := []struct {
		name string
		in   map[string]interface{}
		want *structpb.Struct
	}{
		{
			name: "empty map",
			in:   map[string]interface{}{},
			want: nil,
		}, {
			name: "",
			in: map[string]interface{}{
				"nil":    nil,
				"number": float64(100),
				"string": "name",
				"bool":   true,
				"list":   []interface{}{"abc", "def"},
				"struct": map[string]interface{}{
					"number": float64(100),
					"string": "name",
				},
			},
			want: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"nil": nil,
					"number": &structpb.Value{
						Kind: &structpb.Value_NumberValue{
							NumberValue: float64(100),
						},
					},
					"string": &structpb.Value{
						Kind: &structpb.Value_StringValue{
							StringValue: "name",
						},
					},
					"bool": &structpb.Value{
						Kind: &structpb.Value_BoolValue{
							BoolValue: true,
						},
					},
					"list": &structpb.Value{
						Kind: &structpb.Value_ListValue{
							ListValue: &structpb.ListValue{
								Values: []*structpb.Value{
									{
										Kind: &structpb.Value_StringValue{
											StringValue: "abc",
										},
									}, {
										Kind: &structpb.Value_StringValue{
											StringValue: "def",
										},
									},
								},
							},
						},
					},
					"struct": &structpb.Value{
						Kind: &structpb.Value_StructValue{
							StructValue: &structpb.Struct{
								Fields: map[string]*structpb.Value{
									"number": &structpb.Value{
										Kind: &structpb.Value_NumberValue{
											NumberValue: float64(100),
										},
									},
									"string": &structpb.Value{
										Kind: &structpb.Value_StringValue{
											StringValue: "name",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		st := ToStruct(tc.in)
		assert.Equal(t, tc.want, st)
	}
}

func TestToMap(t *testing.T) {
	testCases := []struct {
		in   *structpb.Struct
		want map[string]interface{}
	}{
		{
			in: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"nil": &structpb.Value{
						Kind: &structpb.Value_NullValue{},
					},
					"number": &structpb.Value{
						Kind: &structpb.Value_NumberValue{
							NumberValue: float64(100),
						},
					},
					"string": &structpb.Value{
						Kind: &structpb.Value_StringValue{
							StringValue: "name",
						},
					},
					"bool": &structpb.Value{
						Kind: &structpb.Value_BoolValue{
							BoolValue: true,
						},
					},
					"list": &structpb.Value{
						Kind: &structpb.Value_ListValue{
							ListValue: &structpb.ListValue{
								Values: []*structpb.Value{
									{
										Kind: &structpb.Value_StringValue{
											StringValue: "abc",
										},
									}, {
										Kind: &structpb.Value_StringValue{
											StringValue: "def",
										},
									},
								},
							},
						},
					},
					"struct": &structpb.Value{
						Kind: &structpb.Value_StructValue{
							StructValue: &structpb.Struct{
								Fields: map[string]*structpb.Value{
									"number": &structpb.Value{
										Kind: &structpb.Value_NumberValue{
											NumberValue: float64(100),
										},
									},
									"string": &structpb.Value{
										Kind: &structpb.Value_StringValue{
											StringValue: "name",
										},
									},
								},
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"nil":    nil,
				"number": float64(100),
				"string": "name",
				"bool":   true,
				"list":   []interface{}{"abc", "def"},
				"struct": map[string]interface{}{
					"number": float64(100),
					"string": "name",
				},
			},
		},
	}

	for _, tc := range testCases {
		m := ToMap(tc.in)
		assert.Equal(t, tc.want, m)
	}
}
