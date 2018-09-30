<template>
    <div>      
      <h3  class="title">Agregar nueva receta</h3>
      <div class="field">
        <label class="label">Nombre de receta</label>
        <div class="control">
        <input class="input" v-model="recipe.name" type="text" placeholder="nombre receta">
        </div>
      </div>                                            

      <div class="field">
        <label class="label">Descripción</label>
        <div class="control">
        <textarea class="textarea" v-model="recipe.description" placeholder="describa su receta"></textarea>
        </div>
      </div>                

      <div class="field is-grouped">
        <div class="control">
        <button class="button is-link" @click="saveRecipe">Guardar</button>
        </div>
        <div class="control">
        <button class="button is-text" @click="clearRecipe">Limpiar</button>
        </div>
      </div>
    </div>
</template>
<script>
import recipeService from "./services/recipe";

export default {
  name: "form-recipe",
  props: ["prop-data"],
  data() {
    return {
      title: "Agregar nueva receta",
      recipe: {
        id: 0,
        name: "",
        description: ""
      },
      index: 0
    };
  },
  created() {
    this.$bus.$on("update", data => {
      this.recipe = {
        id: data.recipe.id,
        name: data.recipe.name,
        description: data.recipe.description
      };
      this.index = data.index;
      this.title = "Editar receta";
      console.log(this.index)
    });
  },
  methods: {
    saveRecipe() {
      let _this = this;
      if (this.recipe.id == 0) {
        recipeService
          .postRecipe(this.recipe)
          .then(res => {
            _this.recipe = res.data;
            _this.$emit("save", _this.recipe);
            _this.clearRecipe();
          })
          .catch(e => {
            console.log(e);
          });
      } else {
        recipeService
          .putRecipe(this.recipe)
          .then(res => {
            _this.recipe = res.data;
            _this.$emit("updaterecipe", {
              recipe: _this.recipe,
              index: _this.index
            });
            _this.clearRecipe();
          })
          .catch(e => {
            console.log(e);
          });
      }
    },
    clearRecipe() {
      this.recipe = {
        id: 0,
        name: "",
        description: "",        
      };
      this.index = 0
    }
  }
};
</script>

