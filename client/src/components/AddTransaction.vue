<template>
    <p v-if="loading">
        Sending Transaction ...
    </p>
    <form v-else id="form" @submit.prevent="onSubmit">
        <div class="form-control">
            <label for="text">Text</label>
            <input v-model="text" type="text" id="text" placeholder="Enter text..." />
        </div>
        <div class="form-control">
            <label for="amount">Amount <br />
                (negative - expense, positive - income)</label>
            <input v-model="amount" type="text" id="amount" placeholder="Enter amount..." />
        </div>
        <button class="btn">Add transaction</button>
    </form>
</template>

<script setup>
import { ref } from 'vue';
import { useToast } from 'vue-toastification';

const text = ref('')
const amount = ref('')
const toast = useToast()

const props = defineProps(
	{
	loading: {
		type: Boolean,
		required: true
	}
})	

const emit = defineEmits(['trans-added'])

const onSubmit = () => {
    if (!text.value || !amount.value) {
        toast.error('You must fill all fields')
        return
    }

    const transData = {
        text: text.value,
        amount: parseFloat(amount.value)
    }

    if (isNaN(transData.amount)) {
        toast.error("Not valid amount")
        return
    }

    emit('trans-added', transData)    
    text.value = ""
    amount.value = ""

}

</script>

