#### SUMMARIZE
如何构建一个自定义的镜像？通常有两种方式，一种是进入运行着的容器中，进行对应的操作之后，通过docker commit命令将容器以镜像的方式保存起来，另一种是创建一个Dockerfile文件，再通过docker build命令来从Dockerfile中构建一个镜像。通常推荐的方式是使用Dockerfile，这样就可以使得镜像的构建过程变得可迁移，对镜像的构建过程进行再调整也会变得更容易。

Dockerfile使用了一种DSL(领域专属语言)，可以在其中指定构建镜像的相关信息。Dockerfile比较清楚的现实了容器修改都执行了哪些相关的指令，采用Dockerfile的方式是可重复，可迁移的，也方便规模化的操作。在Dockerfile中，每个指令都是与相关的命令成对出现的，命令最好用大写字母表示，在Dockerfile中，文件中指令的执行方式是从上到下的。每一个新的指令都会添加一个image layer并且将当前的image提交。


Dockerfiles支持如下语法的命令：

`INSTRUCTION argument`

Dockerfile中的指令不区分大小写。但是，命名约定为全部大写。

关于Dockerfile的更详细的介绍可以参考[官方文档](官方文档：http://docs.docker.com/reference/builder/)


#### FROM

格式：

`FORM <image>`或`FORM<image>:<tag>`

Dockerfile的第一个指令必须为FROM,FROM后面跟的镜像是一个base image，表明后面的操作都应该在这个base image的基础上进行,这里通常是要必须被指明的。所有Dockerfile都必须以FROM命令开始。 FROM命令会指定镜像基于哪个基础镜像创建，接下来的命令也会基于这个基础镜像。FROM命令可以多次使用，表示会创建多个镜像。具体语法如下：`FROM <image name>`
这个指定告诉我们，新的镜像将基于Ubuntu的镜像来构建。

例子：

`FORM ubuntu` 表示使用ubuntu作为基础镜像

####MAINTAINER
格式：

`MAINTAINER <name><mail>`

指定维护者信息。

例子：

`MAINTANINER zju lover@docker.com`

####USER
格式：

`USER daemon`
指定运行容器时的用户名或UID，可以通过该命令指定运行用户。并且可以在之前创建所需要的用户。

####LABEL

格式：

`LABEL <key>=<value> <key>=<value>...`

label指令可以为镜像添加一些元信息，一个`LABEL`标签就是一个key-value的键值对，用户可以对一个镜像指定多个label标记，用户可以用EOL标记将多个换行符区别开。

例子：
<pre>
<code>
LABEL "com.example.vendor"="ACME Incorporated"
LABEL version="1.0"
</code>
</pre>

####ENV
格式：

`ENV <key> <value>`

设置环境变量,会被后续`RUN`指令使用，并在容器运行时保持。它们使用键值对，增加运行程序的灵活性。

例子：
`ENV PATH $GOPATH:$PATH`

####WORKDIR
格式：

`WORKDIR /path/to/workdir`
为后续的RUN、CMD、ENTERPOINT以及ADD指令配置工作目录。用户可以多次使用WORKDIR指令，后续命令如果参数是相对路径，则会基于之前命令指定的路径。


例子：

`WORKDIR /a`
`WORKDIR b`
`WORKDIR c`
则最终路径为/a/b/c

####RUN
格式：

- `RUN <command>`
- `RUN ["executable","param1","param2"]`

采用第一种方式的格式，命令将自动在shell终端中运行命令，即命令运行的时候会在前面自动加上`/bin/sh。 –C`;如果想使用默认的shell方式，则可以使用第二种格式，直接采用`exec`模式执行，比如`RUN ["/bin/bash","-c","echo hello"]`。采用第二种格式，命令会以json的格式进行解析，必须使用""来引用每个命令。
例子：
- `RUN [“apt-get”,”install”,”-y”,”nginx”]`

####COPY

格式：

`COPY <src> <dest>`

复制本地主机的（为Dockerfile所在目录的相对路径，文件或目录）为容器中的 。目标路径不存在时，会自动创建。当使用本地目录为源目录时，推荐使用COPY。

####ADD
格式：

`ADD <src> <dest>`

ADD命令基本功能与COPY命令类似，用来将构建环境下的文件和目录复制到镜像中，复制指定的\<src\>到容器中的<dest> 。与COPY命令不同的是，ADD命令的<src>目录可以为URL，\<src\>可以是Dockerfile所在目录的一个相对路径（文件或目录）；也可以是一个URL；还可以是一个tar文件（自动解压为目录）。比如下面指令`ADD software.lic /opt/application/software.lic`会将构建目录下的`software.lic`文件复制到镜像中的`/opt/application/software.lic`。
如果指令为 `ADD latest.tar.gz /var/www/wordpress/`则`tar.gz`文件会被解开到`/var/www/wordpress/`目录之下。
####VOLUME

格式：

`VOLUME ["/data"]`

volume可以创建本地主机或其他容器挂载的挂载点，在volume命令中声明过的容器目录会以某种形式会绕过容器的AUFS联合文件系统，将容器中的指定目录保存在宿主机上，进行持久化，一般用来存放数据库和需要保持的数据等。比如在ubuntu中，volume的内容具体可以在/var/lib/docker/volume中查询。

例子：

```
FROM ubuntu
RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
VOLUME /myvol
```
####ONBUILD

格式：

`ONBUILD [INSTRUCTION]`
ONBUILD命令会添加一个 trigger 到image中。使用ONBUILD命令的时候有一种层次性的关系再里面，在父镜像构建的时候，使用Dockerfile的ONBUILD命令中指定的内容，会在子镜像构建的时候被执行。如果当前的这个镜像被用作其他镜像的base image那么这个trigger就会被执行，在trigger中可以插入一条新的指令。例如，Dockerfil使用如下的内容创建了镜像test。
```
[...]
ONBUILD ADD . /app/src
ONBUILD RUN /usr/local/bin/python --dir /app/src
[...]
```
如果基于镜像test创建新的镜像时，新的Dockerfile中使用`FROM test`指定基础镜像时，会自动执行ONBUILD指令内容，等价于在后面又添加了两条指令。

####EXPOSE
格式：

`EXPOSE <port> [<port>...]`

指定容器在运行时监听的端口，在容器启动的时候，可以通过-p指令将宿主机上的端口映射到容器所监听的端口上面。

例子：

`EXPOSE 8080`

####ENTRYPOINT
格式：

- `CMD ["executable","param1","param2"]`
- `CMD ["param1","param2"]` 

ENTRYPOINT在镜像构建中，通常是扮演一个初始化的工作，用户可以把对于系统的一些必须要做的初始化工作放在脚本中，通过ENTRYPOINT来执行。而通过ENTRYPOINT来生成的指令可能不会被轻易的重写，从基本功能上来看，用户可以把不需要变动的信息写在ENTRYPOINT中，把需要变动的，最后容器执行的时候要指定的信息写在CMD中。

例子：

```
FROM  ubuntu:14.10
ENTRYPOINT  ["top", "-b"]
CMD  ["-c"]
```

####CMD

CMD指令有三种格式：

- `CMD ["executable","param1","param2"]`使用exec执行，推荐方式
- `CMD ["param1","param2"]` 提供给ENTRYPOINT的默认参数
- `CMD command param1 param2` 在/bin/sh中执行，提供给需要交互的应用

CMD命令的主要功能是用于指定启动容器时执行的命令

*注意：*

- 每个Dockerfile只能有一条CMD命令。如果指定了多条命令，只有最后一条会被执行。如果用户启动容器时候指定了运行的命令，则会覆盖掉CMD指定的命令。
- 注意与`RUN`指令进行区别，`RUN`指令在构建镜像的过程中就会被执行，执行结果会生成一层新的镜像，`CMD`指令在build镜像的过程中不会被执行，在镜像启动的时候才会执行。

例子：

- `CMD ["/bin/ls","-l","-s"]`
- `CMD ["/usr/bin/wc","--help"]`
- `CMD echo "This is a container"`

#### finish





