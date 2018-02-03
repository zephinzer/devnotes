ARG MYSQL_VERSION="5.7"
FROM mysql:${MYSQL_VERSION}
ENV MYSQL_ROOT_PASSWORD "default_root_password" \
    MYSQL_DATABASE "default_db" \
    MYSQL_USER "default_user" \
    MYSQL_PASSWORD "default_password"
EXPOSE 3306
VOLUME [ "/var/lib/mysql" ]