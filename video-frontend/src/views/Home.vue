<template>
    <v-container fluid>
        <v-row dense>
            <v-col
                    v-for="(file,index) in videos"
                    :key="index"
            >
                <video-card :file="file"></video-card>
            </v-col>
        </v-row>
        <v-footer app>
        </v-footer>
    </v-container>
</template>

<script>
    import VideoCard from "../components/VideoCard";

    export default {
        name: "Home",
        components: {VideoCard},
        data() {
            return {
                videos: [],
                dark: false,
            }
        },
        methods: {},
        beforeCreate: function () {
            this.axios.get('/list', {       // 还可以直接把参数拼接在url后边
            }).then(res => {
                    this.videos = res.data['files'];
                    this.$store.commit("update", this.videos)
                }
            ).catch(err => {
                this.text = err
                this.color = "error"
                this.snackbar = true
            });
        },
    }
</script>

<style scoped>

</style>
