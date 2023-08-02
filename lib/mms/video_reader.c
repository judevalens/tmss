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
#define FRAME_BUFF_SIZE 4096

#define VIDEO_EXT "mp4"
#define AUDIO_EXT "aac"
#define N_STREAMS 2

char EXT_MAP[2][20] = {VIDEO_EXT, AUDIO_EXT};

char *getFileName(char *name);

int init_decoder(AVFormatContext *ctx, AVCodecContext **codec_ctx, int stream_index);

char decode_packet(AVCodecContext *dec, const AVPacket *pkt, AVFrame *frame);

AVFormatContext *open_media(char *mediaPath) {
    AVFormatContext *mediaContext = avformat_alloc_context();
    printf("open media: %s,\n", av_err2str(avformat_open_input(&mediaContext, mediaPath, NULL, NULL)));
    printf("file name:: %s, extension %s, n stream: %d, long name: %s\n", mediaContext->url,
           mediaContext->iformat->name,
           mediaContext->nb_streams, mediaContext->iformat->long_name);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        printf("extra data size: %d\n", current_stream->codecpar->extradata_size);
        /*const AVCodecDescriptor *code_desc = avcodec_descriptor_get(current_stream->codecpar->codec_id);
        printf("# %d, codec type %s, media_type %s \n",current_stream->id,code_desc->name,av_get_media_type_string(code_desc->type));*/
        printf("fps or sample rate: %d,nb frames: %ld\n", current_stream->avg_frame_rate.num,
               current_stream->nb_frames);
        AVDictionaryEntry *entry = NULL;
        for (int j = 0; j < av_dict_count(current_stream->metadata); j++) {
            entry = av_dict_get(current_stream->metadata, "", entry, AV_DICT_IGNORE_SUFFIX);
            printf("%s:%s\n", entry->key, entry->value);
        }
        printf("%s\n", avcodec_get_name(current_stream->codecpar->codec_id));
        //av_dump_format(mediaContext, i, mediaContext->url, 0);
    }


    // avformat_close_input(&mediaContext);
    return mediaContext;
}

