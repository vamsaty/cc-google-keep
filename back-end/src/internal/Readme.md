The heart of our application, i.e., all its internal logic, is implemented here.
/internal is not imported into other applications and libraries. 
The code written here is intended solely for internal use within the code base

# Code traversal
1. start with `services/server.go` : implements the HTTP server. It uses services for handling 
   RESTful APIs. Then dig deeper into the method/function calls
2. read `services/svc_iface.go`: Defines the interface for the services.
