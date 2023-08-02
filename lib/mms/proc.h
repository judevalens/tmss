//
// Created by jude on 8/2/23.
//

#ifndef MMS_PROC_H
#define MMS_PROC_H
#include "glib.h"
#define N_DECODE_WORKER 10

typedef struct {
    GAsyncQueue * queue;
    int capacity;
    GRWLock* grwLock;
} *BufferedAsyncQueue;

BufferedAsyncQueue init_buffered_async_queue(int capacity) {
    BufferedAsyncQueue queue = malloc(sizeof(*queue));
    queue->queue = g_async_queue_new();
    queue->capacity = capacity;
}

struct PDecoder {
    int n_worker;
    GThread ** decode_workers;
    GThread *assembler;
};
#endif //MMS_PROC_H
