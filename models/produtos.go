package models

import (
	"log"
	"strconv"

	"github.com/ProgramadorRicardoMeneses/GO_CrieUmaAplicacaoWeb/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBD()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err)
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err)
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao, precoStr, quantidadeStr string) {
	preco, err := strconv.ParseFloat(precoStr, 64)
	if err != nil {
		log.Println("Erro na conversão do preço:", err)
	}
	quantidade, err := strconv.Atoi(quantidadeStr)
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
	}
	db := db.ConectaComBD()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error)
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DetetaProduto(id string) {
	db := db.ConectaComBD()
	deletarOProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBD()
	produtoBD, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizacao := Produto{}

	for produtoBD.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := produtoBD.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err)
		}
		produtoParaAtualizacao.Id = id
		produtoParaAtualizacao.Nome = nome
		produtoParaAtualizacao.Descricao = descricao
		produtoParaAtualizacao.Preco = preco
		produtoParaAtualizacao.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizacao
}

func AlteraProduto(idStr, nome, descricao, precoStr, quantidadeStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
	}
	preco, err := strconv.ParseFloat(precoStr, 64)
	if err != nil {
		log.Println("Erro na conversão do preço:", err)
	}
	quantidade, err := strconv.Atoi(quantidadeStr)
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
	}
	db := db.ConectaComBD()

	AltualizaOProduto, err := db.Prepare("update produtos set nome = $1, descricao=$2, preco=$3, quantidade=$4 where id = $5")
	if err != nil {
		panic(err.Error())
	}
	AltualizaOProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
