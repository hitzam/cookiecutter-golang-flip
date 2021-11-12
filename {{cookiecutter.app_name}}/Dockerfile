FROM registry-intl.ap-southeast-5.aliyuncs.com/flip-default/alpine:latest

# Add packages to set timezone to WIB instead of default UTC value
RUN apk update && \
    apk add --no-cache tzdata

# Set working directory for docker container to /app
WORKDIR /app

# Create directory for storing application logs
RUN mkdir /app/log

# Copy application binary and respective config files to /app
COPY /{{ cookiecutter.app_name }} /app/main

# Run the app
ENTRYPOINT ["/app/main"]
