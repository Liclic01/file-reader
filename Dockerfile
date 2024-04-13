FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend
COPY go.mod go.sum ./
COPY . .
RUN go mod vendor

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so ./cmd/app/main.go

FROM registry.heroiclabs.com/heroiclabs/nakama:3.21.1

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY local.yml /nakama/data/
