package api

import (
	"github.com/monlabs/grpc-mock/pkg/stub"
	pbutils "github.com/monlabs/grpc-mock/pkg/utils/pb"
	mockpb "github.com/monlabs/grpc-mock/proto/mock"
)

func PBStubsToStubs(pbstubs []*mockpb.Stub) ([]*stub.Stub, error) {
	var stubs []*stub.Stub
	for _, pbst := range pbstubs {
		st := PBStubToStub(pbst)
		if err := st.Validate(); err != nil {
			return nil, err
		}
		stubs = append(stubs, st)
	}
	return stubs, nil
}

func PBStubToStub(pbstub *mockpb.Stub) *stub.Stub {
	return &stub.Stub{
		Service: pbstub.Service,
		Method:  pbstub.Method,
		In:      PBInputToInput(pbstub.In),
		Out:     PBOutputToOutput(pbstub.Out),
	}
}

func PBInputToInput(pbin *mockpb.Input) *stub.Input {
	if pbin == nil {
		return nil
	}

	in := &stub.Input{}
	equals := pbin.GetEquals()
	if equals != nil {
		in.Equals = pbutils.ToMap(equals)
	}

	contains := pbin.GetContains()
	if contains != nil {
		in.Contains = pbutils.ToMap(contains)
	}

	matches := pbin.GetMatches()
	if matches != nil {
		in.Matches = pbutils.ToMap(matches)
	}
	return in
}

func PBOutputToOutput(pbout *mockpb.Output) *stub.Output {
	out := &stub.Output{}
	if pbout == nil {
		return out
	}

	out.Data = pbutils.ToMap(pbout.Data)
	out.Code = pbout.Code
	out.Message = pbout.Error
	return out
}

func StubsToPBStubs(stubs []*stub.Stub) []*mockpb.Stub {
	var pbstubs []*mockpb.Stub
	for _, st := range stubs {
		pbstubs = append(pbstubs, StubToPBStub(st))
	}
	return pbstubs
}

func StubToPBStub(stub *stub.Stub) *mockpb.Stub {
	pbst := &mockpb.Stub{
		Service: stub.Service,
		Method:  stub.Method,
		In:      InputToPBInput(stub.In),
		Out:     OutputToPBOutput(stub.Out),
	}
	return pbst
}

func InputToPBInput(in *stub.Input) *mockpb.Input {
	pbIn := &mockpb.Input{}
	if len(in.Equals) != 0 {
		pbIn.Rule = &mockpb.Input_Equals{
			Equals: pbutils.ToStruct(in.Equals),
		}
	}

	if len(in.Contains) != 0 {
		pbIn.Rule = &mockpb.Input_Contains{
			Contains: pbutils.ToStruct(in.Contains),
		}
	}

	if len(in.Matches) != 0 {
		pbIn.Rule = &mockpb.Input_Matches{
			Matches: pbutils.ToStruct(in.Matches),
		}
	}

	return pbIn
}

func OutputToPBOutput(out *stub.Output) *mockpb.Output {
	pbOut := &mockpb.Output{
		Data:  pbutils.ToStruct(out.Data),
		Code:  out.Code,
		Error: out.Message,
	}
	return pbOut
}
