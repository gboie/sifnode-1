<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      polling: undefined as ReturnType<typeof setTimeout> | undefined,
      rowanUSD: "",
    };
  },
  methods: {
    async pollData() {
      function isNumeric(s: any) {
        return s - 0 == s && ("" + s).trim().length > 0;
      }
      const data = await fetch(
        "https://vtdbgplqd6.execute-api.us-west-2.amazonaws.com/default/tokenstats",
      );
      const json = await data.json();
      const rowanPriceInUSDT = json.body ? json.body.rowanUSD : "";

      if (isNumeric(rowanPriceInUSDT)) {
        this.rowanUSD =
          "ROWAN: $" + parseFloat(rowanPriceInUSDT).toPrecision(6);
      }
    },
  },
  created() {
    this.pollData();
    this.polling = setInterval(() => this.pollData(), 10000);
  },
  unmounted() {
    if (!this.polling) return;
    clearInterval(this.polling);
  },
});
</script>

<template>
  <div>
    <div v-if="rowanUSD" class="rowan">
      <img class="image" src="../../../public/images/siflogo.png" />
      <div>
        {{ rowanUSD }}
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.rowan {
  background: transparent;
  color: $c_gray_700;
  border-radius: $br_md;
  letter-spacing: 0.1px;
  font-family: Arial, Helvetica, sans-serif;
  font-style: normal;
  font-weight: bold;
  font-size: 12px;
  padding: 1px 5px;
  height: auto;
  line-height: initial;
  border: 1px solid $c_gray_400;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 2px;
  margin-right: 48px;

  .image {
    height: 16px;
    margin-right: 4px;
  }
}
</style>
