FROM golang:1.21-alpine

# define work directory
WORKDIR /app

# install packages
COPY go.mod go.sum ./
RUN go mod download

# collect all required files
COPY . .

# compile app to binary file
RUN chmod +x build.sh
RUN rm -rf ./dist/smtp-artha && sh build.sh && go build -o ./dist/smtp-artha

# serve the app
RUN chmod +x production-serve.sh

CMD ["sh", "production-serve.sh"]
