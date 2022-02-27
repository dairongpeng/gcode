FROM alpine:latest
RUN mkdir -p /home/app/
ADD ./echo-go /home/app/
CMD /home/app/echo-go -host 0.0.0.0 -port 9090