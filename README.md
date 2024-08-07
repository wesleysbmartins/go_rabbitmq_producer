# Golang RabbitMQ Producer
[![My Skills](https://skillicons.dev/icons?i=golang,rabbitmq)](https://skillicons.dev)

O RabbitMQ é um message broker altamente consolidado e utilizado por quem trabalha com comunicação entre sistemas. Operando de forma assíncrona, ele age como um intermediário que processa as nossas mensagens entre produtores e consumidores, além de contar com filas que possuem diversas opções de encaminhamento.

Para entender melhor o que é e como ele funciona [acesse](https://github.com/wesleysbmartins/rabbitmq).

Esta aplicação é um exemplo de Producer do RabbitMQ em Golang, nela utilizamos o pacote [**amqp**](https://pkg.go.dev/github.com/streadway/amqp), como um driver para integrar a aplicação ao message broker RabbitMQ.

### Pré-Requisitos
Para executar esta aplicação será necessário ter em seu ambiente o [Docker](https://docs.docker.com/engine/install/) para os containers da própria aplicação quanto do RabbitMQ. Com o Docker instalado podemos seguir o passo a passo.

### Caso de Uso
Esta aplicação foi desenvolvida para receber **Sales** ou Vendas, que serão enviadas as exchanges e filas configuradas.

Entidade Sale:
```go
package entities

type Sale struct {
	SellingCompany     string `json:"sellingCompany"`
	Product            string `json:"product"`
	Price              string `json:"price"`
	DeliveryCompany    string `json:"deliveryCompany"`
	OriginAddress      string `json:"originAddress"`
	DestinationAddress string `json:"destinationAddress"`
	ClientName         string `json:"clientName"`
	Order              int64  `json:"order"`
}
```
Este é o objeto que será processado pela aplicação e enviado aoMessage Broker.
### Container RabbitMQ
Com o Docker presente em seu ambiente, podemos realizar o download da imagem do RabbitMQ:
```shell
$ docker pull rabbitmq:management
```
Execute um conteiner do RabbitMQ passando suas credenciais de usuário, senha e portas:
```shell
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=<user> -e RABBITMQ_DEFAULT_PASS=<password> rabbitmq:management
```
Pronto, você tem um RabbitMQ operando em seu ambiente!

### Configuração da Aplicação
Antes de executar a aplicação deverá ser preenchido arquivos de configuração, semelhantes aos conhecidos arquivos **.env**, mas neste caso utilizaremos arquivos de extensão **yml**.

Na pasta config presente na raiz do projeto, você deverá criar um arquivo chamado **rabbitmq.yml**, e inserir os seguintes valores:
```yml
host:     "<host do rabbitmq>"
port:     5672
user:     "<usuário do rabbitmq>"
password: "<senha do rabbitmq>"
```
Estas seriam as credenciais necessárias para conexão com o RabbitMQ.

Também será necessário criar um arquivo que contém as configurações de seu Producer, contento informações de sua Exchange, Queue e Message, nomeie como **sale-producer.yml**:
```yml
exchange: {
  name: "exchange-name", //nome da exchange
  kind: "direct", //tipo da exchange
  durable: false,
  auto-delete: false,
  internal: false,
  no-wait: false,
  args: [], // argumentos de configuração mais especificos como x-message-ttl em milisegundos
  bind: "", // nome da exchange que deseja criar um bind, ou seja, a mensagem recebida será enviada para esta exchange também
}
queue: {
  name: "queue-name", // nome da fila
  durable: false,
  exclusive: false,
  auto-delete: false,
  no-wait: false,
  args: [], // argumentos de configuração mais especificos como x-message-ttl em milisegundos
  bind: "", // nome da fila que deseja criar um bind, ou seja, a mensagem recebida será enviada para esta fila também
}
message: {
  headers: [], // argumentos de configuração mais especificos como x-message-ttl em milisegundos
  content-type: "application/json",
  delivery-mode: 0,
  priority: 0,
  correlation-id: "",
  reply-to: "",
  exiration: "",
  type: ""
}
```

Substitua os valores genéricos por suas configurações de preferência, eles serão lidos e convertidos em configurações de seu **Producer**, com base neles serão criados suas **Exchanges, Binds e Queues**.

#### OBS!!
Se você criar mais de um arquivo não quer dizer que todos eles serão executados, esta aplicação só executa operações ligadas as regras de negócio de **Sales**, seria necessário implementar novos Usecases para suas novas abordagens.

Caso queira adicionar novos endpoints, producers e etc, eu recomendo a análise do código em controllers, usecases e factory, acredito que será necessário adicionar somente novos endpoints, controllers e usecases.

Por fim, devemos adicionar o arquivo **server.yml**, que conterá as informações de configuração do servidor:
```yml
port: 8082 // porta da aplicação
allowed-origins: ["*"] // lista de origens habilitadas a realizar operações no seu serviço
```

### Execução
Com as configurações implementadas, e RabbitMQ rodando no seu ambiente, você pode executar sua aplicação.

Atente-se ao Dockerfile e escolha sua porta de preferência, ou continue usano a padrão **8082**.

Crie a imagem no Docker:
```shell
docker build -t producer_service .
```

Execute o container:
```shell
docker run -p 8082:8082 producer_service
```

Você deve ter um resultado semelhante a este:
```
SERVER LISTENNING ON PORT: 808231 <nil>
```

Assim você poderá alcançar o endpoint para enviar mensagens, utilize o Postman para testar o endpoint.
![alt text](./src/image.png)

Utilizando a seguinte estrutura no **BODY** da requisição:
```json
{
	"sale":{
        "sellingCompany": "Magalu",
        "product": "Notebook Lenovo",
        "price": "R$2.255,90",
        "deliveryCompany": "Correios",
        "originAddress": "Alameda Santos - Cerqueira César, São Paulo - SP, 01418-970",
        "destinationAddress": "Parque da Independência - Ipiranga, São Paulo - SP, 04263-000",
        "clientName": "João da Silva",
        "order": 8080800808
    },
    "appId": "postman collection",
    "userId": "postman runner"
}
```
Resultado esperado:
```
"Sale Received With Success!"
```
Assim você tem uma API preparada para receber solicitações de Sales, que envia seus dados para a exchange configurada e suas respectivas filas.