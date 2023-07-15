<script setup>
const props = defineProps({
	title: {
		type: String,
		required: true
	},
	subTitle: {
		type: String,
		required: true
	},
	action: {
		type: String,
		required: true
	},
	icon: {
		type: String,
		required: false
	},
	bgColor: {
		type: String,
		required: false
	}
})

defineEmits(['click']);

</script>
<template>
	<v-card :color="bgColor" max-width="200" class="pt-1">
		<div class="d-flex flex-no-wrap justify-space-between">
			<div>
				<v-card-title class="text-h6">
					<v-icon class="mb-1" v-if="icon" size="x-small" color="gray">{{ icon }}</v-icon>
					{{ title }}
				</v-card-title>

				<v-card-subtitle>{{ subTitle }}</v-card-subtitle>

				<v-card-actions>
					<v-btn class="ma-2" size="small" variant="outlined" color="white" @click="$emit('click')">
						{{ action.toUpperCase() }}
					</v-btn>
				</v-card-actions>
			</div>
		</div>
	</v-card>
</template>

<style scoped>
.v-card-subtitle {
	overflow: hidden;
	display: -webkit-box;
	-webkit-line-clamp: 2;
	/* 设置所需的行数 */
	-webkit-box-orient: vertical;
}
</style>