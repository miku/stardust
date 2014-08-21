package stardust

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

// ColumnSpec contains two column indexes
type ColumnSpec struct {
	left  int
	right int
}

// Record represents a single input (fields) and two highlighted fields, left and right, that are used for comparison
type Record struct {
	left   int
	right  int
	Fields []string
}

// Left returns one of the highlighted rows
func (r *Record) Left() string {
	return r.Fields[r.left]
}

// Right returns another one of the highlighted rows
func (r *Record) Right() string {
	return r.Fields[r.right]
}

// ParseColumnSpec parses a string like "2,3" into a ColumnSpec struct
func ParseColumnSpec(s string) (*ColumnSpec, error) {
	if len(s) == 0 {
		return nil, errors.New("columnspec cannot be empty")
	}
	fields := strings.Split(s, ",")
	if len(fields) != 2 {
		return nil, errors.New("columnspec must be of form 'c1,c2'")
	}
	left, err := strconv.ParseInt(fields[0], 10, 0)
	if err != nil {
		return nil, err
	}
	right, err := strconv.ParseInt(fields[1], 10, 0)
	if err != nil {
		return nil, err
	}

	return &ColumnSpec{left: int(left) - 1, right: int(right) - 1}, nil
}

// RecordGeneratorFileDelimiter will produce pair values, that are extracted according to a given column specification
// and a custom field delimiter
func RecordGeneratorFileDelimiter(reader io.ReadCloser, c *ColumnSpec, delim string) chan *Record {
	records := make(chan *Record)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fields := strings.Split(scanner.Text(), delim)
			if c.left < 0 || c.left >= len(fields) || c.right < 0 || c.right >= len(fields) {
				log.Fatalf("columnspec mismatch: %+v, %+v\n", fields, c)
			}
			records <- &Record{Fields: fields, left: c.left, right: c.right}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		defer close(records)
		if reader != os.Stdin {
			err := reader.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()
	return records
}

// RecordGeneratorFile will produce pair values, that are extracted according to a given column specification
// and tab delimiter.
func RecordGeneratorFile(reader io.ReadCloser, c *ColumnSpec) chan *Record {
	return RecordGeneratorFileDelimiter(reader, c, "\t")
}

// RecordGenerator abstracts from the way the strings are specified, e.g. via
// stdin a file or directly on the command line
func RecordGenerator(c *cli.Context) chan *Record {
	columnSpec, err := ParseColumnSpec(c.GlobalString("f"))
	if err != nil {
		log.Fatal(err)
	}
	// stdin
	if len(c.Args()) == 0 || (len(c.Args()) == 1 && c.Args()[0] == "-") {
		return RecordGeneratorFileDelimiter(os.Stdin, columnSpec, c.GlobalString("delimiter"))
	}
	// from filename
	if len(c.Args()) == 1 {
		filename := c.Args()[0]
		if _, err := os.Stat(filename); err == nil {
			file, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			return RecordGeneratorFile(file, columnSpec)
		}
		log.Fatalf("no such file: %s\n", filename)

	}
	// direct
	if len(c.Args()) == 2 {
		records := make(chan *Record)
		go func() {
			records <- &Record{Fields: c.Args(), left: 0, right: 1}
			close(records)
		}()
		return records
	}
	return nil
}
