<template>
  <Header></Header>
  <div class="container">
    <Balance :total="+total"/>
    <IncomeExpenses :income="+income" :expense="+expense" />
    <TransactionList :transactions="transactions" @trans-deleted="deleteTrans" :loading="loadingTransactions"/>
    <AddTransaction @trans-added="addTrans" :loading = "loadingPost"  />
  </div>
</template>

<script setup lang="js">
import Header from './components/Header.vue'
import Balance from './components/Balance.vue';
import IncomeExpenses from './components/IncomeExpenses.vue';
import TransactionList from './components/TransactionList.vue';
import AddTransaction from './components/AddTransaction.vue';
import { computed, onMounted } from 'vue';
import { useToast } from 'vue-toastification';
import { useAPI } from './composables/useAPI';
const toast = useToast()


const {
    transactions, 
    error, //Que hacer con este error ???
    loadingTransactions, 
    getAllTransactions,
    loadingPost,
    postError,//Que hacer con este error ???
    postTransaction
  } = useAPI()

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

const expense = computed(() => {
  if (transactions.value) {
    return transactions.value.filter(t => t.amount < 0).reduce(
      (acc, t) => acc + t.amount, 0
    )
  }
  return 0
})

const addTrans = async (transData) => { 
  const newTrans = {
   "description": transData.text,
    "amount": transData.amount
  }

  await postTransaction(newTrans)
  toast.success("Transaction Added")
  await getAllTransactions()

}

const deleteTrans = (id) => {
  console.log("Call the API to delete transaction ...")
}

onMounted(async () => {
  await getAllTransactions()
})

</script>
