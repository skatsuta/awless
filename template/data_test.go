package template

var DefsExample = map[string]Definition{
	"createinstance": {
		Action:         "create",
		Entity:         "instance",
		Api:            "ec2",
		RequiredParams: []string{"image", "count", "count", "type", "subnet"},
		ExtraParams:    []string{"keypair", "ip", "userdata", "securitygroup", "lock"},
	},
	"createkeypair": {
		Action:         "create",
		Entity:         "keypair",
		Api:            "ec2",
		RequiredParams: []string{"name"},
		ExtraParams:    []string{},
	},
	"createtag": {
		Action:         "create",
		Entity:         "tag",
		Api:            "ec2",
		RequiredParams: []string{"resource", "key", "value"},
		ExtraParams:    []string{},
	},
}
