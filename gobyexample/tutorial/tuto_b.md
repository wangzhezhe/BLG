####Dockerfile综述
#Dockerfile模板(用于构建tomcat容器)
#选择所使用的基础镜像
FROM ubuntu:14.04
#添加维护者的信息
MAINTAINER dockerlover@zju.edu.cn
#指定镜像运行时候的用户
USER root
#指定镜像的LABEL信息
LABEL version="1.0"
#在镜像中运行一些基本的指令
RUN apt-get update
RUN apt-get install tomcat7 default-jdk

#在镜像中添加环境变量
ENV CATALINA_HOME=/usr/share/tomcat7
ENV CATALINA_BASE=/var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp

#添对应的指令 在构建镜像的时候运行
RUN mkdir -p $CATALINA_TMPDIR

#为容器添加volumn 对该目录进行持久化操作
VOLUMN ["/var/lib/tomcat7/webapps/"]

#添加容器所暴露的端口
EXPOSE 8080

#指定容器在启动的时候所要运行的命令
CMD ["/usr/share/tomcat7/bin/catalina.sh","run"]

####FROM
FROM ubuntu:14.04

####MAINTAINER
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn

####USER
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root

####LABEL
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"

####ENV
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp

####WORKDIR
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat

####RUN
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk

####COPY
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk

####ADD
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
####VOLUME
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
VOLUMN ["/var/lib/tomcat7/webapps/"]

####ONBUILD
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
VOLUMN ["/var/lib/tomcat7/webapps/"]
EXPOSE 8080

####EXPOSE
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
VOLUMN ["/var/lib/tomcat7/webapps/"]
EXPOSE 8080

####ENTRYPOINT
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
VOLUMN ["/var/lib/tomcat7/webapps/"]
EXPOSE 8080

####CMD
FROM ubuntu:14.04
MAINTAINER dockerlover@zju.edu.cn
USER root
LABEL version="1.0"
ENV CATALINA_HOME /usr/share/tomcat7
ENV CATALINA_BASE /var/lib/tomcat7
ENV CATALINA_PID /var/run/tomcat7.pid
ENV CATALINA_SH /usr/share/tomcat7/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat7-tomcat7-tmp
WORKDIR /usr/local/tomcat
RUN apt-get update
RUN apt-get install tomcat7 default-jdk
ADD ./start.sh /var/lib/tomcat7/webapps/
VOLUMN ["/var/lib/tomcat7/webapps/"]
EXPOSE 8080
CMD ["/usr/share/tomcat7/bin/catalina.sh","run"]