MediaBuffer init_media_buffer(char *mediaPath, int bufferByteSize) {

    int err;
    MediaBuffer mediaBuffer;
    AVFormatContext *mediaContext = avformat_alloc_context();

    char *fileUrl = malloc(sizeof(char *) * MAX_URL_LEN);
    if (snprintf(fileUrl, MAX_URL_LEN, "%s%s", FILE_URL_SCHEME, mediaPath) < 0) {
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
    int res = 0;

    const char *mediaBaseName = av_basename(mediaContext->url);
    const char *dirname = av_dirname(mediaContext->url);
    printf("dirname: %s\n", dirname);
    char *fileName = getFileName((char *) mediaBaseName);
    printf("input url: %s\n", mediaContext->url);

    char **streamUrls = malloc(sizeof(char *) * N_STREAMS);
    char **streamBaseNames = malloc(sizeof(char *) * N_STREAMS);
    AVFormatContext **outCtx = malloc(sizeof(AVFormatContext **) * N_STREAMS);
    AVCodecContext **codecCtx = malloc(sizeof(AVCodecContext **) * N_STREAMS);
    int videoIdx = av_find_best_stream(mediaContext, AVMEDIA_TYPE_VIDEO, -1, -1, NULL, 0);
    int audioIdx = av_find_best_stream(mediaContext, AVMEDIA_TYPE_AUDIO, -1, -1, NULL, 0);
    int idxMap[2] = {videoIdx, audioIdx};

    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        int mediaType = current_stream->codecpar->codec_type;
        if (mediaType != AVMEDIA_TYPE_VIDEO && mediaType != AVMEDIA_TYPE_AUDIO) {
            continue;
        }
        int stream_idx;

        if (current_stream->codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
            stream_idx = 0;
        } else if (current_stream->codecpar->codec_type == AVMEDIA_TYPE_AUDIO) {
            stream_idx = 1;
        }

        char *file_url = av_append_path_component(dirname, fileName);
        int file_url_len = snprintf(NULL, 0, "%s%s_%s_stream.%s", FILE_URL_SCHEME, file_url,
                                    av_get_media_type_string(mediaType), EXT_MAP[stream_idx]);
        streamUrls[i] = malloc(sizeof(char) * file_url_len);
        snprintf(streamUrls[i], MAX_URL_LEN, "%s%s_%s_stream.%s", FILE_URL_SCHEME, file_url,
                 av_get_media_type_string(mediaType), EXT_MAP[stream_idx]);

        streamBaseNames[i] = (char *) av_basename(streamUrls[i]);

        printf("stream url: %s\n", streamUrls[i]);
        res = avformat_alloc_output_context2(&outCtx[i], NULL, NULL, streamUrls[i]);
        if (res < 0) {
            printf("failed to create out ctx for stream: %d", i);
            return NULL;
        }
        res = init_decoder(mediaContext, &codecCtx[stream_idx], i);
        if (res < 0) {
            fprintf(stderr, "failed to init decoder for stream: %d", stream_idx);
        }
        AVStream *newStream = avformat_new_stream(outCtx[i], NULL);
        avcodec_parameters_copy(newStream->codecpar, current_stream->codecpar);

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
    AVFrame *frame = av_frame_alloc();
    int i = 0;
    // demux the media file and decode each avpacket from the video stream to extract the picture type and save it as a side data before writing it disk

    while (av_read_frame(mediaContext, packet) >= 0) {
        printf("%d: writing frame: stream# %d\n", i++, packet->stream_index);
        AVStream *src_stream = mediaContext->streams[packet->stream_index];
        int mediaType = src_stream->codecpar->codec_type;
        int stream_idx;
        if (mediaType == AVMEDIA_TYPE_VIDEO) {
            stream_idx = 0;
        } else if (mediaType == AVMEDIA_TYPE_AUDIO) {
            stream_idx = 1;
        } else {
            continue;
        }

        int streamIndex = packet->stream_index;
        packet->stream_index = 0;
        decode_packet(codecCtx[stream_idx], packet, frame);
        av_packet_rescale_ts(packet, src_stream->time_base,
                             outCtx[streamIndex]->streams[0]->time_base);
        av_interleaved_write_frame(outCtx[streamIndex], packet);
    }

    for (int j = 0; j < N_STREAMS; j++) {
        res = av_write_trailer(outCtx[j]);
        if (res < 0) {
            printf("failed to write trailer for stream: %d", i);
            return NULL;
        }
    }
    return streamBaseNames;
}

char decode_packet(AVCodecContext *dec, const AVPacket *pkt, AVFrame *frame) {
    int ret = 0;
    char c;
    // submit the packet to the decoder
    ret = avcodec_send_packet(dec, pkt);
    if (ret < 0) {
        fprintf(stderr, "Error submitting a packet for decoding (%s)\n", av_err2str(ret));
        return c;
    }
    // get all the available frames from the decoder
    while (ret >= 0) {
        ret = avcodec_receive_frame(dec, frame);
        if (ret < 0) {
            // those two return values are special and mean there is no output
            // frame available, but there were no errors during decoding
            if (ret == AVERROR_EOF || ret == AVERROR(EAGAIN))
                return c;
            fprintf(stderr, "Error during decoding (%s)\n", av_err2str(ret));
            return c;
        }
        c = av_get_picture_type_char(frame->pict_type);
        int pkt_metadata_size = 0;
        uint8_t *metadata_str;


        AVDictionary *metadata = NULL;

        av_dict_set(&metadata, "frame_type", &c, 0);
        av_dict_get_string(metadata, (char **) &metadata_str, '=', ':');
        printf("str pair: %s\n", metadata_str);
        av_packet_add_side_data((AVPacket *) pkt, AV_PKT_DATA_METADATA_UPDATE, metadata_str,
                                strlen(((char *) metadata_str)) - 1);


        printf("frame type: %c\n", c);
        av_frame_unref(frame);
    }

    return c;
}

char *decode(char *filepath) {
    AVFormatContext *ctx = open_media(filepath);
    AVStream *videoStream = ctx->streams[0];
    AVCodec *videoCodec = avcodec_find_decoder(videoStream->codecpar->codec_id);
    AVCodecContext *videoCodecCtx = avcodec_alloc_context3(videoCodec);
    AVCodecParserContext *parser;
    AVPacket *pkt = av_packet_alloc();
    FILE *inputFile;
    uint8_t inBuff[FRAME_BUFF_SIZE + AV_INPUT_BUFFER_PADDING_SIZE];
    uint8_t *data;
    size_t data_size;
    int eof = 0;
    int remaining = 0;
    char *pict_types = malloc(sizeof(char) * videoStream->nb_frames);
    char *h_pic_types = pict_types;
    if (videoCodec != NULL) {
        printf("video codec name: %s\n", avcodec_get_name(videoStream->codecpar->codec_id));
    }

    /* Copy codec parameters from input stream to output codec context */
    if ((avcodec_parameters_to_context(videoCodecCtx, videoStream->codecpar)) < 0) {
        perror("");
        exit(2);
    }

    if (avcodec_open2(videoCodecCtx, videoCodec, NULL) < 0) {
        perror("");
        exit(1);
    }

    parser = av_parser_init(videoStream->codecpar->codec_id);

    if (!parser) {
        perror("");
        exit(1);
    }
    (&ctx);

    inputFile = fopen(filepath, "rb");
    if (!inputFile) {
        perror("");
        return NULL;
    }
    int count = 0;
    AVFrame *frame = av_frame_alloc();
    while (av_read_frame(ctx, pkt) >= 0) {

        if (pkt->size) {
            //printf("----new frame----\n");
            *h_pic_types = decode_packet(videoCodecCtx, pkt, frame);
            printf("----new frame----: %c\n", *h_pic_types);
            h_pic_types++;
            count++;
        }
    }

    if (av_seek_frame(ctx, 0, 0, AVSEEK_FLAG_BACKWARD) < 0) {
        exit(9);
    }
    while (av_read_frame(ctx, pkt) >= 0) {

        if (pkt->size) {
            printf("print side data\n");
            //printf("----new frame----\n");
            AVDictionary *dict = NULL;
            int metadata_size = 0;
            uint8_t *metadata_st = av_packet_get_side_data(pkt, AV_PKT_DATA_METADATA_UPDATE, &metadata_size);
            if (metadata_size) {
                av_dict_parse_string(&dict, (char *) metadata_st, "=", ":", 0);
                if (dict) {
                    AVDictionaryEntry *entry = NULL;
                    while ((entry = av_dict_get(dict, "", entry, AV_DICT_IGNORE_SUFFIX))) {
                        printf("Key: %s, Value: %s\n", entry->key, entry->value);
                    }
                }
            }
        }
    }

    printf("%s\n", pict_types);
    return pict_types;
}

int init_decoder(AVFormatContext *ctx, AVCodecContext **codec_ctx, int stream_index) {
    AVStream *stream = ctx->streams[stream_index];
    AVCodec *streamCodec = avcodec_find_decoder(stream->codecpar->codec_id);
    if (!streamCodec) {
        fprintf(stderr, "failed to find codec");
        return -1;
    }
    *codec_ctx = avcodec_alloc_context3(streamCodec);
    /* Copy codec parameters from input stream to output codec context */
    if (avcodec_parameters_to_context(*codec_ctx, stream->codecpar) < 0) {
        perror("");
        exit(2);
    }
    if (avcodec_open2(*codec_ctx, streamCodec, NULL) < 0) {
        perror("");
        exit(1);
    }
    return -1;
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