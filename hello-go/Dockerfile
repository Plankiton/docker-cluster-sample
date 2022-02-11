FROM golang

RUN echo Hello, World!

ADD . /server
WORKDIR /server

RUN go mod tidy
RUN go build -o /server/server .

CMD /server/server
