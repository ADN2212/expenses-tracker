import { ref } from 'vue'

const BASE_URL = 'http://localhost:8080/transactions'

export function useAPI() {
    const transactions = ref(null)
    const error = ref(null)
    const loadingTransactions = ref(false)


    const getAllTransactions = async () => {
        loadingTransactions.value = true
        error.value = null
        try {
            const response = await fetch(BASE_URL)
            transactions.value = await response.json()
        } catch (err) {
            error.value = err
        } finally {
            loadingTransactions.value = false
        }
    }

    const loadingPost = ref(false)
    const postError = ref(null)

    const postTransaction = async (newTrans) => {
        loadingPost.value = true
        try {
            await fetch(BASE_URL, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(newTrans)
            })
        } catch (err) {
            console.log(err.message)
            postError.value = err.message
        }
        loadingPost.value = false
    }

    const loadingDelete = ref(false)
    const deleteError = ref(null)

    const deleteTransaction = async (id) => {
        loadingDelete.value = true
        try {

            const response = await fetch(BASE_URL + "/" + id, {
                method: "DELETE",
            })

            const responseStatus = response.status
            if (responseStatus >= 400 && responseStatus <= 599) {
                deleteError.value = "Someting went wrong while deleting"
            }
            
        //Ojo: Si a API responde con un 400 o 500 el bloque del cacth no se va a ejecutar.
        } catch (err) {
            console.log(err.message)
            deleteError.value = err.message
        }
        loadingDelete.value = false
    }

    return {
        transactions,
        error,
        loadingTransactions,
        getAllTransactions,
        loadingPost,
        postError,
        postTransaction,
        loadingDelete,
        deleteError,
        deleteTransaction
    }
}
