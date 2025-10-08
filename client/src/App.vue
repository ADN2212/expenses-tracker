<template>
  <Header></Header>
  <div class="container">
    <Balance :total="+total"/>
    <!-- +"3" = 3 -->
    <IncomeExpenses :income="+income" :expense="+expense" />
    <TransactionList :transactions="transactions" @trans-deleted="deleteTrans" :loading="loading"/>
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

//Que hacer con este error ???
const {transactions, error, loading, getAllTransactions} = useAPI()

const total = computed(() => {  
  if (transactions.value) {
    return transactions.value.reduce((acc, t) => acc + t.amount, 0)
  }
  return 0
})

const income = computed(() => {
  if (transactions.value) {
    return transactions.value.filter(t => t.amount > 0).reduce(
      (acc, t) => acc + t.amount, 0
    )
  }
  return 0
})

//Sumar solo los gastos:
const expense = computed(() => {
  if (transactions.value) {
    return transactions.value.filter(t => t.amount < 0).reduce(
      (acc, t) => acc + t.amount, 0
    )
  }
  return 0
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
onMounted(async () => {
  getAllTransactions()
})

</script>
