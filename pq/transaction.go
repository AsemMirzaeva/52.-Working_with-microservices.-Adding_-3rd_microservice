package pq

import (
	"database/sql"
	"errors"
	pb "clientstream/salespb"
)

type SalesTransactionRepo struct {
	DB *sql.DB
}

func NewSalesTransactionRepo(db *sql.DB) *SalesTransactionRepo {
	return &SalesTransactionRepo{db}
}

func(db *SalesTransactionRepo) SaveTransaction(req *pb.SalesTransaction) (error) {

	sqlResult, err := db.DB.Exec("insert into sales_transactions(transaction_id, product_id, quantity, price) values($1, $2, $3, $4)", req.TransactionId, req.ProductId, req.Quantity, req.Price)
	if err != nil {
		return err
	}
	number, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if number == 0 {
		return errors.New("error while insert: no rows affected")
	}
	return nil
}

func(db *SalesTransactionRepo) SaveSummary(req *pb.SalesSummary) (error) {

	sqlResult, err := db.DB.Exec("insert into sales_summary(total_amount, total_transactions) values($1, $2)", req.TotalAmount, req.TotalTransactions)
	if err != nil {
		return err
	}
	number, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if number == 0 {
		return errors.New("error while insert: no rows affected")
	}
	return nil
}