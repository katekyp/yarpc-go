// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package atomic

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// ReadOnlyStore_Integer_Args represents the arguments for the ReadOnlyStore.integer function.
//
// The arguments for integer are sent and received over the wire as this struct.
type ReadOnlyStore_Integer_Args struct {
	Key *string `json:"key,omitempty"`
}

// ToWire translates a ReadOnlyStore_Integer_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ReadOnlyStore_Integer_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ReadOnlyStore_Integer_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ReadOnlyStore_Integer_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ReadOnlyStore_Integer_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ReadOnlyStore_Integer_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ReadOnlyStore_Integer_Args
// struct.
func (v *ReadOnlyStore_Integer_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}

	return fmt.Sprintf("ReadOnlyStore_Integer_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ReadOnlyStore_Integer_Args match the
// provided ReadOnlyStore_Integer_Args.
//
// This function performs a deep comparison.
func (v *ReadOnlyStore_Integer_Args) Equals(rhs *ReadOnlyStore_Integer_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Key, rhs.Key) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ReadOnlyStore_Integer_Args.
func (v *ReadOnlyStore_Integer_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", *v.Key)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *ReadOnlyStore_Integer_Args) GetKey() (o string) {
	if v != nil && v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *ReadOnlyStore_Integer_Args) IsSetKey() bool {
	return v != nil && v.Key != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "integer" for this struct.
func (v *ReadOnlyStore_Integer_Args) MethodName() string {
	return "integer"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ReadOnlyStore_Integer_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ReadOnlyStore_Integer_Helper provides functions that aid in handling the
// parameters and return values of the ReadOnlyStore.integer
// function.
var ReadOnlyStore_Integer_Helper = struct {
	// Args accepts the parameters of integer in-order and returns
	// the arguments struct for the function.
	Args func(
		key *string,
	) *ReadOnlyStore_Integer_Args

	// IsException returns true if the given error can be thrown
	// by integer.
	//
	// An error can be thrown by integer only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for integer
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// integer into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by integer
	//
	//   value, err := integer(args)
	//   result, err := ReadOnlyStore_Integer_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from integer: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(int64, error) (*ReadOnlyStore_Integer_Result, error)

	// UnwrapResponse takes the result struct for integer
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if integer threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ReadOnlyStore_Integer_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ReadOnlyStore_Integer_Result) (int64, error)
}{}

func init() {
	ReadOnlyStore_Integer_Helper.Args = func(
		key *string,
	) *ReadOnlyStore_Integer_Args {
		return &ReadOnlyStore_Integer_Args{
			Key: key,
		}
	}

	ReadOnlyStore_Integer_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *KeyDoesNotExist:
			return true
		default:
			return false
		}
	}

	ReadOnlyStore_Integer_Helper.WrapResponse = func(success int64, err error) (*ReadOnlyStore_Integer_Result, error) {
		if err == nil {
			return &ReadOnlyStore_Integer_Result{Success: &success}, nil
		}

		switch e := err.(type) {
		case *KeyDoesNotExist:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ReadOnlyStore_Integer_Result.DoesNotExist")
			}
			return &ReadOnlyStore_Integer_Result{DoesNotExist: e}, nil
		}

		return nil, err
	}
	ReadOnlyStore_Integer_Helper.UnwrapResponse = func(result *ReadOnlyStore_Integer_Result) (success int64, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ReadOnlyStore_Integer_Result represents the result of a ReadOnlyStore.integer function call.
//
// The result of a integer execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ReadOnlyStore_Integer_Result struct {
	// Value returned by integer after a successful execution.
	Success      *int64           `json:"success,omitempty"`
	DoesNotExist *KeyDoesNotExist `json:"doesNotExist,omitempty"`
}

// ToWire translates a ReadOnlyStore_Integer_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ReadOnlyStore_Integer_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueI64(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ReadOnlyStore_Integer_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _KeyDoesNotExist_Read(w wire.Value) (*KeyDoesNotExist, error) {
	var v KeyDoesNotExist
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a ReadOnlyStore_Integer_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ReadOnlyStore_Integer_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ReadOnlyStore_Integer_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ReadOnlyStore_Integer_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _KeyDoesNotExist_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ReadOnlyStore_Integer_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ReadOnlyStore_Integer_Result
// struct.
func (v *ReadOnlyStore_Integer_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}

	return fmt.Sprintf("ReadOnlyStore_Integer_Result{%v}", strings.Join(fields[:i], ", "))
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this ReadOnlyStore_Integer_Result match the
// provided ReadOnlyStore_Integer_Result.
//
// This function performs a deep comparison.
func (v *ReadOnlyStore_Integer_Result) Equals(rhs *ReadOnlyStore_Integer_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_I64_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ReadOnlyStore_Integer_Result.
func (v *ReadOnlyStore_Integer_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddInt64("success", *v.Success)
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ReadOnlyStore_Integer_Result) GetSuccess() (o int64) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *ReadOnlyStore_Integer_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *ReadOnlyStore_Integer_Result) GetDoesNotExist() (o *KeyDoesNotExist) {
	if v != nil && v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// IsSetDoesNotExist returns true if DoesNotExist is not nil.
func (v *ReadOnlyStore_Integer_Result) IsSetDoesNotExist() bool {
	return v != nil && v.DoesNotExist != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "integer" for this struct.
func (v *ReadOnlyStore_Integer_Result) MethodName() string {
	return "integer"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ReadOnlyStore_Integer_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
