import { ref } from 'vue'
import axios from 'axios'

export function useAPI(baseURL) {

    const data = ref(null)
    const error = ref(null)
    const loading = ref(false)

    const api = axios.create({
        baseURL,
        timeout: 10000,//A que se refiere esto ???
        headers: {
            'Access-Control-Allow-Origin': 'http://localhost:8080/',
            'Content-Type': 'application/json'
        }     
    })

    const getAllTransactions = async (endpoint) => {
        loading.value = true
        error.value = null
        try {
            const res = await api.get(endpoint)
            data.value = res.data
        } catch (err) {
            error.value = err
        } finally {
            loading.value = false
        }
    }

    return {data, error, loading, getAllTransactions}
}
