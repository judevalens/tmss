#include <stdio.h>
#include <string.h>
#include "libavformat/avformat.h"
#include "libavutil/timestamp.h"
#include "libavutil/dict.h"
#include "libavcodec/avcodec.h"
#include "video_reader.h"

#define VIDEO_SAMPLE_URL "file:/home/jude/Downloads/buck_bunny.mp4"
#define OUT_VIDEO_URL "file:home/jude/Downloads/buck_bunny.mkv"
#define OUT_VIDEO_PATH "home/jude/Downloads/buck_bunny.mkv"
#define OUT_FORMAT "MKV"

AVFormatContext* openMedia();

int main() {
    printf("codec name: %s\n", avcodec_get_name( AV_CODEC_ID_AAC_LATM));
   AVFormatContext  *ctx = open_media("/home/jude/Desktop/amnis_server/big_buck_bunny.mp4");
   // printf("demuxing\n");
 //  demux_file(ctx);
 int buffSize =  1024;
//    MediaBuffer  b = init_media_buffer("/home/jude/Desktop/amnis_server/big_buck_bunny.mp4",buffSize);
    return 0;
}

static void print_dict(const AVDictionary *m);

static void transcode(AVFormatContext *input_fmt_ctx, char *out_fmt, char *out_path);

AVFormatContext* openMedia() {
    AVFormatContext *mediaContext = avformat_alloc_context();
    printf("open media: %s,\n", av_err2str(avformat_open_input(&mediaContext, VIDEO_SAMPLE_URL, NULL, NULL)));
    printf("file name:: %s, extension %s, n stream: %d\n", mediaContext->url, mediaContext->iformat->name,
           mediaContext->nb_streams);
    for (int i = 0; i < mediaContext->nb_streams; i++) {
        AVStream *current_stream = mediaContext->streams[i];
        /*const AVCodecDescriptor *code_desc = avcodec_descriptor_get(current_stream->codecpar->codec_id);
        printf("# %d, codec type %s, media_type %s \n",current_stream->id,code_desc->name,av_get_media_type_string(code_desc->type));*/
        av_dump_format(mediaContext, i, mediaContext->url, 0);
    }
    return mediaContext;
}

void transcode(AVFormatContext *input_fmt_ctx, char *out_fmt, char *out_path) {
    printf("transcoding to %s\n", out_fmt);
    int err = 0;
    AVPacket *av_packet = av_packet_alloc();
    AVFormatContext *out_ctx;
    AVOutputFormat *av_out_fmt = av_guess_format(out_fmt, out_path, NULL);

    err = avformat_alloc_output_context2(&out_ctx, av_out_fmt, out_fmt, out_path);

    if (err < 0) {
        printf("failed to allocate output ctx\nerr: %s\n", av_err2str(err));
    }
    for (int i = 0; i < input_fmt_ctx->nb_streams; i++) {
        if (!avformat_new_stream(out_ctx, NULL)) {
            printf("failed to add stream to output ctx\n");
            return;
        }
        err = avcodec_parameters_copy(out_ctx->streams[i]->codecpar, input_fmt_ctx->streams[i]->codecpar);
        if (err) {
            printf("failed to copy codec param for stream # %d\nerr:%s", i, av_err2str(err));
            return;
        }
        out_ctx->streams[i]->codecpar->codec_tag = 0;
        av_dump_format(out_ctx, i, OUT_VIDEO_URL, 1);
    }

    if (!av_out_fmt) {
        printf("err: could not create AVOutputFormat struct\n");
        return;
    }

    err = avio_open(&out_ctx->pb, OUT_VIDEO_URL, AVIO_FLAG_WRITE);
    if (err) {
        printf("failed to create IO context for out_fmt_ctx\nerr: %s\n", av_err2str(err));
        return;
    }
    err = avformat_write_header(out_ctx, NULL);
    if (err < 0) {
        printf("failed to write header:\nerr: %s", av_err2str(err));
    }
    while (1) {
        AVStream *in_stream, *out_stream;
        err = av_read_frame(input_fmt_ctx, av_packet);

        if (err < 0) {
            printf("err: %s\n", av_err2str(err));
            break;
        }
        printf("stream index %d, out nb stream %d\n", av_packet->stream_index, out_ctx->nb_streams);

        av_packet_rescale_ts(av_packet, input_fmt_ctx->streams[av_packet->stream_index]->time_base,
                             out_ctx->streams[av_packet->stream_index]->time_base);
        av_packet->pos = -1;
        printf("packet from stream #%d, at %ld\n", av_packet->stream_index, av_packet->pts);

        err = av_interleaved_write_frame(out_ctx, av_packet);
        if (err < 0) {
            printf("err: %s", av_err2str(err));
            break;
        }
    }
    err = av_write_trailer(out_ctx);

    if (err < 0) {
        printf("err: %s\n", av_err2str(err));
    }
}

static void print_dict(const AVDictionary *m) {
    AVDictionaryEntry *t = NULL;
    while ((t = av_dict_get(m, "", t, AV_DICT_IGNORE_SUFFIX)))
        printf("%s %s   ", t->key, t->value);
    printf("\n");
}