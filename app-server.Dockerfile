FROM opensuse/tumbleweed

WORKDIR /boredomboard

RUN zypper refresh && \
    zypper install -y go
    
COPY go.mod go.sum ./
COPY ./projects/boredomboard/ projects/boredomboard/

RUN go mod tidy -e

RUN go build ./projects/boredomboard/cmd/boredomboard/

CMD [ "./boredomboard" ]