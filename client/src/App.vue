<template>
  <Header></Header>
  <div class="container">
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
const toast = useToast()

const transactions = ref([])
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
  console.log('A new transaction was added')  
  //console.log(transData)
  transactions.value.push(
    {
      id: createId(),
      text: transData.text,
      amount: transData.amount 
    }  
  )
  toast.success("A new trasnaction was added!")    
  console.log(transactions.value)
}

function createId() {
  return Math.floor(Math.random() * 1000000)
}

//Esta funcion se ejecuta cuando se ejecuta el evento.
const deleteTrans = (id) => {
  //console.log(`Deleting trans of id = ${id}`)
  transactions.value = transactions.value.filter(t => t.id !== id)
  toast.success(`Trans of id = ${id} deleted`)
}

//Optener los datos del local storage:
onMounted(() => {
  const savedTransactions = JSON.parse(localStorage.getItem('transactions'))
  console.log(savedTransactions)
  if (savedTransactions) {
    console.log('They were transactions')
    transactions.value = savedTransactions
    return
  }
  console.log('They were not transactions')
})

//Este se ejecuta cade vez que array de transacsiones cambie
//En el video tutotial esto se hizo de otra forma.
watch(transactions, () => {
  console.log('Updating trans arrays in local storage')
  localStorage.setItem('transactions', JSON.stringify(transactions.value))
}, {deep: true})//El deep permite que el escuche todos los cambios.

</script>
