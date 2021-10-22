FROM registry.access.redhat.com/ubi8-minimal

EXPOSE 8442

RUN microdnf update -y && rm -rf /var/cache/yum && microdnf install git go make -y && microdnf clean all

COPY . /opt/app
WORKDIR /opt/app

RUN make build

CMD ["/opt/app/bin/jump-app-golang-grpc"]