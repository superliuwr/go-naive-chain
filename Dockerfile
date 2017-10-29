# iron/go is the alpine image with only ca-certificates added
FROM iron/go

WORKDIR /app

ADD ./go-naive-chain /app/

ENTRYPOINT ["./go-naive-chain"]