package common

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func StdinInt() (int, error) {
	var (
		in string
		e error
		num int64
	)
	reader := bufio.NewReader(os.Stdin)
	if in, e = reader.ReadString('\n'); e != nil {
		return 0, e
	}
	in = strings.TrimSpace(in)
	if num, e = strconv.ParseInt(in, 10, 0); e != nil {
		return 0, e
	}
	return int(num), nil
}

func StdinInt64() (int64, error) {
	var (
		in string
		e error
		num int64
	)
	reader := bufio.NewReader(os.Stdin)
	if in, e = reader.ReadString('\n'); e != nil {
		return 0, e
	}
	in = strings.TrimSpace(in)
	if num, e = strconv.ParseInt(in, 10, 0); e != nil {
		return 0, e
	}
	return num, nil
}

func StdinString() (string, error) {
	var (
		in string
		e error
		out string
	)
	reader := bufio.NewReader(os.Stdin)
	if in, e = reader.ReadString('\n'); e != nil {
		return "", e
	}
	out = strings.TrimSpace(in)
	return out, nil
}