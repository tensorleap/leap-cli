from typing import Union, Dict

import numpy as np
from code_loader.visualizers.default_visualizers import default_image_visualizer
from code_loader.inner_leap_binder.leapbinder_decorators import *




@tensorleap_preprocess()
def preprocess_func_leap() -> List[PreprocessResponse]:
    """
    Prepares dataset splits for Tensorleap integration.

    Called once before training or evaluation to create
    train, validation, test, and unlabeled subsets.

    Example:
        train = PreprocessResponse(len(train_df), train_df, DataStateType.training)
        val = PreprocessResponse(len(val_df), val_df, DataStateType.validation)
        return [train, val]

    Returns:
        List[PreprocessResponse]: Dataset partitions; must include
        at least training and validation sets.
    """
    pass



@tensorleap_input_encoder(name="<place input name here>", channel_dim="<place channel_dim here (i.e: -1)> ")
def input_encoder(idx: int, preprocess: PreprocessResponse) -> np.ndarray:
    """
    Generates an input sample from the preprocessed dataset.

    Called for each sample index during training or evaluation.
    Each input encoder corresponds to one network input.

    Example:
        return preprocess.data.iloc[idx]['samples'].astype('float32')

    Args:
        idx (int): Sample index.
        preprocess (PreprocessResponse): Preprocessed dataset slice.

    Returns:
        np.ndarray: Encoded input sample.
    """
    pass

@tensorleap_gt_encoder(name="<place gt name here>")
def gt_encoder(idx: int, preprocess: Union[PreprocessResponse, list]) -> np.ndarray:
    """
    Generates the ground truth for a given sample index.

    Called for each sample to provide the label or target value
    used by the loss function during training or evaluation.

    Example:
        return preprocess.data.iloc[idx]['ground_truth'].astype('float32')

    Args:
        idx (int): Sample index.
        preprocess (PreprocessResponse | list): Preprocessed dataset slice.

    Returns:
        np.ndarray: Ground truth value for the sample.
    """
    pass

@tensorleap_custom_loss(name="<place loss name here>")
def custom_loss(y_true: np.ndarray, y_pred: np.ndarray) -> np.ndarray:
    """
    Defines a custom loss function for Tensorleap.

    Computes per-sample loss values used during training.
    Supports multiple input arrays exposed in the CustomLoss UI.

    Example:
        y_pred = np.clip(y_pred / np.sum(y_pred, axis=-1, keepdims=True), 1e-7, 1 - 1e-7)
        loss = -np.sum(y_true * np.log(y_pred), axis=-1)
        return loss

    Args:
        y_true (np.ndarray): Ground truth labels.
        y_pred (np.ndarray): Model predictions.

    Returns:
        np.ndarray: Batch loss values.
    """
    pass



@tensorleap_metadata(name="<place metadata name here>")
def custom_metadata(idx: int, preprocess: PreprocessResponse) -> Dict[str, Union[int, bool]]:
    """
    Generates sample-level metadata for Tensorleap analysis.

    Called for each sample to provide additional descriptive data
    such as labels or custom attributes.

    Example:
        return {
            'label': int_metadata_creator(preprocess, idx),
            'is_circle': bool_metadata_creator(preprocess, idx),
        }

    Args:
        idx (int): Sample index.
        preprocess (PreprocessResponse): Preprocessed dataset slice.

    Returns:
        Dict[str, Union[int, bool]]: Flat dictionary of metadata attributes.
    """
    pass

@tensorleap_custom_metric(name="<place metric name here>",
                          direction="< place MetricDirection.<direction name> here >")
def custom_metric(y_true: np.ndarray[np.float32],y_pred: np.ndarray[np.float32]) -> dict[str, np.ndarray[np.float32]]:
    """
    Defines a custom per-sample metric for Tensorleap analysis.

    Computes user-defined metrics from predictions and ground truth.
    Metrics are stored per sample for later inspection and insights.

    Example:
        diff = y_true - y_pred
        return {
            "mean_difference": np.mean(diff, axis=(1,)),
            "mean_absolute_difference": np.mean(np.abs(diff), axis=(1,))
        }

    Args:
        y_true (np.ndarray): Ground truth values.
        y_pred (np.ndarray): Model predictions.

    Returns:
        dict[str, np.ndarray]: Dictionary of computed metric values.
    """
    pass

@tensorleap_custom_visualizer(name="<place metadata name here>", visualizer_type="<place LeapDataType.<datatype> here>")
def custom_visualizer(data: np.float32):
    """
    Defines a custom visualizer for Tensorleap outputs.

    Converts raw tensors into interpretable formats for display
    in the Tensorleap UI (e.g., text, images, or annotated visuals).

    Example:
        return default_image_visualizer(data)

    Args:
        input_ids (np.ndarray): Model tensor to visualize.

    Returns:
        LeapDataType.<datatype>: Visualized representation of the input tensor.
    """
    pass



if __name__ == '__main__':
    leap_binder.check()
