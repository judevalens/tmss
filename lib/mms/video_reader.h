//
// Created by judev on 4/29/2022.
//

#ifndef MMS_VIDEO_READER_H
#define MMS_VIDEO_READER_H
#include "libavutil/avutil.h"
#include "libavformat/avformat.h"
#include "libavformat/avio.h"
#include "libavutil/avstring.h"
#include "libavcodec/avcodec.h"
#endif //MMS_VIDEO_READER_H

 struct MediaBuffer{
    AVFormatContext *mediaContext;
    struct FrameBuffer *videoBuffer;
    int videoStreamIndex;
    int audioStreamIndex;
};

struct FrameBuffer {
    void **packets;
    int start;
    int end;
    int count;
    int maxSize;
};




AVFormatContext* open_media (char *mediaPath);
struct MediaBuffer* init_media_buffer(char *mediaPath);
void buffer(struct MediaBuffer *buffer, int percent);
int bufferUP(struct MediaBuffer *buffer);
void *circularBuffGet(struct FrameBuffer *buffer);
void circularBufferAdd(struct FrameBuffer *buffer, void *packet);
void demux_file(AVFormatContext *mediaContext, char *OutUrl);