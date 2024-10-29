# 🔗 API de Encurtamento de Links

Bem-vindo à API de Encurtamento de Links! Com esta API, você pode encurtar URLs de forma prática e rápida, transformando links longos em links curtos e fáceis de compartilhar. 🖇️

## 🚀 Passo a Passo para Utilizar a API

### 1. Inicie o Servidor

Primeiro, inicie o servidor para que a API fique disponível. No terminal, execute o comando abaixo:

```bash
go run .
```

Isso inicializará o servidor localmente, pronto para receber suas requisições de encurtamento de links.

### 2. Encurte um Link

Para encurtar um link, faça uma requisição POST para o endpoint `/api/shorten`. Aqui está um exemplo usando `curl` para encurtar o link do Google:

```bash
curl -X POST http://localhost:8080/api/shorten -d "{\"url\":\"https://www.google.com\"}"
```

### 3. Visualize a Resposta da API

Se a requisição for bem-sucedida, o servidor responderá com um JSON contendo o link encurtado. Exemplo de resposta:

```json
{"data":"hU0wmTsH"}
```

### 4. Acesse o Link Encurtado

Para acessar o link original, basta abrir o link encurtado no navegador. Use o código fornecido no JSON de resposta e adicione-o ao final da URL do servidor:

```bash
http://localhost:8080/hU0wmTsH
```

Ao acessar essa URL, você será redirecionado para o link original - neste caso, `https://www.google.com`.

---

**🎉 Pronto! Agora você pode transformar links longos em links curtos e redirecioná-los com facilidade. Divirta-se!**
