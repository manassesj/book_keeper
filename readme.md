- Implementar todos os verbos HTTP (GET, POST, PUT, PATCH, DELETE, OPTIONS)
- O que é REST?
- O que é o modelo de maturidade REST?
- Pra que usamos REST?
- O que é HATEOAS?
- Qual é a vantagem de ter HATEOAS?
--------------------------------------------------------------------------------------------
- Metodos
  - Metodo "Options" : O método HTTP OPTIONS solicita opções de comunicação permitidas para uma determinada URL ou servidor. Um cliente pode especificar uma URL com este método ou um asterisco (*) para se referir a todo o servidor.

  - Metodo "PATCH" : E considerada um conjunto de instruções sobre como modificar um recurso. Um pouco difeene do metodo PUT que é uma representação completa de um recurso.

- REST
  - REST define um conjunto de restrições de como a arquitetura de um sistema de hipermídia distribuído em escala da Internet, como a Web, deve se comportar.

- Maturidade REST
  - Nivel 0: O nível zero de maturidade não utiliza recursos de URI, HTTP Methods e HATEOAS.
  - nivel 1:  O nível um de maturidade já considera a utilização eficiente de URIs. Os recursos são mapeados, mas ainda não emprega o uso eficiente dos verbos. Geralmente utilizam apenas GET e POST.
  - nivel 2: O nível dois de maturidade faz o uso eficiente de URIs e verbos HTTP.
  - nivel 3: O nível três de maturidade faz o uso eficiente dos três fatores. URIs, HTTP e HATEOAS.

- HATEOAS
  - HATEOAS é uma restrição que faz parte da arquitetura de aplicações REST, cujo objetivo é ajudar os clientes a consumirem o serviço sem a necessidade de conhecimento prévio profundo da API.
  - A vantagem é que  o cliente não precisa ter um conhecimento profundo da API, basta conhecer a URL de inicial e partir dos links fornecidos poderá acessar todos os recursos de forma circular, se guiando através das requisições realizadas.