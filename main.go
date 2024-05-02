package main

import (
    "fmt"
    "strings"

    "github.com/jdkato/prose/v2"
)

type Conversa struct {
    Mensagens []string
}

func (c *Conversa) AdicionarMensagem(mensagem string) {
    c.Mensagens = append(c.Mensagens, mensagem)
}

func ObterContexto(conversa Conversa) string {
    if len(conversa.Mensagens) == 0 {
        return "Não há mensagens disponíveis para gerar contexto."
    }

    conteudo := strings.Join(conversa.Mensagens, " ")

    doc, err := prose.NewDocument(conteudo)
    if err != nil {
        fmt.Println("Erro ao criar documento:", err)
        return ""
    }

    entidades := doc.Entities()

    if len(entidades) == 0 {
        return "Não foram encontradas entidades para definir um contexto."
    }

    resumo := "O contexto da conversa está baseado nesse assunto: " + entidades[0].Text

    if len(entidades) > 1 {
        resumo += " e com essas palavras-chave: "
        for _, entidade := range entidades[1:] {
            resumo += entidade.Text + ", "
        }
        resumo = strings.TrimSuffix(resumo, ", ")
    }

    return resumo
}

func main() {
    conversa := Conversa{}

    conversa.AdicionarMensagem("Hoje, João e Maria vão ao Museu de Arte de São Paulo.")
    conversa.AdicionarMensagem("Eles planejam se encontrar lá às 15h.")
    conversa.AdicionarMensagem("Depois, vão jantar no restaurante Cantina Italiana.")

    contexto := ObterContexto(conversa)
    fmt.Println(contexto)
}