FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o checklist ./cmd/morning-check

# Production image
FROM alpine

COPY ./templates /templates
COPY --from=build /app/checklist /checklist

EXPOSE 8080

CMD [ "/checklist" ]
