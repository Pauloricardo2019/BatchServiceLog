package service

import (
	"batch-service/internal/model"
	"errors"
	"fmt"
	"regexp"
)

type regexService struct {
}

func NewRegexService() *regexService {
	return &regexService{}
}

func (r *regexService) SeparateByGroups(regex *regexp.Regexp, logRow string) (*model.Log, error) {
	webLog := &model.Log{}

	matches := regex.FindStringSubmatch(logRow)
	if matches == nil {
		return nil, errors.New("no match regex")
	}

	//Procura por todos os grupos que incrementei na minha regex; e retorna os grupos
	groups := regex.SubexpNames()
	fmt.Println(matches, groups)

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	ip := r.getMatchedValueByIdentifier(regex, "IPV4", matches, groups)
	if ip == "" {
		ip = r.getMatchedValueByIdentifier(regex, "IPV6", matches, groups)
	}

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	date := r.getMatchedValueByIdentifier(regex, "date", matches, groups)

	//Entra na função e pega qual valor de determinado grupo eu quero pegar
	verb := r.getMatchedValueByIdentifier(regex, "Verb", matches, groups)

	webLog.IP = ip
	webLog.Date = date
	webLog.Verb = verb

	return webLog, nil
}

func (r *regexService) getMatchedValueByIdentifier(regex *regexp.Regexp, id string, matches []string, groups []string) string {
	for _, v := range groups {
		if v == id {
			idx := regex.SubexpIndex(v)
			return matches[idx]
		}
	}
	return ""
}
