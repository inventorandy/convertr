package adapt

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// M global marshaler
func m() protojson.MarshalOptions {
	return protojson.MarshalOptions{
		AllowPartial:  true,
		UseProtoNames: true,
	}
}

// U global unmarshaler
func u() protojson.UnmarshalOptions {
	return protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
}

// FromProto converts a proto message to a struct using the JSON struct tags
func FromProto(from proto.Message, to interface{}) error {
	// Marshal to JSON
	b, err := m().Marshal(from)
	if err != nil {
		return fmt.Errorf("unable to unmarshal proto object to json: %s", err.Error())
	}

	// Unmarshal into the Out Object
	if err := json.Unmarshal(b, to); err != nil {
		return fmt.Errorf("unable to unmarshal object from json: %s", err.Error())
	}

	// Return nil by default
	return nil
}

// ToProto converts a struct to a proto message using the JSON struct tags
func ToProto(from interface{}, to proto.Message) error {
	// Marshal to JSON
	b, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("unable to marshal object to json: %s", err.Error())
	}

	// Unmarshal into the Out Object
	if err := u().Unmarshal(b, to); err != nil {
		return fmt.Errorf("unable to unmarshal json to proto object: %s", err.Error())
	}

	// Return nil by default
	return nil
}
