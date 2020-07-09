<template>
    <v-app>
        <v-app-bar
                app
                dark
                color="primary"
                hide-on-scroll
        >
            <v-app-bar-nav-icon></v-app-bar-nav-icon>
            <v-btn text link to="/">
                <v-icon x-large>
                    mdi-video
                </v-icon>
                <div class="hidden-xs-only">Local Video Streaming Platform</div>
            </v-btn>

            <v-spacer></v-spacer>
            <v-menu
                    bottom
                    left
                    transition="scale-transition"
            >
                <template v-slot:activator="{ on, attrs }">
                    <v-btn
                            icon
                            v-bind="attrs"
                            v-on="on"
                    >
                        <v-icon>mdi-dots-vertical</v-icon>
                    </v-btn>
                </template>

                <v-list>
                    <v-list-item>
                        <v-switch v-model="dark"
                                  dense
                                  :label="$vuetify.theme.dark?'Light Theme':'Dark Theme'"
                                  class="font-weight-bold my-2"/>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-app-bar>

        <v-main>
            <router-view @reload="key=(key+1)%4" :key="key"/>
        </v-main>
    </v-app>
</template>

<script>
    import Home from "./views/Home";

    export default {
        name: 'App',

        components: {
            Home,
        },

        data: () => ({
            dark: false,
            items: [],
            key: 0
        }),
        computed: {
            shouldBeDark: function () {
                let hour = new Date().getHours()
                return hour <= 7 || hour >= 19
            }
        },
        mounted: function () {
            this.dark = this.shouldBeDark
        },
        watch: {
            dark: function (val) {
                this.$vuetify.theme.dark = val
            }
        }
    };
</script>
