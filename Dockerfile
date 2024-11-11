# Etapa 1: Construir el binario
FROM golang:1.20-alpine AS build

WORKDIR /app

# Copiar el código fuente
COPY . .

# Compilar el binario
RUN go mod init proyectogo && go build -o proyectogo

# Etapa 2: Ejecutar la aplicación
FROM alpine

WORKDIR /app

# Copiar el binario compilado y las plantillas
COPY --from=build /app/proyectogo /app/proyectogo
COPY --from=build /app/templates /app/templates

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/app/proyectogo"]
