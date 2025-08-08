https://github.com/Garhlz/my_music_player/commits/feature/backend-rewrite-go/
swag init -g ./cmd/server/main.go


# Proxy configuration
export HOST_IP=$(cat /etc/resolv.conf | grep nameserver | awk '{print $2}')
export PROXY_PORT=7890
export http_proxy="http://${HOST_IP}:${PROXY_PORT}"
export https_proxy="http://${HOST_IP}:${PROXY_PORT}"
export all_proxy="socks5://${HOST_IP}:${PROXY_PORT}"

# 设置不走代理的地址 (非常重要)
# 确保对 WSL 内部、本地和局域网的访问不走代理
export no_proxy="localhost,127.0.0.1,::1,*.local,wsl.localhost,${HOST_IP}"


- 导出数据
mysqldump -u root -p[old_password] --default-character-set=utf8mb4 music_db1 > db_backup.sql

- 导入数据(docker 中)
mysql -h 127.0.0.1 -P 3307 -u elaine -p123456 --default-character-set=utf8mb4 music_db1 < db_backup.sql

