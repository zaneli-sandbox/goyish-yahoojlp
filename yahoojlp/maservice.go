package yahoojlp

import (
	"./maservice"
	"strconv"
	"strings"
)

type MAService struct {
	Service
}

func (c *Client) MAService(sentence string, results ...string) *MAService {
	errors := []error{}
	for _, result := range results {
		if !containStr(result, maservice.ResultTypes...) {
			errors = append(errors, newInvalidArgumentError("results", result))
		}
	}

	params := map[string]string{
		"sentence": sentence,
		"results":  strings.Join(results, ","),
	}
	return &MAService{Service{c, errors, "MAService", params}}
}

func (s *MAService) withResponse(paramName string, responses ...string) *MAService {
	for _, response := range responses {
		if !containStr(response, maservice.ResponseTypes...) {
			s.errors = append(s.errors, newInvalidArgumentError(paramName, response))
		}
	}
	s.params[paramName] = strings.Join(responses, ",")
	return s
}

func (s *MAService) WithResponse(paramName string, responses ...string) *MAService {
	return s.withResponse("response", responses...)
}

func (s *MAService) WithMqResponse(paramName string, responses ...string) *MAService {
	return s.withResponse("ma_response", responses...)
}

func (s *MAService) WithUniqResponse(paramName string, responses ...string) *MAService {
	return s.withResponse("uniq_response", responses...)
}

func (s *MAService) withFilter(paramName string, filters ...int) *MAService {
	filterParam := ""
	for i, filter := range filters {
		if !containInt(filter, maservice.FilterTypes...) {
			s.errors = append(s.errors, newInvalidArgumentError(paramName, strconv.Itoa(filter)))
		}
		filterParam += strconv.Itoa(filter)
		if i != len(filters)-1 {
			filterParam += "|"
		}
	}
	s.params[paramName] = filterParam
	return s
}

func (s *MAService) WithFilter(filters ...int) *MAService {
	return s.withFilter("filter", filters...)
}

func (s *MAService) WithMaFilter(filters ...int) *MAService {
	return s.withFilter("ma_filter", filters...)
}

func (s *MAService) WithUniqFilter(filters ...int) *MAService {
	return s.withFilter("uniq_response", filters...)
}

func (s *MAService) WithUniqByBaseform(flag bool) *MAService {
	s.params["uniq_by_baseform"] = strconv.FormatBool(flag)
	return s
}

func (s *MAService) Parse() (*MAParseResultSet, error) {
	if err := s.getError(); err != nil {
		return nil, err
	}

	result := &MAParseResultSet{}
	if err := s.client.callApi(get, s.name, "V1", "parse", s.params, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MAService) String() string {
	return "yahoojlp.MAService"
}
