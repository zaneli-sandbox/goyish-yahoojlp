package yahoojlp

import (
	"fmt"
	"strings"
)

type Service struct {
	client *Client
	errors []error
	name   string
	params map[string]string
}

func (s *Service) getError() error {
	switch len(s.errors) {
	case 0:
		return nil
	case 1:
		return s.errors[0]
	default:
		return newMultiErrors(s.errors...)
	}
}

type MAParseResultSet struct {
	MAResult   MAResult   `xml:"ma_result"`
	UniqResult UniqResult `xml:"uniq_result"`
}

func (rs MAParseResultSet) string() string {
	ss := []string{}
	if rs.MAResult.TotalCount > 0 {
		ss = append(ss, fmt.Sprintf("ma_result=%s", rs.MAResult.String()))
	}
	if rs.UniqResult.TotalCount > 0 {
		ss = append(ss, fmt.Sprintf("uniq_result=%s", rs.UniqResult.String()))
	}
	return strings.Join(ss, ", ")
}

func (rs MAParseResultSet) String() string {
	return fmt.Sprintf("{%s}", rs.string())
}

type MAResult struct {
	TotalCount    int      `xml:"total_count"`
	FilteredCount int      `xml:"filtered_count"`
	WordList      WordList `xml:"word_list"`
}

func (r MAResult) string() string {
	return fmt.Sprintf(
		"total_count=%d, filtered_count=%d, word_list=%s",
		r.TotalCount, r.FilteredCount, r.WordList.String())
}

func (r MAResult) String() string {
	return fmt.Sprintf("{%s}", r.string())
}

type UniqResult struct {
	MAResult
	WordList WordWithCountList `xml:"word_list"`
}

func (r UniqResult) string() string {
	return fmt.Sprintf(
		"total_count=%d, filtered_count=%d, word_list=%s",
		r.TotalCount, r.FilteredCount, r.WordList.String())
}

func (r UniqResult) String() string {
	return fmt.Sprintf("{%s}", r.string())
}

type WordList struct {
	Words []Word `xml:"word"`
}

func (ws WordList) string() string {
	ss := []string{}
	for _, w := range ws.Words {
		ss = append(ss, w.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ", "))
}

func (ws WordList) String() string {
	return fmt.Sprintf("{%s}", ws.string())
}

type WordWithCountList struct {
	Words []WordWithCount `xml:"word"`
}

func (ws WordWithCountList) string() string {
	ss := []string{}
	for _, w := range ws.Words {
		ss = append(ss, w.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ", "))
}

func (ws WordWithCountList) String() string {
	return fmt.Sprintf("{%s}", ws.string())
}

type Word struct {
	Surface  string `xml:"surface"`
	Reading  string `xml:"reading"`
	Pos      string `xml:"pos"`
	Baseform string `xml:"baseform"`
	Feature  string `xml:"feature"`
}

func (w Word) string() string {
	ws := []string{}
	if w.Surface != "" {
		ws = append(ws, fmt.Sprintf("surface=%s", w.Surface))
	}
	if w.Reading != "" {
		ws = append(ws, fmt.Sprintf("reading=%s", w.Reading))
	}
	if w.Pos != "" {
		ws = append(ws, fmt.Sprintf("pos=%s", w.Pos))
	}
	if w.Baseform != "" {
		ws = append(ws, fmt.Sprintf("baseform=%s", w.Baseform))
	}
	if w.Feature != "" {
		ws = append(ws, fmt.Sprintf("feature=%s", w.Feature))
	}
	return strings.Join(ws, ", ")
}

func (w Word) String() string {
	return fmt.Sprintf("{%s}", w.string())
}

type WordWithCount struct {
	Word
	Count int `xml:"count"`
}

func (w WordWithCount) string() string {
	return fmt.Sprintf("%s, count=%d", w.Word.string(), w.Count)
}

func (w WordWithCount) String() string {
	return fmt.Sprintf("{%s}", w.string())
}

type InvalidArgumentError struct {
	name  string
	value string
}

func newInvalidArgumentError(name string, value string) InvalidArgumentError {
	return InvalidArgumentError{name, value}
}

func (e InvalidArgumentError) Error() string {
	return fmt.Sprintf("Invalid value: %s=%s", e.name, e.value)
}

type MultiErrors struct {
	errors []error
}

func newMultiErrors(errors ...error) MultiErrors {
	return MultiErrors{errors}
}

func (m MultiErrors) Error() string {
	xs := make([]string, len(m.errors))
	for i, e := range m.errors {
		xs[i] = e.Error()
	}
	return strings.Join(xs, ", ")
}
