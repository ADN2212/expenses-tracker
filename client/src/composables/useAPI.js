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

    return {transactions, error, loadingTransactions, getAllTransactions}
}
