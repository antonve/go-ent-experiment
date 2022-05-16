FROM golang:1.18

RUN go install entgo.io/ent/cmd/ent@latest

CMD ["/bin/bash"]
