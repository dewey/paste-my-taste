<template>
  <div>
    <form class="mw7 center pa4 br2-ns" @submit.prevent="getPopularArtists">
      <fieldset class="cf bn ma0 pa0">
        <div class="cf">         
          <div class="flex">
            <div class="w-100 pa3">
              <label class="clip" for="email-address">Last.FM username</label>
              <input class="f6 f5-l tc input-reset bn fl black-80 bg-white pa3 w-100 br2-ns" placeholder="Last.FM Username" type="text" name="username" v-model="username" autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false">
            </div>
          </div>
          <div class="flex pa3">
            <div class="w-25 pr3">
              <v-select v-model="period" :options="[{label: 'Overall', value: 'overall'},{label: '7 days', value: '7day'},{label: '1 Month', value: '1month'},{label: '3 Months', value: '3month'},{label: '6 Months', value: '6month'}, {label: '12 Months', value: '12month'}]"></v-select>
            </div> 
            <div class="w-25 pr3">
              <v-select v-model="limit" :options="[{label: '10 days', value: '10'},{label: '15 days', value: '15'},{label: '20 days', value: '20'},{label: '25 days', value: '25'}]"></v-select>
            </div>
            <div class="w-30 pr3" >
              <div class="flex items-center pv2">
                <input class="mr2 pa3" type="checkbox" id="linkArtists" value="false" v-model="linkArtists">
                <label for="linkArtists" class="lh-copy">Add artist links</label>
              </div>
            </div>
            <div class="w-20">
              <input class="f6 f5-l button-reset fl pv2 tc bn bg-animate bg-black-70 hover-bg-black white pointer w-100 br2-ns" type="submit" value="Generate">
            </div>
          </div>
          <div class="flex" v-if="error.length > 0">
            <div class="w-100 pa3">
              There was an error processing your request: {{ error }} 😔
            </div>
          </div>
        </div>
      </fieldset>
    </form>
    <div v-if="popularArtists.length > 0">
        <h2 class="f3" id="result">Result 🥳</h2>
        <div class="pmt-result" @focus="$event.target.select()">
          <p>I'm into {{ popularTags }} including:
  {{ popularArtists }}
          </p>
          <p>Check out my music taste: <a v-bind:href="'https://www.last.fm/user/' + username">https://www.last.fm/user/{{username}}</a></p>
        </div>
    </div>
  </div>
</template>

<script>
import UserApi from "@/services/User";
import vueSlider from "vue-slider-component";
import vSelect from "vue-select";

export default {
  components: {
    vueSlider,
    vSelect
  },
  name: "PopularArtists",
  data() {
    return {
      username: "",
      period: { label: "Period" },
      limit: { label: "Limit" },
      linkArtists: false,
      loading: true,
      popularArtists: "",
      popularTags: "",
      error: ""
    };
  },
  methods: {
    getPopularArtists() {
      UserApi.getPopularArtists({
        username: this.username,
        period: this.period.value,
        limit: this.limit.value
      })
        .then(popularArtists => {
          var pal = [];
          var ptl = [];
          var outStr = "";

          popularArtists.forEach(element => {
            if (this.linkArtists) {
            pal.push(`[${element.name}](${element.url})`);
            } else {
              pal.push(element.name);
            }
            ptl.push(element.genre);
          });
          if (pal.length === 1) {
            outStr = pal[0];
          } else if (pal.length === 2) {
            outStr = pal.join(" and ");
          } else if (pal.length > 2) {
            outStr =
              pal.slice(0, -1).join(", ") + " and " + pal.slice(-1) + ".";
          }
          this.popularArtists = outStr;
          this.popularTags = ptl.filter(val => val).join(", ");
        })
        .catch(error => (this.error = error.response.data))
        .finally(() => {
          // Once this is done we set the loading to be over, even if we had an error
          this.loading = false;
        });
    }
  }
};
</script>

<style scoped>
.ba {
  border-style: solid;
  border-width: 1px;
}

.bn {
  border-style: none;
  border-width: 0;
}

.b--black-10 {
  border-color: rgba(0, 0, 0, 0.1);
}

.pmt-result {
  background-color: white;
  padding: 20px;
}

.pmt-result p {
  color: black;
}

div.v-select {
  background-color: white;
  border-radius: 4px;
}
</style>