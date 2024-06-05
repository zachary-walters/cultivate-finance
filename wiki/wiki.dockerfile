FROM --platform=linux/amd64 bitnami/dokuwiki:latest

## Change user to perform privileged actions
USER 0
RUN install_packages vim 
RUN install_packages zip
RUN install_packages unzip
RUN install_packages bash
## Revert to the original non-root user
USER 1001

## Enable mod_ratelimit module
RUN sed -i -r 's/#LoadModule ratelimit_module/LoadModule ratelimit_module/' /opt/bitnami/apache/conf/httpd.conf

ENV APACHE_HTTP_PORT_NUMBER=8181
ENV APACHE_HTTPS_PORT_NUMBER=8143
EXPOSE 8181 8143

COPY cfbackups_prd /cfbackups_prd
COPY scripts/prodrestore.sh /prodrestore.sh