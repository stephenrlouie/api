FROM alpine:3.7
MAINTAINER Megan O'Keefe <meokeefe@cisco.com>

# create working directory
RUN mkdir -p /optikon-api

# set the working directory
WORKDIR /optikon-api

# add binary
COPY bin/optikon-api /bin

# set the entrypoint
ENTRYPOINT ["/bin/optikon-api"]
