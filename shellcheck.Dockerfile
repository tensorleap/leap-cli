# Use Alpine Linux for a minimal image
FROM alpine:latest

# Install shellcheck
RUN apk add --no-cache shellcheck

# Set the working directory inside the container
WORKDIR /repo

# By default, run shellcheck on all shell scripts in the working directory
CMD find . -type f -name "*.sh" -exec shellcheck {} +
