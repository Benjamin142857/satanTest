package protocol

import "bytes"

type StProtocolType byte
const (
	Byte StProtocolType = iota
	Bool
	Int
	Long
	Float
	Double
	List
	Map
	Struct
)

type StStruct interface {
	WriterDataBuf(bf *StBuffer)
}

type StBuffer struct {
	buf *bytes.Buffer
}

func (bf *StBuffer) WriteTag(tag byte) error {
	return nil
}

func (bf *StBuffer) WriteDataType(tp StProtocolType) error {
	return nil
}

func (bf *StBuffer) WriteDataBuf(tp StProtocolType, d interface{}, sTps ...StProtocolType) error {
	switch tp {
	case Byte:
		_d, ok := d.(byte)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeByte(_d)
	case Bool:
		_d, ok := d.(bool)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeBool(_d)
	case Int:
		_d, ok := d.(int)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeInt(int32(_d))
	case Long:
		_d, ok := d.(int64)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeLong(_d)
	case Float:
		_d, ok := d.(float32)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeFloat(_d)
	case Double:
		_d, ok := d.(float64)
		if !ok {
			// TODO: Error类型错误
		}
		return bf.writeDouble(_d)
	case List:
		_d, ok := d.([]interface{})
		if !ok {
			// TODO: Error类型错误
		}
		for _, e := range _d {
			err := bf.WriteDataBuf(sTps[0], e, sTps[1:]...)
			if err != nil {
				return err
			}
		}
	case Map:
		_d, ok := d.(map[interface{}]interface{})
		if !ok {
			// TODO: Error类型错误
		}
		for k, v := range _d {
			err1 := bf.WriteDataBuf(sTps[0], k)
			if err1 != nil {
				return err1
			}

			err2 := bf.WriteDataBuf(sTps[1], v, sTps[2:]...)
			if err2 != nil {
				return err2
			}
		}
	case Struct:
		_d, ok := d.(StStruct)
		if !ok {
			// TODO: Error类型错误
		}
		_d.WriterDataBuf(bf)
	default:
		// TODO: Error无效编码
	}
	return nil
}

func (bf *StBuffer) writeByte(d byte) error {
	return nil
}

func (bf *StBuffer) writeBool(d bool) error {
	return nil
}

func (bf *StBuffer) writeInt(d int32) error {
	return nil
}

func (bf *StBuffer) writeLong(d int64) error {
	return nil
}

func (bf *StBuffer) writeFloat(d float32) error {
	return nil
}

func (bf *StBuffer) writeDouble(d float64) error {
	return nil
}

func (bf *StBuffer) Bytes() []byte {
	return bf.buf.Bytes()
}