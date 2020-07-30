
###################################
#Build stage
FROM golang:1.14-alpine AS build-env

ARG VERSION

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/TZ /etc/localtime && echo TZ > /etc/timezone

#Build deps
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add build-base git

#Setup repo
COPY . /goStudy
WORKDIR /goStudy

#Checkout version if set
RUN if [ -n "${VERSION}" ]; then git checkout "${VERSION}"; fi \
 && go build -mod=vendor -a -ldflags='-linkmode external -extldflags "-static" -s -w -X main.Version=${VERSION}'

EXPOSE 30000

ENV TZ=Asia/Shanghai
VOLUME ["/app/goStudy/data"]

ENTRYPOINT ["/app/goStudy/goStudy"]

COPY --from=build-env /case_collector/case_collector /app/case_collector/case_collector
COPY --from=build-env /etc/timezone /etc/timezone
