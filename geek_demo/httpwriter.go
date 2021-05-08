package geek_demo

import (
	"fmt"
	"io"
)

type Header struct {
	Key,Value string
}

type Status struct {
	Code int
	Reason string
}

func WriteResponse1(w io.Writer, st Status, header []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range header {
		_,err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _,err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}


type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error)  {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

// WriteResponse 实现2, 封装了Writer接口
func WriteResponse(w io.Writer, st Status, header []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range header {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}