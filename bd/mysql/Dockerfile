FROM mysql

ENV MYSQL_ROOT_PASSWORD 123  
ENV MYSQL_DATABASE test  
ENV MYSQL_USER users_service  
ENV MYSQL_PASSWORD 123

ADD setup.sql /docker-entrypoint-initdb.d
