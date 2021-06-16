package context

// Service interface defines what each service needs to expose to be included
// within the service context.
// Id is the UNIQUE id of the given service used for internal context discovery.
// Configure is used to handle any non-service orientated tasks such as parsing config/env required to start
// Start is used to start the physical service. Any blocking services should be moved to a worker to not block
// Shutdown is called upon SIG to graciously close down each service.
type Service interface {
	Id() string
	Configure(ctx *Context) error
	Start() error
	Shutdown()
}
