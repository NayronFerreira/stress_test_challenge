# Stress Test Challenge

## Sobre

`stress_test_challenge` é uma ferramenta CLI para teste de carga em aplicações web. Ela permite enviar múltiplas requisições a partir de multiplos canais para uma URL especificada e medir os tempos de resposta.

Além da url fa aplicação, o CLI também oferece a possibilidade de receber Headers para que sejam adicionados na requisição.

## Parâmetros

--url (obrigatório): A URL da aplicação Web que será testada.
--requests: O número total de requisições a serem enviadas. Se esse valor não for definido, o padrão é 100.
--concurrency: Número de requisições concorrentes. Se esse valor não for definido, o padrão é 10.
--header: Cabeçalho a ser incluído na requisição. Pode ser especificado várias vezes para múltiplos cabeçalhos.

Todos os comandos acima possuem atalhos, conforme evidenciado abaixo na mesma ordem:

-u
-r
-c
-h

## Relatório

Após a conclusão do teste de carga, um relatório será gerado no console, incluindo:

- Total de requisições enviadas
- Quantidade de requisições bem-sucedidas e com erros
- Duração total do teste
- Duração média, mínima e máxima das requisições
- Distribuição dos códigos de status HTTP recebidos

Segue abaixo ilustração do relatório apresentado após a execução do teste:


```bash
Teste de carga concluído para http://google.com

    SSSS   TTTTT  RRRR   EEEE  SSSS  SSSS   TTTTT  EEEE  SSSS  TTTTT
   S        T    R   R  E    S      S        T    E    S        T  
    SSS     T    RRRR   EEE   SSS    SSS     T    EEE   SSS     T  
       S    T    R  R   E        S      S    T    E        S    T  
   SSSS     T    R   R  EEEE  SSSS   SSSS    T    EEEE  SSSS    T  
    
=========================================
                 RELATORY
=========================================
URL Testada: http://google.com
Total de Requisições: 100
Requisições Bem-Sucedidas (200 OK): 100
Requisições com Erros: 0
Duração Total do Teste: 14080.00 ms
Duração Média por Requisição: 140.80 ms
Duração Mínima da Requisição: 215.00 ms
Duração Máxima da Requisição: 534.00 ms

Distribuição dos Códigos de Status HTTP:
  200: 100
=========================================
```

## Começando o Uso

Certifique-se de ter o Docker instalado em sua máquina. Se você ainda não tem, poderá baixá-lo a partir do link:

- Docker: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

### Construção e Execução

Com o Docker instalado, no terminal navegue até o diretório do projeto e execute o seguinte comando para construir a imagem.
Note que voce só precisa alterar o "nome-imagem" para o nome de imagem desejado.

```bash
docker build -t nome-imagem -t Dockerfile .
```

- Com a imagem criada, basta rodar o Docker Run com as variáveis do stress_test. Segue abaixo o camando lietal e o comando com atalhos.

```bash
docker run nome-imagem stressTest --url=http://google.com --requests=120 --concurrency=2
```
O mesmo comando com atalho:
```bash
docker run nome-imagem stressTest -u=http://google.com -r=120 -c=2
```
- Caso você precise passar headers junto com a URL, voce pode fazer da sequinte forma:

```bash
docker run nome-imagem stressTest --url=http://google.com --requests=120 --concurrency=2 --header='API_KEY: TOKEN_1'
```
O mesmo comando com atalho:
```bash
docker run nome-imagem stressTest -u=http://google.com -r=120 -c=2 -H='API_KEY: TOKEN_1'