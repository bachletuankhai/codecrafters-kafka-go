package response

import (
	"encoding/binary"
	"net"
)


type ResponseHeader struct {
	CorrelationID int32
}

func (h ResponseHeader) Size() int32 {
	return 4
}

type Response struct {
	conn net.Conn

	Size   int32
	Header ResponseHeader
	Body   []byte
}

func NewResponse(conn net.Conn) *Response {
	return &Response{
		conn: conn,
		Size: 0,
		Body: []byte{},
		Header: ResponseHeader{
			CorrelationID: 7,
		},
	}
}

func (r *Response) Write(header *ResponseHeader, body []byte) error {
	if header != nil {
		r.Header = *header
	}
	if body != nil {
		r.Body = append(r.Body, body...)
	}

	err := binary.Write(r.conn, binary.BigEndian, r.Size)
	if err != nil {
		return err
	}
	err = binary.Write(r.conn, binary.BigEndian, r.Header.CorrelationID)
	if err != nil {
		return err
	}
	
	return nil
}