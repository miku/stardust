package stardust

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestLeft(t *testing.T) {
	var tests = []struct {
		record Record
		out    string
	}{
		{Record{left: 0, right: 1, Fields: []string{"0", "1"}}, "0"},
		{Record{left: 1, right: 0, Fields: []string{"0", "1"}}, "1"},
	}
	for _, tt := range tests {
		out := tt.record.Left()
		if out != tt.out {
			t.Errorf("%+v.Left() => %s, want: %s", tt.record, out, tt.out)
		}
	}
}

func TestRight(t *testing.T) {
	var tests = []struct {
		record Record
		out    string
	}{
		{Record{left: 0, right: 1, Fields: []string{"0", "1"}}, "1"},
		{Record{left: 1, right: 0, Fields: []string{"0", "1"}}, "0"},
	}
	for _, tt := range tests {
		out := tt.record.Right()
		if out != tt.out {
			t.Errorf("%+v.Right() => %s, want: %s", tt.record, out, tt.out)
		}
	}
}

func TestParseColumnSpec(t *testing.T) {
	var tests = []struct {
		s   string
		out *ColumnSpec
	}{
		{"1,2", &ColumnSpec{left: 0, right: 1}},
	}

	for _, tt := range tests {
		out, _ := ParseColumnSpec(tt.s)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("ParseColumnSpec(%s) => %+v, want: %+v", tt.s, out, tt.out)
		}
	}
}

func TestRecordGeneratorFile(t *testing.T) {
	fixture := "fixtures/strings.tsv"
	if _, err := os.Stat(fixture); os.IsNotExist(err) {
		t.Skipf("skipping since fixture is missing: %s\n", fixture)
		return
	}
	file, err := os.Open(fixture)
	if err != nil {
		t.Errorf("could not open fixture: %v\n", err)
	}
	columnSpec := &ColumnSpec{left: 0, right: 1}
	records := RecordGeneratorFile(file, columnSpec)
	select {
	case r := <-records:
		if len(r.Fields) != 2 {
			t.Errorf("RecordGeneratorFile(%v, %v) => unexpected length of record.fields: %v", file, columnSpec, r)
		}
		if r.left != 0 {
			t.Errorf("RecordGeneratorFile(%v, %v) => unexpected value in record.left: %v", file, columnSpec, r)
		}
		if r.right != 1 {
			t.Errorf("RecordGeneratorFile(%v, %v) => unexpected value in record.field : %v", file, columnSpec, r)
		}
		if r.Left() != "Somewhere over the rainbow" {
			t.Errorf("RecordGeneratorFile(%v, %v) => unexpected record.Left(): %v, got: %v", file, columnSpec, r, r.Left())
		}
		if r.Right() != "Somewhere near Japan" {
			t.Errorf("RecordGeneratorFile(%v, %v) => unexpected record.Right(): %v, got: %v", file, columnSpec, r, r.Right())
		}
	case <-time.After(1 * time.Millisecond):
		t.Errorf("RecordGeneratorFile(%v, %v) => timeout! ", file, columnSpec)
	}
}
