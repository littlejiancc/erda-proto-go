// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/erda-project/erda-proto-go/oap/common/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *EntityRow) Validate() error {
	if this.RowData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RowData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RowData", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
