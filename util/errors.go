package util

import "errors"

var ErrNoMappingFound = errors.New("cloud product mapping not found")
var ErrNoMapperFound = errors.New("cloud mapper not found")
var ErrInvalidImpactMode = errors.New("invalid impact mode")
