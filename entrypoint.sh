#!/bin/sh

# Obter o GID do docker.sock
SOCKET_GID=$(stat -c '%g' /var/run/docker.sock)

# Verificar se o grupo com o GID especificado já existe
if ! getent group $SOCKET_GID; then
    # Criar um grupo com o GID especificado se ele não existir
    addgroup -g $SOCKET_GID docker_dynamic
fi

# Adicionar o usuário 'appuser' ao grupo
addgroup appuser docker_dynamic

# Trocar para o usuário 'appuser' e executar o comando passado ao container
exec su -s /bin/sh -c "$@" appuser
