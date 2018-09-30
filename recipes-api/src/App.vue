<template>
  <div id="app"> 
    <div class="container">
      <nav class="navbar is-transparent has-shadow">
        <div class="navbar-brand">
          <a class="navbar-item" href="/">
            RECETAS DE COCINA
          </a>
          <div class="navbar-burger burger" data-target="navbarExampleTransparentExample">
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>
        <!--
        <div id="navbarExampleTransparentExample" class="navbar-menu">
          <div class="navbar-start">            
            <div class="navbar-item has-dropdown is-hoverable">
              <a class="navbar-link" href="#">  Recetas </a>
              <div class="navbar-dropdown is-boxed">
                <a class="navbar-item" href="#"> Listar </a>
                <a class="navbar-item" href="#"> Agregar </a>                
              </div>
            </div>
          </div>    
        </div>
        -->
      </nav>
      <br>
      <br>
      <section>
        <header>
          
          <div class="columns">
            <div class="column is-4">
              <h1 class="title is-2">Recetas de cocina</h1>
              <h2 class="subtitle">
                Aquí encontrarás todas las recetas que te puedas imaginar                   
              </h2>
                
            </div>
            <div class="column is-8">
                <p class="is-6">Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias, nam impedit. Iste iusto ratione impedit, nesciunt id sed porro cum saepe harum nulla recusandae, eligendi dicta facere, voluptate architecto? Aut!
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias, nam impedit. Iste iusto ratione impedit, nesciunt id sed porro cum saepe harum nulla recusandae, eligendi dicta facere, voluptate architecto? Aut!
                </p>
                
            </div>

          </div>
          
          <!--<a href="#" class="button is-primary">Agregar nueva receta</a>-->
          <hr/>          
        </header>
        <div class="container">
          <div class="columns" > 
            <div class="column is-4">
              <form-recipe @save="getRecipeSave" @updaterecipe="getRecipeUpdated"/>              
            </div>
            <div class="column is-8"> 
              <recipe @delete="deleteRecipe" v-for="(recipe, index) of recipes" v-bind:key="recipe.id" :prop-data="{recipe:recipe, index:index}"/>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div> 
</template>

<script>  
  import recipeService from './services/recipe'
  import Recipe from './Recipe'
  import FormRecipe from './FormRecipe'

  export default {
    name: 'app',
    componets: { Recipe, FormRecipe },
    data () {
      return {        
        recipes: [ ]
      }
    },
    created(){
      this.getRecipes()
    },
    methods:{
      getRecipes(){
        recipeService.getRecipes().then(res => {
          this.recipes = res.data
        }).catch(e => {
          console.log(e)
        })
      },
      getRecipeSave(recipe){        
        this.recipes.unshift(recipe)
      },
      deleteRecipe(index){
        this.recipes.splice(index,1);
      },
      getRecipeUpdated(data){   
        this.recipes.splice(data.index,1)
        this.recipes.unshift(data.recipe) 
      }
    }
  }
</script>

<style lang='scss'>
  @import url('./scss/main.scss')

</style>
