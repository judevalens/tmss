//
// Created by Jude Paulemon on 4/29/2022.
//
#include "video_reader.h"
#include "rtp.h"

#define FRAME_BUFFER_SIZE  5
#define AUDIO_SUFFIX  "_AUDIO"
#define VIDEO_SUFFIX  "_VIDEO"
#define MAX_URL_LEN     500
#define FILE_URL_SCHEME "file:"
#define DEFAULT_DIR "/home/jude/Desktop/amnis server/"
#define BUFF_INIT_SIZE 10

char *getFileName(char *name);

AVFormatContext *open_media(char *mediaPath) {
    AVFormatContext *mediaContext = avformat_alloc_context();
    printf("open media: %s,\n", av_err2str(avformat_open_input(&mediaContext, mediaPath, NULL, NULL)));
    printf("file name:: %s, extension %s, n stream: %d, long name: %s\n", mediaContext->url,
           mediaContext->iformat->name,
           mediaContext->nb_streams, mediaContext->iformat->long_name);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        printf("extra data size: %d\n",current_stream->codecpar->extradata_size);


        /*const AVCodecDescriptor *code_desc = avcodec_descriptor_get(current_stream->codecpar->codec_id);
        printf("# %d, codec type %s, media_type %s \n",current_stream->id,code_desc->name,av_get_media_type_string(code_desc->type));*/
        printf("fps or sample rate: %d,nb frames: %ld\n", current_stream->codecpar->sample_rate,
               current_stream->nb_frames);
        AVDictionaryEntry *entry = NULL;
        for (int j = 0; j < av_dict_count(current_stream->metadata); j++) {
            entry = av_dict_get(current_stream->metadata, "", entry, AV_DICT_IGNORE_SUFFIX);
            printf("%s:%s\n", entry->key, entry->value);
        }
        // av_dump_format(mediaContext, i, mediaContext->url, 0);
    }


    // avformat_close_input(&mediaContext);
    return mediaContext;
}

MediaBuffer init_media_buffer(char *mediaPath, int bufferByteSize) {
    
    int err;
    MediaBuffer mediaBuffer;
    AVFormatContext *mediaContext = avformat_alloc_context();

    char *fileUrl = malloc(sizeof(char *) * MAX_URL_LEN);
    if (snprintf(fileUrl,MAX_URL_LEN,"%s%s",FILE_URL_SCHEME,mediaPath) < 0) {
        return NULL;
    }
    err = avformat_open_input(&mediaContext, fileUrl, NULL, NULL);
    if (err < 0) {
        printf("could not open media file\nerr: %s", av_err2str(err));
    }
     mediaBuffer = malloc(sizeof(MediaBuffer));
    mediaBuffer->mediaContext = mediaContext;
    mediaBuffer->packetBuffers = malloc(sizeof(PacketBuffer *) * 2);
    for (int i = 0; i < 2; i++) {
        mediaBuffer->packetBuffers[i] = malloc(sizeof(struct PacketBuffer));
      //  mediaBuffer->packetBuffers[i]->totalByteSize = 100;
        mediaBuffer->packetBuffers[i]->size = BUFF_INIT_SIZE;
        mediaBuffer->packetBuffers[i]->currentIdx = 0;
        mediaBuffer->packetBuffers[i]->currentByteSize = 0;
        mediaBuffer->packetBuffers[i]->totalByteSize = bufferByteSize;
        mediaBuffer->packetBuffers[i]->packets = malloc(sizeof(AVPacket *) * BUFF_INIT_SIZE);
    }
    return mediaBuffer;
}

void buffer_2(MediaBuffer mediaBuffer, int bufferIdx) {
    PacketBuffer buffer = mediaBuffer->packetBuffers[bufferIdx];

    for (int i = buffer->currentByteSize; i < buffer->totalByteSize; i++) {
        AVPacket *pkt = av_packet_alloc();

        if (buffer->currentIdx == buffer->size) {
            buffer->packets = realloc(buffer->packets, sizeof(AVPacket *) * (buffer->currentIdx * 2));
        }
        int res = av_read_frame(mediaBuffer->mediaContext, pkt);
        if (res < 0) {
            break;
        }
        buffer->currentByteSize += pkt->size;
        buffer->packets[buffer->currentIdx++] = pkt;
    }
}

