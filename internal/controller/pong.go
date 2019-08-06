package controller

//NOTE: keep controllers clean with just modular calls for steps
//put all logic in helper methods
func (c *Controller) pong(request PingRequest) *PingResponse {
	response := PingResponse{}

	//1. process pong's response message
	response.Message = c.panicProcessPong(request.Message)

	/* 2. for example call graphql -> call helper method
	*  3. for example store new values -> call helper method
	*  ...
	*  N. update DB -> call helper
	 */

	return &response
}
