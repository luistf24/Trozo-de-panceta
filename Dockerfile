FROM golang:1.19-alpine
LABEL maintainer="luis.tormo24@gmail.com"

# Creamos la carpeta /app/test
RUN mkdir -p /app/test

RUN adduser --disabled-password tests \
	&& chown -R tests:tests /app/ \
	&& apk update \
	&& apk add build-base

WORKDIR /app/test

# Cambio al nuevo usuario
USER tests

# Modificamos la variable global GOCACHE ya que si no, los test se ejecutan en .cache
# y el nuevo usario tiene el acceso denegado
ENV GOCACHE=/tmp/

COPY go.mod /app/test

RUN go mod download

ENTRYPOINT ["make", "test"]
