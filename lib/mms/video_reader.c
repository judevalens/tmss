//
// Created by Jude Paulemon on 4/29/2022.
//
#include "video_reader.h"

AVFormatContext* open_media (char* mediaPath){
    AVFormatContext* mediaFmtCx = avformat_alloc_context();
    avformat_open_input(&mediaFmtCx,mediaPath,NULL,NULL);
    return mediaFmtCx;
}