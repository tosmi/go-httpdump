FROM scratch

COPY ./httpdump /bin/httpdump

ENTRYPOINT ["/bin/httpdump"]
