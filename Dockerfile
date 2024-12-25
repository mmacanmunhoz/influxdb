# Usa a imagem oficial do Ubuntu
FROM ubuntu:latest

# Atualiza pacotes e instala Docker CLI
RUN apt-get update && \
    apt-get install -y docker.io && \
    apt-get clean

# Verifica se o grupo 'docker' existe antes de criá-lo e adiciona um usuário
RUN groupadd -f -g 999 docker && \
    useradd -ms /bin/bash -G docker ubuntu-user

# Define o diretório de trabalho
WORKDIR /home/ubuntu-user

# Define o usuário padrão como 'ubuntu-user'
USER ubuntu-user

# Copia arquivos necessários (se houver)
COPY . .

# Configura o comando padrão
CMD ["bash"]
