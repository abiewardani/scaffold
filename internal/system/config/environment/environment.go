package environment

type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
	Staging     Environment = "staging"
)
