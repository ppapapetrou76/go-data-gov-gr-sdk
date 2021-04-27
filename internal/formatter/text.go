package formatter

import (
	"fmt"
	"io"
	"reflect"
	"text/tabwriter"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/util/numbers"
)

const textName = "text"

// Text is the formatter to format data into readable text.
type Text struct {
	writer io.Writer
}

// NewText acts as the factory for formatter.Text.
func NewText(writer io.Writer) *Text {
	return &Text{
		writer: writer,
	}
}

// Format formats data as simple text output.
func (f *Text) Format(data interface{}) error {
	t := reflect.TypeOf(data).Kind()
	//nolint:exhaustive //covered by default switch
	switch t {
	case reflect.Slice:
		v := reflect.ValueOf(data)
		for i := 0; i < v.Len(); i++ {
			if i == 0 {
				d := v.Index(i).Interface()
				if reflect.TypeOf(d).Kind() == reflect.Struct {
					w := tabwriter.NewWriter(f.writer, 15, 0, 0, ' ', tabwriter.TabIndent)

					v := reflect.ValueOf(d)
					typeOfS := v.Type()
					for i := 0; i < v.NumField(); i++ {
						fieldName := typeOfS.Field(i).Name
						min := numbers.Min(len(fieldName), 14)
						fmt.Fprintf(w, "%s\t", typeOfS.Field(i).Name[0:min])
					}
					fmt.Fprintln(w)
					w.Flush()
				}
			}
			if err := f.Format(v.Index(i).Interface()); err != nil {
				return err
			}
		}
	case reflect.Struct:
		w := tabwriter.NewWriter(f.writer, 15, 1, 0, ' ', tabwriter.TabIndent)
		v := reflect.ValueOf(data)
		for i := 0; i < v.NumField(); i++ {
			fmt.Fprintf(w, "%v\t", v.Field(i).Interface())
		}
		fmt.Fprintln(w)
		w.Flush()
	default:

		return fmt.Errorf("not supported type %+v", t)
	}

	return nil
}

// Name obtains the name of the formatter.
func (f *Text) Name() string {
	return textName
}
