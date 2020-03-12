FROM debian:latest
RUN mkdir -p /app
WORKDIR /app

# 将当前目录下的货物信息文件 consignment.json 拷贝到 /app 目录下
ADD user-cli /app/user-cli

#CMD ["./user-cli --name 'cao liangcheng' --email 'defsky@qq.com' --password 'password' --company 'sunhoo'"]
CMD ["/bin/bash"]