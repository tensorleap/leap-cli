/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// NodeType the model 'NodeType'
type NodeType string

// List of NodeType
const (
	NODETYPE_ACTIVATION                     NodeType = "Activation"
	NODETYPE_ACTIVITY_REGULARIZATION        NodeType = "ActivityRegularization"
	NODETYPE_ADD                            NodeType = "Add"
	NODETYPE_ADDITIVE_ATTENTION             NodeType = "AdditiveAttention"
	NODETYPE_ALPHA_DROPOUT                  NodeType = "AlphaDropout"
	NODETYPE_AVERAGE                        NodeType = "Average"
	NODETYPE_AVERAGE_POOLING1_D             NodeType = "AveragePooling1D"
	NODETYPE_AVERAGE_POOLING2_D             NodeType = "AveragePooling2D"
	NODETYPE_AVERAGE_POOLING3_D             NodeType = "AveragePooling3D"
	NODETYPE_BATCH_NORMALIZATION            NodeType = "BatchNormalization"
	NODETYPE_BIDIRECTIONAL                  NodeType = "Bidirectional"
	NODETYPE_CATEGORY_ENCODING              NodeType = "CategoryEncoding"
	NODETYPE_CENTER_CROP                    NodeType = "CenterCrop"
	NODETYPE_CONCATENATE                    NodeType = "Concatenate"
	NODETYPE_CONV1_D                        NodeType = "Conv1D"
	NODETYPE_CONV1_D_TRANSPOSE              NodeType = "Conv1DTranspose"
	NODETYPE_CONV2_D                        NodeType = "Conv2D"
	NODETYPE_CONV2_D_TRANSPOSE              NodeType = "Conv2DTranspose"
	NODETYPE_CONV3_D                        NodeType = "Conv3D"
	NODETYPE_CONV3_D_TRANSPOSE              NodeType = "Conv3DTranspose"
	NODETYPE_CONV_LSTM1_D                   NodeType = "ConvLSTM1D"
	NODETYPE_CONV_LSTM2_D                   NodeType = "ConvLSTM2D"
	NODETYPE_CONV_LSTM3_D                   NodeType = "ConvLSTM3D"
	NODETYPE_CROPPING1_D                    NodeType = "Cropping1D"
	NODETYPE_CROPPING2_D                    NodeType = "Cropping2D"
	NODETYPE_CROPPING3_D                    NodeType = "Cropping3D"
	NODETYPE_CU_DNNGRU                      NodeType = "CuDNNGRU"
	NODETYPE_CU_DNNLSTM                     NodeType = "CuDNNLSTM"
	NODETYPE_DENSE                          NodeType = "Dense"
	NODETYPE_DEPTHWISE_CONV1_D              NodeType = "DepthwiseConv1D"
	NODETYPE_DEPTHWISE_CONV2_D              NodeType = "DepthwiseConv2D"
	NODETYPE_DISCRETIZATION                 NodeType = "Discretization"
	NODETYPE_DOT                            NodeType = "Dot"
	NODETYPE_DROPOUT                        NodeType = "Dropout"
	NODETYPE_ELU                            NodeType = "ELU"
	NODETYPE_EMBEDDING                      NodeType = "Embedding"
	NODETYPE_FLATTEN                        NodeType = "Flatten"
	NODETYPE_GRU                            NodeType = "GRU"
	NODETYPE_GRU_CELL                       NodeType = "GRUCell"
	NODETYPE_GRU_CELL_V1                    NodeType = "GRUCellV1"
	NODETYPE_GRU_CELL_V2                    NodeType = "GRUCellV2"
	NODETYPE_GRUV1                          NodeType = "GRUV1"
	NODETYPE_GRUV2                          NodeType = "GRUV2"
	NODETYPE_GAUSSIAN_DROPOUT               NodeType = "GaussianDropout"
	NODETYPE_GAUSSIAN_NOISE                 NodeType = "GaussianNoise"
	NODETYPE_GLOBAL_AVERAGE_POOLING1_D      NodeType = "GlobalAveragePooling1D"
	NODETYPE_GLOBAL_AVERAGE_POOLING2_D      NodeType = "GlobalAveragePooling2D"
	NODETYPE_GLOBAL_AVERAGE_POOLING3_D      NodeType = "GlobalAveragePooling3D"
	NODETYPE_GLOBAL_MAX_POOLING1_D          NodeType = "GlobalMaxPooling1D"
	NODETYPE_GLOBAL_MAX_POOLING2_D          NodeType = "GlobalMaxPooling2D"
	NODETYPE_GLOBAL_MAX_POOLING3_D          NodeType = "GlobalMaxPooling3D"
	NODETYPE_GROUP_NORMALIZATION            NodeType = "GroupNormalization"
	NODETYPE_INTEGER_LOOKUP                 NodeType = "IntegerLookup"
	NODETYPE_LSTM                           NodeType = "LSTM"
	NODETYPE_LSTM_CELL                      NodeType = "LSTMCell"
	NODETYPE_LSTM_CELL_V1                   NodeType = "LSTMCellV1"
	NODETYPE_LSTM_CELL_V2                   NodeType = "LSTMCellV2"
	NODETYPE_LSTMV1                         NodeType = "LSTMV1"
	NODETYPE_LSTMV2                         NodeType = "LSTMV2"
	NODETYPE_LAYER_NORMALIZATION            NodeType = "LayerNormalization"
	NODETYPE_LEAKY_RE_LU                    NodeType = "LeakyReLU"
	NODETYPE_LOCALLY_CONNECTED1_D           NodeType = "LocallyConnected1D"
	NODETYPE_LOCALLY_CONNECTED2_D           NodeType = "LocallyConnected2D"
	NODETYPE_MASKING                        NodeType = "Masking"
	NODETYPE_MAX_POOLING1_D                 NodeType = "MaxPooling1D"
	NODETYPE_MAX_POOLING2_D                 NodeType = "MaxPooling2D"
	NODETYPE_MAX_POOLING3_D                 NodeType = "MaxPooling3D"
	NODETYPE_MAXIMUM                        NodeType = "Maximum"
	NODETYPE_MINIMUM                        NodeType = "Minimum"
	NODETYPE_MULTI_HEAD_ATTENTION           NodeType = "MultiHeadAttention"
	NODETYPE_MULTIPLY                       NodeType = "Multiply"
	NODETYPE_NORMALIZATION                  NodeType = "Normalization"
	NODETYPE_PRE_LU                         NodeType = "PReLU"
	NODETYPE_PERMUTE                        NodeType = "Permute"
	NODETYPE_RANDOM_BRIGHTNESS              NodeType = "RandomBrightness"
	NODETYPE_RANDOM_CONTRAST                NodeType = "RandomContrast"
	NODETYPE_RANDOM_CROP                    NodeType = "RandomCrop"
	NODETYPE_RANDOM_FLIP                    NodeType = "RandomFlip"
	NODETYPE_RANDOM_FOURIER_FEATURES        NodeType = "RandomFourierFeatures"
	NODETYPE_RANDOM_HEIGHT                  NodeType = "RandomHeight"
	NODETYPE_RANDOM_ROTATION                NodeType = "RandomRotation"
	NODETYPE_RANDOM_TRANSLATION             NodeType = "RandomTranslation"
	NODETYPE_RANDOM_WIDTH                   NodeType = "RandomWidth"
	NODETYPE_RANDOM_ZOOM                    NodeType = "RandomZoom"
	NODETYPE_RE_LU                          NodeType = "ReLU"
	NODETYPE_REPEAT_VECTOR                  NodeType = "RepeatVector"
	NODETYPE_RESHAPE                        NodeType = "Reshape"
	NODETYPE_RESIZING                       NodeType = "Resizing"
	NODETYPE_SEPARABLE_CONV1_D              NodeType = "SeparableConv1D"
	NODETYPE_SEPARABLE_CONV2_D              NodeType = "SeparableConv2D"
	NODETYPE_SIMPLE_RNN                     NodeType = "SimpleRNN"
	NODETYPE_SIMPLE_RNN_CELL                NodeType = "SimpleRNNCell"
	NODETYPE_SOFTMAX                        NodeType = "Softmax"
	NODETYPE_SPATIAL_DROPOUT1_D             NodeType = "SpatialDropout1D"
	NODETYPE_SPATIAL_DROPOUT2_D             NodeType = "SpatialDropout2D"
	NODETYPE_SPATIAL_DROPOUT3_D             NodeType = "SpatialDropout3D"
	NODETYPE_STRING_LOOKUP                  NodeType = "StringLookup"
	NODETYPE_SUBTRACT                       NodeType = "Subtract"
	NODETYPE_SYNC_BATCH_NORMALIZATION       NodeType = "SyncBatchNormalization"
	NODETYPE_TEXT_VECTORIZATION             NodeType = "TextVectorization"
	NODETYPE_THRESHOLDED_RE_LU              NodeType = "ThresholdedReLU"
	NODETYPE_TIME_DISTRIBUTED               NodeType = "TimeDistributed"
	NODETYPE_UNIT_NORMALIZATION             NodeType = "UnitNormalization"
	NODETYPE_UP_SAMPLING1_D                 NodeType = "UpSampling1D"
	NODETYPE_UP_SAMPLING2_D                 NodeType = "UpSampling2D"
	NODETYPE_UP_SAMPLING3_D                 NodeType = "UpSampling3D"
	NODETYPE_ZERO_PADDING1_D                NodeType = "ZeroPadding1D"
	NODETYPE_ZERO_PADDING2_D                NodeType = "ZeroPadding2D"
	NODETYPE_ZERO_PADDING3_D                NodeType = "ZeroPadding3D"
	NODETYPE_BINARY_CROSSENTROPY            NodeType = "BinaryCrossentropy"
	NODETYPE_BINARY_FOCAL_CROSSENTROPY      NodeType = "BinaryFocalCrossentropy"
	NODETYPE_CATEGORICAL_CROSSENTROPY       NodeType = "CategoricalCrossentropy"
	NODETYPE_CATEGORICAL_HINGE              NodeType = "CategoricalHinge"
	NODETYPE_COSINE_SIMILARITY              NodeType = "CosineSimilarity"
	NODETYPE_HINGE                          NodeType = "Hinge"
	NODETYPE_HUBER                          NodeType = "Huber"
	NODETYPE_KL_DIVERGENCE                  NodeType = "KLDivergence"
	NODETYPE_LOG_COSH                       NodeType = "LogCosh"
	NODETYPE_MEAN_ABSOLUTE_ERROR            NodeType = "MeanAbsoluteError"
	NODETYPE_MEAN_ABSOLUTE_PERCENTAGE_ERROR NodeType = "MeanAbsolutePercentageError"
	NODETYPE_MEAN_SQUARED_ERROR             NodeType = "MeanSquaredError"
	NODETYPE_MEAN_SQUARED_LOGARITHMIC_ERROR NodeType = "MeanSquaredLogarithmicError"
	NODETYPE_POISSON                        NodeType = "Poisson"
	NODETYPE_SQUARED_HINGE                  NodeType = "SquaredHinge"
	NODETYPE_ADADELTA                       NodeType = "Adadelta"
	NODETYPE_ADAGRAD                        NodeType = "Adagrad"
	NODETYPE_ADAM                           NodeType = "Adam"
	NODETYPE_ADAMAX                         NodeType = "Adamax"
	NODETYPE_FTRL                           NodeType = "Ftrl"
	NODETYPE_NADAM                          NodeType = "Nadam"
	NODETYPE_RM_SPROP                       NodeType = "RMSprop"
	NODETYPE_SGD                            NodeType = "SGD"
	NODETYPE_ONNX_ABS                       NodeType = "OnnxAbs"
	NODETYPE_ONNX_ERF                       NodeType = "OnnxErf"
	NODETYPE_ONNX_HARD_SIGMOID              NodeType = "OnnxHardSigmoid"
	NODETYPE_ONNX_LSTM                      NodeType = "OnnxLSTM"
	NODETYPE_ONNX_REDUCE_MEAN               NodeType = "OnnxReduceMean"
	NODETYPE_ONNX_SQRT                      NodeType = "OnnxSqrt"
	NODETYPE_DATASET                        NodeType = "Dataset"
	NODETYPE_INPUT                          NodeType = "Input"
	NODETYPE_REPRESENTATION_BLOCK           NodeType = "RepresentationBlock"
	NODETYPE_GROUND_TRUTH                   NodeType = "GroundTruth"
	NODETYPE_CUSTOM_LAYER                   NodeType = "CustomLayer"
	NODETYPE_CUSTOM_LOSS                    NodeType = "CustomLoss"
	NODETYPE_VISUALIZER                     NodeType = "Visualizer"
	NODETYPE_METRIC                         NodeType = "Metric"
	NODETYPE_LAMBDA                         NodeType = "Lambda"
	NODETYPE_TFOP_LAMBDA                    NodeType = "TFOpLambda"
	NODETYPE_SLICING_OP_LAMBDA              NodeType = "SlicingOpLambda"
	NODETYPE_REPEAT                         NodeType = "Repeat"
	NODETYPE_VARIABLE                       NodeType = "Variable"
	NODETYPE_GATHER                         NodeType = "Gather"
	NODETYPE_FIXED_DROPOUT                  NodeType = "FixedDropout"
)