int buffer(MediaBuffer mediaBuffer, int bufferIdx) {
    PacketBuffer buffer = mediaBuffer->packetBuffers[bufferIdx];

    for (int i = 0; i < buffer->size; i++) {
        AVPacket *pkt = av_packet_alloc();
        int res = av_read_frame(mediaBuffer->mediaContext, pkt);
        printf("buffering\n");
        if (res < 0) {
            perror(av_err2str(res));
            buffer->size = i;
            buffer->eof = 1;
            break;
        }
        buffer->packets[i] = pkt;
    }
    return 0;
}


int seek(MediaBuffer mediaBuffer, int64_t position) {
    av_seek_frame(mediaBuffer->mediaContext, 0, position, AVSEEK_FLAG_FRAME);
}

/**
 * splits a video file into singular streams
 * @parameter mediaContext
 */
char **demux_file(AVFormatContext *mediaContext) {
    char *defaultVideoContainer = "mp4";
    char *defaultAudioContainer = "mp4";
    int res = 0;

    const char *mediaBaseName = av_basename(mediaContext->url);
    const char *dirname = av_dirname(mediaContext->url);
    printf("dirname: %s\n", dirname);
    char *fileName = getFileName((char *) mediaBaseName);
    printf("input url: %s\n", mediaContext->url);
    char **streamUrls = malloc(sizeof(char *) * mediaContext->nb_streams);
    char **streamBaseNames = malloc(sizeof(char *) * mediaContext->nb_streams);
    AVFormatContext **outCtx = malloc(sizeof(AVFormatContext *) * mediaContext->nb_streams);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        char *fileUrl = av_append_path_component(dirname, fileName);
        int fileUrlLen = snprintf(NULL, 0, "%s%s_stream_%d.%s", FILE_URL_SCHEME, fileUrl, i,
                                  defaultAudioContainer);
        streamUrls[i] = malloc(sizeof(char) * fileUrlLen);

        snprintf(streamUrls[i], MAX_URL_LEN, "%s%s_stream_%d.%s", FILE_URL_SCHEME, fileUrl, i,
                 defaultAudioContainer);


        streamBaseNames[i] = (char *) av_basename(streamUrls[i]);

        AVStream *currentStream = mediaContext->streams[i];
        printf("stream url: %s\n", streamUrls[i]);
        res = avformat_alloc_output_context2(&outCtx[i], NULL, NULL, streamUrls[i]);

        if (res < 0) {
            printf("failed to create out ctx for stream: %d", i);
            return NULL;
        }
        AVStream *newStream = avformat_new_stream(outCtx[i], NULL);
        avcodec_parameters_copy(newStream->codecpar, currentStream->codecpar);

        res = avio_open(&outCtx[i]->pb, outCtx[i]->url, AVIO_FLAG_WRITE);

        if (res < 0) {
            printf("failed to open write file for stream: %d", i);
            return NULL;
        }
        res = avformat_write_header(outCtx[i], NULL);
        if (res < 0) {
            printf("failed to write header for stream: %d", i);
            return NULL;
        }
    }
    free(fileName);
    free((char *) dirname);
    AVPacket *packet = av_packet_alloc();
    int i = 0;
    while (av_read_frame(mediaContext, packet) >= 0) {
        printf("%d: writing frame: stream# %d\n", i++, packet->stream_index);

        AVStream *srcStream = mediaContext->streams[packet->stream_index];
        int streamIndex = packet->stream_index;
        packet->stream_index = 0;
        av_packet_rescale_ts(packet, srcStream->time_base,
                             outCtx[streamIndex]->streams[0]->time_base);
        av_interleaved_write_frame(outCtx[streamIndex], packet);
    }
    for (int j = 0; j < mediaContext->nb_streams; j++) {
        res = av_write_trailer(outCtx[j]);
        if (res < 0) {
            printf("failed to write trailer for stream: %d", i);
            return NULL;
        }
    }
    return streamBaseNames;
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