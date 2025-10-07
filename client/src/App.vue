<template>
  <Header></Header>
  <div class="container">

    
  {{ data }}
  {{ error }}
  {{ loading }}

    <Balance :total="+total"/>
    <!-- +"3" = 3 -->
    <IncomeExpenses :income="+income" :expense="+expense" />
    <TransactionList :transactions="transactions" @trans-deleted="deleteTrans"/>
    <!--@eventName agrega un event listener que ejeuca la funcion al lado del igual-->
    <AddTransaction @trans-added="addTrans" />
  </div>
</template>

<script setup lang="js">
import Header from './components/Header.vue'
import Balance from './components/Balance.vue';
import IncomeExpenses from './components/IncomeExpenses.vue';
import TransactionList from './components/TransactionList.vue';
import AddTransaction from './components/AddTransaction.vue';
import { ref, computed, onMounted, watch} from 'vue';
import { useToast } from 'vue-toastification';
import { useAPI } from './composables/useAPI';

const toast = useToast()
const transactions = ref([])

const {data, error, loading, getAllTransactions} = useAPI('http://localhost:8080/')

//Sera reactivo al value de transactions, es decir, cuando la lista cambie este valor tambien cambiara.
const total = computed(() => {
  return transactions.value.reduce((acc, t) => acc + t.amount, 0) 
})

//Sumar solo el income, es decir las cantidades positivas:
const income = computed(() => {
  return transactions.value.filter(t => t.amount > 0).reduce(
    (acc, t) => acc + t.amount, 0
  )
})

//Sumar solo los gastos:
const expense = computed(() => {
  return transactions.value.filter(t => t.amount < 0).reduce(
    (acc, t) => acc + t.amount, 0
  )
})

//Lo que esta en el argumento de la lambda es la data que se envia "a traves" del evento
const addTrans = (transData) => {
  console.log('Call the API to add transaction ...')  
}

//Esta funcion se ejecuta cuando se ejecuta el evento.
const deleteTrans = (id) => {
  console.log("Call the API to delete transaction ...")
}

//Optener los datos del local storage:
onMounted(() => {
  console.log('Get trans-actions from API.')
  getAllTransactions('/transactions')
})

</script>
