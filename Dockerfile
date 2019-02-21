FROM golang:1.10
ADD . .
CMD './start.sh'
EXPOSE 9000
