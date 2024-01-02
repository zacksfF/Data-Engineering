package wine

import (
	"github.com/zacksfF/Data-Engineering/Data-pipeline/internal/api"
	"google.golang.org/protobuf/proto"
)

// loading data ...
func Load(raw []byte) (out proto.Message, err error) {
	var data api.Processed
	if err := proto.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
