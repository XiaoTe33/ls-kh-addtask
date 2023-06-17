FROM busybox

EXPOSE 3306

COPY ./ls-kh-addtask /
COPY ./conf.yaml /etc/

ENTRYPOINT ["/ls-kh-addtask"]