## Inicialização do Ambiente de Desenvolvimento

Para configurar e iniciar o ambiente de desenvolvimento do projeto "API Gateway Knative-Docker", siga os passos abaixo:

### 1. Pré-requisitos:
- Certifique-se de ter o [Docker](https://www.docker.com/get-started) e o [Docker Compose](https://docs.docker.com/compose/install/) instalados em sua máquina.

### 2. Clonar o Repositório:
Clone o repositório para sua máquina local usando o comando:
```bash
git clone https://github.com/caiomarcatti12/api-gateway-knative-docker.git
cd api-gateway-knative-docker
```

### 3. Configuração das Rotas:
- Antes de iniciar o projeto, é necessário configurar as rotas. Para isso, crie um arquivo chamado `config.yaml` na raiz do projeto.
- Use o arquivo `config.example.yaml` como referência para a estrutura do arquivo `config.yaml`.
- Configure as rotas conforme sua necessidade. Se precisar adicionar mais rotas ou entender a configuração das existentes, consulte o guia [Configuração de Rotas](./route_configuration.md).

### 4. Iniciar o Projeto:
- Com o arquivo `config.yaml` configurado, você pode iniciar o projeto usando o Docker Compose. Execute o seguinte comando no diretório raiz do projeto:
```bash
docker-compose up -d
```
Este comando iniciará todos os serviços definidos no arquivo `docker-compose.yaml` em modo "detached", ou seja, em segundo plano.

### 5. Verificação:
- Após iniciar os serviços, você pode verificar se todos os contêineres estão em execução usando o comando:
```bash
docker ps
```
- Agora, o projeto deve estar em execução e pronto para receber solicitações conforme as rotas configuradas.

### 6. Parar o Projeto:
- Quando quiser parar o projeto, use o seguinte comando:
```bash
docker-compose down
```
