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
#include "rtp.h"

#endif //MMS_VIDEO_READER_H

typedef struct PacketBuffer {
    AVPacket **packets;
    int start;
    int end;
    int currentIdx;
    int size;
} *PacketBuffer;


typedef struct  MediaBuffer{
    AVFormatContext *mediaContext;
    PacketBuffer *packetBuffers;
} *MediaBuffer;


AVFormatContext *open_media(char *mediaPath);

MediaBuffer init_media_buffer(char *mediaPath, int bufferSize);

void buffer(MediaBuffer mediaBuffer, int bufferIdx);
int seek(MediaBuffer mediaBuffer, int64_t position);
void demux_file(AVFormatContext *mediaContext, char *OutUrl);