FROM golang:1.19-alpine
LABEL maintainer="luis.tormo24@gmail.com"

# Creamos la carpeta /app/test
RUN mkdir -p /app/test

# Creamos el usuario test_obj-5 para ejecutar los test desde el modo usuario
# Añadimos los permisos al nuevo usario en /app/test
RUN adduser --disabled-password test_obj-5 \
	&& chown -R test_obj-5:test_obj-5 /app/

RUN apk add build-base

WORKDIR /app/test

# Cambio al nuevo usuario
USER test_obj-5

COPY go.mod /app/test
COPY go.sum /app/test 

RUN go mod download

# Modificamos la variable global GOCACHE ya que si no, los test se ejecutan en .cache
# y el nuevo usario tiene el acceso denegado
ENV GOCACHE=/tmp/


ENTRYPOINT ["make", "test"]
