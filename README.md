# API de Encurtamento de Links

Esta API permite encurtar links de maneira aleatória. Siga os passos abaixo para testá-la:

## Passo a Passo

1. **Inicie o servidor**

   Execute o seguinte comando no terminal para iniciar o servidor:

   ```bash
   go run .

Envie uma requisição POST para o servidor, informando o link que deseja encurtar. Neste exemplo, utilizamos o comando curl para encurtar o link do Google:
curl -X POST http://localhost:8080/api/shorten -d "{\"url\":\"https://www.google.com\"}"

O servidor responderá com um código JSON contendo o link encurtado. Exemplo de resposta:
{"data":"hU0wmTsH"}

Utilize o código retornado para acessar o link original. Abra o endereço encurtado no navegador:
http://localhost:8080/hU0wmTsH
Isso redirecionará para a página original, neste caso, https://www.google.com.

