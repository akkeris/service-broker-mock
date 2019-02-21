FROM quay.octanner.io/base/oct-golang:1.7
ADD . .
CMD './start.sh'
EXPOSE 9000
