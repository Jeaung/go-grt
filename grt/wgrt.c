//
// Created by Jeaung on 11/29/18.
//

#include "wgrt.h"
#include "GRT/GRT.h"
#include <vector>

GRT::TimeSeriesClassificationData trainingSet;
GRT::GestureRecognitionPipeline pipeline;

#define MODE_MODEL 1
#define MODE_PREDICT 2

#define MODEL_NAME "DTWModel.grt"

void init(int mode) {
    if (mode == MODE_MODEL) {
        trainingSet.setNumDimensions(3);

        //Initialize the DTW classifier
        GRT::DTW dtw;

        //Turn on null rejection, this lets the classifier output the predicted class label of 0 when the likelihood of a gesture is low
        dtw.enableNullRejection(true);

        //Set the null rejection coefficient to 3, this controls the thresholds for the automatic null rejection
        //You can increase this value if you find that your real-time gestures are not being recognized
        //If you are getting too many false positives then you should decrease this value
        dtw.setNullRejectionCoeff(2);

        //Turn on the automatic data triming, this will remove any sections of none movement from the start and end of the training samples
        dtw.enableTrimTrainingData(true, 0.1, 90);

        //Offset the timeseries data by the first sample, this makes your gestures (more) invariant to the location the gesture is performed
        dtw.setOffsetTimeseriesUsingFirstSample(true);

        pipeline.addPostProcessingModule(GRT::ClassLabelTimeoutFilter(300));

        //Add the classifier to the pipeline (after we do this, we don't need the DTW classifier anymore)
        pipeline.setClassifier(dtw);
    } else if (mode == MODE_PREDICT) {
        if (pipeline.load(MODEL_NAME)) {
            printf("pipeline loaded successfully %d classes\n", pipeline.getNumClasses());
        } else {
            printf("failed to load pipeline\n");
        }
    }
}

void addSample(unsigned int label, Point sample[], int pointNum) {
    GRT::MatrixFloat matrix(pointNum, 3);

    for (int i = 0; i < pointNum; i++) {
        Point p = sample[i];

        printf("%.2f %.2f %.2f added label %d\n", p.x, p.y, p.z, label);

        std::vector<double> data;
        data.push_back(p.x);
        data.push_back(p.y);
        data.push_back(p.z);

        matrix.push_back(data);
    }

    trainingSet.addSample(label, matrix);
}

void train() {
    trainingSet.saveDatasetToFile("trainingData");
    if(pipeline.train(trainingSet)){
    	printf("training %d samples successfully\n", trainingSet.getNumSamples());
    	trainingSet.printStats();
	} else {
	    printf("Failed to train classifier!\n");
	}
}

void saveModel() {
    if(pipeline.save(MODEL_NAME)){
	    printf("classifier pipeline saved successfully!\n");
    } else {
        printf("Failed to save the classifier pipeline!\n");
    }
}

float predict(Point *p) {
    GRT::VectorFloat data;
    data.push_back(p->x);
    data.push_back(p->y);
    data.push_back(p->z);

    if (pipeline.getTrained()) {
        if (pipeline.predict(data)) {
            float likelihood = pipeline.getMaximumLikelihood();
            unsigned int label = pipeline.getPredictedClassLabel();
            return likelihood / 10 + label;
        } else {
            printf("failed to predict %.2f %.2f %.2f\n", p->x, p->y, p->z);
        }
    } else {
        printf("pipeline not trained");
    }

    return 0;
}
