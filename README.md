# FATEC DEV DAY 2021 - ISMAEL APOLINARIO - LINGUAGEM GO & OUTROS

```go
package main

import "fmt"

const (
	contatoIsmael = "ismaelapolinario@hotmail"
)

func main() {
	agradecer := []string{"Todos os professores FATEC MM", "Todos que estiveram presentes na live"}
	agradecimentosEspeciais := []string{"Maromo", "Jeferson Linux Tips"}

	fmt.Printf("Gostaria muito de agradecer: \n")

	for _, agradecido := range agradecer {
		fmt.Printf("%v \n", agradecido)
	}
	fmt.Println()

	fmt.Printf("E um agradecimento especial para: \n")

	for _, agradecidoEspecial := range agradecimentosEspeciais {
		fmt.Printf("%v \n", agradecidoEspecial)
	}

	fmt.Println()

	fmt.Printf("Caso tenham alguma duvida, enviem um email para: %v \n\n", contatoIsmael)
	fmt.Printf("Sem voces esse evento nao teria acontecido! Abaixo deixo os links que citei durante a live. Aproveitem! \n")

}


```

## Meus contatos

- Github: [iapolinario](https://github.com/IAPOLINARIO)
- Twitter: [@iapolinario](https://twitter.com/iapolinario)
- Email: ismaelapolinario@hotmail.com
- Linkedin: [ismaelapolinario](https://www.linkedin.com/in/ismaelapolinario/)
- Instagram: [iapolinario](https://www.instagram.com/iapolinario/)
- Website: [ismaelapolinario.com](https://ismaelapolinario.com)

## Perfis para seguir

- [Ellen Korbes](https://twitter.com/ellenkorbes)
- [Linux Tips](https://twitter.com/LINUXtipsBR)
- [Jef](https://twitter.com/badtux_)
- [Golang Weekly](https://twitter.com/golangweekly)

### O que foi criado durante a LIVE

#### [Bruteforce em arquivos ZIP](https://github.com/fatecmm-devday/palestra_linguagem_go/tree/devday21/brute-force)

Codigo escrito em Go para quebrar senhas de arquivos zip utilizando o metodo de brute force.

### [Twitter Bot](https://github.com/fatecmm-devday/palestra_linguagem_go/tree/devday21/twitter-bot)

Aplicacao que busca todos os usuarios que twittaram uma determinada hashtag nos ultimos minutos. Baseado na lista de usuarios obtidos, um sera sorteado aleatoriamente.

## Ferramentas

### [WSL](https://www.treinaweb.com.br/blog/o-que-e-windows-subsystem-for-linux-wsl/)

O [WSL](https://www.treinaweb.com.br/blog/o-que-e-windows-subsystem-for-linux-wsl/) ?? um recurso opcional dispon??vel no Windows 10 (a partir da vers??o 1607) que te permite executar bin??rios e scripts em Linux diretamente no Windows, traduzindo as instru????es enviadas para o sistema (as chamadas de sistema ou system calls) para uma instru????o v??lida para o kernel do Windows. Com ele ?? poss??vel ter um ambiente id??ntico a de uma distribui????o Linux que voc?? j?? esteja acostumado sem precisar usar uma m??quina virtual ou algo do tipo para isso.

### [OH-MY-ZSH](https://medium.com/@rgdev/como-instalar-oh-my-zsh-c0f96218fd90#:~:text=Oh%20my%20ZSH%20%C3%A9%20um%20framework%20open%2Dsource%20dirigido%20pela,e%20mais%20de%20140%20temas.)

[Oh my ZSH](https://medium.com/@rgdev/como-instalar-oh-my-zsh-c0f96218fd90#:~:text=Oh%20my%20ZSH%20%C3%A9%20um%20framework%20open%2Dsource%20dirigido%20pela,e%20mais%20de%20140%20temas.) ?? um framework open-source dirigido pela comunidade para gerenciar a configura????o do ZSH (Terminal feito para uso interativo) e melhorar o workflow de desenvolvimento. Ele inclui mais de 200 plugins opcionais (rails, git, OSX, hub, capistrano, brew, ant, php, python, etc), e mais de 140 temas.

### [HAVE I BEEN PWNED?](https://haveibeenpwned.com/)

[Have I Been Pwned?](https://haveibeenpwned.com/) (HIBP, sendo "Pwned" pronunciado em ingl??s como "poned",geralmente associado ?? frase "have i been pwned") ?? um site que permite aos usu??rios da Internet verificar se seus dados pessoais foram comprometidos por viola????es de dados. O servi??o coleta e analisa centenas de dumps de banco de dados e pastas contendo informa????es sobre bilh??es de contas vazadas e permite que os usu??rios pesquisem suas pr??prias informa????es digitando seu nome de usu??rio ou endere??o de e-mail.

### Repositorios no Github

- [100 dias de codigo](https://github.com/IAPOLINARIO/100-days-of-code)
- [Golang](https://github.com/golang/go)
- [Awesome-Go](https://github.com/avelino/awesome-go)

## Materiais para Estudo

### Plataformas e Sites

- [Golang.org](http://www.golangbr.org/)
- [Codecademy](https://www.codecademy.com/learn/learn-go)

### Livros

- [Programando em Go](https://www.casadocodigo.com.br/products/livro-google-go)
- [Introdu????o ?? linguagem Go](https://novatec.com.br/livros/introducao-linguagem-go/)
- [Go em A????o](https://novatec.com.br/livros/go-em-acao/)
- [Go Bootcamp - gratis (em ingles)](http://www.golangbootcamp.com/book)
- [An Introduction to Programming in Go - gratis(em ingles)](https://www.golang-book.com/books/intro)

### YOUTUBE

- [Curso completo de Go da Ellen Korbes](https://www.youtube.com/channel/UCxD5EE0H7qOhRr0tIVsOZPQ)
- [Linux Tips](https://www.youtube.com/user/linuxtipscanal)

![Namaria quebrando pc](brute-force/assets/namaria.jpg?raw=true)
