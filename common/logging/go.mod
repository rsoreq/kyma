module github.com/kyma-project/kyma/common/logging

go 1.18

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/go-logr/zapr v0.4.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.9.0
	go.uber.org/zap v1.27.0
	k8s.io/klog/v2 v2.130.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	golang.org/x/crypto => golang.org/x/crypto v0.24.0
	golang.org/x/net => golang.org/x/net v0.26.0
	golang.org/x/text => golang.org/x/text v0.16.0
	golang.org/x/tools => golang.org/x/tools v0.22.0
)
