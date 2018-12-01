//
// Created by Jeaung on 11/29/18.
//

#ifndef CGOTEST_WGRT_H
#define CGOTEST_WGRT_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    float x;
    float y;
    float z;
} Point;

typedef Point* Sample;

typedef Sample* TrainingSet;

void init(int);

void addSample(unsigned int, Point[], int);

void train();

float predict(Point*);

#ifdef __cplusplus
}
#endif

#endif //CGOTEST_WGRT_H
