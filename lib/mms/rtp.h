//
// Created by jude on 11/24/22.
//

#ifndef MMS_RTP_H
#define MMS_RTP_H
#include "libavcodec/avcodec.h"
char* get_rtp_payload_format(enum AVCodecID codec_id);
int get_rtp_clock_rate(enum AVCodecID codec_id);
#endif //MMS_RTP_H
