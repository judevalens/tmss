//
// Created by Jude Paulemon on 4/29/2022.
//
#include "video_reader.h"
#include "rtp.h"

#define FRAME_BUFFER_SIZE  5
#define AUDIO_SUFFIX  "_AUDIO"
#define VIDEO_SUFFIX  "_VIDEO"
#define MAX_URL_LEN     250
#define FILE_URL_SCHEME "file:"
#define DEFAULT_DIR "/home/jude/Desktop/amnis server/"

char *getFileName(char *name);

AVFormatContext *open_media(char *mediaPath) {
    AVFormatContext *mediaContext = avformat_alloc_context();
    printf("open media: %s,\n", av_err2str(avformat_open_input(&mediaContext, mediaPath, NULL, NULL)));
    printf("file name:: %s, extension %s, n stream: %d, long name: %s\n", mediaContext->url,
           mediaContext->iformat->name,
           mediaContext->nb_streams, mediaContext->iformat->long_name);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        /*const AVCodecDescriptor *code_desc = avcodec_descriptor_get(current_stream->codecpar->codec_id);
        printf("# %d, codec type %s, media_type %s \n",current_stream->id,code_desc->name,av_get_media_type_string(code_desc->type));*/
        av_dump_format(mediaContext, i, mediaContext->url, 0);
    }


    // avformat_close_input(&mediaContext);
    return mediaContext;
}
MediaBuffer init_media_buffer(char *mediaPath, int bufferSize) {
    int err;
    MediaBuffer mediaBuffer;

    AVFormatContext *mediaContext = avformat_alloc_context();
    err = avformat_open_input(&mediaContext, mediaPath, NULL, NULL);
    if (err) {
        printf("could not open media file\nerr: %s", av_err2str(err));
    }
    mediaBuffer = malloc(sizeof(MediaBuffer));
    mediaBuffer->mediaContext = mediaContext;
    mediaBuffer->packetBuffers = malloc(sizeof(PacketBuffer *)*2);
    for (int i = 0; i < 2; i++) {
        mediaBuffer->packetBuffers[i] = malloc(sizeof(PacketBuffer));
        mediaBuffer->packetBuffers[i]->size = bufferSize;
        mediaBuffer->packetBuffers[i]->packets = malloc(sizeof(AVPacket*)*bufferSize);
    }
    return mediaBuffer;
}

void buffer(MediaBuffer mediaBuffer, int bufferIdx) {
    PacketBuffer buffer  = mediaBuffer->packetBuffers[bufferIdx];
    AVPacket *pkt = av_packet_alloc();
    for (int i = 0; i < buffer->size; i++) {
        int res = av_read_frame(mediaBuffer->mediaContext,pkt);
        if (res < 0 ) {
            break;
        }
        buffer->packets[i] = pkt;
    }
}


int seek(MediaBuffer mediaBuffer, int64_t position) {
    AV_PKT_FLAG_KEY
    av_seek_frame(mediaBuffer->mediaContext,0,position,AVSEEK_FLAG_FRAME);
}
/**
 * splits a video file into singular streams
 * @parameter mediaContext
 */
void demux_file(AVFormatContext *mediaContext, char *OutUrl) {
    char *defaultVideoContainer = "mp4";
    char *defaultAudioContainer = "mp4";
    int res = 0;
    AVFormatContext *audioOutCtx = NULL;
    AVFormatContext *videoOutCtx = NULL;
    AVFormatContext **mediaMap = malloc(sizeof(AVFormatContext *) * 2);
    const char *mediaBaseName = av_basename(mediaContext->url);
    char *fileName = getFileName((char *) mediaBaseName);
    printf("input url: %s\n", mediaContext->url);

    char *audioOutName = malloc(MAX_URL_LEN);
    char *videoOutName = malloc(MAX_URL_LEN);
    char *audioOutUrl = malloc(MAX_URL_LEN);
    char *videoOutUrl = malloc(MAX_URL_LEN);

    sprintf(audioOutName, "%s%s.%s", fileName, AUDIO_SUFFIX, defaultAudioContainer);
    sprintf(videoOutName, "%s%s.%s", fileName, VIDEO_SUFFIX, defaultVideoContainer);
    sprintf(audioOutUrl, "%s%s/%s", FILE_URL_SCHEME, OutUrl, audioOutName);
    sprintf(videoOutUrl, "%s%s/%s", FILE_URL_SCHEME, OutUrl, videoOutName);
    free(fileName);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        if (audioOutCtx != NULL && videoOutCtx != NULL) {
            printf("There should not be more than two streams\n");
            return;
        }
        AVFormatContext *currentOutCtx;
        AVStream *currentStream = mediaContext->streams[i];
        if (currentStream->codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
            res = avformat_alloc_output_context2(&videoOutCtx, NULL, NULL, videoOutUrl);
            currentOutCtx = videoOutCtx;
            if (res < 0) {
                return;
            }
        } else if (currentStream->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
            res = avformat_alloc_output_context2(&audioOutCtx, NULL, NULL, audioOutUrl);
            if (res < 0) {
                return;
            }
            currentOutCtx = audioOutCtx;
        }
        AVStream *newStream = avformat_new_stream(currentOutCtx, NULL);
        avcodec_parameters_copy(newStream->codecpar, currentStream->codecpar);
        mediaMap[currentStream->index] = currentOutCtx;
    }

    AVPacket *packet = av_packet_alloc();
    printf("video url: %s\n", videoOutCtx->url);
    printf("audio url: %s\n", audioOutCtx->url);
    res = avio_open(&audioOutCtx->pb, audioOutCtx->url, AVIO_FLAG_WRITE);
    if (res < 0) {
        printf("failed to open audio AVIO context\n err: %s\n", av_err2str(res));
        return;
    }
    res = avio_open(&videoOutCtx->pb, videoOutCtx->url, AVIO_FLAG_WRITE);
    if (res < 0) {
        printf("failed to open video AVIO context\n err: %s", av_err2str(res));
        return;
    }
    res = avformat_write_header(videoOutCtx, NULL);

    if (res < 0) {
        perror(av_err2str(res));
        return;
    }

    res = avformat_write_header(audioOutCtx, NULL);

    if (res < 0) {
        perror(av_err2str(res));
        return;
    }
    int i = 0;
    while (av_read_frame(mediaContext, packet) >= 0) {
        printf("%d: writing frame: stream# %d\n", i++, packet->stream_index);

        AVStream *srcStream = mediaContext->streams[packet->stream_index];
        int streamIndex = packet->stream_index;
        packet->stream_index = 0;
        av_packet_rescale_ts(packet, srcStream->time_base,
                             mediaMap[streamIndex]->streams[0]->time_base);
        av_interleaved_write_frame(mediaMap[streamIndex], packet);
    }
    av_write_trailer(audioOutCtx);
    av_write_trailer(videoOutCtx);
}

char *getFileName(char *name) {

    int lastDotI = 0;

    for (int i = 0; i < strlen(name); i++) {

        if (name[i] == '.') {
            lastDotI = i;
        }
    }

    char *fileName = malloc(sizeof(char) * (lastDotI + 1));
    memcpy(fileName, name, sizeof(char) * (lastDotI));
    fileName[lastDotI] = '\0';
    printf("file name: %s\n", fileName);
    return fileName;
}