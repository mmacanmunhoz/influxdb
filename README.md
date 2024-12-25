Projeto de Monitoramento com InfluxDB em Golang
Descrição do Projeto

Este projeto demonstra como configurar e usar o InfluxDB para coletar métricas de containers Docker e enviar métricas personalizadas usando uma aplicação escrita em Golang. O projeto usa Docker Compose para orquestrar os serviços necessários.


Requisitos

- Docker e Docker Compose instalados
- Golang (opcional, para testar a aplicação localmente)


Configuração
Estrutura do Projeto

```
├── docker-compose.yml   # Arquivo para orquestrar os serviços
├── telegraf.conf        # Configuração do Telegraf para monitoramento do Docker
│── main.go              # Código principal para enviar métricas
│── go.mod               # Gerenciamento de dependências
│── go.sum               # Checksum das dependências
└── README.md            # Este arquivo
```

Importante obter o GID do grupo docker para passar para o container do telegraf no docker-compose

```
getent group docker | cut -d: -f3
```

```
user: "0:125"
```