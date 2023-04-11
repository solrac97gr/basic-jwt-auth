# Basic JWT Auth
I create this repository as a simple example of how to use JWT in a Go Fiber project.

This repository come with a explanation of how to use JWT in a Fiber project in my blog.

This is very basic and this is a little resume of how the project is build.
## Parts
- Config: The config folder use for simulate a packages that extract the config from a file or a database.

- Handlers: The handlers folder is the place where the function of Login and Protected are defined the login will generate the JWT and the Protected will read the information.

- Middlewares: The middlewares folder is the place where the function of JWT is defined.

- Models: The models folder is the place where the struct of the user and login (Request and Response) is defined.

- Repository: The repository folder is the place where the function of find the user is defined this simulate a call to a Database.

Thats all, I hope the project can help you a little bit in your journey of Go developer.