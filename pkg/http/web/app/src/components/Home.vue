<template>
  <div>
    <v-container grid-list-md fluid class="grey lighten-4" >
        <v-tabs
            slot="extension"
            v-model="tabs"
            centered
            color="teal"
            text-color="white"
            slider-color="white"
        >
            <v-tab class="white--text" :key="2">
            KUDOS
            </v-tab>
            <v-tab class="white--text" :key="1">
            SEARCH
            </v-tab>
        </v-tabs>
    </v-container>
  </div>
</template>