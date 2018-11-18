<template>
  <div class="hello">
    <form @submit.prevent="getPopularArtists">
    <input type="text" placeholder="Last.FM username" v-model="username">
    <select v-model="period">
      <option disabled value="">Select timespan</option>
      <option value="overall">overall</option>
      <option value="7day">7 days</option>
      <option value="1month">1 month</option>
      <option value="3month">3 months</option>
      <option value="6month">6 months</option>
      <option value="12month">12 months</option>
    </select>
    <div>
      <vue-slider ref="slider" v-model="limit" :min=0 :max=50 :interval=5></vue-slider>
    </div>
    <div>
        <span>I'm into {{ popularTags }} including:</span>
        <span>{{ popularArtists }}</span>
    </div>
    </form>
  </div>
</template>

<script>
import UserApi from "@/services/User";
import vueSlider from "vue-slider-component";

export default {
  components: {
    vueSlider
  },
  name: "PopularArtists",
  data() {
    return {
      username: "",
      period: "",
      limit: "",
      loading: true,
      popularArtists: "",
      popularTags: ""
    };
  },
  methods: {
    getPopularArtists() {
      console.log("getPopularArtists");
      UserApi.getPopularArtists({
        username: this.username,
        period: this.period,
        limit: this.limit
      })
        .then(popularArtists => {
          var pal = [];
          var ptl = [];
          var outStr = "";

          popularArtists.forEach(element => {
            pal.push(element.name);
            ptl.push(element.genre);
          });
          if (pal.length === 1) {
            outStr = pal[0];
          } else if (pal.length === 2) {
            outStr = pal.join(" and ");
          } else if (pal.length > 2) {
            outStr = pal.slice(0, -1).join(", ") + " and " + pal.slice(-1);
          }
          this.popularArtists = outStr;
          this.popularTags = ptl.join(", ")
        })
        .catch(error => console.log(error))
        .finally(() => {
          // Once this is done we set the loading to be over, even if we had an error
          this.loading = false;
        });
    }
  }
};
</script>