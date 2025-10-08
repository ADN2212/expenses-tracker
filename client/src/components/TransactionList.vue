<template>
	<h3>History</h3>
	<p v-if="loading">
		Cargando transactiones ...
	</p>
	<ul v-else id="list" class="list">
		<li v-for="t in transactions" v-bind="t.id" v-bind:class="t.amount > 0 ? 'plus': 'minus'">
			{{ t.description }} <span>${{ t.amount }}</span>
			<button 
				@click="emitTransDeleteion(t.id)"
				class="delete-btn"> 
				delete 
			</button>
		</li>
	</ul>
</template>

<script setup>
import { defineProps } from 'vue';

const emit = defineEmits(['trans-deleted'])

const props = defineProps(
	{
	transactions: {
		type: Array,
		required: true	
	},
	loading: {
		type: Boolean,
		required: true
	}
}	
)

function emitTransDeleteion(id) {
	console.log('Sending event to delete trans')
	emit('trans-deleted', id)
}

</script>
