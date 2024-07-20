FROM golang:1.22-alpine as builder

## The RUN command is used to run commands as you would in a shell.
RUN mkdir /app

## Here, you are copying the whole content of your current directory (your Go code) to the /app directory.
COPY . /app

# This changes the working directory for any subsequent COPY, CMD, ENTRYPOINT, or RUN instructions
WORKDIR /app

# This runs the go build command which compiles your Go code into a single, statically linked binary
# called small-app.
RUN go build -o small-app .

RUN chmod +x /app/small-app

# build a tiny docker image
#starts a new, much smaller build stage. This helps to keep the final Docker image size down.
# If we didn't do this, the final Docker image would include the entire Go SDK, not just our static binary.
FROM alpine:latest

RUN mkdir /app


# Here, you're copying the compiled binary file small-app from the "builder" stage to the /app directory in the current stage.
COPY --from=builder /app/small-app /app

# This provides the command to be executed when Docker runs your image.
CMD [ "/app/small-app" ]
