# KMC - Kenoby Mind Cracker Suite

Suite de mini aplicativos para ajudar a resolver as questões de matemática dos testes de RH online da Kenoby/Mindsight

## Aplicações


### kmcBee

kmcBee - é um aplicativo de linha de comando que ajuda a resolver o "problema das abelhas" e similares, quando uma grandeza é inversamente proporcional à outra. 

Por exemplo:  Quanto mais recursos, menos tempo para realizar uma tarefa e vice versa. 

#### Funcionamento

1. Abra um terminal (prompt de comando) no diretório onde baixou e descompactou o kmcBee. 
2. Execute o comando kmcBee (linux) ou kmcBee.exe (windows) seguido dos quatro parâmetros do problema, separados por espaço. 
3. O programa retornará a variável x devidamente calculada. 
Vide exemplos abaixo

Exemplo 1:
	3000 abelhas --> 50 dias
	4000 abelhas -->  x dias

-->	Execute: `kmcBee 3000 50 6000 x`

Exemplo 2:
	3000 abelhas --> 50 dias
	x    abelhas --> 37.5 dias

-->	Execute: `kmcBee 3000 50 x 37.5`

#### pacotes: 
* kmcBee --> binário linux. Não esqueça de dar permissão de execução com o comando `chmod +x kmcBee`
* kmcBee.zip --> binário windows. Descompacte e execute: `kmcBee.exe `
* kmcBee.go --> source code
