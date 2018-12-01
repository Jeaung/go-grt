GRT_PIPELINE_FILE_V3.0
PipelineMode: CLASSIFICATION_MODE
NumPreprocessingModules: 0
NumFeatureExtractionModules: 0
NumPostprocessingModules: 0
Trained: 1
Info: 
PreProcessingModuleDatatypes:
FeatureExtractionModuleDatatypes:
ClassificationModuleDatatype:	DTW
PostProcessingModuleDatatypes:
GRT_DTW_Model_File_V2.0
Trained: 1
UseScaling: 0
NumInputDimensions: 3
NumOutputDimensions: 0
NumTrainingIterationsToConverge: 0
MinNumEpochs: 0
MaxNumEpochs: 100
ValidationSetSize: 20
LearningRate: 0.1
MinChange: 1e-05
UseValidationSet: 0
RandomiseTrainingOrder: 1
UseNullRejection: 1
ClassifierMode: 1
NullRejectionCoeff: 3
NumClasses: 2
NullRejectionThresholds:  55.9195 61.4887
ClassLabels:  1 2
DistanceMethod: 1
UseSmoothing: 0
SmoothingFactor: 5
UseZNormalisation: 0
OffsetUsingFirstSample: 1
ConstrainWarpingPath: 1
Radius: 0.2
RejectionMode: 0
NumberOfTemplates: 2
OverallAverageTemplateLength: 7
***************TEMPLATE***************
Template: 1
ClassLabel: 1
TimeSeriesLength: 5
TemplateThreshold: 55.9195
TrainingMu: 30.6307
TrainingSigma: 8.4296
AverageTemplateLength: 7
TimeSeries: 
0	0	0	
-4.94255	0.0784531	-2.00056	
-3.76575	16.9851	2.03978	
-12.4348	27.0271	5.53095	
-3.96189	27.0271	7.57073	
***************TEMPLATE***************
Template: 2
ClassLabel: 2
TimeSeriesLength: 5
TemplateThreshold: 61.4887
TrainingMu: 25.8699
TrainingSigma: 11.8729
AverageTemplateLength: 7
TimeSeries: 
0	0	0	
-17.6127	6.74698	5.17791	
0.156907	3.64807	-16.2398	
0.902212	12.8271	-25.7719	
1.84365	11.6503	-25.968	
