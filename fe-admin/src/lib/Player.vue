<template>
    <div class="video-player-container">
        <video ref="videoPlayer" class="video-js vjs-default-skin" controls>
            <source :src="videoSource" type="application/x-mpegURL" />
            <track v-for="(subtitle, index) in subtitles" :key="index" :label="subtitle.language" :kind="subtitle.kind"
                :src="subtitle.src" default />
        </video>
    </div>
</template>

<script>
import videojs from 'video.js';
import 'video.js/dist/video-js.css';

export default {
    name: 'M3uPlayer',
    props: {
        videoSource: {
            type: String,
            required: true,
        },
        subtitles: {
            type: Array,
            required: false,
            default: () => []
        }
    },
    mounted() {
        this.initializePlayer();
    },
    methods: {
        initializePlayer() {
            const videoElement = this.$refs.videoPlayer;
            this.player = videojs(videoElement, {
                autoplay: true,
                controls: true,
                preload: 'auto',
                fluid: true,
            });
        },
    },
    beforeDestroy() {
        if (this.player) {
            this.player.dispose();
        }
    }
};
</script>

<style scoped>
.video-player-container {
    max-width: 100%;
    margin: 0 auto;
}
</style>