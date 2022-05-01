package utils

import "regexp"

var locationRegex = regexp.MustCompile(`^(\w)-(\w)-(\w)$`)

type LocationSymbolParts struct {
	Region   string
	Island   string
	Location string
}

func ParseLocationSymbol(locationSymbol string) LocationSymbolParts {
	parts := locationRegex.FindStringSubmatch(locationSymbol)
	return LocationSymbolParts{
		Region:   parts[0],
		Island:   parts[0] + "-" + parts[1],
		Location: parts[0] + "-" + parts[1] + "-" + parts[2],
	}
}
