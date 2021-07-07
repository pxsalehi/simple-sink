FROM alpine
COPY ./simple-sink /
RUN chmod +x /simple-sink
EXPOSE 8080
ENTRYPOINT ["/simple-sink"]
