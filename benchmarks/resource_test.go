package benchmarks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	"google.golang.org/protobuf/proto"

	resourcegogo "github.com/bogdandrutu/lazyproto/gogo/resource"
	resourcelp "github.com/bogdandrutu/lazyproto/testproto/resource"
)

func TestOneofPrimitives(t *testing.T) {
	res := &resourcepb.Resource{}
	res.Attributes = append(res.Attributes,
		&commonpb.KeyValue{Key: "int", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_IntValue{IntValue: 5}}},
		&commonpb.KeyValue{Key: "bool", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BoolValue{BoolValue: true}}},
		&commonpb.KeyValue{Key: "double", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 7.3}}},
		&commonpb.KeyValue{Key: "string", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_StringValue{StringValue: "test"}}},
		&commonpb.KeyValue{Key: "bytes", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BytesValue{BytesValue: []byte{0, 1, 2}}}},
	)

	buf, err := proto.Marshal(res)
	assert.NoError(t, err)

	lazyResource := resourcelp.Resource{}
	assert.NoError(t, lazyResource.UnmarshalVT(buf))

	lazyBuf, err := lazyResource.MarshalVT()
	assert.NoError(t, err)

	newRes := resourcepb.Resource{}
	assert.NoError(t, proto.Unmarshal(lazyBuf, &newRes))

	require.Len(t, newRes.Attributes, 5)
	assert.Equal(t, "int", newRes.Attributes[0].Key)
	assert.Equal(t, int64(5), newRes.Attributes[0].GetValue().GetIntValue())

	assert.Equal(t, "bool", newRes.Attributes[1].Key)
	assert.Equal(t, true, newRes.Attributes[1].GetValue().GetBoolValue())

	assert.Equal(t, "double", newRes.Attributes[2].Key)
	assert.Equal(t, 7.3, newRes.Attributes[2].GetValue().GetDoubleValue())

	assert.Equal(t, "string", newRes.Attributes[3].Key)
	assert.Equal(t, "test", newRes.Attributes[3].GetValue().GetStringValue())

	assert.Equal(t, "bytes", newRes.Attributes[4].Key)
	assert.Equal(t, []byte{0, 1, 2}, newRes.Attributes[4].GetValue().GetBytesValue())
}

func BenchmarkMarshalUnmarshal_LazyProto(b *testing.B) {
	buf := getBuffer(b)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lazyResource := resourcelp.Resource{}
		assert.NoError(b, lazyResource.UnmarshalVT(buf))

		lazyBuf, err := lazyResource.MarshalVT()
		assert.NoError(b, err)

		assert.Equal(b, len(buf), len(lazyBuf))
	}
}

func BenchmarkMarshalUnmarshal_Google(b *testing.B) {
	buf := getBuffer(b)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res := resourcepb.Resource{}
		assert.NoError(b, proto.Unmarshal(buf, &res))
		newBuf, err := proto.Marshal(&res)
		assert.NoError(b, err)

		assert.Equal(b, len(buf), len(newBuf))
	}
}

func BenchmarkMarshalUnmarshal_GoGo(b *testing.B) {
	buf := getBuffer(b)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res := resourcegogo.Resource{}
		assert.NoError(b, res.Unmarshal(buf))
		newBuf, err := res.Marshal()
		assert.NoError(b, err)

		assert.Equal(b, len(buf), len(newBuf))
	}
}

func getBuffer(bt testing.TB) []byte {
	res := &resourcepb.Resource{}
	res.Attributes = append(res.Attributes,
		&commonpb.KeyValue{Key: "int", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_IntValue{IntValue: 5}}},
		&commonpb.KeyValue{Key: "bool", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BoolValue{BoolValue: true}}},
		&commonpb.KeyValue{Key: "double", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 7.3}}},
		&commonpb.KeyValue{Key: "string", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_StringValue{StringValue: "test"}}},
		&commonpb.KeyValue{Key: "bytes", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BytesValue{BytesValue: []byte{0, 1, 2}}}},

		&commonpb.KeyValue{Key: "int", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_IntValue{IntValue: 5}}},
		&commonpb.KeyValue{Key: "bool", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BoolValue{BoolValue: true}}},
		&commonpb.KeyValue{Key: "double", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 7.3}}},
		&commonpb.KeyValue{Key: "string", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_StringValue{StringValue: "test"}}},
		&commonpb.KeyValue{Key: "bytes", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BytesValue{BytesValue: []byte{0, 1, 2}}}},

		&commonpb.KeyValue{Key: "int", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_IntValue{IntValue: 5}}},
		&commonpb.KeyValue{Key: "bool", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BoolValue{BoolValue: true}}},
		&commonpb.KeyValue{Key: "double", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 7.3}}},
		&commonpb.KeyValue{Key: "string", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_StringValue{StringValue: "test"}}},
		&commonpb.KeyValue{Key: "bytes", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BytesValue{BytesValue: []byte{0, 1, 2}}}},

		&commonpb.KeyValue{Key: "int", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_IntValue{IntValue: 5}}},
		&commonpb.KeyValue{Key: "bool", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BoolValue{BoolValue: true}}},
		&commonpb.KeyValue{Key: "double", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_DoubleValue{DoubleValue: 7.3}}},
		&commonpb.KeyValue{Key: "string", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_StringValue{StringValue: "test"}}},
		&commonpb.KeyValue{Key: "bytes", Value: &commonpb.AnyValue{Value: &commonpb.AnyValue_BytesValue{BytesValue: []byte{0, 1, 2}}}},
	)

	buf, err := proto.Marshal(res)
	assert.NoError(bt, err)
	return buf
}
