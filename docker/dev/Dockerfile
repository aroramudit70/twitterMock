FROM ubuntu:21.10

MAINTAINER aroramudit70@gmail.com

ADD ./bin/twittermock /

COPY ./configs/ /configs/

USER root
RUN mkdir -p /logs
RUN chmod 777 -R /logs

HEALTHCHECK NONE

ENTRYPOINT ["/twittermock"]
CMD ["-config=configs/dev.env"]