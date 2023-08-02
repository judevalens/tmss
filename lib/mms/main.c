#include <stdio.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "libavformat/avformat.h"
#include "libavutil/timestamp.h"
#include "libavutil/dict.h"
#include "libavcodec/avcodec.h"
#include "video_reader.h"
#define VIDEO_SAMPLE_URL "/home/jude/Desktop/amnis_server/big_buck_bunny.mp4"
#define OUT_VIDEO_URL "file:home/jude/Downloads/buck_bunny.mkv"
#define OUT_VIDEO_PATH "home/jude/Downloads/buck_bunny.mkv"
#define OUT_FORMAT "MKV"
#define VID "/home/jude/Desktop/amnis_server/big_buck_bunny_244e2a14a22_stream_0.mp4"
AVFormatContext* openMedia();
void* worker(gpointer data);
int main() {
    g_path_is_absolute("sk");
  // AVFormatContext *media_ctx = open_media(VIDEO_SAMPLE_URL);
   //demux_file(media_ctx);
   int nthreads = 100;
    int *thread_ids = malloc(sizeof (int)*nthreads);
    GThread **threads = malloc(sizeof (GThread)*nthreads);
   for (int i = 0; i < nthreads; i++) {
       thread_ids[i] = i;
       threads[i] = g_thread_new("my thread",worker,&thread_ids[i]);
   }



   for (int i = 0; i < nthreads; i++) {
       g_thread_join(threads[i]);
   }
    printf("waiting for threads");
    return 0;
}

void* worker(gpointer data) {
    int val = *(int*)(data);
    GRand *rand = g_rand_new();
    int rand_val;
    do {
        rand_val = g_rand_int_range(rand,0,2000);
        g_usleep(1000*1);
        //printf("rand val: %d\n",rand_val);
    } while (rand_val <= 99 || rand_val >= 101);

    printf("thread: %d\n",val);

}

static void print_dict(const AVDictionary *m);
static void print_dict(const AVDictionary *m) {
    AVDictionaryEntry *t = NULL;
    while ((t = av_dict_get(m, "", t, AV_DICT_IGNORE_SUFFIX)))
        printf("%s %s   ", t->key, t->value);
    printf("\n");
}