# KMC - Kenoby Mind Cracker Suite

Suite de mini aplicativos para ajudar a resolver as questões de matemática dos testes de RH online da Kenoby/Mindsight

## Aplicações


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
