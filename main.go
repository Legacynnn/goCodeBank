package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Legacynnn/goCodeBank/domain"
	"github.com/Legacynnn/goCodeBank/infra/repository"
	"github.com/Legacynnn/goCodeBank/useCase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Daniel"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0
	
	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
	
}

func setupTransactionUseCase(db *sql.DB) useCase.UseCaseTransaction{
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := useCase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDb() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable","localhost", "5432", "pguser", "pgpass", "codebank")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error to connect to db")
	}

	return db
}