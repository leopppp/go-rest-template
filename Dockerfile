FROM golang:1.10 
RUN mkdir /app 
ADD ./cmd/rest /app/ 
WORKDIR /app 
RUN go build -o rest . 
EXPOSE 8080
CMD [ "/app/rest" ]