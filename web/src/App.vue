<template>
<div class="w-100 bg-light-red white avenir" id="app"> 
    <section class="mw5 mw7-ns center">
        <nav class="db dt-l w-100 border-box pv4">
        </nav>
        <section class="mw7 center">
            <h1 class="f2">Paste My Taste 🎶</h1>
            <article class="lh-copy">
                <p>Paste My Taste generates an easily shareable list of top artists and genre information from your Last.FM account for a given timespan.</p>
            </article>
        </section>
        <section class="mw7 center">
          <PopularArtists msg="Paste Your Taste for Last.FM"/>
        </section>
        <section class="mw7 center lh-copy">
            <h2 class="f3" id="about">About 👋</h2>
            I built this replacement as the old "Paste My Taste" site went offline, if you have any feature requests or bug reports feel free to ping me on <a href="https://twitter.com/tehwey" class="link dim white underline">Twitter</a>.
            <p>
                If you want to support further development tell your friends or
                <a class="link dim white underline" target="_blank" href="https://www.buymeacoffee.com/dewey">
                    <span style="margin-right:5px">buy me a tea</span>
                    <img src="https://www.buymeacoffee.com/assets/img/BMC-btn-logo.svg" alt="Buy me a tea">
                </a>
            </p>
        </section>
        <footer class="pv6 center white">
            <p class="f6 db lh-solid">
                <a href="https://github.com/dewey" class="pr2 link dim white">Source on Github</a>
                <a href="https://twitter.com/tehwey" class="pr2 link dim white">Twitter</a>
            </p>
        </footer>
    </section>
</div>
</template>

<script>
import PopularArtists from "./components/PopularArtists.vue";

export default {
  name: "app",
  components: {
    PopularArtists
  }
};
</script>
