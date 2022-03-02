# Fiber Demo

Small golang POC for using Fiber to handle requests. Demo is set up to do the following:

- Serve Open API Spec via Swagger.
- Generate Open API Spec based on convention and using husky.
- Handle validation on for models.
- Demonstrate middleware & logging.
- Make use of existing fiber middleware.
- Demonstrate routing available via fiber.

## Trying Out

In order to run this [husky](https://github.com/typicode/husky) needs to be installed and prepared:

```
npm install husky -D
npm run prepare
```

Additionally we need to install [SwagGo](https://github.com/swaggo/swag):

```
go install github.com/swaggo/swag/cmd/swag@latest
```

If isn't working, try making sure the gopath is correct:
```
export PATH=$(go env GOPATH)/bin:$PATH
```
