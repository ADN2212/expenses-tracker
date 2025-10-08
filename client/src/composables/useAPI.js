import { ref } from 'vue'

const BASE_URL = 'http://localhost:8080/transactions'

export function useAPI() {
    const transactions = ref(null)
    const error = ref(null)
    const loading = ref(false)
    
    const getAllTransactions = async () => {
        loading.value = true
        error.value = null
        try {   
            const response = await fetch(BASE_URL)  
            transactions.value = await response.json() 
        } catch (err) {
            error.value = err
        } finally {
            loading.value = false
        }
    }

    return {transactions, error, loading, getAllTransactions}
}
