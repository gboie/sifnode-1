<template>
  <div class="layout">
    <Panel dark>
      <template v-slot:header v-if="header">
        <PanelNav />
      </template>
      <div class="header" v-if="backLink || title">
        <div v-if="backLink">
          <router-link class="back-link" :to="backLink"
            ><Icon icon="back"
          /></router-link>
        </div>
        <div v-if="emitBack">
          <span @click="$emit('back')" class="back-link"
            ><Icon icon="back"
          /></span>
        </div>
        <div class="title">
          {{ title }}
        </div>
      </div>
      <slot></slot>
    </Panel>
    <Panel v-if="!!$slots.after" class="after">
      <slot name="after"></slot>
    </Panel>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Panel from "@/components/shared/Panel.vue";
import PanelNav from "@/components/navigation/PanelNav.vue";
import Icon from "@/components/shared/Icon.vue";
export default defineComponent({
  components: { Panel, PanelNav, Icon },
  props: {
    backLink: String,
    header: { type: Boolean, default: true },
    emitBack: {
      type: Boolean,
      default: false,
    },
    title: String,
  },
});
</script>

<style lang="scss" scoped>
.layout {
  background: url("../../assets/World_Background_opt.jpg");
  background-size: cover;
  background-position: bottom center;
  box-sizing: border-box;
  padding-top: $header_height;
  width: 100%;
  height: 100vh; /* TODO: header height */
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.after {
  margin-top: 15px;
  padding: 25px;
  background: linear-gradient(180deg, $c_gray_50 0%, $c_gray_200 100%);
}

.header {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}
.back-link {
  text-align: left;
  display: block;
  text-decoration: none;
  position: relative;
  top: 2px;
  cursor: pointer;
}
.title {
  @include title16;
  flex-grow: 1;
  text-align: center;
}
</style>
