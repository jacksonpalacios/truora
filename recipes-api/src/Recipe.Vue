<template>  
      <div style="background-color:#eee; margin-bottom:0px; padding:10px;">
        <div class="box">
          <article class="media">
            <div class="media-left">
              <figure class="image is-100x100">
                <img src="https://cdn.kiwilimon.com/pages/temporadas/mundial-futbol/img_blog_alitas_mundial.png" alt="Image">
              </figure>
            </div>
            <div class="media-content">
              <div class="content">
                <p>
                  <strong>{{propData.recipe.name }} </strong> <small>31m</small>
                  <br>
                  {{propData.recipe.description}}
                </p>
              </div>
              <nav class="level is-mobile">
                <div class="level-left">                            
                  <a @click="updatedRecipe(propData.recipe, propData.index)" class="level-item" aria-label="retweet">
                    <span class="icon is-small">
                      <i class="fas fa-edit" aria-hidden="true"></i>
                    </span>
                  </a>
                  <a @click="deleteRecipe(propData.recipe.id, propData.index)" class="level-item" aria-label="like">
                    <span class="icon is-small">
                      <i class="fas fa-trash-alt" aria-hidden="true"></i>
                    </span>
                  </a>
                </div>
              </nav>
            </div>
          </article>
        </div>                          
    </div>  
</template>
<script>
import recipeService from "./services/recipe";

export default {
  name: "recipe",
  props: ["prop-data"],
  methods: {
    deleteRecipe(id, index) {
      let _this = this;
      recipeService
        .deleteRecipe(id)
        .then(res => {
          console.log(res.data);
          _this.$emit("delete", index);
        })
        .catch(e => {
          console.log(e);
        });
    },
    updatedRecipe(recipe, index){
      this.$bus.$emit('update', {recipe, index})
    }

  }
};
</script>
<style lang="sass">

</style>


