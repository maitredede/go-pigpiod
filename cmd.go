package pigpiod

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Cmd struct {
	cmd  uint32
	p1   uint32
	p2   uint32
	p3   uint32
	data []byte
}

var endianess = binary.LittleEndian

func (c Cmd) ExecuteRes(stream io.ReadWriter) (Cmd, error) {

	buf := &bytes.Buffer{}
	if err := binary.Write(buf, endianess, c.cmd); err != nil {
		return c, fmt.Errorf("failed to encode data: %w", err)
	}
	if err := binary.Write(buf, endianess, c.p1); err != nil {
		return c, fmt.Errorf("failed to encode data: %w", err)
	}
	if err := binary.Write(buf, endianess, c.p2); err != nil {
		return c, fmt.Errorf("failed to encode data: %w", err)
	}
	if err := binary.Write(buf, endianess, c.p3); err != nil {
		return c, fmt.Errorf("failed to encode data: %w", err)
	}
	if len(c.data) > 0 {
		if _, err := buf.Write(c.data); err != nil {
			return c, fmt.Errorf("failed to encode data: %w", err)
		}
	}
	bin := buf.Bytes()
	if _, err := stream.Write(bin); err != nil {
		return c, fmt.Errorf("failed to send data: %w", err)
	}

	var response Cmd
	if err := binary.Read(stream, endianess, &response.cmd); err != nil {
		return response, fmt.Errorf("failed to read response cmd: %w", err)
	}
	if err := binary.Read(stream, endianess, &response.p1); err != nil {
		return response, fmt.Errorf("failed to read response p1: %w", err)
	}
	if err := binary.Read(stream, endianess, &response.p2); err != nil {
		return response, fmt.Errorf("failed to read response p2: %w", err)
	}
	res, p3, err := readBoth(stream)
	if err != nil {
		return response, fmt.Errorf("failed to read response p3: %w", err)
	}
	response.p3 = p3
	if res < 0 {
		if err, ok := pigsErrors[res]; ok {
			return response, fmt.Errorf("cmd failed: %w", err)
		} else {
			return response, fmt.Errorf("cmd failed: %v", res)
		}
	}
	return response, nil
}

func readBoth(stream io.Reader) (int32, uint32, error) {
	buff := &bytes.Buffer{}
	tee := io.TeeReader(stream, buff)

	var iVal int32
	var uVal uint32
	if err := binary.Read(tee, endianess, &uVal); err != nil {
		return iVal, uVal, err
	}
	if err := binary.Read(buff, endianess, &iVal); err != nil {
		return iVal, uVal, err
	}
	return iVal, uVal, nil
}

func (c Cmd) ExecuteResData(stream io.ReadWriter) (Cmd, error) {
	response, err := c.ExecuteRes(stream)
	if err != nil {
		return response, err
	}
	if response.p3 > 0 {
		data := make([]byte, response.p3)
		buf := bytes.Buffer{}
		remaining := int(response.p3)
		buf.Grow(remaining)
		for remaining > 0 {
			n, err := stream.Read(data)
			if err != nil {
				return response, fmt.Errorf("failed to read response data: %w", err)
			}
			if _, err := buf.Write(data[:n]); err != nil {
				panic(fmt.Errorf("failed to write to read buffer"))
			}
			remaining = remaining - n
		}
		response.data = buf.Bytes()
	}
	return response, nil
}

func makeData(items ...any) []byte {
	buff := &bytes.Buffer{}
	for _, i := range items {
		if err := binary.Write(buff, endianess, i); err != nil {
			panic(err)
		}
	}
	return buff.Bytes()
}
