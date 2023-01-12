package service

import (
	"errors"
	"fmt"
	"regexp"
)

type regexService struct {
	*regexp.Regexp
}

func NewRegexService(regex *regexp.Regexp) *regexService {
	return &regexService{
		regex,
	}
}

func (r *regexService) SeparateByGroups(logRow string) (*string, error) {
	matches := r.FindStringSubmatch(logRow)
	if matches == nil {
		return nil, errors.New("no match regex")
	}

	//Procura por todos os grupos que incrementei na minha regex; e retorna os grupos
	groups := r.SubexpNames()
	fmt.Println(matches, groups)

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	ip := r.getMatchedValueByIdentifier("IPV4", matches, groups)
	if ip == "" {
		ip = r.getMatchedValueByIdentifier("IPV6", matches, groups)
	}

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	date := r.getMatchedValueByIdentifier("date", matches, groups)

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	verb := r.getMatchedValueByIdentifier("Verb", matches, groups)

	values := fmt.Sprintf("%s %s %s", ip, date, verb)

	return &values, nil
}

func (r *regexService) getMatchedValueByIdentifier(id string, matches []string, groups []string) string {
	for _, v := range groups {
		if v == id {
			idx := r.SubexpIndex(v)
			return matches[idx]
		}
	}
	return ""
}
