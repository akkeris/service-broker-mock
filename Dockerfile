FROM golang:1.10
ADD . .
RUN go get "github.com/go-martini/martini"
RUN go get "github.com/martini-contrib/render"

CMD './start.sh'
EXPOSE 9000
