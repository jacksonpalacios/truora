import Vue from 'vue'
import App from './App.vue'

import RecipeComponent from './Recipe.vue'
import FormRecipeComponent from './FormRecipe.vue'

import EventBus from './plugins/event-bus'

Vue.use(EventBus)
Vue.component('recipe', RecipeComponent)
Vue.component('form-recipe', FormRecipeComponent)

new Vue({
  el: '#app',
  render: h => h(App)
})
