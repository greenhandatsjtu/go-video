<template>
    <v-container fill-height fluid>
        <v-row justify="center" class="mx-auto">
            <v-card class="mx-2" id="card">
                <vue-d-player :options="playerOptions" ref="player"></vue-d-player>
                <v-card-title>{{$route.query.video.split(".")[0]}}</v-card-title>
            </v-card>
        </v-row>
        <v-btn fixed fab left bottom x-large icon @click="toPrevious" color="blue accent-2">
            <v-icon>
                mdi-arrow-left-bold
            </v-icon>
        </v-btn>
        <v-spacer></v-spacer>
        <v-btn fixed fab right bottom x-large icon @click="toNext" color="red">
            <v-icon>
                mdi-arrow-right-bold
            </v-icon>
        </v-btn>
    </v-container>
</template>

<script>
    import VueDPlayer from 'vue-dplayer'
    import 'vue-dplayer/dist/vue-dplayer.css'

    export default {
        name: "VideoPage",
        components: {
            VueDPlayer
        },
        data() {
            return {
                playerOptions: {
                    container: document.querySelector('#card'),
                    loop: true,
                    autoplay: true,
                    video: {
                        url: "/video/" + this.$route.query.video,
                        pic: ""
                    }
                }
            }
        },
        methods: {
            toPrevious() {
                let file = this.$route.query.video
                let ind = this.$store.state.videos.indexOf(file)
                if (ind === 0) {
                    return
                }
                this.$router.push({name: 'video', query: {video: this.$store.state.videos[ind-1]}})
                this.$emit("reload")
            },
            toNext() {
                let file = this.$route.query.video
                let ind = this.$store.state.videos.indexOf(file)
                if (ind === this.$store.state.videos.length-1) {
                    return
                }
                this.$router.push({name: 'video', query: {video: this.$store.state.videos[ind+1]}})
                this.$emit("reload")
            }
        },
        // mounted() {
        //     this.$refs.player.dp.switchVideo({
        //         url: "/video/" + this.$route.query.video,
        //     })
        // }
    }
</script>

<style scoped>

</style>
