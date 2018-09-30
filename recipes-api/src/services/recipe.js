import goApi from './api-rest'

const recipeService = {}

recipeService.getRecipes = () => {
  return goApi.get('/recipes')
}

recipeService.postRecipe = (data) => {
  return goApi.post('/recipe', data)
}

recipeService.deleteRecipe = (id) => {
  return goApi.delete('/recipe/' + id)
}

recipeService.putRecipe = (data) => {
  return goApi.put('/recipe/' + data.id, data)
}

export default recipeService
