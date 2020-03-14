# KMC - Kenoby Mind Cracker Suite

Suite de mini aplicativos para ajudar a resolver as questões de matemática dos testes de RH online da Kenoby/Mindsight.

## Filosofia de desenvolvimento
Como os testes da Kenoby são cronometrados, os mini-aplicativos foram desenvolvidos para execução em linha de comando, com intuito de dar agilidade ao candidato.

O funcionamento de cada mini-aplicativo foi pensado para ser o mais direto possível, evitando formulários, uso do mouse e qualquer coisa que possa atrasar o candidato.

Seguindo esta lógica, preferencialmente os mini-aplicativos são executados da seguinte forma:

`<comando> <pr1> <pr2> <pr n...>`

Onde:
* <comando> = nome do mini-aplicativo
* <pr> = parâmetros da questão.

## Dicas de uso:

### Antes de fazer o seu teste online na Kenoby recomendamos:
- leia as instruções de cada mini-aplicativo.
- se possível, execute cada mini-aplicativo com seus exemplos.
- lembre-se que a ordem das questões podem variar no teste.
- Deixe a janela de terminal aberta na pasta dos mini-aplicativos
- Não deixe outras janelas abertas além do terminal e do navegador que usará no teste.
- Tenha lápis e papel em branco para anotar os parametros das questões.

### Durante o teste:
- leia atentamente e com muita calma a questão.
- Anote os parâmetros na folha de papel.
- utilize as teclas `Alt-tab` para alternar entre janelas.
- execute o mini-aplicativo adequado usando os parâmetros da questão.

## Contribua:
Caso apareça uma questão diferente, tente decorar rapidamente o enunciado, anote os dados e depois abra uma issue no projeto com sugestão para inclusão na suite.


## Mini Aplicativos


### kmcBee

kmcBee - é um aplicativo de linha de comando que ajuda a resolver o "problema das abelhas" e similares, quando uma grandeza é inversamente proporcional à outra.

Por exemplo:  Quanto mais recursos, menos tempo para realizar uma tarefa e vice versa.

#### Funcionamento

1. Baixe o pacote do programa conforme sistema utilizado.
2. Descompacte o pacote em um diretório
3. Abra um terminal e execute o aplicativo conforme instruções abaixo.


##### Exemplo 1:
	3000 abelhas --> 50 dias
	4000 abelhas -->  x dias

* No linux execute: `kmcBee 3000 50 6000 x`
* No windows Execute `kmcBee.exe 3000 50 6000 x`

##### Exemplo 2:
	3000 abelhas --> 50 dias
	x    abelhas --> 37.5 dias

* Linux ---->	Execute: `kmcBee 3000 50 x 37.5`
* Windows --> Execute: `kmcBee.exe 3000 50 x 37.5`

#### pacotes:
* kmcBee.tar.gz --> binário linux. Não esqueça de dar permissão de execução com o comando `chmod +x kmcBee`
* kmcBee.zip --> binário windows.
* kmcBee.go --> source code
