//
// Created by Jude Paulemon on 4/29/2022.
//
#include "video_reader.h"

#define FRAME_BUFFER_SIZE  5

AVFormatContext *open_media(char *mediaPath) {
    AVFormatContext *mediaContext = avformat_alloc_context();
    printf("open media: %s,\n", av_err2str(avformat_open_input(&mediaContext, mediaPath, NULL, NULL)));
    printf("file name:: %s, extension %s, n stream: %d\n", mediaContext->url, mediaContext->iformat->name,
           mediaContext->nb_streams);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        /*const AVCodecDescriptor *code_desc = avcodec_descriptor_get(current_stream->codecpar->codec_id);
        printf("# %d, codec type %s, media_type %s \n",current_stream->id,code_desc->name,av_get_media_type_string(code_desc->type));*/
        av_dump_format(mediaContext, i, mediaContext->url, 0);
    }

    // avformat_close_input(&mediaContext);
    return mediaContext;
}


struct MediaBuffer *init_media_buffer(char *mediaPath) {
    int err;
    struct MediaBuffer *mediaBuffer;
    AVFormatContext *mediaContext = avformat_alloc_context();
    err = avformat_open_input(&mediaContext, mediaPath, NULL, NULL);
    if (err) {
        printf("could not open media file\nerr: %s", av_err2str(err));
    }
    mediaBuffer = malloc(sizeof(struct MediaBuffer));
    mediaBuffer->mediaContext = mediaContext;

    //set max buffer size
    AVStream *videoStream;

    for (int i = 0; i < mediaBuffer->mediaContext->nb_streams; i++) {
        AVStream *stream = mediaBuffer->mediaContext->streams[i];

        if (stream->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
            mediaBuffer->audioStreamIndex = i;
        } else if (stream->codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
            mediaBuffer->videoStreamIndex = i;
            videoStream = stream;
        }
    }
    size_t videoBuffSize = sizeof(void*) * videoStream->r_frame_rate.num * FRAME_BUFFER_SIZE;
    printf("video buffer size: %zu\n",videoBuffSize);
    mediaBuffer->videoBuffer = malloc(sizeof (struct FrameBuffer));
    mediaBuffer->videoBuffer->count = 0;
    mediaBuffer->videoBuffer->maxSize = videoStream->r_frame_rate.num * FRAME_BUFFER_SIZE;
    printf("buffer count %d, maxsize %d\n",mediaBuffer->videoBuffer->count,mediaBuffer->videoBuffer->maxSize);

    return mediaBuffer;
}

void buffer(struct MediaBuffer *buffer, int percent) {
    int err = 0;
    AVPacket *pkt = av_packet_alloc();
    int nReadVideoFrame = 0;
    int nReadAudioFrame = 0;
    int64_t nbFrames = buffer->mediaContext->streams[buffer->videoStreamIndex]->nb_frames;
    while ((((long double) nReadVideoFrame) / nbFrames) * 100 <= percent) {

        err = av_read_frame(buffer->mediaContext, pkt);
        if (err == 0) {
            int streamIndex = pkt->stream_index;

            enum AVMediaType mediaType = buffer->mediaContext->streams[streamIndex]->codecpar->codec_type;

            if (mediaType == AVMEDIA_TYPE_VIDEO) {
                AVStream *videoStream = buffer->mediaContext->streams[streamIndex];
                if (nReadVideoFrame == 0) {
                    printf("real framerate: %d/%d fps | average framerate: %d/%d fps\n", videoStream->r_frame_rate.num,
                           videoStream->r_frame_rate.den, videoStream->avg_frame_rate.num,
                           videoStream->avg_frame_rate.den);
                    printf("\ntimebase: num :%d, den: %d\n", videoStream->time_base.num, videoStream->time_base.den);
                }
                buffer->videoBuffer->packets++;
                buffer->videoBuffer->packets = malloc(sizeof (AVPacket));
                * (AVPacket*)buffer->videoBuffer->packets = *pkt;
                AVPacket* pkt2 =  (AVPacket*)buffer->videoBuffer->packets;
                pkt = av_packet_alloc();

                nReadVideoFrame++;
                printf("#%d: video dts %ld, pts %ld  max frame: %ld\n", nReadVideoFrame, pkt2->dts,
                       pkt2->pts, buffer->mediaContext->streams[streamIndex]->nb_frames);

            } else if (mediaType == AVMEDIA_TYPE_AUDIO) {
                nReadAudioFrame++;
               printf("#%d: audio dts %ld, pts %ld  max frame: %ld\n", nReadAudioFrame, pkt->dts,
                                     pkt->pts, buffer->mediaContext->streams[streamIndex]->nb_frames);
            }
        } else {
            printf("err while reading frame\nerr: %s", av_err2str(err));
            break;
        }
    }
}

int bufferUP(struct MediaBuffer *buffer) {
    printf("hello world from buffer2\n");
    int err = 0;
    AVPacket *pkt = av_packet_alloc();
    int64_t nbFrames = buffer->mediaContext->streams[buffer->videoStreamIndex]->nb_frames;
    printf("buffer count %d, maxsize %d",buffer->videoBuffer->count,buffer->videoBuffer->maxSize);
    while (buffer->videoBuffer->count < buffer->videoBuffer->maxSize) {

        err = av_read_frame(buffer->mediaContext, pkt);

        if (err < 0 ){
            if (pkt->data == NULL){
                printf("err while reading frame\nerr: %s", av_err2str(err));
                return -2;
            }else{
                printf("END OF STREAM\n");
                return -1;
            }
        }
            int streamIndex = pkt->stream_index;

            enum AVMediaType mediaType = buffer->mediaContext->streams[streamIndex]->codecpar->codec_type;

            if (mediaType == AVMEDIA_TYPE_VIDEO) {
                AVStream *videoStream = buffer->mediaContext->streams[streamIndex];
                circularBufferAdd(buffer->videoBuffer,pkt);
                printf("buffer start %d, buffer end %d\n",buffer->videoBuffer->start,buffer->videoBuffer->end);
                printf("pts %ld\n",pkt->pts);
                pkt = av_packet_alloc();
            } else if (mediaType == AVMEDIA_TYPE_AUDIO) {
                //TODO
            }
    }
}

void read_frame(AVFormatContext *mediaContext) {
    AVPacket *pkt = av_packet_alloc();
    av_read_frame(mediaContext, pkt);
}

void *circularBuffGet(struct FrameBuffer *buffer) {
    if (buffer->count == 0) {
        return NULL;
    }
    void *item = buffer->packets[buffer->start];
    buffer->start++;
    buffer->start %= buffer->maxSize;
    return item;
}

void circularBufferAdd(struct FrameBuffer *buffer, void *packet){
    buffer->packets[buffer->end] = packet;
    buffer->end++;
    buffer->end %= buffer->maxSize;
    buffer->count++;
}
