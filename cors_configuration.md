## Configuração CORS (Cross-Origin Resource Sharing)

O CORS é um mecanismo de segurança implementado pelos navegadores web que permite controlar o acesso a recursos em um servidor, com base na origem da solicitação.

Pode-se configurar regras CORS no nível global ou especificamente para uma rota.

### Parâmetros CORS

1. **allowedOrigins**: Uma lista de origens permitidas para acessar os recursos. Use `"*"` para permitir qualquer origem.
2. **allowedMethods**: Métodos HTTP permitidos.
3. **allowedHeaders**: Uma lista de cabeçalhos HTTP que podem ser usados quando se faz a solicitação real.
4. **allowCredentials**: Um booleano que indica se os cookies e credenciais devem ser enviados nas solicitações.
5. **exposedHeaders**: Cabeçalhos que podem ser expostos como parte da resposta.
6. **maxAge**: O tempo máximo que as informações de CORS podem ser armazenadas pelo navegador, em segundos.

### Exemplo de Configuração CORS

Configuração global de CORS:

```yaml
cors:
  allowedOrigins:
    - "*"
  allowedMethods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
  allowedHeaders:
    - "Authorization"
    - "Content-Type"
  allowCredentials: true
  exposedHeaders:
    - "X-Custom-Header"
  maxAge: 3600
routes:
  ...
```

Configuração CORS específica para uma rota:

```yaml
routes:
  - path: "/properties/"
    ...
    cors:
      allowedOrigins:
        - "https://specific-domain.com"
      allowedMethods:
        - "GET"
        - "POST"
      allowedHeaders:
        - "Authorization"
```

Neste exemplo, a rota `/properties/` tem uma configuração CORS específica que permite apenas solicitações do domínio `https://specific-domain.com` usando os métodos GET e POST e permite apenas o cabeçalho `Authorization`. Todas as outras rotas que não possuem uma configuração CORS específica usarão a configuração global definida na raiz do arquivo YAML.
