package enpoints

import (
	"expenses-tracker.com/DB"
	"expenses-tracker.com/IO"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTransactions(ctx *gin.Context) {
	trans, err := DB.GetAllTransactions()
	if err != nil {
		//Que estatus deberia devolver si hay un error a nivel de Basse de Datos ???
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusFound, trans)
}

func AddTransaction(ctx *gin.Context) {

	var newTransaction IO.TransactionInput

	err := ctx.BindJSON(&newTransaction)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	insertError := DB.CreateTransaction(newTransaction.Description, newTransaction.Amount)

	if insertError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": insertError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "new transaction added"})

}

func DeleteTransaction(ctx *gin.Context) {
	id, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	rowsAffected, deleteError := DB.DeleteTransactionById(id) 

	if deleteError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
		return
	}

	if rowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "succesful deleted"})	

}
