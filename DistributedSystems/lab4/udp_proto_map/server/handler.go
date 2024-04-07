package main

import (
	"errors"
	"log"

	"github.com/ssyrota/university/DistributedSystems/lab4/proto_rpc/client/map_proto"
	"google.golang.org/protobuf/proto"
)

func handler(data []byte) ([]byte, error) {
	log.Printf("received msg: %s", string(data))
	h := Handler{data: data}
	return h.Handle()

}

type Handler struct {
	data []byte
}

func (h *Handler) Handle() ([]byte, error) {
	data, err, ok := h.tryGet()
	if ok {
		return data, err
	}
	data, err, ok = h.trySet()
	if ok {
		return data, err
	}
	return nil, errors.New("unknown request")
}

func (h *Handler) tryGet() ([]byte, error, bool) {
	var get map_proto.MapGetRequest
	if err := proto.Unmarshal(h.data, &get); err == nil {
		res := map_proto.MapGetResponse{Values: make([]*map_proto.MapValue, 0)}
		for _, k := range get.GetKeys() {
			v, ok := ma.Get(k)
			if !ok {
				res.Values = append(res.Values, &map_proto.MapValue{Key: k, Value: nil})
				continue
			}
			res.Values = append(res.Values, &map_proto.MapValue{Key: k, Value: &v})
		}
		d, err := proto.Marshal(&res)
		return d, err, true
	}
	return nil, nil, false
}
func (h *Handler) trySet() ([]byte, error, bool) {
	var set map_proto.MapSetRequest
	if err := proto.Unmarshal(h.data, &set); err == nil {
		res := map_proto.MapSetResponse{}
		for _, value := range set.Values {
			ma.Set(value.GetKey(), value.GetValue())
		}
		d, err := proto.Marshal(&res)
		return d, err, true
	}
	return nil, nil, false
}
