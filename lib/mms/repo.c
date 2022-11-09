//
// Created by jude on 11/6/22.
//

#include "repo.h"
#include "cjson/cJSON.h"
#include <ctype.h>
#include "video_reader.h"

#define SERVERDIR  "Desktop/amnis_server"

const char *getHomeDir();

int addFile(struct mediaInfo mediaInfo) {
    AVFormatContext *mediaCtx = open_media(mediaInfo.filePath);

    FILE *mediaRepo = fopen(getHomeDir(), "rw+");
    if (!mediaRepo) {
        perror(strerror(errno));
        return -1;
    }
    fseek(mediaRepo, 0, SEEK_END);
    long fileSize = ftell(mediaRepo) + 1;
    fseek(mediaRepo,0,SEEK_SET);
    char *mediaRepoJson = malloc(sizeof(char)*fileSize);
    fread(mediaRepoJson,1,fileSize,mediaRepo);
    cJSON_Parse()

    cJSON *mediaJson = cJSON_CreateObject();
    cJSON_AddStringToObject(mediaJson, "name", av_basename(mediaInfo.filePath));
}

const char *getHomeDir() {
    char *home = getenv("HOME");
    int strLen = strlen(home) + strLen(SERVERDIR);
    char *dir = malloc(sizeof(char) * strLen);
    strcat(dir, home);
    strcat(dir, SERVERDIR);
    return dir;
}