package dataframe

import (
	"errors"
)

type Dataframe struct {
	headers []string
	rows    []map[string]*string
}

func (c *Dataframe) Columns() []string {
	return c.headers
}

func (c *Dataframe) Rows() []map[string]*string {
	return c.rows
}

func NewDataframe(records [][]string) (*Dataframe, error) {
	if len(records) <= 1 {
		return nil, errors.New("no data provided for csv")
	}

	df := &Dataframe{
		headers: make([]string, len(records[0])),
		rows:    make([]map[string]*string, len(records)-1),
	}

	df.headers = records[0]
	set := map[string]bool{}

	// make sure no duplicate headers
	for _, header := range df.headers {
		_, ok := set[header]
		if ok {
			return nil, errors.New("duplicate header")
		}
		set[header] = true
	}

	rows := records[1:]
	for idx, row := range rows {
		df.rows[idx] = map[string]*string{}

		for j, data := range row {
			dataValue := data
			headerValue := df.headers[j]

			if data == "" {
				df.rows[idx][headerValue] = nil
			} else {
				df.rows[idx][headerValue] = &dataValue
			}
		}
	}
	return df, nil
}
