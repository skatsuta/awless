package ast

type Entity string

var entities = map[Entity]struct{}{
	"none": {},

	"accesskey":           {},
	"alarm":               {},
	"scalinggroup":        {},
	"bucket":              {},
	"database":            {},
	"dbsubnetgroup":       {},
	"elasticip":           {},
	"function":            {},
	"group":               {},
	"instance":            {},
	"image":               {},
	"internetgateway":     {},
	"keypair":             {},
	"launchconfiguration": {},
	"listener":            {},
	"loadbalancer":        {},
	"policy":              {},
	"queue":               {},
	"record":              {},
	"role":                {},
	"route":               {},
	"routetable":          {},
	"s3object":            {},
	"scalingpolicy":       {},
	"securitygroup":       {},
	"snapshot":            {},
	"subnet":              {},
	"subscription":        {},
	"tag":                 {},
	"targetgroup":         {},
	"topic":               {},
	"user":                {},
	"volume":              {},
	"vpc":                 {},
	"zone":                {},
}

func IsInvalidEntity(s string) bool {
	_, ok := entities[Entity(s)]
	return !ok
}
