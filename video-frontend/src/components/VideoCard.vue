<template>
    <v-card
            class="mx-auto my-2"
            raised
            dark
            color="light-blue darken-4"
    >
        <v-list-item three-line>
            <v-list-item-content>
                <v-list-item-title class="headline mb-1">{{file}}</v-list-item-title>
                <v-list-item-subtitle v-clipboard:copy="videoRoot+file" v-clipboard:success="onCopy">
                    {{videoRoot}}{{file}}
                </v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>
        <v-card-actions>
            <v-btn text large v-clipboard:copy="videoRoot+file" v-clipboard:success="onCopy"
                   color="orange">
                <v-icon>
                    mdi-content-copy
                </v-icon>
                Copy
            </v-btn>
            <v-btn text large @click="openVideo(videoRoot+file)" color="orange">
                <v-icon>
                    mdi-download
                </v-icon>
                download
            </v-btn>
            <v-btn v-if="file.toString().endsWith('.mp4')" text large link
                   :to="{name:'video',query:{video:file.toString()}}" color="orange">
                <v-icon>
                    mdi-play-circle-outline
                </v-icon>
                play
            </v-btn>
        </v-card-actions>
        <v-snackbar
                v-model="snackbar"
                top
                :color="color"
                rounded="pill"
                :timeout="1000"
        >{{text}}
            <template v-slot:action="{ attrs }">
                <v-btn
                        dark
                        icon
                        @click="snackbar = false"
                >
                    <v-icon>
                        mdi-close
                    </v-icon>
                </v-btn>
            </template>
        </v-snackbar>
    </v-card>
</template>

<script>
    export default {
        name: "VideoCard",
        props: ['file'],
        data() {
            return {
                snackbar: false,
                videoRoot: window.location.href + "video/",
                text: "",
                color: "",
            }
        },
        methods: {
            onCopy: function () {
                this.text = "复制成功!"
                this.color = "success"
                this.snackbar = true
            },
            openVideo: function (videoPath) {
                window.open(videoPath, "_blank")
            }
        },
    }
</script>

<style scoped>

</style>
