package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/clbanning/mxj/v2"
)

const (
	defaultPrefix = ""
	defaultIndent = "  "
)

type xmlNode struct {
	Attr     []xml.Attr
	XMLName  xml.Name
	Children []xmlNode `xml:",any"`
	Text     string    `xml:",chardata"`
}

var (
	errorLog *log.Logger
	infoLog  *log.Logger
)

func main() {
	// input := []byte(`{"data":[{"type":"articles","id":"1","attributes":{"title":"JSON:API paints my bikeshed!","body":"The shortest article. Ever."},"relationships":{"author":{"data":{"id":"42","type":"people"}}}}],"included":[{"type":"people","id":"42","attributes":{"name":"John"}}]}`)
	// input := []byte("<1xml><test>blah</test></xml>")
	// input := []byte(`<book> <author>Fred</author>
	// <!--
	// <price>20</price><currency>USD</currency>
	// -->
	//  <isbn>23456</isbn> </book>`)

	initializeLogging()

	input, errRead := io.ReadAll(os.Stdin)
	if errRead != nil {
		errorLog.Fatal(errRead)
	}

	var out []byte
	var err error

	switch string(input[0:1]) {
	case "{", "[":
		out, err = formatJSON(input)
	case "<":
		out, err = formatXML2(input)
	default:
		err = errors.New("unsupported content")
	}

	if err != nil {
		errorLog.Fatal(err)
		return
	}

	infoLog.Println(string(out))
}

func formatJSON(input []byte) ([]byte, error) {
	var out bytes.Buffer
	if !json.Valid((input)) {
		return nil, errors.New("invalid json")
	}

	err := json.Indent(&out, input, defaultPrefix, defaultIndent)
	if err != nil {
		return nil, fmt.Errorf("failed to indent json, error: %w", err)
	}

	return out.Bytes(), nil
}

func formatXML(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	buf := new(bytes.Buffer)
	encoder := xml.NewEncoder(buf)
	encoder.Indent(defaultPrefix, defaultIndent)

tokenize:
	for {
		token, err := decoder.Token()

		switch {
		case err == io.EOF:
			encoder.Flush()
			break tokenize
		case err != nil:
			return nil, fmt.Errorf("failed to tokenize xml, error: %w", err)
		}

		errEncode := encoder.EncodeToken(token)
		if errEncode != nil {
			fmt.Printf("failed to encode xml, error: %s\n", errEncode)
		}
	}

	return buf.Bytes(), nil
}

func formatXML2(input []byte) ([]byte, error) {
	return mxj.BeautifyXml(input, defaultPrefix, defaultIndent)
}

func initializeLogging() {
	infoLog = log.New(os.Stdout, "", 0)
	errorLog = log.New(os.Stderr, "error: ", 0)
}
