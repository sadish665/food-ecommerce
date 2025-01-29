FROM golang:1.23.5-alpine3.21
#define working directory
WORKDIR /app
#copy all the files from here to app directory inside container
COPY . .
#fetch all the dependencies inside container
RUN go get -d -v ./...
#build the application inside container
RUN go build -o api .
#export the port inside container
EXPOSE 8080
#run command when the container starts
CMD [ "./api" ]