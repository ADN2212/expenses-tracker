package DB

import (
	"expenses-tracker.com/IO"
)

func GetAllTransactions() ([]IO.TransactionOutput, error) {
	dbCon := createDBConnection()
	defer dbCon.Close()

	var trans []IO.TransactionOutput
	transactionsRows, err := dbCon.Query("SELECT * FROM transactions;") //Las querys podrian estar en un file a parte.

	if err != nil {
		return trans, err
	}

	var currTrans IO.TransactionOutput

	defer transactionsRows.Close() //Esto es una conexion a parte, que tambien debe ser cerrada
	for transactionsRows.Next() {
		singleScanErr := transactionsRows.Scan(&currTrans.Id, &currTrans.Description, &currTrans.Amount)
		if singleScanErr != nil {
			return trans, singleScanErr //Deberia retornar alguna transaccion si han ocurrido errores ???
		}
		trans = append(trans, currTrans)
	}

	return trans, nil
}

func CreateTransaction(description string, amount float32) error {
	dbCon := createDBConnection()
	defer dbCon.Close()
	
	_, insertError := dbCon.Exec(`INSERT INTO transactions (description, amount) VALUES ($1, $2);`, description, amount)

	if insertError != nil {
		return insertError
	}

	return nil
}

func DeleteTransactionById(id uint64) (int64, error) {
	dbCon := createDBConnection()
	defer dbCon.Close()

	res, deleteError := dbCon.Exec(`DELETE FROM transactions WHERE id = $1;`, id)

	if deleteError != nil {
		return 0, deleteError
	}

	rowAfected, raErr := res.RowsAffected()

	if raErr != nil {
		return 0, raErr
	}

	return rowAfected, nil
}
