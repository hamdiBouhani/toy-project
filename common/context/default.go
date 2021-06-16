package context

//DefaultService that should be extended for each service. handles the internal context routing.
type DefaultService struct {
	ctx *Context
}

//Configure Base that will be called for EVERY service extending default service.
func (ds *DefaultService) Configure(ctx *Context) error {
	ds.ctx = ctx

	return nil
}

// Start the service
func (ds *DefaultService) Start() error {
	return nil
}

//Shutdown Perform any shutdown procedures
func (ds *DefaultService) Shutdown() {
	//
}

//Service Helper function to return the inner context of default service
func (ds *DefaultService) Service(id string) Service {
	return ds.ctx.Service(id)
}

//Service Helper function to return the inner context list of service keys
func (ds *DefaultService) Services() []string {
	return ds.ctx.Services()
}
