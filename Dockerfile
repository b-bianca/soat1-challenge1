FROM golang

WORKDIR /
COPY go.mod .
COPY go.sum .

RUN go install github.com/b-bianca/soat1-challenge1@latest
RUN go build -gcflags="all=-N -l" -o /feedme
RUN echo $(ls /go/bin)

FROM gcr.io/distroless/base-debian10
WORKDIR /

EXPOSE 2345

COPY --from=build /go/bin/soat1 /soat1
COPY --from=build /feedme ~/feedme
#ENTRYPOINT [ "/dlv" ]
CMD ["/soat1", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "~/feedme"]