// All allowed values of NodeType enum
var AllowedNodeTypeEnumValues = []NodeType{
	"Activation",
	"ActivityRegularization",
	"Add",
	"AdditiveAttention",
	"AlphaDropout",
	"Average",
	"AveragePooling1D",
	"AveragePooling2D",
	"AveragePooling3D",
	"BatchNormalization",
	"Bidirectional",
	"CategoryEncoding",
	"CenterCrop",
	"Concatenate",
	"Conv1D",
	"Conv1DTranspose",
	"Conv2D",
	"Conv2DTranspose",
	"Conv3D",
	"Conv3DTranspose",
	"ConvLSTM1D",
	"ConvLSTM2D",
	"ConvLSTM3D",
	"Cropping1D",
	"Cropping2D",
	"Cropping3D",
	"CuDNNGRU",
	"CuDNNLSTM",
	"Dense",
	"DepthwiseConv1D",
	"DepthwiseConv2D",
	"Discretization",
	"Dot",
	"Dropout",
	"ELU",
	"Embedding",
	"Flatten",
	"GRU",
	"GRUCell",
	"GRUCellV1",
	"GRUCellV2",
	"GRUV1",
	"GRUV2",
	"GaussianDropout",
	"GaussianNoise",
	"GlobalAveragePooling1D",
	"GlobalAveragePooling2D",
	"GlobalAveragePooling3D",
	"GlobalMaxPooling1D",
	"GlobalMaxPooling2D",
	"GlobalMaxPooling3D",
	"GroupNormalization",
	"IntegerLookup",
	"LSTM",
	"LSTMCell",
	"LSTMCellV1",
	"LSTMCellV2",
	"LSTMV1",
	"LSTMV2",
	"LayerNormalization",
	"LeakyReLU",
	"LocallyConnected1D",
	"LocallyConnected2D",
	"Masking",
	"MaxPooling1D",
	"MaxPooling2D",
	"MaxPooling3D",
	"Maximum",
	"Minimum",
	"MultiHeadAttention",
	"Multiply",
	"Normalization",
	"PReLU",
	"Permute",
	"RandomBrightness",
	"RandomContrast",
	"RandomCrop",
	"RandomFlip",
	"RandomFourierFeatures",
	"RandomHeight",
	"RandomRotation",
	"RandomTranslation",
	"RandomWidth",
	"RandomZoom",
	"ReLU",
	"RepeatVector",
	"Reshape",
	"Resizing",
	"SeparableConv1D",
	"SeparableConv2D",
	"SimpleRNN",
	"SimpleRNNCell",
	"Softmax",
	"SpatialDropout1D",
	"SpatialDropout2D",
	"SpatialDropout3D",
	"StringLookup",
	"Subtract",
	"SyncBatchNormalization",
	"TextVectorization",
	"ThresholdedReLU",
	"TimeDistributed",
	"UnitNormalization",
	"UpSampling1D",
	"UpSampling2D",
	"UpSampling3D",
	"ZeroPadding1D",
	"ZeroPadding2D",
	"ZeroPadding3D",
	"BinaryCrossentropy",
	"BinaryFocalCrossentropy",
	"CategoricalCrossentropy",
	"CategoricalHinge",
	"CosineSimilarity",
	"Hinge",
	"Huber",
	"KLDivergence",
	"LogCosh",
	"MeanAbsoluteError",
	"MeanAbsolutePercentageError",
	"MeanSquaredError",
	"MeanSquaredLogarithmicError",
	"Poisson",
	"SquaredHinge",
	"Adadelta",
	"Adagrad",
	"Adam",
	"Adamax",
	"Ftrl",
	"Nadam",
	"RMSprop",
	"SGD",
	"OnnxAbs",
	"OnnxErf",
	"OnnxHardSigmoid",
	"OnnxLSTM",
	"OnnxReduceMean",
	"OnnxSqrt",
	"Dataset",
	"Input",
	"RepresentationBlock",
	"GroundTruth",
	"CustomLayer",
	"CustomLoss",
	"Visualizer",
	"Metric",
	"Lambda",
	"TFOpLambda",
	"SlicingOpLambda",
	"Repeat",
	"Variable",
	"Gather",
	"FixedDropout",
}

func (v *NodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NodeType(value)
	for _, existing := range AllowedNodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NodeType", value)
}

// NewNodeTypeFromValue returns a pointer to a valid NodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNodeTypeFromValue(v string) (*NodeType, error) {
	ev := NodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NodeType: valid values are %v", v, AllowedNodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NodeType) IsValid() bool {
	for _, existing := range AllowedNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeType value
func (v NodeType) Ptr() *NodeType {
	return &v
}

type NullableNodeType struct {
	value *NodeType
	isSet bool
}

func (v NullableNodeType) Get() *NodeType {
	return v.value
}

func (v *NullableNodeType) Set(val *NodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeType(val *NodeType) *NullableNodeType {
	return &NullableNodeType{value: val, isSet: true}
}

func (v NullableNodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
