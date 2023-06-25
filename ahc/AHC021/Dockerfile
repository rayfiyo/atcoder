FROM ubuntu

WORKDIR /root

COPY ./* ./

# "RUN" は、ビルド時に コマンドを実行する。
RUN apt update; apt upgrade -y;
#RUN apt install -y curl; \
#	apt-get install -y gcc; \
#	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh ;\
#	source ~/.profile;
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
RUN source ~/.profile

# "CMD" は、ビルド後の コンテナ生成時に実行する。
CMD ["/usr/bin/pwd"]
